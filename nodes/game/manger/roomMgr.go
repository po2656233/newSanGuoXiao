package manger

import (
	"fmt"
	log "github.com/po2656233/superplace/logger"
	"superman/internal/hints"
	protoMsg "superman/internal/protocol/gofile"
	"sync"
	"time"
)

type No int32
type GameID int64

// ----------房间----------------------

// Room 房间信息
type Room struct {
	*protoMsg.RoomInfo
	State     uint8             //房间状态 [0:无效] [1:Open] [2:Close] [3:Other] [4:Clear]
	Tables    map[GameID]*Table //桌子集合
	MaxPeople int32             //最大承载人数

	nowPeople int32 // 当前已允许最大人数
	sync.RWMutex
}

// Table 桌牌
type Table struct {
	*protoMsg.TableInfo
	GameHandle *IGameOperate
	Sitters    []*Player // 座位上的玩家
	Lookers    []*Player // 旁观者
	sync.RWMutex
}

// Chair 椅子
type Chair struct {
	*protoMsg.PlayerInfo
	ID       int32       // 椅子号(仅表示序号) 0:空
	Code     int8        // 牌值状态码(用作是否看牌|听牌|碰牌|杠牌 )
	HandCard []byte      // 手牌
	Score    int64       // 当前一局的消费额(斗地主叫分阶段为叫分分值)
	Total    int64       // 消费总额
	Gain     int64       // 获利
	Gold     int64       // 金币
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
	Create(info *protoMsg.RoomInfo) (*Room, bool)        //创建房间[房间ID和钥匙配对]
	Join(tid int64, person *Player) (*Table, int)        //加入
	Match(gid int64, person *Player) (*Table, int)       //配桌
	ChangeTable(gid int64, person *Player) (*Table, int) //换桌
	Check(rid int64) (*Room, bool)                       //查找房间
	Open(rid int64, key string) (*Room, bool)            //开启房间
	Close(rid int64, key string) bool                    //关闭房间
	Clear(rid int64, key string) bool                    //清理房间
	Del(rid int64, key string) bool                      //删除房间[房间ID和钥匙配对成功后,才能删除]

}

// RoomManger 管理房间
type RoomManger struct {
	sync.Map
}

// /////////////房间和子游戏管理////////////////////////
var lock sync.Mutex

// GetRoomManger 由于新增平台，则这里的管理不应该再是单例模式
func GetRoomManger() *RoomManger {
	return &RoomManger{
		sync.Map{},
	}
}

// AddTable 新增桌子(tid == gid 即桌子ID和游戏ID共用)
func (self *Room) AddTable(table *protoMsg.TableInfo) error {
	self.Lock()
	defer self.Unlock()
	if self.Tables == nil {
		self.Tables = make(map[GameID]*Table)
	}
	gid := GameID(table.Gid)
	if _, ok := self.Tables[gid]; ok {
		return fmt.Errorf(hints.StatusText[hints.Table01])
	}

	if self.MaxPeople < self.nowPeople+table.SitterCount {
		return fmt.Errorf(hints.StatusText[hints.Table02])
	}
	self.nowPeople += table.SitterCount
	self.Tables[gid] = &Table{
		TableInfo: table,
	}
	return nil
}

// Join 加入
func (self *Room) Join(tid int64, person *Player) *Table {
	self.Lock()
	self.Unlock()
	// 玩家还在其他桌且该桌的游戏还没结束
	if 0 < person.InTableId {
		return nil
	}
	for _, table := range self.Tables {
		if table.Id == tid {
			table.Sitters = append(table.Sitters, person)
			return table
		}
	}
	return nil
}

func (self *Room) Match(gid int64, person *Player) *Table {
	self.Lock()
	self.Unlock()
	for _, table := range self.Tables {
		if table.Gid == gid && table.SitterCount+1 == len(table.Sitters) {
			table.Sitters = append(table.Sitters, person)
			return table
		}
	}
	return nil
}

func (self *Room) ChangeTable(gid int64, person *Player) (*Table, int) {

}

// GetTable 获取桌牌
func (self *Room) GetTable(tid int64) *Table {
	self.RLock()
	defer self.RUnlock()
	if self.Tables == nil {
		return nil
	}
	for _, table := range self.Tables {
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
	if self.Tables == nil {
		return nil
	}
	for _, table := range self.Tables {
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
	if self.Tables == nil {
		return
	}
	for i, item := range self.Tables {
		if item.Id == tid {
			self.MaxPeople -= item.SitterCount
			if self.MaxPeople < 0 {
				self.MaxPeople = 0
			}
			delete(self.Tables, i)
			return
		}
	}
}

// DelTableForGid  移除指定游戏
func (self *Room) DelTableForGid(gid int64) {
	self.Lock()
	defer self.Unlock()
	if self.Tables == nil {
		return
	}
	for i, item := range self.Tables {
		if item.Gid == gid {
			self.MaxPeople -= item.SitterCount
			if self.MaxPeople < 0 {
				self.MaxPeople = 0
			}
			delete(self.Tables, i)
		}
	}
}

/////////////////////桌牌功能//////////////////////////////////////////////

func (tb *Table) SitDown(player *Player) error {

}

/////////////////////游戏功能//////////////////////////////////////////////

// Ready 准备
func (g *Game) Ready() bool {

	return true
}

// Start 开始
func (tb *Game) Start(time int64) bool {

	return true
}

// Playing 运作
func (tb *Game) Playing() bool {
	return true
}

// Maintain 维护
func (tb *Game) Maintain(time int64) bool {
	return true
}

// Clear 清桌
func (tb *Game) Clear(time int64) bool {
	return true
}

// Close 关闭
func (tb *Game) Close() bool {
	return true
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
