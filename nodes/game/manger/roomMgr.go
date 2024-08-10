package manger

import (
	"context"
	"fmt"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	"superman/nodes/leaf/jettengame/gamedata/goclib/util"
	"sync"
	"sync/atomic"
)

// ----------房间----------------------

// Room 房间信息
type Room struct {
	*protoMsg.RoomInfo
	State         uint8    // 房间状态 [0:无效] [1:Open] [2:Close] [3:Other] [4:Clear]
	tables        []*Table // 桌子集合
	cancelUidList []int64
	onlineCount   int32
	waiting       chan Waiter // 等待匹配的玩家队列
	tbLk          sync.RWMutex
	cancelLk      sync.RWMutex
}

// Waiter 等待玩家
type Waiter struct {
	Uid int64
	Gid int64
}

// IRoomHandle 房间接口
type IRoomHandle interface {
	Join(tid int64, person *Player) *Table        //加入
	Match(gid int64, person *Player) *Table       //配桌
	ChangeTable(gid int64, person *Player) *Table //换桌
	Leave(person *Player)                         // 离开房间
}

// RoomManger 管理房间
type RoomManger struct {
	sync.Map
}

// /////////////房间管理////////////////////////
var rOnce sync.Once
var roomMgr *RoomManger

// GetRoomMgr 由于新增平台，则这里的管理不应该再是单例模式
func GetRoomMgr() *RoomManger {
	rOnce.Do(func() {
		roomMgr = &RoomManger{
			sync.Map{},
		}
	})

	return roomMgr
}

/////////////////////////////////////////[manger]////////////////////////////////////////////

// AddRoom 添加房间
func (self *RoomManger) AddRoom(info *protoMsg.RoomInfo) *Room {
	val, ok := self.Load(info.Id)
	if ok && val != nil {
		log.Infof("[RoomManger][AddRoom]已经存在房间:%d", info.Id)
		return val.(*Room)
	}
	rm := Create(info)
	self.Store(info.Id, rm)
	return rm
}

func (self *RoomManger) AddRooms(infos []*protoMsg.RoomInfo) {
	for _, info := range infos {
		self.AddRoom(info)
	}
}

func (self *RoomManger) GetRoom(rid int64) *Room {
	val, ok := self.Load(rid)
	if ok && val != nil {
		return val.(*Room)
	}
	return nil
}

//////////////////////////////////////[room]////////////////////////////////////////////////////////////

// Create 创建房间[房间ID和钥匙配对]
func Create(info *protoMsg.RoomInfo) *Room {
	rm := &Room{
		RoomInfo: info,
		State:    0,
		tables:   make([]*Table, 0),
		//waiting:  make(chan Waiter, 100),
	}
	if rm.MaxPerson == Unlimited {
		rm.waiting = make(chan Waiter, 1000)
	} else {
		rm.waiting = make(chan Waiter, rm.MaxPerson)
	}
	log.Infof("[%v:%v] >>>开启游戏队列检测<<< 牌桌数:%v 最大牌桌数:%v 最大人数:%v", rm.Id, rm.Name, rm.TableCount, rm.MaxTable, rm.MaxPerson)
	// 创建一个独立的上下文用于这个Room的处理
	ctx, cancel := context.WithCancel(context.Background())
	// 启动协程处理Room的等待队列
	go func(room *Room, ctx context.Context) {
		defer cancel() // 在协程结束时取消上下文
		room.ProcessWaitingQueue(ctx)
	}(rm, ctx)
	return rm
}

// 获取随机桌牌
func getRandomTable(gid, butTid int64, tables []*Table) *Table {
	list := make([]*Table, 0)
	for _, table := range tables {
		// 不能是之前桌
		if table.Gid == gid && table.Id != butTid && table.GameHandle != nil && (table.MaxSitter == Unlimited || table.SitCount()+1 <= table.MaxSitter) {
			list = append(list, table)
		}
	}
	// 随机换桌
	size := len(list)
	if 0 == size {
		return nil
	}
	return list[util.RandIntn(size)]
}

