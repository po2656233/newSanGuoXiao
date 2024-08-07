package manger

import (
	log "github.com/po2656233/superplace/logger"
	"google.golang.org/protobuf/proto"
	"strings"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	"sync"
	"time"
)

// IGameOperate 子游戏接口 在检测到没真实玩家时,回归空闲状态
type IGameOperate interface {
	Scene(args []interface{}) bool        // 场 景
	Start(args []interface{}) bool        // 开始
	Ready(args []interface{}) bool        // 玩家准备
	Playing(args []interface{}) bool      // 游 戏(下分|下注)
	UpdateInfo(args []interface{}) bool   // 更新信息 如玩家进入、准备或离开等操作
	SuperControl(args []interface{}) bool // 超级控制
}

// IAgainst 对战类
type IAgainst interface {
	Ready(args []interface{}) bool     //准备
	DispatchCard() bool                //发牌
	CallScore(args []interface{}) bool //叫分
	OutCard(args []interface{}) bool   //出牌
	Discard(args []interface{}) bool   //操作
}

// IMultiPlayer 百人类
type IMultiPlayer interface {
	DispatchCard()                              //发牌
	DeduceWin() []byte                          //开奖区域
	BonusArea(area int32, betScore int64) int64 //区域赔额
}

type IDevice interface {
	Start(time int64) bool    // 开启(游戏)
	Maintain(time int64) bool // 维护(暂停)
	Over() bool               // 结束
	Close(time int64) bool    // 关闭
}

// NewGameFunc 实例回调(根据桌子号创建游戏)
type NewGameFunc func(gid int64, table *Table) IGameOperate
type NewSingleGameFunc func(gid, tid int64) IGameOperate
type CalculateFunc func(info CalculateInfo) (nowMoney, factDeduct int64, isOK bool)

// ClearFunc 清场回调
type ClearFunc func() *Table

type WorkTask struct {
	sceneInfo proto.Message
	before    func()      // 事件之前需要做的
	job       func() bool // 当前场景下需要执行的事情,返回false,则不执行after
	after     func()      // 事件之后需要处理的事情
	timeout   int32
}

// CalculateInfo 结算信息
type CalculateInfo struct {
	UserID   int64  // 结算对象(接受充值的ID)
	ByUID    int64  // 该用户发起
	PreMoney int64  // 结算前的钱
	Payment  int64  // 付款金额
	Code     int32  // 操作码(即为什么付款)
	Order    string // 订单号(或牌局号)
	Remark   string // 备注
}
type Game struct {
	*protoMsg.GameInfo
	IsStart    bool   // 第一次启动
	IsClear    bool   // 是否清场
	ReadyCount int32  // 已准备人数
	RunCount   int32  // 运行次数
	Inning     string // 牌局号
	TimeStamp  int64  // 当前状态时刻的时间戳
	PlayerList []int64
	Timer      *time.Timer
	openTime   int64 // 开局时间戳
	Event      map[protoMsg.GameScene]WorkTask
	playlistLK sync.RWMutex
}

type GameMgr struct {
	sync.Map
}

var gOnce sync.Once

// ///////////////////游戏功能//////////////////////////////////////////////
var gameMgr *GameMgr

func GetGameInfoMgr() *GameMgr {
	gOnce.Do(func() {
		gameMgr = &GameMgr{}
	})
	return gameMgr
}

func (gmr *GameMgr) AddGames(games []*protoMsg.GameInfo) {
	for _, game := range games {
		gmr.Store(game.Id, game)
	}
}

func (gmr *GameMgr) AddGame(game *protoMsg.GameInfo) {
	gmr.Store(game.Id, game)
}

func (gmr *GameMgr) GetGame(gid int64) *protoMsg.GameInfo {
	val, ok := gmr.Load(gid)
	if !ok {
		return nil
	}
	return val.(*protoMsg.GameInfo)
}

///////////////////////////GAME///////////////////////////////////////////////////////////
//////////////////////////基础功能//////////////////////////////////////////////////////////

// Init 重置信息
func (g *Game) Init() {
	g.RunCount++
	g.ReadyCount = 0
	g.IsStart = false
	if g.Timer != nil {
		g.Timer.Stop()
	}
	g.TimeStamp = time.Now().Unix()
	g.Inning = strings.ToUpper(utils.Md5Sum(g.Name + time.Now().String()))
	return
}

