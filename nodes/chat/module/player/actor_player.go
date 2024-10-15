package player

import (
	exString "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	superActor "github.com/po2656233/superplace/net/actor"
	"github.com/po2656233/superplace/net/parser/simple"
	"google.golang.org/protobuf/proto"
	"strconv"
	. "superman/internal/constant"
	event2 "superman/internal/event"
	"superman/internal/rpc"
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

// NotifyAll 通知玩家
func (p *ActorPlayer) NotifyAll(message proto.Message) bool {
	parent, ok := p.System().GetIActor(ActIdChat)
	if !ok {
		rpc.SendResult(&p.ActorBase, Chat028)
		return false
	}
	chatAct, ok1 := parent.(*superActor.Actor)
	if !ok1 {
		rpc.SendResult(&p.ActorBase, Chat029)
		return false
	}

	chatAct.Child().Each(func(i cfacade.IActor) {
		p.SendTo(i.SID(), message)
	})
	return true
}

// NotifyTo 通知玩家
func (p *ActorPlayer) NotifyTo(uidList []int64, message proto.Message) bool {
	parent, ok := p.System().GetIActor(ActIdChat)
	if !ok {
		rpc.SendResult(&p.ActorBase, Chat028)
		return false
	}
	chatAct, ok1 := parent.(*superActor.Actor)
	if !ok1 {
		rpc.SendResult(&p.ActorBase, Chat029)
		return false
	}
	for _, uid := range uidList {
		if act, ok2 := chatAct.Child().Get(exString.ToString(uid)); ok2 {
			rpc.SendToGate(act.App(), uid, message)
		}
	}
	return true
}

// GetOnline 通知玩家
func (p *ActorPlayer) GetOnline() []int64 {
	parent, ok := p.System().GetIActor(ActIdChat)
	if !ok {
		rpc.SendResult(&p.ActorBase, Chat028)
		return nil
	}
	chatAct, ok1 := parent.(*superActor.Actor)
	if !ok1 {
		rpc.SendResult(&p.ActorBase, Chat029)
		return nil
	}
	uidList := make([]int64, 0)
	chatAct.Child().Each(func(i cfacade.IActor) {
		i.ActorID()
		uid, _ := strconv.ParseInt(i.ActorID(), 10, 64)
		if 0 < uid {
			uidList = append(uidList, uid)
		}
	})
	return uidList
}

////////////////////////////////////////////////////////////////////////////////////////////

func (p *ActorPlayer) OnStop() {
	clog.Infof("OnStop onlineCount = %d", online.Count())
}
