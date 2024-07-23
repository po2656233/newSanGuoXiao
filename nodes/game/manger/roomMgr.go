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
	"time"
)

// ----------房间----------------------

// Room 房间信息
type Room struct {
	*protoMsg.RoomInfo
	State       uint8    // 房间状态 [0:无效] [1:Open] [2:Close] [3:Other] [4:Clear]
	tables      []*Table // 桌子集合
	onlineCount int32
	waiting     []*WaitPlayer // 等待匹配的玩家队列
	tblock      sync.RWMutex
	wplock      sync.RWMutex
}

// WaitPlayer 等待玩家
type WaitPlayer struct {
	*Player
	Gid int64
}

// Table 桌牌
type Table struct {
	*protoMsg.TableInfo
	GameHandle IGameOperate
	sitters    []*Chair  // 座位上的玩家
	lookers    []*Player // 旁观者
	sitCount   int32
	sync.RWMutex
}

// Chair 椅子
type Chair struct {
	*protoMsg.PlayerInfo
	ID       int32       // 椅子号(仅表示序号) 0:空
	HandCard []byte      // 手牌
	Score    int64       // 当前一局的消费额(斗地主叫分阶段为叫分分值)
	Total    int64       // 消费总额
	Gain     int64       // 获利
	Timer    *time.Timer // 座位上的定时器
}

// ZjhChair zjh桌位
type ZjhChair struct {
	*Chair
	Multiple  int     //倍数|番数
	Challenge []int64 //挑战过的玩家
}

// MahjongChair 麻将桌位
type MahjongChair struct {
	*Chair
	Multiple int //倍数|番数
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

// /////////////房间和子游戏管理////////////////////////
var settingFixTime int64

var rOnce sync.Once
var roomMgr *RoomManger

// GetRoomMgr 由于新增平台，则这里的管理不应该再是单例模式
func GetRoomMgr() *RoomManger {
	rOnce.Do(func() {
		roomMgr = &RoomManger{
			sync.Map{},
		}
		go roomMgr.CheckQueue()
	})

	return roomMgr
}

func (self *RoomManger) CheckQueue() {
	// 定期检查等待队列
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Infof("正在定期检测游戏队列")
			self.Range(func(key, value any) bool {
				val, ok := value.(*Room)
				if ok {
					return true
				}
				// 创建一个独立的上下文用于这个Room的处理
				ctx, cancel := context.WithCancel(context.Background())
				// 启动协程处理Room的等待队列
				go func(room *Room, ctx context.Context) {
					defer cancel() // 在协程结束时取消上下文
					room.ProcessWaitingQueue(ctx)
				}(val, ctx)
				return true
			})
		}
	}
}

/////////////////////////////////////////[manger]////////////////////////////////////////////

// AddRoom 添加房间
func (self *RoomManger) AddRoom(info *protoMsg.RoomInfo) {
	val, ok := self.Load(info.Id)
	if ok && val != nil {
		log.Infof("[RoomManger][AddRoom]已经存在房间:%d", info.Id)
		return
	}
	self.Store(info.Id, Create(info))
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
		waiting:  make([]*WaitPlayer, 0),
	}
	return rm
}

