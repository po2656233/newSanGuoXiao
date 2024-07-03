package internal

import (
	"fmt"
	"reflect"
	protoMsg "superman/internal/protocol/gofile"
	token2 "superman/internal/token"
	. "superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/game/internal/category/CardGame/sanguoxiao"
	"superman/nodes/leaf/jettengame/game/internal/category/MahjongGame/mahjong"
	"superman/nodes/leaf/jettengame/game/internal/category/SlotGame/zhaocaimiao"
	. "superman/nodes/leaf/jettengame/manger"
	"superman/nodes/leaf/jettengame/sql/redis"
	"time"

	"github.com/po2656233/goleaf/log"

	"superman/nodes/leaf/jettengame/sql/mysql"
)

var redisClient = redis.RedisHandle()

// 初始化
func init() {
	handlerBehavior()

}

// 注册传输消息
func handlerMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// 玩家行为处理
func handlerBehavior() {
	//系统
	handlerMsg(&protoMsg.SuggestReq{}, handleSuggest)           //意见反馈
	handlerMsg(&protoMsg.NotifyNoticeReq{}, handleNotifyNotice) //意见反馈
	//玩家行为
	handlerMsg(&protoMsg.EnterGameReq{}, enter)
	handlerMsg(&protoMsg.ExitGameReq{}, exit)
	handlerMsg(&protoMsg.DisbandedGameReq{}, disbandedGame)

	handlerMsg(&protoMsg.TrusteeReq{}, trustee)

	handlerMsg(&protoMsg.ChangeTableReq{}, changeTable)
	handlerMsg(&protoMsg.GetInningsInfoReq{}, getInnings)
	handlerMsg(&protoMsg.GameRecord{}, getGameRecords)
	handlerMsg(&protoMsg.GetRecordReq{}, getRecords)
	handlerMsg(&protoMsg.GetBackPasswordReq{}, getBackPassword)

	handlerMsg(&protoMsg.UpdateGoldReq{}, updateGold)

	//slotGame
	// zhaocaimiao
	handlerMsg(&protoMsg.ZhaocaimiaoBetReq{}, playing)
}

// 校验合法性
func verifyMsg(args []interface{}) (*Player, IGameOperate, bool) {
	_ = args[1]
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if game := person.PtrTable; nil != game && nil != game.Instance {
			return person, game.Instance, true
		}
	}
	GetClientManger().SendError(agent)
	return nil, nil, false
}

// --------------------------------------------------------------------------------
// 创建游戏对象
func productGame(game *Game) IGameOperate {
	log.Debug("[[创建游戏实例]]\t[%v:%v] KIND:%v Type:%v Level:%v HostID:%v ", game.G.Name, game.ID, game.G.KindID, game.G.Type, game.G.Level, game.T.HostID)
	switch game.G.KindID {
	case Zhaocaimiao:
		return zhaocaimiao.NewGame(game)
	case Mahjong:
		return mahjong.NewMahjong(game)
	case Sanguoxiao:
		return superman.NewGame(game)
	}

	return nil
}

//---------------------------------------------------------------------------------
//

// 意见反馈
func handleSuggest(args []interface{}) {
	_ = args[0].(*protoMsg.SuggestReq)
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		msg := &protoMsg.SuggestResp{
			UserID: person.UserID,
			Feedback: &protoMsg.EmailInfo{
				EmailID:    GetPChatID(),
				TimeStamp:  time.Now().Unix(),
				AcceptName: person.Name,
				Topic:      "感谢信",
				Content:    "由衷感谢您的反馈,您的反馈是我们持续的动力!",
				Sender:     "系统邮件",
			},
		}
		GetClientManger().SendData(agent, msg)
		GetClientManger().SendPopResult(agent, SUCCESS, StatusText[Title007], StatusText[User22])
	}
}
func handleNotifyNotice(args []interface{}) {
	m := args[0].(*protoMsg.NotifyNoticeReq)
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if m.Timeout <= INVALID {
			m.Timeout = Minute
		}
		msg := &protoMsg.NotifyNoticeResp{
			UserID: person.UserID,
			GameID: m.GameID,
			Level:  m.Level,
			TimeInfo: &protoMsg.TimeInfo{
				OutTime:   INVALID,
				WaitTime:  INVALID,
				TotalTime: m.Timeout,
				TimeStamp: time.Now().Unix(),
			},
			Title:   m.Title,
			Content: m.Content,
		}

		//写到缓存
		redisClient.Set(KeyNormalNotify, msg, time.Duration(m.Timeout)*time.Second)

		//if person.UserID == SYSTEMID {
		//	msg.TimeInfo.TotalTime = Date
		//}
		GetClientManger().NotifyAll(msg)
		//GetClientManger().SendPopResult(agent, SUCCESS, StatusText[Title001], StatusText[User22])
	}
}

