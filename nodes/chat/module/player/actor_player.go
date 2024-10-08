package player

import (
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	. "superman/internal/constant"
	event2 "superman/internal/event"
	"superman/nodes/chat/db"
	"superman/nodes/game/module/online"
)

type (
	// ActorPlayer 玩家总管理actor
	ActorPlayer struct {
		//pomelo.ActorBase
		simple.ActorBase
		isOnline    bool // 玩家是否在线
		playerId    int64
		uid         int64
		session     *cproto.Session
		userData    interface{}
		dbComponent *db.Component
	}
)

//	func (p *ActorPlayer) AliasID() string {
//		return "game"
//	}

func (p *ActorPlayer) OnInit() {
	// 注册 session关闭的remote函数(网关触发连接断开后，会调用RPC发送该消息)
	p.Remote().Register(p.sessionClose)
	p.registerLocalMsg()
	p.dbComponent = p.App().Find(ChatDb).(*db.Component)
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
}