// 获取随机桌牌
func getRandomTable(gid, butTid int64, tables []*Table) *Table {
	list := make([]*Table, 0)
	for _, table := range tables {
		// 不能是之前桌
		if table.Gid == gid && table.Id != butTid && table.GameHandle != nil && table.MaxSitter <= table.SitCount()+1 {
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
func (self *Room) AddTable(table *protoMsg.TableInfo, f NewGameCallback) (*Table, error) {
	self.tblock.Lock()
	defer self.tblock.Unlock()
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
		TableInfo:  table,
		GameHandle: f(table.Id, table.Gid), //创建游戏句柄
		sitters:    make([]*Chair, 0),
		lookers:    make([]*Player, 0),
		RWMutex:    sync.RWMutex{},
	}
	if tb.GameHandle == nil {
		return nil, fmt.Errorf("%s 游戏ID:%d", StatusText[TableInfo03], table.Gid)
	}
	self.tables = append(self.tables, tb)
	return tb, nil
}

// AddWait 添加玩家到等待列表
func (self *Room) AddWait(person *WaitPlayer) bool {
	self.wplock.Lock()
	defer self.wplock.Unlock()
	if person.GameHandle != nil {
		return false
	}
	self.waiting = append(self.waiting, person)
	return true
}

// RemoveWaits 移除等待列表
func (self *Room) RemoveWaits() bool {
	self.wplock.Lock()
	defer self.wplock.Unlock()
	for i, player := range self.waiting {
		if player == nil || player.GameHandle != nil {
			self.waiting = append(self.waiting[:i], self.waiting[i+1:]...)
		}
	}
	return true
}

// ProcessWaitingQueue 匹配玩家列表
func (self *Room) ProcessWaitingQueue(ctx context.Context) {
	// 执行一些操作
	if err := ctx.Err(); err != nil {
		return
	}
	self.wplock.Lock()
	defer self.wplock.Unlock()
	// 检查等待队列中的玩家，并尝试将他们加入房间
	for len(self.waiting) > 0 {
		waiter := self.waiting[0]       // 获取队列中的第一个玩家
		self.waiting = self.waiting[1:] // 移除队列中的第一个玩家

		// 尝试加入房间，如果失败则重新加入队列
		tb := getRandomTable(waiter.Gid, Unlimited, self.tables) // 不限制
		if tb != nil {
			err := tb.AddChair(waiter.Player)
			if err != nil {
				// 尝试添加新桌子
				if self.Id == SYSTEMID {

				}
				self.waiting = append(self.waiting, waiter)
				continue
			}
			log.Infof("[room][ProcessWaitingQueue]玩家:(%d) 配桌成功 rid:%d  gid:%d ", waiter.UserID, self.Id, waiter.Gid)
		} else {
			self.waiting = append(self.waiting, waiter)
		}
	}
}

// Join 加入
func (self *Room) Join(tid int64, person *Player) *Table {
	self.tblock.RLock()
	defer self.tblock.RUnlock()
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
		if err := tb.AddChair(person); err != nil {
			return nil
		}
		atomic.AddInt32(&self.onlineCount, 1)
	}
	return tb
}

func (self *Room) Match(gid int64, person *Player, isChange bool) *Table {
	self.tblock.RLock()
	defer self.tblock.RUnlock()
	strOp := "配桌"
	tid := int64(Unlimited)
	if isChange {
		strOp = "换桌"
		tid = person.InTableId
	}
	tb := getRandomTable(gid, tid, self.tables) // 不限制
	if tb != nil {
		if err := tb.AddChair(person); err != nil {
			log.Warnf("[room][Match]玩家:(%d) %s失败 rid:%d  gid:%d err:%v", person.UserID, strOp, self.Id, gid, err)
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
	self.tblock.Lock()
	defer self.tblock.Unlock()
	var tb *Table
	for _, table := range self.tables {
		if table.Id == person.InTableId {
			tb = table
			break
		}
	}
	if tb != nil {
		tb.RemoveChair(person)
	}
	atomic.AddInt32(&self.onlineCount, -1)
}

// CheckGameFull  检测游戏是否满员
func (self *Room) CheckGameFull(gid int64) bool {
	self.tblock.Lock()
	defer self.tblock.Unlock()
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
	self.tblock.RLock()
	defer self.tblock.RUnlock()
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
	self.tblock.RLock()
	defer self.tblock.RUnlock()
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
	self.tblock.RLock()
	defer self.tblock.RUnlock()
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
	self.tblock.Lock()
	defer self.tblock.Unlock()
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
	self.tblock.Lock()
	defer self.tblock.Unlock()
	if self.tables == nil {
		return
	}
	for i, item := range self.tables {
		if item.Gid == gid {
			// todo 控制关闭游戏
			//item.GameHandle.SuperControl()
			self.tables = append(self.tables[:i], self.tables[i+1:]...)
			break
		}
	}
}

/////////////////////桌牌功能//////////////////////////////////////////////

// Init 加载配置 [废弃:不加载公用配置,拆分至子游戏内部实现]
func (tb *Table) Init() {

	//解析游戏
	////获取文件的
	//fInfo, _ := os.Stat(settingFile)
	//lastFixTime := fInfo.ModTime().Unix()
	//// 在文件被修改或配置内容为空时，则读取配置信息
	//if lastFixTime != settingFixTime  {
	//	//读取配置信息
	//	yamlFile, err := os.ReadFile(settingFile)
	//	//log.Debug("yamlFile:%v", yamlFile)
	//	if err != nil {
	//		log.Fatal("yamlFile.Get err #%v ", err)
	//	}
	//
	//	err = yaml.Unmarshal(yamlFile, self.Config)
	//	// err = yaml.Unmarshal(yamlFile, &resultMap)
	//	if err != nil {
	//		log.Error("Unmarshal: %v", err)
	//	}
	//}
	//settingFixTime = lastFixTime

	//log.Debug("conf:%v", self.Config)
}

func (tb *Table) AddChair(player *Player) error {
	tb.Lock()
	defer tb.Unlock()
	if tb.MaxSitter != Unlimited && tb.MaxSitter < int32(len(tb.sitters)+1) {
		return fmt.Errorf("已经满员")
	}
	player.PlayerInfo.State = protoMsg.PlayerState_PlayerSitDown
	tb.sitters = append(tb.sitters, &Chair{
		PlayerInfo: player.PlayerInfo,
		ID:         0,
		HandCard:   make([]byte, 0),
		Score:      0,
		Total:      0,
		Gain:       0,
		Timer:      &time.Timer{},
	})
	tb.sitters = utils.SliceRemoveDuplicate(tb.sitters).([]*Chair)
	atomic.SwapInt32(&tb.sitCount, int32(len(tb.sitters)))
	return nil
}

func (tb *Table) GetChair(uid int64) *Chair {
	tb.RLock()
	defer tb.RUnlock()
	for _, sitter := range tb.sitters {
		if sitter.PlayerInfo.UserID == uid {
			return sitter
		}
	}
	return nil
}

func (tb *Table) GetChairInfos() []*protoMsg.PlayerInfo {
	tb.RLock()
	defer tb.RUnlock()
	list := make([]*protoMsg.PlayerInfo, 0)
	for _, sitter := range tb.sitters {
		list = append(list, sitter.PlayerInfo)
	}
	return list
}

func (tb *Table) RemoveChair(player *Player) {
	tb.Lock()
	defer tb.Unlock()
	player.PlayerInfo.State = protoMsg.PlayerState_PlayerStandUp
	for i, sitter := range tb.sitters {
		if sitter.PlayerInfo.UserID == player.UserID {
			tb.sitters = append(tb.sitters[:i], tb.sitters[i+1:]...)
			break
		}
	}
	tb.sitters = utils.SliceRemoveDuplicate(tb.sitters).([]*Chair)
	atomic.SwapInt32(&tb.sitCount, int32(len(tb.sitters)))
}

// NextChair 轮换规则
func (tb *Table) NextChair(curUid int64) *Chair {
	tb.RLock()
	defer tb.RUnlock()
	size := len(tb.sitters)
	if 0 == size {
		return nil
	}
	index := 0
	for i, sitter := range tb.sitters {
		if sitter.PlayerInfo.UserID == curUid {
			index = i
			break
		}
	}
	index++
	if index <= size {
		index = 0
	}
	for i := 0; i < size; i++ {
		idx := index + i
		if idx <= size {
			idx -= size
		}
		if tb.sitters[idx].State < protoMsg.PlayerState_PlayerTrustee && tb.sitters[idx].State > protoMsg.PlayerState_PlayerAgree {
			return tb.sitters[idx]
		}
	}
	return nil
}
func (tb *Table) SitCount() int32 {
	return atomic.LoadInt32(&tb.sitCount)
}

// Reset 重置座位上的玩家
func (tb *Table) Reset() int64 {
	log.Debugf("[%v:%v]   \t扫地僧出来干活了...List %v", tb.Name, tb.Id, tb.sitters)
	// 遍历过程中存在slice的数据删除
	tb.Lock()
	for _, sit := range tb.sitters {
		sit.State = protoMsg.PlayerState_PlayerSitDown
		sit.Gain = INVALID
		sit.Score = INVALID
		sit.Total = INVALID
		sit.HandCard = make([]byte, 0)
		if sit.Timer != nil {
			sit.Timer.Stop()
			sit.Timer = nil
		}
	}
	tb.Unlock()

	tb.CorrectCommission()
	tb.Init()
	return tb.Id
}

// CorrectCommission 修正税收
func (tb *Table) CorrectCommission() {
	if tb.Commission < 50 {
		tb.Commission = 100 - tb.Commission
	}
	if tb.Commission <= 0 || tb.Commission == 100 {
		tb.Commission = 100
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
