package player

import (
	"github.com/po2656233/goleaf/gate"
	clog "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	mgr "superman/nodes/game/manger"
	"time"
)

// enter 进入
func enter(args []interface{}) {
	//查找玩家
	m := args[0].(*protoMsg.EnterGameReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[enter]params %+v  uid:%+v", m, uid)

	person := mgr.GetPlayerMgr().Get(uid)
	// 玩家仍在游戏中
	if person.GameHandle != nil {
		person.GameHandle.Scene([]interface{}{person})
		return
	}

	room := mgr.GetRoomMgr().GetRoom(m.RoomID)
	if room == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Room16])
		return
	}
	var tb *mgr.Table
	tb = room.GetTable(m.TableID)
	if tb == nil {
		data, errCode := rpc.SendData(agent.App(), SourcePath, DBActor, NodeTypeCenter, &protoMsg.GetTableReq{Tid: m.TableID})
		if errCode == 0 {
			if tbResp, ok := data.(*protoMsg.GetTableResp); ok {
				var err error
				tb, err = room.AddTable(tbResp.Info, NewGame)
				if err != nil {
					clog.Warnf("[enter] params %+v  uid:%+v AddTable err:%v", m, uid, err)
					agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Room17])
					return
				}
			}
		}
	}
	if tb == nil {
		clog.Warnf("[enter] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[TableInfo09])
		return
	}

	person.GameHandle = tb.GameHandle
	person.Enter(args)
}

func join(args []interface{}) {
	m := args[0].(*protoMsg.JoinGameReadyQueueReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[join]params %+v  uid:%+v", m, uid)

	person := mgr.GetPlayerMgr().Get(uid)
	// 玩家仍在游戏中
	if person.GameHandle != nil {
		clog.Warnf("[join] [GameHandle] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Login09])
		person.GameHandle.Scene([]interface{}{person})
		return
	}
	// 获取房间句柄
	room := mgr.GetRoomMgr().GetRoom(m.RoomID)
	if room == nil {
		clog.Warnf("[join] [room] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Room16])
		return
	}

	// 检测游戏是否可用
	gInfo := mgr.GetGameInfoMgr().GetGame(m.GameID)
	if gInfo == nil {
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[TableInfo03])
		return
	}
	if gInfo.State == protoMsg.GameState_InitTB {
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Game16])
		return
	}
	if gInfo.State == protoMsg.GameState_CloseTB {
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Game19])
		return
	}
	// 玩家加入游戏准备列表
	if ok := room.AddWait(m.GameID, person); !ok {
		clog.Warnf("[join] [GameHandle] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Login09])
		return
	}
	// 检测房间的游戏是否满员
	if ok := mgr.GetRoomMgr().GetRoom(m.RoomID).CheckGameFull(m.GameID); ok {
		if m.RoomID != SYSTEMID {
			// 牌桌不足
			agent.SendResult(FAILED, StatusText[TableInfo10])
			return
		}
		// 系统房主动添加牌桌。 如果没有加入队列,并且是系统房,则创建牌桌
		// 请求数据库创建牌桌
		data, errCode := rpc.SendDataToDB(agent.App(), &protoMsg.CreateTableReq{
			Rid:        SYSTEMID,
			Gid:        m.GameID,
			PlayScore:  Unlimited,
			Name:       "",
			Opentime:   time.Now().Unix(),
			Taxation:   0, //系统房没有税收
			Commission: 0,
			MaxRound:   Unlimited,
		})
		if errCode != SUCCESS {
			clog.Warnf("[join] [SendDataToDB] params %+v  uid:%+v FAIL", m, uid)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
		resp, ok1 := data.(*protoMsg.CreateTableResp)
		if !ok1 || resp == nil || resp.Table == nil {
			clog.Warnf("[join]  params %+v  uid:%+v FAIL", m, uid)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
		// 给系统房添加刚创建的牌桌
		t, err := room.AddTable(resp.Table, NewGame)
		if err != nil || t == nil {
			clog.Warnf("[join] [AddTable] params %+v  uid:%+v FAIL. err:%v", m, uid, err)
			agent.SendResultPop(FAILED, StatusText[Title001], StatusText[User29])
			return
		}
	}

	// 成功添加队列
	agent.SendResult(SUCCESS, StatusText[User30])
}

// 退出游戏
func exit(args []interface{}) {
	_ = args[1]
	m := args[0].(*protoMsg.ExitGameReq)
	agent := args[1].(*ActorPlayer)
	clog.Debugf("[exit]params %+v  uid:%+v", m, agent.Session.Uid)
	mgr.GetPlayerMgr().Get(agent.Session.Uid).Exit()
}

func ready(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if m == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	uid := agent.Session.Uid
	person := mgr.GetPlayerMgr().Get(uid)
	if ok := person.UpdateState(protoMsg.PlayerState_PlayerReady, args); !ok {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User02])
	}
}

// 设置时长
func setTime(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if m == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	uid := agent.Session.Uid
	person := mgr.GetPlayerMgr().Get(uid)
	if ok := person.UpdateState(protoMsg.PlayerState_PlayerSetTime, args); !ok {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User02])
	}
}

// 玩法操作
func playing(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if m == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	person := mgr.GetPlayerMgr().Get(agent.Session.Uid)
	if person.GameHandle != nil {
		person.GameHandle.Playing(args)
	}
}

// 意见反馈
func handleSuggest(args []interface{}) {
	_ = args[0].(*protoMsg.SuggestReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*mgr.Player)
		msg := &protoMsg.SuggestResp{
			UserID: person.UserID,
			Feedback: &protoMsg.EmailInfo{
				//EmailID:    GetPChatID(),
				TimeStamp:  time.Now().Unix(),
				AcceptName: person.Name,
				Topic:      "感谢信",
				Content:    "由衷感谢您的反馈,您的反馈是我们持续的动力!",
				Sender:     "系统邮件",
			},
		}
		mgr.GetClientMgr().SendData(agent, msg)
		mgr.GetClientMgr().SendPopResult(agent, SUCCESS, StatusText[Title007], StatusText[User22])
	}
}
func handleNotifyNotice(args []interface{}) {
	m := args[0].(*protoMsg.NotifyNoticeReq)
	agent := args[1].(mgr.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*mgr.Player)
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
		//RedisClient.Set("mini_Notice", msg, time.Duration(m.Timeout)*time.Second)

		//if person.UserID == SYSTEMID {
		//	msg.TimeInfo.TotalTime = Date
		//}
		mgr.GetClientMgr().NotifyAll(msg)
		//GetClientManger().SendPopResult(agent, SUCCESS, StatusText[Title001], StatusText[User22])
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
