package player

import (
	clog "github.com/po2656233/superplace/logger"
	protoMsg "superman/internal/protocol/gofile"
)

// enter 进入
func enter(args []interface{}) {
	//查找玩家
	//_ = args[1]
	m := args[0].(*protoMsg.EnterGameReq)
	agent := args[1].(*ActorPlayer)
	clog.Debugf("params %+v  uid:%+v", m, agent.Session.Uid)

	agent.WriteMsg(&protoMsg.EnterGameResp{
		GameID:   1,
		ChairNum: 1,
		UserInfo: &protoMsg.PlayerSimpleInfo{
			Uid:    agent.Session.Uid,
			Name:   agent.Session.GetSid(),
			HeadId: 0,
			Score:  0,
			RankNo: 0,
		},
	})
}

// 退出游戏
func exit(args []interface{}) {
	_ = args[1]
	m := args[0].(*protoMsg.ExitGameReq)
	agent := args[1].(*ActorPlayer)
	clog.Debugf("params %+v  uid:%+v", m, agent.Session.Uid)
}