///////////////////////////////////////////////////////////////////////////////////////////

// enter 进入
func enter(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0].(*protoMsg.EnterGameReq)
	agent := args[1].(*ActorPlayer)
	//var agent gate.Agent

	// 001 获取并解析token信息
	tk, st := token2.ValidateBase64(m.Token)
	if st != SUCCESS || tk == nil {
		agent.SendResult(FAILED, StatusText[Room05])
		return
	}

	uid := agent.Session.GetUid()
	if _, ok := GetClientManger().Get(uid); !ok {
		GetClientManger().Append(uid, agent)
	}

	// 002 构建玩家实例 并添加保存
	person := GetPlayerManger().Get(uid)
	if person == nil {
		person = &Player{
			PlayerInfo: &protoMsg.PlayerInfo{},
			Knapsack:   &protoMsg.KnapsackInfo{},
			PtrRoom:    &Room{},
			PtrTable:   &Table{},
		}
	}

	if GetPlayerManger().Exist(uid) {
		person = GetPlayerManger().Get(uid)
	} else if info := mysql.SqlHandle().CheckUserInfo(uid); info != nil {
		// 数据转换
		person.Sex = info.Gender
		person.Money = info.Money
		person.Name = info.Name
		person.Account = info.Account
		person.FaceID = info.FaceID
		person.Level = info.Level
		person.Age = info.Age
	}

	// 反推房间号
	if person.RoomNum == 0 {
		num, err := mysql.SqlHandle().CheckRoomNum(m.GameID)
		if err != nil {
			agent.SendResult(FAILED, StatusText[Game47])
			return
		}
		person.RoomNum = num
	}

	person.PtrRoom = &Room{
		Num: person.RoomNum,
	}
	//agent.SetUserData(person)

	GetPlayerManger().Append(person)

	// 玩家是否存在没有结束的游戏
	if person.PtrTable != nil && person.GameID != m.GameID && protoMsg.PlayerState_PlayerAgree < person.State {
		hints := fmt.Sprintf("您所参与的游戏(ID:%v)本轮还没结束!", person.GameID)
		GetClientManger().SendResult(agent, FAILED, hints)
		//agent.SendResult(FAILED, hints)
		return
	}

	//003 查找平台是否包含该房间
	pid := mysql.SqlHandle().CheckPlatformInfo(uid)
	platform := GetPlatformManger().Get(pid)
	if platform == nil {
		agent.SendResult(FAILED, StatusText[Room12])
		//GetClientManger().SendResult(agent, FAILED, StatusText[Room12])
		return
	}
	var isHaveRoom = false
	for _, v := range platform.ClassIDs {
		if v == person.RoomNum {
			isHaveRoom = true
			break
		}
	}
	if !isHaveRoom {
		agent.SendResult(FAILED, StatusText[Room13])
		//GetClientManger().SendResult(agent, FAILED, StatusText[Room13])
		return
	}

	//004 获取游戏
	game, ok := GetGamesManger().GetGame(m.GameID)
	if !ok {
		agent.SendResult(FAILED, StatusText[TableInfo04])
		//GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo04])
		return
	}

	if game.T.HostID != SYSTEMID && game.T.Password != m.Password {
		agent.SendResult(FAILED, StatusText[Room01])
		//GetClientManger().SendResult(agent, FAILED, StatusText[Room01])
		return
	}

	//005 进行配桌
	_, code := GetGamesManger().Match(game, person, productGame)
	if code != SUCCESS {
		agent.SendResult(FAILED, StatusText[code])
		//GetClientManger().SendResult(agent, FAILED, StatusText[code])
		return
	}

	//006 进入游戏场景
	log.Debug("[%v:%v]\t[配桌] 玩家:%v 成功配桌!!", game.G.Name, game.ID, uid)
	var enterArgs []interface{}
	enterArgs = append(enterArgs, game, agent)
	person.Enter(enterArgs)
	log.Debug("[%v:%v]\t[场景:%v] UID:%v", game.G.Name, game.ID, protoMsg.GameScene_name[int32(game.G.Scene)], person.UserID)

}

func changeTable(args []interface{}) {
	//查找玩家
	_ = args[1]
	m := args[0].(*protoMsg.ChangeTableReq)
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if table, code := GetGamesManger().ChangeTable(m.GameID, person, productGame); code == SUCCESS {
			msg := &protoMsg.ChangeTableResp{
				GameID:   table.GameID,
				TableNum: table.Num,
				UserID:   person.UserID,
			}
			GetClientManger().SendData(agent, msg)
		} else {
			GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo05])
		}
	}

}