// AddTable 新增桌子(tid == gid 即桌子ID和游戏ID共用)
func (self *Room) AddTable(table *protoMsg.TableInfo, f NewGameFunc) (*Table, error) {
	self.tbLk.Lock()
	defer self.tbLk.Unlock()
	if self.tables == nil {
		self.tables = make([]*Table, 0)
	}
	for _, t := range self.tables {
		if t.Id == table.Id {
			return t, fmt.Errorf(StatusText[TableInfo07])
		}
	}
	if self.MaxTable < int32(len(self.tables)+1) {
		return nil, fmt.Errorf(StatusText[TableInfo08])
	}
	tb := &Table{
		TableInfo: table,
		sitters:   sync.Map{},
		isStart:   false,
	}
	var gameName string
	tb.GameHandle, gameName = f(table.Gid, tb) //创建游戏句柄
	if tb.GameHandle == nil {
		return nil, fmt.Errorf("房间ID:%d 牌桌ID:%d 游戏ID:%d err:%s!!! ", table.Rid, table.Id, table.Gid, StatusText[Room17])
	}
	log.Infof("[%v:%v]房->(%v:%v)桌 <%v:%v> 游戏創建成功!", self.Id, self.Name, tb.Id, tb.Name, table.Gid, gameName)
	self.tables = append(self.tables, tb)
	return tb, nil
}

// AddWait 添加玩家到等待列表
func (self *Room) AddWait(gid int64, person *Player) bool {
	if person.GameHandle != nil {
		return false
	}
	person.State = protoMsg.PlayerState_PlayerWaiting
	self.waiting <- Waiter{
		Gid: gid,
		Uid: person.UserID,
	}
	return true
}

// CancelWait 玩家取消等待
func (self *Room) CancelWait(person *Player) bool {
	if person.GameHandle != nil {
		return false
	}
	person.State = protoMsg.PlayerState_PlayerGiveUp
	self.cancelLk.Lock()
	self.cancelUidList = utils.RemoveValue(self.cancelUidList, person.UserID)
	self.cancelLk.Unlock()
	return true
}

// ProcessWaitingQueue 匹配玩家列表
func (self *Room) ProcessWaitingQueue(ctx context.Context) {
	// 执行一些操作
	if err := ctx.Err(); err != nil {
		return
	}
	for waiter := range self.waiting {
		// 尝试加入房间，如果失败则重新加入队列
		if tb := getRandomTable(waiter.Gid, Unlimited, self.tables); tb != nil {
			if player := GetPlayerMgr().Get(waiter.Uid); player != nil {
				code := tb.AddChair(player)
				if code == TableInfo12 {
					continue
				}
				if code != SUCCESS {
					self.waiting <- waiter
				}
				log.Infof("[ProcessWaitingQueue]玩家:(%d) [gid:%d]配桌成功  [rid:%v][tid:%d] [chair:%d] ", waiter.Uid, waiter.Gid, self.Id, tb.Id, player.InChairId)
			}
		} else {
			self.waiting <- waiter
		}
	}
}

// Join 加入
func (self *Room) Join(tid int64, person *Player) *Table {
	self.tbLk.RLock()
	defer self.tbLk.RUnlock()
	// 玩家还在其他桌且该桌的游戏还没结束
	if 0 < person.InTableId {
		return nil
	}
	var tb *Table
	for _, table := range self.tables {
		if table.Id == tid {
			tb = table
			break
		}
	}
	if tb != nil {
		if code := tb.AddChair(person); code != SUCCESS {
			return nil
		}
		atomic.AddInt32(&self.onlineCount, 1)
	}
	return tb
}

func (self *Room) Match(gid int64, person *Player, isChange bool) *Table {
	self.tbLk.RLock()
	defer self.tbLk.RUnlock()
	strOp := "配桌"
	tid := int64(Unlimited)
	if isChange {
		strOp = "换桌"
		tid = person.InTableId
	}
	tb := getRandomTable(gid, tid, self.tables) // 不限制
	if tb != nil {
		if code := tb.AddChair(person); code != SUCCESS {
			log.Warnf("[room][Match]玩家:(%d) %s失败 rid:%d  gid:%d err:%v", person.UserID, strOp, self.Id, gid, StatusText[code])
			return nil
		}
		atomic.AddInt32(&self.onlineCount, 1)
		log.Infof("[room][Match]玩家:(%d) %s成功 rid:%d  gid:%d ", person.UserID, strOp, self.Id, gid)
	} else {
		log.Warnf("[room][Match]玩家:(%d) %s失败 rid:%d  gid:%d ", person.UserID, strOp, self.Id, gid)
	}
	return tb
}

// Leave 离开房间
func (self *Room) Leave(person *Player) {
	var tb *Table
	self.tbLk.Lock()
	for _, table := range self.tables {
		if table.Id == person.InTableId {
			tb = table
			break
		}
	}
	self.tbLk.Unlock()
	if tb != nil {
		tb.RemoveChair(person.UserID)
	}
	atomic.AddInt32(&self.onlineCount, -1)
}

