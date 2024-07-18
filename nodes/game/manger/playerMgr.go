package manger

import (
	"github.com/po2656233/goleaf/gate"
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	. "superman/internal/hints"
	protoMsg "superman/internal/protocol/gofile"
	"sync"
)

// IUserBehavior 玩家行为
type IUserBehavior interface {
	Enter(args []interface{})     //入场
	Out(args []interface{})       //出场
	Offline(args []interface{})   //离线
	Reconnect(args []interface{}) //重入
	Ready(args []interface{})     //准备
	CallScore(args []interface{}) //叫分
	Host(args []interface{})      //抢庄/地主叫分
	SuperHost(args []interface{}) //超级抢庄
	Look(args []interface{})      //看牌
	Roll(args []interface{})      //摇骰子
	GiveUp(args []interface{})    //放弃
	Trustee(args []interface{})   //玩家托管
	Disbanded(args []interface{}) //解散游戏(房主权限)
}

// Player 玩家属性
type Player struct {
	//基本信息
	*protoMsg.PlayerInfo
	//用户包裹
	Knapsack *protoMsg.KnapsackInfo
	// 用于追踪用户行为
	GameHandle IGameOperate // 所在房间 时刻更新最新房间 状态
}

// PlayerManger 管理玩家
type PlayerManger struct {
	sync.Map // 即  map[uid]*Player
}

var manger *PlayerManger = nil
var once sync.Once

// GetPlayerManger 玩家管理对象(单例模式)//manger.persons = make(map[int64]*Player)
func GetPlayerManger() *PlayerManger {
	once.Do(func() {
		manger = &PlayerManger{
			sync.Map{},
		}
	})
	return manger
}

// Append 添加玩家
func (playerSelf *PlayerManger) Append(play *Player) (*Player, bool) {
	v, ok := playerSelf.Load(play.UserID)
	if !ok || v == nil {
		//log.Debug("[新增玩家]\tID:%v 账号:%v isRobot:%v", play.UserID, play.Account, play.Sex == 0x0F)
		playerSelf.Store(play.UserID, play)
		return play, true
	}
	//log.Debug("[已經存在玩家]\tID:%v 账号:%v isRobot:%v", play.UserID, play.Account, play.Sex == 0x0F)
	return v.(*Player), false
}

// Get 获取指定玩家[根据索引,即userID]
func (playerSelf *PlayerManger) Get(userID int64) *Player {
	value, ok := playerSelf.Load(userID)
	if ok {
		return value.(*Player)
	}
	return nil
}

// GetName 获取指定玩家昵称[根据索引,即userID]
func (playerSelf *PlayerManger) GetName(userID int64) string {
	value, ok := playerSelf.Load(userID)
	if ok {
		return value.(*Player).Name
	}
	return ""
}

// GetAccount 获取指定玩家账号[根据索引,即userID]
func (playerSelf *PlayerManger) GetAccount(userID int64) string {
	value, ok := playerSelf.Load(userID)
	if ok {
		return value.(*Player).Account
	}
	return ""
}

// /////////////////////////行为接口////////////////////////////////////////////

// Enter 进入场景(限制条件由外部转入)
func (itself *Player) Enter(args []interface{}) { //入场
	game := args[0].(*Game)
	agent := args[1].(gate.Agent)
	if itself == nil {
		GetClientManger().SendResult(agent, FAILED, StatusText[User03])
		return
	}

	//游戏句柄
	userID := itself.UserID
	gameHandle := itself.GameHandle
	if gameHandle == nil {
		log.Debug("[%v:%v] 玩家%v 进入出错:%v", game.Name, game.Id, userID, StatusText[Game18])
		GetClientManger().SendResult(agent, FAILED, StatusText[Game18])
		return
	}

	var sceneArgs []interface{}
	sceneArgs = append(sceneArgs, agent)
	gameHandle.Scene(sceneArgs) // 【进入-> 游戏场景】
}

func (itself *Player) Exit() { //退出
	itself.State = protoMsg.PlayerState_PlayerStandUp
	itself.InTableId = INVALID
	itself.InChairId = INVALID
	itself.InRooId = INVALID
	itself.GameHandle = nil //游戏正常结束退出时

}

func (itself *Player) UpdateState(flag protoMsg.PlayerState, args []interface{}) bool {
	if game := itself.GameHandle; game != nil { //[1-0
		var updateArgs []interface{}
		updateArgs = append(updateArgs, flag, itself.UserID, args)
		//log.Release("\t[%v:%v]玩家:%v-> 操作:%v\n", game.Info.Name, game.GameID, itself.UserID, RecvMessage[int(flag)])
		game.UpdateInfo(updateArgs)
		return true
	}

	log.Debug("[error:GameUpdate]\t ->:%v userID:%v", flag, itself.UserID)
	return false
}
