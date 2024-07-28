package manger

import (
	log "github.com/po2656233/superplace/logger"
	"strings"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/utils"
	"sync"
	"time"
)

// IGameOperate 子游戏接口
type IGameOperate interface {
	Scene(args []interface{}) //场 景
	Ready(args []interface{})
	Start(args []interface{})             //开 始
	Playing(args []interface{})           //游 戏(下分|下注)
	Over(args []interface{})              //结 算
	UpdateInfo(args []interface{}) bool   //更新信息 如玩家进入或离开
	SuperControl(args []interface{}) bool //超级控制 在检测到没真实玩家时,且处于空闲状态时,自动关闭
}

// IAgainst 对战类
type IAgainst interface {
	Ready(args []interface{})     //准备
	DispatchCard()                //发牌
	CallScore(args []interface{}) //叫分
	OutCard(args []interface{})   //出牌
	Discard(args []interface{})   //操作
}

// IMultiPlayer 百人类
type IMultiPlayer interface {
	DispatchCard()                              //发牌
	DeduceWin() []byte                          //开奖区域
	BonusArea(area int32, betScore int64) int64 //区域赔额
}

type IDevice interface {
	Ready() bool              //准备
	Start(time int64) bool    //开启(游戏)
	Maintain(time int64) bool //维护(暂停)
	Clear(time int64) bool    //清场
	Close() bool              //关闭
}

// NewGameFunc 实例回调(根据桌子号创建游戏)
type NewGameFunc func(gid, tid int64) IGameOperate
type NewSingleGameFunc func(gid, tid int64) IGameOperate

type CalculateSQL func(info CalculateInfo) (nowMoney, factDeduct int64, isOK bool)

// ClearCallback 清场回调
type ClearCallback func() (gameID int64)

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
	IsStart    bool    // 第一次启动
	IsClear    bool    // 是否清场
	ReadyCount int32   // 已准备人数
	RunCount   int32   // 运行次数
	Inning     string  // 牌局号
	TimeStamp  int64   // 当前状态时刻的时间戳
	PlayIDList []int64 // 玩家列表
	lk         sync.RWMutex
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

///////////////////////////GAME//////////////////////////////////////////////////////

// Reset 重置信息
func (g *Game) Reset() bool {
	g.ChangeState(protoMsg.GameScene_Free)
	g.RunCount++
	g.ReadyCount = 0
	g.IsStart = false
	g.TimeStamp = time.Now().Unix()
	g.Inning = strings.ToUpper(utils.Md5Sum(g.Name + time.Now().String()))
	return true
}

// ChangeState 改变游戏状态
func (self *Game) ChangeState(state protoMsg.GameScene) {
	self.GameInfo.Scene = state
	self.TimeStamp = time.Now().Unix()
	log.Debugf("[%v:%v]   \t当前状态:%v ", self.Name, self.Id, protoMsg.GameScene_name[int32(self.GameInfo.Scene)])
}

// Ready 准备
func (g *Game) Ready(args []interface{}) {
	play := args[0].(*Player)
	if g.MaxPlayer != Unlimited && g.MaxPlayer < int32(len(g.PlayIDList)) {
		return
	}
	g.PlayIDList = append(g.PlayIDList, play.UserID)
	g.PlayIDList = utils.Unique(g.PlayIDList)
}

// Scene 游戏场景
func (g *Game) Scene(args []interface{}) {
	_ = args
}

// Start 开 始
func (g *Game) Start(args []interface{}) {
	_ = args
	if g.IsStart {
		return
	}
	g.IsStart = true
	g.ChangeState(protoMsg.GameScene_Start)
}

// Playing 游 戏(下分|下注)
func (g *Game) Playing(args []interface{}) {
	_ = args
}

// Over 结 算
func (g *Game) Over(args []interface{}) {
	_ = args
}

// UpdateInfo 更新信息 如玩家进入或离开
func (g *Game) UpdateInfo(args []interface{}) bool {
	if len(args) < 2 {
		return false
	}
	flag, ok := args[0].(protoMsg.PlayerState)
	uid, ok1 := args[1].(int64)
	if !ok || !ok1 {
		return false
	}

	switch flag {
	case protoMsg.PlayerState_PlayerSitDown: //新增玩家
		{
			g.lk.Lock()
			g.PlayIDList = append(g.PlayIDList, uid)
			g.lk.Unlock()
		}
	case protoMsg.PlayerState_PlayerStandUp: //删除玩家
		{
			scene := g.GameInfo.Scene
			// 空闲  结算 关闭 除了这三种场景，一般不允许玩家离开
			if scene != protoMsg.GameScene_Free && scene != protoMsg.GameScene_Over && scene != protoMsg.GameScene_Closing {
				return false
			}
			g.lk.Lock()
			g.PlayIDList = utils.RemoveValue(g.PlayIDList, uid)
			g.lk.Unlock()
		}
	default:
		return false
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

//////////////////////////////////////////////////////////////////////////