func (g *Game) SetPlaylist(uids []int64) {
	g.playlistLK.Lock()
	defer g.playlistLK.Unlock()
	g.PlayerList = uids
}

// AddPlayer [注]一般不用，因为都是等玩家坐满之后，直接SetPlaylist
func (g *Game) AddPlayer(uid int64) {
	g.playlistLK.Lock()
	defer g.playlistLK.Unlock()
	g.PlayerList = append(g.PlayerList, uid)
	g.PlayerList = utils.Unique(g.PlayerList)
}

func (g *Game) RemovePlayer(uid int64) {
	g.playlistLK.Lock()
	defer g.playlistLK.Unlock()
	g.PlayerList = utils.RemoveValue(g.PlayerList, uid)
}
func (g *Game) ClearPlayer() {
	g.playlistLK.Lock()
	defer g.playlistLK.Unlock()
	g.PlayerList = make([]int64, 0)
}

func (g *Game) GetPlayerCount() int {
	g.playlistLK.RLock()
	defer g.playlistLK.RUnlock()
	return len(g.PlayerList)
}

func (g *Game) PlayerWork(f func(int64)) {
	g.playlistLK.RLock()
	defer g.playlistLK.RUnlock()
	for _, uid := range g.PlayerList {
		f(uid)
	}
}

// ChangeState 改变游戏状态
func (self *Game) ChangeState(state protoMsg.GameScene) {
	self.GameInfo.Scene = state
	self.TimeStamp = time.Now().Unix()
	log.Infof("[%v:%v]   \t当前场景:%v ", self.Name, self.Id, protoMsg.GameScene_name[int32(state)])
}

// ChangeStateAndWork 改变游戏状态
func (self *Game) ChangeStateAndWork(state protoMsg.GameScene) {
	nowScene, ok := self.Event[state]
	if !ok {
		log.Warnf("[%v:%v]   \t 场景:%v 未注册!!! 可能导致游戏无法运行. ", self.Name, self.Id, protoMsg.GameScene_name[int32(state)])
		return
	}

	if nowScene.before != nil {
		nowScene.before()
	}
	if self.GameInfo.Scene != state {
		if self.GameInfo.MaxPlayer != Unlimited {
			log.Infof("[%v:%v]   \t当前场景:%v 当前玩家列表%v", self.Name, self.Id, protoMsg.GameScene_name[int32(state)], self.PlayerList)
		} else {
			log.Infof("[%v:%v]   \t当前场景:%v ", self.Name, self.Id, protoMsg.GameScene_name[int32(state)])
		}
	}
	// 变更状态
	self.GameInfo.Scene = state
	self.TimeStamp = time.Now().Unix()
	canAfter := true
	if nowScene.job != nil {
		// 变更时,执行的任务
		canAfter = nowScene.job()
	}
	// 是否执行后续任务
	if !canAfter {
		return
	}
	// 继续执行任务
	if nowScene.timeout > 0 {
		if self.Timer != nil {
			self.Timer.Stop()
		}
		self.Timer = time.AfterFunc(time.Duration(nowScene.timeout)*time.Second, func() {
			if self.IsClear { // 如果清场了,则不再执行后续任务
				return
			}
			if nowScene.after != nil {
				nowScene.after()
			}
			self.ToNextScene()
		})
	}
	if nowScene.sceneInfo != nil && 0 < self.GetPlayerCount() {
		GetClientMgr().NotifyOthers(self.PlayerList, nowScene.sceneInfo)
	}

}

///////////////////////////////游戏操作//////////////////////////////////////////////

// Scene 游戏场景
func (g *Game) Scene(args []interface{}) bool {
	_ = args
	person, ok := args[0].(*Player)
	if !ok {
		return false
	}
	g.AddPlayer(person.UserID)
	now := time.Now().Unix()
	wait := g.openTime - now
	if 0 < wait {
		GetClientMgr().SendTo(person.UserID, &protoMsg.WaitGameStartResp{
			GameID:      g.Id,
			WaitTimeOut: wait,
		})
		return false
	}
	return true
}

// Ready 准备
func (g *Game) Ready(args []interface{}) bool {
	_ = args
	return true
}

