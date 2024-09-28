package player

import (
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	event2 "superman/internal/event"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/online"
	"time"
)

type (
	// ActorPlayer 玩家总管理actor
	ActorPlayer struct {
		//pomelo.ActorBase
		simple.ActorBase
		isOnline bool // 玩家是否在线
		playerId int64
		uid      int64
		Session  *cproto.Session
		userData interface{}
	}
)

//	func (p *ActorPlayer) AliasID() string {
//		return "game"
//	}

func (p *ActorPlayer) OnInit() {
	// 注册 session关闭的remote函数(网关触发连接断开后，会调用RPC发送该消息)
	p.Remote().Register(p.sessionClose)
	p.registerLocalMsg()
}

// sessionClose 接收角色session关闭处理
func (p *ActorPlayer) sessionClose() {
	online.UnBindPlayer(p.uid)
	p.isOnline = false
	p.Exit()

	logoutEvent := event2.NewPlayerLogout(p.ActorID(), p.playerId)
	p.PostEvent(&logoutEvent)

}

////////////////////////////////////////////////////////////////////////////////////////////

func (p *ActorPlayer) OnStop() {
	clog.Infof("OnStop onlineCount = %d", online.Count())
	if p.Session == nil {
		clog.Warnf("OnStop uid = %d", p.uid)
		return
	}
	// 移除客户端，避免再向其反馈消息
	mgr.GetClientMgr().DeleteClient(p.Session.Uid)
	// 玩家退出 并且从玩家管理中删除信息
	person := mgr.GetPlayerMgr().Get(p.Session.Uid)
	if person != nil {
		person.LeaveTime = time.Now().Unix()
		if person.Exit() {
			mgr.GetPlayerMgr().Delete(person.UserID)
			person = nil
		}
	}
}