// CheckGameFull  检测游戏是否满员
func (self *Room) CheckGameFull(gid int64) bool {
	self.tbLk.Lock()
	defer self.tbLk.Unlock()
	for _, table := range self.tables {
		if gid == table.Gid {
			if table.MaxSitter == Unlimited {
				return false
			}
			if table.SitCount() < table.MaxSitter {
				return false
			}
		}
	}
	return true
}

// OnlineCount 当前在线人数
func (self *Room) OnlineCount() int32 {
	return atomic.LoadInt32(&self.onlineCount)
}

// GetTable 获取桌牌
func (self *Room) GetTable(tid int64) *Table {
	self.tbLk.RLock()
	defer self.tbLk.RUnlock()
	if self.tables == nil {
		return nil
	}
	for _, table := range self.tables {
		if table.TableInfo.Id == tid {
			return table
		}
	}
	return nil
}

// GetTableForGid 获取桌牌
func (self *Room) GetTableForGid(gid int64) *Table {
	self.tbLk.RLock()
	defer self.tbLk.RUnlock()
	if self.tables == nil {
		return nil
	}
	for _, table := range self.tables {
		if table.TableInfo.Gid == gid {
			return table
		}
	}
	return nil
}
func (self *Room) GetTablesForGid(gid int64) []*Table {
	self.tbLk.RLock()
	defer self.tbLk.RUnlock()
	tbls := make([]*Table, 0)
	if self.tables == nil {
		return nil
	}
	for _, table := range self.tables {
		if table.TableInfo.Gid == gid {
			tbls = append(tbls, table)
		}
	}
	return tbls
}

// DelTable 删除桌子
func (self *Room) DelTable(tid int64) {
	self.tbLk.Lock()
	defer self.tbLk.Unlock()
	if self.tables == nil {
		return
	}
	for i, item := range self.tables {
		if item.Id == tid {
			if item.GameHandle != nil {
				//  todo 控制关闭游戏
				//item.GameHandle.SuperControl()
			}
			self.tables = append(self.tables[:i], self.tables[i+1:]...)
			return
		}
	}
}

// DelTableForGid  移除指定游戏
func (self *Room) DelTableForGid(gid int64) {
	self.tbLk.Lock()
	defer self.tbLk.Unlock()
	if self.tables == nil {
		return
	}
	for i, item := range self.tables {
		if item.Gid == gid {
			// todo 控制关闭游戏
			//item.GameHandle.SuperControl()
			self.tables = append(self.tables[:i], self.tables[i+1:]...)
			return
		}
	}
}

/////////////////////房间管理功能//////////////////////////////////////////////

// PrintfAll 测试
func (self *RoomManger) PrintfAll() { //打印当前所房间
	log.Debug("所有房间号:->")
	self.Range(func(index, value interface{}) bool {
		log.Debug("index:%v, 房间号码:%v", index, value)
		return true
	})
}

func (self *RoomManger) Create(room *protoMsg.RoomInfo) (*Room, bool) { //新增
	info := &Room{}
	bRet := true
	if v, ok := self.Load(room.Id); !ok {
		log.Debug("创建房间:%v", room.Id)
		info.RoomInfo = room
		self.Store(info.Id, info)
	} else {
		log.Debug("房间:%v 已经存在", room.Id)
		info = v.(*Room)
		bRet = false
	}
	return info, bRet
}

func (self *RoomManger) Del(rid int64, key string) bool { //删除
	bRet := false
	self.Range(func(index, value interface{}) bool {
		if rid == index.(int64) && key == value.(*Room).RoomKey {
			bRet = true
			// 移除桌牌

			//for _,gameHandle :=range value.(*Room).Things.GameSet{
			//}
			self.Delete(index)
			return false
		}
		return true
	})
	return bRet
}

func (self *RoomManger) Check(rid int64) (*Room, bool) { //查找
	if value, ok := self.Load(rid); ok {
		return value.(*Room), ok
	}
	return nil, false
}
func (self *RoomManger) Open(rid int64, key string) (*Room, bool) { //开启

	info := &Room{}

	return info, true
}
func (self *RoomManger) Close(rid int64, key string) bool { //关闭

	return true
}
func (self *RoomManger) Clear(rid int64, key string) bool { //清理房间

	return true
}