// Start 准备
func (g *Game) Start(args []interface{}) bool {
	_ = args
	if g.IsStart {
		return false
	}
	g.IsStart = true
	return true
}

// Playing 游 戏(下分|下注)
func (g *Game) Playing(args []interface{}) bool {
	_ = args
	agent := args[1].(Agent)
	if g.GameInfo.Scene != protoMsg.GameScene_Playing {
		GetClientMgr().SendResult(agent, FAILED, StatusText[Game04])
		return false
	}
	userData := agent.UserData()
	if userData == nil { //[0
		GetClientMgr().SendResult(agent, FAILED, StatusText[Game37])
		return false
	}
	return true
}

// UpdateInfo 更新信息 如玩家进入或离开
func (g *Game) UpdateInfo(args []interface{}) bool {
	flag, ok := args[0].(protoMsg.PlayerState)
	if !ok {
		return false
	}
	uid, ok1 := args[1].(int64)
	if !ok1 {
		return false
	}
	switch flag {
	case protoMsg.PlayerState_PlayerSitDown:
		g.AddPlayer(uid)
	case protoMsg.PlayerState_PlayerStandUp:
		if protoMsg.GameScene_Start <= g.GameInfo.Scene && g.GameInfo.Scene <= protoMsg.GameScene_Over {
			return false
		}
	case protoMsg.PlayerState_PlayerGiveUp:
		g.RemovePlayer(uid)

	}
	return true
}

// SuperControl 超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭
func (g *Game) SuperControl(args []interface{}) bool {
	if len(args) == 0 {
		return false
	}
	flag, ok := args[0].(int32)
	if !ok {
		return false
	}
	switch protoMsg.GameScene(flag) {
	case protoMsg.GameScene_Closing:
		{
			// 移除所有玩家
		}
	default:
		return false
	}
	return true
}

///////////////////////////////////////////////////////////////////////////////////////
// 									注册事件
//////////////////////////////////////////////////////////////////////////////////////

func (g *Game) RegisterEvent(scene protoMsg.GameScene, info proto.Message, timeout int32, work func() bool, before, after func()) {
	if g.Event == nil {
		g.Event = make(map[protoMsg.GameScene]WorkTask)
	}
	task := WorkTask{}
	task.timeout = timeout
	task.sceneInfo = info
	task.before = before
	task.job = work
	task.after = after
	g.Event[scene] = task
}
func (g *Game) Running() {
	for i := protoMsg.GameScene_Free; i < protoMsg.GameScene_Closing; i++ {
		if _, ok := g.Event[i]; ok {
			g.ChangeStateAndWork(i)
			return
		}
	}
}
func (g *Game) ToNextScene() bool {
	if g.IsClear { // 如果清场了,则不再执行后续任务
		return false
	}
	next := int32(g.GameInfo.Scene) + 1
	if next > int32(protoMsg.GameScene_Closing) {
		next = int32(protoMsg.GameScene_Free)
	}
	max := int32(protoMsg.GameScene_Closing) + next
	for i := next; i < max; i++ {
		next = i
		if int32(protoMsg.GameScene_Closing) < next+1 {
			next = i - int32(protoMsg.GameScene_Closing)
		}
		if _, ok := g.Event[protoMsg.GameScene(next)]; ok {
			break
		}
	}
	if int32(protoMsg.GameScene_Closing) <= next {
		return false
	}
	// 如果当前没有玩家,游戏回归空闲状态
	if int32(protoMsg.GameScene_Start) == next && g.GetPlayerCount() == INVALID {
		next = int32(protoMsg.GameScene_Free)
	}
	g.ChangeStateAndWork(protoMsg.GameScene(next))
	return true
}

////////////////////////////////////////////////////////////////////////////////////

// Close 关闭游戏
func (g *Game) Close(f ClearFunc) {
	g.ChangeState(protoMsg.GameScene_Closing)
	//
	tb := f()
	if tb != nil {
		GetClientMgr().NotifyOthers(g.PlayerList, &protoMsg.DisbandedTableResp{
			RoomID:  tb.Rid,
			TableID: tb.Id,
		})
		if rm := GetRoomMgr().GetRoom(tb.Rid); rm != nil {
			rm.DelTable(tb.Id)
		}
		tb.Close()
		*tb = Table{}
		tb = nil
	}

}

//////////////////////////////////////////////////////////////////////////
