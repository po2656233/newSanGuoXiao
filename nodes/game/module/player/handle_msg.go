package player

import (
	clog "github.com/po2656233/superplace/logger"
	"superman/internal/constant"
	. "superman/internal/hints"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	"superman/nodes/game/manger"
	"superman/nodes/game/module/online"
)

// enter 进入
func enter(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0].(*protoMsg.EnterGameReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[enter]params %+v  uid:%+v", m, uid)

	person := manger.GetPlayerMgr().Get(uid)
	if person == nil {
		// 获取玩家信息
		data, errCode := rpc.SendData(agent.App(), rpc.SourcePath, rpc.DBActor, rpc.CenterType, &protoMsg.GetUserInfoReq{Uid: uid})
		if errCode == 0 && data != nil {
			resp, ok := data.(*protoMsg.GetUserInfoResp)
			if ok && resp.Info != nil {
				person = manger.SimpleToPlayer(resp.Info)
				online.BindPlayer(uid, person.UserID, agent.PathString())
				manger.GetPlayerMgr().Append(person)
			}
		}
	}

	if person == nil {
		clog.Warnf("[enter]params %+v  uid:%+v err:%s", m, uid, StatusText[User03])
		agent.SendResultPop(constant.FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	agent.SetUserData(person)
	manger.GetClientMgr().Append(person.UserID, agent)

	// 玩家仍在游戏中
	if person.GameHandle != nil {
		person.GameHandle.Scene([]interface{}{agent})
		return
	}

	room := manger.GetRoomMgr().GetRoom(m.RoomID)
	if room == nil {
		agent.SendResultPop(constant.FAILED, StatusText[Title004], StatusText[Room16])
		return
	}
	var tb *manger.Table
	tb = room.GetTable(m.TableID)
	if tb == nil {
		data, errCode := rpc.SendData(agent.App(), rpc.SourcePath, rpc.DBActor, rpc.CenterType, &protoMsg.GetTableReq{Tid: m.TableID})
		if errCode == 0 {
			if tbResp, ok := data.(*protoMsg.GetTableResp); ok {
				tb, _ = room.AddTable(tbResp.Info, NewGame)
			}
		}
	}
	if tb == nil {
		clog.Warnf("[enter] params %+v  uid:%+v FAIL", m, uid)
		agent.SendResultPop(constant.FAILED, StatusText[Title004], StatusText[TableInfo09])
		return
	}

	person.GameHandle = tb.GameHandle
	person.Enter(args)
}

// 退出游戏
func exit(args []interface{}) {
	_ = args[1]
	m := args[0].(*protoMsg.ExitGameReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("[exit]params %+v  uid:%+v", m, uid)
	person := manger.GetPlayerMgr().Get(uid)
	if person == nil {
		agent.SendResultPop(constant.FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	person.Exit()
}
