package manger

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	"strings"
	. "superman/internal/constant"
	"superman/internal/hints"
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
	sync.RWMutex
}

// Table 桌牌
type Table struct {
	*protoMsg.TableInfo
	GameHandle IGameOperate
	InningInfo *Inning
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

// Inning 每局信息
type Inning struct {
	Number   string //牌局号(游戏名+桌号+时间戳)
	OpenData []byte //开奖纪录(路单)
}

// IRoomHandle 房间接口
type IRoomHandle interface {
	Join(tid int64, person *Player) *Table        //加入
	Match(gid int64, person *Player) *Table       //配桌
	ChangeTable(gid int64, person *Player) *Table //换桌
	Leave(person *Player)                         // 离开房间
	//Check(rid int64) (*Room, bool)                       //查找房间
	//Open(rid int64, key string) (*Room, bool)            //开启房间
	//Close(rid int64, key string) bool                    //关闭房间
	//Clear(rid int64, key string) bool                    //清理房间
	//Del(rid int64, key string) bool                      //删除房间[房间ID和钥匙配对成功后,才能删除]

}

// RoomManger 管理房间
type RoomManger struct {
	sync.Map
}

// /////////////房间和子游戏管理////////////////////////
var lock sync.Mutex
var settingFixTime int64

// GetRoomManger 由于新增平台，则这里的管理不应该再是单例模式
func GetRoomManger() *RoomManger {
	return &RoomManger{
		sync.Map{},
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
	return &Room{
		RoomInfo: info,
		State:    0,
		tables:   make([]*Table, 0),
	}
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
func (self *Room) AddTable(table *protoMsg.TableInfo, f NewGameCallback) error {
	self.Lock()
	defer self.Unlock()
	if self.tables == nil {
		self.tables = make([]*Table, 0)
	}
	for _, t := range self.tables {
		if t.Id == table.Id {
			return fmt.Errorf(hints.StatusText[hints.Table01])
		}
	}
	if self.MaxTable < int32(len(self.tables)+1) {
		return fmt.Errorf(hints.StatusText[hints.Table02])
	}
	tb := &Table{
		TableInfo:  table,
		GameHandle: f(table.Gid),
		sitters:    make([]*Chair, 0),
		lookers:    make([]*Player, 0),
		RWMutex:    sync.RWMutex{},
	}
	self.tables = append(self.tables, tb)
	return nil
}

// Join 加入
func (self *Room) Join(tid int64, person *Player) *Table {
	// 玩家还在其他桌且该桌的游戏还没结束
	if 0 < person.InTableId {
		return nil
	}
	var tb *Table
	self.RLock()
	for _, table := range self.tables {
		if table.Id == tid {
			tb = table
			break
		}
	}
	self.RUnlock()
	if tb != nil {
		if err := tb.AddChair(person); err != nil {
			return nil
		}
		atomic.AddInt32(&self.onlineCount, 1)
	}
	return tb
}

func (self *Room) Match(gid int64, person *Player) *Table {
	var tb *Table
	self.RLock()
	tb = getRandomTable(gid, Unlimited, self.tables)
	self.Unlock()
	if tb != nil {
		if err := tb.AddChair(person); err != nil {
			log.Warnf("[room][Match]玩家:(%d) 配桌失败 rid:%d  gid:%d err:%v", person.UserID, self.Id, gid, err)
			return nil
		}
		atomic.AddInt32(&self.onlineCount, 1)
		log.Infof("[room][Match]玩家:(%d) 配桌成功 rid:%d  gid:%d ", person.UserID, self.Id, gid)
	} else {
		log.Warnf("[room][Match]玩家:(%d) 配桌失败 rid:%d  gid:%d ", person.UserID, self.Id, gid)
	}
	return tb
}

func (self *Room) ChangeTable(gid int64, person *Player) *Table {
	var tb *Table
	self.RLock()
	tb = getRandomTable(gid, person.InTableId, self.tables)
	self.RUnlock()
	if tb != nil {
		if err := tb.AddChair(person); err != nil {
			log.Warnf("[room][Match]玩家:(%d) 换桌失败 rid:%d tid:%d gid:%d  err:%v", person.UserID, tb.Rid, tb.Id, gid, err)
			return nil
		}
		log.Infof("[room][Match]玩家:(%d) 换桌成功 rid:%d tid:%d gid:%d  ", person.UserID, tb.Rid, tb.Id, gid)
	} else {
		log.Warnf("[room][Match]玩家:(%d) 换桌失败 rid:%d tid:%d gid:%d 找不到同游戏的其他桌", person.UserID, tb.Rid, tb.Id, gid)
	}
	return tb
}

// Leave 离开房间
func (self *Room) Leave(person *Player) {
	var tb *Table
	self.RLock()
	for _, table := range self.tables {
		if table.Id == person.InTableId {
			tb = table
			break
		}
	}
	self.RUnlock()
	if tb != nil {
		tb.RemoveChair(person)
	}
	atomic.AddInt32(&self.onlineCount, -1)
}

// OnlineCount 当前在线人数
func (self *Room) OnlineCount() int32 {
	return atomic.LoadInt32(&self.onlineCount)
}

// GetTable 获取桌牌
func (self *Room) GetTable(tid int64) *Table {
	self.RLock()
	defer self.RUnlock()
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
	self.RLock()
	defer self.RUnlock()
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

// DelTable 删除桌子
func (self *Room) DelTable(tid int64) {
	self.Lock()
	defer self.Unlock()
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
	self.Lock()
	defer self.Unlock()
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
	tb.GetInnings()
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

// GetInnings 获取牌局信息
func (tb *Table) GetInnings() *Inning {
	if tb.InningInfo == nil {
		info := &Inning{}
		info.OpenData = make([]byte, 0)
		tb.InningInfo = info
	}

	tb.InningInfo.Number = strings.ToUpper(util.Md5Sum(time.Now().String())) //GenerateGameNum(, self.Level, int32(self.TID))
	return tb.InningInfo
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
	lock.Lock()
	defer lock.Unlock()
	info := &Room{}
	return info, true
}
func (self *RoomManger) Close(rid int64, key string) bool { //关闭
	lock.Lock()
	defer lock.Unlock()

	return true
}
func (self *RoomManger) Clear(rid int64, key string) bool { //清理房间
	lock.Lock()
	defer lock.Unlock()

	return true
}