// 获取牌局信息
func getInnings(args []interface{}) {
	_ = args[1]
	agent := args[1].(*ActorPlayer)
	m := args[0].(*protoMsg.GetInningsInfoReq)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)

		//判断玩家是否是房主
		if false == GetGamesManger().CheckHostID(m.GameID, person.UserID) {
			GetClientManger().SendResult(agent, FAILED, StatusText[Room11])
			return
		}

		innings := mysql.SqlHandle().CheckInnings(person.UserID, m.GameID)
		msg := &protoMsg.GetInningsInfoResp{
			PageNum: INVALID,
		}
		if MaxInning < len(innings) {
			inningXs := make([]*protoMsg.InningInfo, 0)
			for i := 0; i < len(innings); i++ {
				if 0 == (i+1)%MaxInning {
					if 0 < len(inningXs) {
						msg.Innings = inningXs
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					inningXs = make([]*protoMsg.InningInfo, 0)
				}
				inningXs = CopyInsert(inningXs, len(inningXs), innings[i]).([]*protoMsg.InningInfo)
			}
			if 0 < len(inningXs) {
				msg.Innings = inningXs
				GetClientManger().SendData(agent, msg)
			}
		} else {
			msg.Innings = innings
			GetClientManger().SendData(agent, msg)
		}
		return
	}
	GetClientManger().SendError(agent)
}

// 获取牌局记录
func getGameRecords(args []interface{}) {

}

// 获取牌局记录
func getRecords(args []interface{}) {
	_ = args[1]
	agent := args[1].(*ActorPlayer)
	m := args[0].(*protoMsg.GetRecordReq)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		innings := mysql.SqlHandle().CheckRecords(person.UserID, m.KindID, m.Level, m.StartTimeStamp, m.EndTimeStamp)
		msg := &protoMsg.GetRecordResp{
			PageNum: INVALID,
			//KindID:         m.KindID,
			//Level:          m.Level,
			//StartTimeStamp: m.StartTimeStamp,
			//EndTimeStamp:   m.EndTimeStamp,
		}

		if MaxInning < len(innings) {
			inningXs := make([]*protoMsg.InningInfo, 0)
			for i := 0; i < len(innings); i++ {
				if 0 == (i+1)%MaxInning {
					if 0 < len(inningXs) {
						msg.Innings = inningXs
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					inningXs = make([]*protoMsg.InningInfo, 0)
				}
				inningXs = CopyInsert(inningXs, len(inningXs), innings[i]).([]*protoMsg.InningInfo)
			}
			if 0 < len(inningXs) {
				msg.Innings = inningXs
				GetClientManger().SendData(agent, msg)
			}
		} else {
			msg.Innings = innings
			GetClientManger().SendData(agent, msg)
		}
		return
	}
}

// 取回游戏密码
func getBackPassword(args []interface{}) {
	agent := args[1].(*ActorPlayer)
	m := args[0].(*protoMsg.GetBackPasswordReq)
	if userData := agent.UserData(); userData != nil { //[0
		if g, ok := GetGamesManger().GetGame(m.GameID); ok {
			person := userData.(*Player)
			if g.T.HostID == person.UserID {
				msg := &protoMsg.GetBackPasswordResp{
					GameID:   m.GameID,
					Password: g.T.Password,
				}
				GetClientManger().SendData(agent, msg)
				return
			}
		}
	}
	GetClientManger().SendResult(agent, FAILED, StatusText[User02])
}

// / 更新金币
func updateGold(args []interface{}) {
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		person.Gold = mysql.SqlHandle().CheckMoney(person.UserID)
		msg := &protoMsg.UpdateGoldResp{
			UserID: person.UserID,
			Gold:   person.Gold,
		}
		GetClientManger().SendData(agent, msg)
		return
	}
	GetClientManger().SendError(agent)
}

// 托管
func trustee(args []interface{}) {
	if person, _, ok := verifyMsg(args); ok {
		person.State = protoMsg.PlayerState_PlayerPickUp
		person.Trustee(args)
	}
}

// 游戏 (下注)
func playing(args []interface{}) {
	if person, gameHandle, ok := verifyMsg(args); ok {
		person.State = protoMsg.PlayerState_PlayerPlaying
		gameHandle.Playing(args)
	}
}

