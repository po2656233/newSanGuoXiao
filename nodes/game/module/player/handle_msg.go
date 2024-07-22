package player

import (
	"github.com/po2656233/goleaf/gate"
	clog "github.com/po2656233/superplace/logger"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/online"
	"time"
)

// enter 进入
func enter(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0].(*protoMsg.EnterGameReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[enter]params %+v  uid:%+v", m, uid)

	person := mgr.GetPlayerMgr().Get(uid)
	if person == nil {
		// 获取玩家信息
		data, errCode := rpc.SendData(agent.App(), SourcePath, DBActor, NodeTypeCenter, &protoMsg.GetUserInfoReq{Uid: uid})
		if errCode == 0 && data != nil {
			resp, ok := data.(*protoMsg.GetUserInfoResp)
			if ok && resp.Info != nil {
				person = mgr.SimpleToPlayer(resp.Info)
				online.BindPlayer(uid, person.UserID, agent.PathString())
				mgr.GetPlayerMgr().Append(person)
			}
		}
	}

	if person == nil {
		clog.Warnf("[enter]params %+v  uid:%+v err:%s", m, uid, StatusText[User03])
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	agent.SetUserData(person)
	mgr.GetClientMgr().Append(person.UserID, agent)

	// 玩家仍在游戏中
	if person.GameHandle != nil {
		person.GameHandle.Scene([]interface{}{agent})
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
				tb, _ = room.AddTable(tbResp.Info, NewGame)
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
	_ = args[1]
	m := args[0].(*protoMsg.JoinGameReadyQueueReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[join]params %+v  uid:%+v", m, uid)
	person := mgr.GetPlayerMgr().Get(uid)
	// 玩家仍在游戏中
	if person.GameHandle != nil {
		clog.Warnf("[enter] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(FAILED, StatusText[Title001], StatusText[Login09])
		return
	}
	// 玩家加入游戏准备列表
	person.Join(args)
}

// 退出游戏
func exit(args []interface{}) {
	_ = args[1]
	m := args[0].(*protoMsg.ExitGameReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[exit]params %+v  uid:%+v", m, uid)
	person := mgr.GetPlayerMgr().Get(uid)
	if person == nil {
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	person.Exit()
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
