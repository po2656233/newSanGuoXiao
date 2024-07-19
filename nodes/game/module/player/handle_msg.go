package player

import (
	clog "github.com/po2656233/superplace/logger"
	"superman/internal/constant"
	. "superman/internal/hints"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	"superman/nodes/game/manger"
)

// enter 进入
func enter(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0].(*protoMsg.EnterGameReq)
	agent := args[1].(*ActorPlayer)
	uid := agent.Session.Uid
	clog.Debugf("params %+v  uid:%+v", m, uid)

	person := manger.GetPlayerManger().Get(uid)
	if person == nil {
		agent.SendResultPop(constant.FAILED, StatusText[Title004], StatusText[User03])
		return
	}
	if person.GameHandle != nil {
		agent.SendResultPop(constant.FAILED, StatusText[Title004], StatusText[User03])
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
		if data, errCodfe := rpc.SendData(agent.App(), rpc.SourcePath, rpc.DBActor, rpc.CenterType, &protoMsg.GetTableReq{Tid: m.TableID}); errCodfe == 0 {
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
	clog.Debugf("params %+v  uid:%+v", m, agent.Session.Uid)
}
