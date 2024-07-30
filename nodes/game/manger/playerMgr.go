package manger

import (
	log "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"sync"
)

//// IUserBehavior 玩家行为[废弃] 由玩家状态PlayerState取代
//type IUserBehavior interface {
//	Enter(args []interface{})     //入场
//	Out(args []interface{})       //出场
//	Offline(args []interface{})   //离线
//	Reconnect(args []interface{}) //重入
//	Ready(args []interface{})     //准备
//	CallScore(args []interface{}) //叫分
//	Host(args []interface{})      //抢庄/地主叫分
//	SuperHost(args []interface{}) //超级抢庄
//	Look(args []interface{})      //看牌
//	Roll(args []interface{})      //摇骰子
//	GiveUp(args []interface{})    //放弃
//	Trustee(args []interface{})   //玩家托管
//	Disbanded(args []interface{}) //解散游戏(房主权限)
//}

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

// GetPlayerMgr 玩家管理对象(单例模式)//manger.persons = make(map[int64]*Player)
func GetPlayerMgr() *PlayerManger {
	once.Do(func() {
		manger = &PlayerManger{
			sync.Map{},
		}
	})
	return manger
}

func ToPlayer(info *protoMsg.UserInfo) *Player {
	return &Player{
		PlayerInfo: &protoMsg.PlayerInfo{
			UserID:    info.UserID,
			Account:   info.Account,
			Name:      info.Name,
			FaceID:    info.FaceID,
			Age:       info.Age,
			Sex:       info.Gender,
			YuanBao:   info.YuanBao,
			Coin:      info.Coin,
			Level:     info.Level,
			Ranking:   0,
			State:     0,
			Gold:      0,
			Money:     info.Money,
			InRooId:   0,
			InTableId: 0,
			InChairId: 0,
		},
	}
}

func SimpleToPlayer(info *protoMsg.UserSimpleInfo) *Player {
	return &Player{
		PlayerInfo: &protoMsg.PlayerInfo{
			UserID:    info.UserID,
			Account:   info.Account,
			Name:      info.Name,
			FaceID:    info.FaceID,
			Age:       info.Age,
			Sex:       info.Gender,
			YuanBao:   info.YuanBao,
			Coin:      info.Coin,
			Level:     info.Level,
			Ranking:   0,
			State:     0,
			Gold:      0,
			Money:     info.Money,
			InRooId:   0,
			InTableId: 0,
			InChairId: 0,
		},
	}
}

// Append 添加玩家
func (playerSelf *PlayerManger) Append(play *Player) (*Player, bool) {
	v, ok := playerSelf.Load(play.UserID)
	if !ok || v == nil {
		log.Debugf("[新增玩家]\tID:%v 账号:%v isRobot:%v", play.UserID, play.Account, play.Sex == 0x0F)
		playerSelf.Store(play.UserID, play)
		return play, true
	}
	log.Debugf("[已經存在玩家]\tID:%v 账号:%v isRobot:%v", play.UserID, play.Account, play.Sex == 0x0F)
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
	agent := args[1].(Agent)
	if itself == nil {
		GetClientMgr().SendResult(agent, FAILED, StatusText[User03])
		return
	}

	//游戏句柄
	if itself.GameHandle == nil {
		log.Debugf("[%v:%v] 玩家%v 进入出错:%v", itself.InRooId, itself.InTableId, itself.UserID, StatusText[Game18])
		GetClientMgr().SendResult(agent, FAILED, StatusText[Game18])
		return
	}

	itself.GameHandle.Scene([]interface{}{agent}) // 【进入-> 游戏场景】
}

func (itself *Player) Exit() bool { //退出
	if itself.GameHandle == nil || false == itself.GameHandle.UpdateInfo([]interface{}{protoMsg.PlayerState_PlayerStandUp, itself.UserID}) {
		return false
	}
	itself.State = protoMsg.PlayerState_PlayerStandUp
	if rm := GetRoomMgr().GetRoom(itself.InRooId); rm != nil {
		if t := rm.GetTable(itself.InTableId); t != nil {
			t.RemoveChair(itself.UserID)
		}
	}
	itself.InRooId = INVALID
	itself.InTableId = INVALID
	itself.InChairId = INVALID
	itself.GameHandle = nil
	return true
}

func (itself *Player) UpdateState(flag protoMsg.PlayerState, args []interface{}) bool {
	if game := itself.GameHandle; game != nil { //[1-0
		log.Infof("玩家[%v]:%v-> 操作:%+v", itself.UserID, itself.Name, flag)
		tempArgs := make([]interface{}, 0)
		tempArgs = append(tempArgs, flag, itself.UserID)
		tempArgs = append(tempArgs, args...)
		return game.UpdateInfo(tempArgs)
	}

	log.Debugf("[error:GameUpdate]\t ->:%v userID:%v", flag, itself.UserID)
	return false
}