// 游戏
func over(args []interface{}) {
	if person, gameHandle, ok := verifyMsg(args); ok {
		person.State = protoMsg.PlayerState_PlayerStandUp
		gameHandle.Over(args)
	}
}

// 退出游戏
func exit(args []interface{}) {
	_ = args[1]
	m := args[0].(*protoMsg.ExitGameReq)
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if game := person.PtrTable; nil != game && m.GameID != INVALID && nil != game.Instance {
			person.Out(args)
		} else {
			//退出
			agent.WriteMsg(&protoMsg.ExitGameResp{
				GameID: m.GameID,
				UserID: person.UserID,
			})
		}

		return
	}
	GetClientManger().SendError(agent)
}

// 解散游戏
func disbandedGame(args []interface{}) {
	_ = args[1]
	m := args[0].(*protoMsg.DisbandedGameReq)
	agent := args[1].(*ActorPlayer)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if game := person.PtrTable; nil != game && m.GameID != INVALID && nil != game.Instance && game.Info.HostID == person.UserID {
			person.Disbanded(args)
			//退出
			agent.WriteMsg(&protoMsg.DisbandedGameResp{
				GameID: m.GameID,
				UserID: person.UserID,
			})
		} else {
			agent.WriteMsg(&protoMsg.ResultResp{
				State: FAILED,
				Hints: StatusText[User20],
			})
		}

		return
	}
	GetClientManger().SendError(agent)
}

// /////////////////////////////////////////////////////////////////////////////////
// 掷骰子
func roll(args []interface{}) {
	_ = args[1]
	agent := args[1].(*ActorPlayer)
	if person, _, ok := verifyMsg(args); ok {
		person.State = protoMsg.PlayerState_PlayerPickUp
		person.Roll(args)
		return
	}
	GetClientManger().SendError(agent)
}

// ///////////////////////对战类//////////////////////////////////////////
// 玩家准备
func ready(args []interface{}) {
	if person, handle, ok := verifyMsg(args); ok {
		if gameHandle, ok1 := handle.(IAgainst); ok1 {
			person.State = protoMsg.PlayerState_PlayerAgree
			gameHandle.Ready(args)
			return
		}
	}
}

// 对战类下分
func call(args []interface{}) {
	if person, gameHandle, ok := verifyMsg(args); ok {
		if gameHandel, ok := gameHandle.(IAgainst); ok {
			person.State = protoMsg.PlayerState_PlayerCall
			gameHandel.CallScore(args)
			return
		}
	}
}

// 对战类出牌
func outCard(args []interface{}) {

	if person, gameHandle, ok := verifyMsg(args); ok {
		if gameHandel, ok := gameHandle.(IAgainst); ok {
			person.State = protoMsg.PlayerState_PlayerOutCard
			gameHandel.OutCard(args)
			return
		}
	}
}

// 对战类 过牌
func disCard(args []interface{}) {
	if person, gameHandle, ok := verifyMsg(args); ok {
		if gameHandel, ok := gameHandle.(IAgainst); ok {
			person.State = protoMsg.PlayerState_PlayerOutCard
			gameHandel.Discard(args)
			return
		}
	}

}

// ///////////////////////////////用户行为（百人类）////////////////////////////////////////////////
// 抢庄
func host(args []interface{}) {
	if person, _, ok := verifyMsg(args); ok {
		person.State = protoMsg.PlayerState_PlayerCall
		person.Host(args)
	}
}

// 超级抢庄
func superHost(args []interface{}) {
	if person, _, ok := verifyMsg(args); ok {
		person.State = protoMsg.PlayerState_PlayerCall
		person.SuperHost(args)
	}
}

func hostlist(args []interface{}) {
	if person, _, ok := verifyMsg(args); ok {
		person.HostList(args)
	}
}

// ///////////////////////////////用户行为////////////////////////////////////////////////
// 跟注
func follow(args []interface{}) {

	if person, gameHandle, ok := verifyMsg(args); ok {
		if _, ok := gameHandle.(IAgainst); ok {
			person.State = protoMsg.PlayerState_PlayerFollow
			person.Follow(args)
			return
		}
	}

}

// 看牌
func look(args []interface{}) {
	if person, gameHandle, ok := verifyMsg(args); ok {
		if _, ok := gameHandle.(IAgainst); ok {
			person.State = protoMsg.PlayerState_PlayerLook
			person.Look(args)
			return
		}
	}
}

func compare(args []interface{}) {
	if person, gameHandle, ok := verifyMsg(args); ok {
		if _, ok := gameHandle.(IAgainst); ok {
			person.State = protoMsg.PlayerState_PlayerCompare
			person.Compare(args)
			return
		}
	}

}
