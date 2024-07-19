package player

import (
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	event2 "superman/internal/event"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	"superman/nodes/game/manger"
	"superman/nodes/game/module/online"
	"time"
)

type (
	// ActorGame 每位登录的玩家对应一个子actor
	ActorGame struct {
		//pomelo.ActorBase
		simple.ActorBase
		checkTimeId   uint64
		childExitTime time.Duration
	}
)

func (p *ActorGame) AliasID() string {
	return "game"
}
func (p *ActorGame) OnInit() {
	clog.Debugf("[ActorGame] path = %s init!", p.PathString())
	p.childExitTime = time.Minute * 30

	// 注册角色登陆事件
	p.Event().Register(p.onLoginEvent)
	p.Event().Register(p.onLogoutEvent)
	p.Event().Register(p.onPlayerCreateEvent)
	p.Remote().Register(p.checkChild)

	p.Timer().RemoveAll()
	p.Timer().AddOnce(time.Second, p.everySecondTimer)
}

func (p *ActorGame) OnFindChild(msg *cfacade.Message) (cfacade.IActor, bool) {
	// 动态创建 player child actor
	childID := msg.TargetPath().ChildID
	childActor, err := p.Child().Create(childID, &ActorPlayer{
		isOnline: false,
	})

	if err != nil {
		return nil, false
	}

	return childActor, true
}

func (p *ActorGame) OnStop() {
	clog.Debugf("[ActorGame] path = %s exit!", p.PathString())
}

// cron
func (p *ActorGame) everySecondTimer() {
	//p.Call(p.PathString(), "checkChild", nil)
	data, errCode := rpc.SendData(p.App(), rpc.SourcePath, rpc.DBActor, rpc.CenterType, &pb.GetGameListReq{
		Kid: -1,
	})
	if errCode == 0 {
		resp, ok := data.(*pb.GetGameListResp)
		if ok && resp != nil && resp.Items != nil {
			manger.GetGameMgr().AddGames(resp.Items.Items)
			return
		}
	}
	p.Timer().AddOnce(3*time.Second, p.everySecondTimer)
}

func (p *ActorGame) checkChild() {
	// 扫描所有玩家actor
	p.Child().Each(func(iActor cfacade.IActor) {
		child, ok := iActor.(*ActorPlayer)
		if !ok {
			return
		}

		if child.isOnline {
			return
		}

		// 玩家下线，并且超过childExitTime时间没有收发消息，则退出actor
		deadline := time.Now().Add(-p.childExitTime).Unix()
		if child.LastAt() < deadline {
			child.Exit() //actor退出
		}
	})
}

// onLoginEvent 玩家登陆事件处理
func (p *ActorGame) onLoginEvent(e cfacade.IEventData) {
	evt, ok := e.(*event2.PlayerLogin)
	if ok == false {
		return
	}

	clog.Infof("[PlayerLoginEvent] [playerId = %d, onlineCount = %d]",
		evt.PlayerId,
		online.Count(),
	)
}

// onLoginEvent 玩家登出事件处理
func (p *ActorGame) onLogoutEvent(e cfacade.IEventData) {
	evt, ok := e.(*event2.PlayerLogout)
	if !ok {
		return
	}

	clog.Infof("[PlayerLogoutEvent] [playerId = %d, onlineCount = %d]",
		evt.PlayerId,
		online.Count(),
	)
}

// onPlayerCreateEvent 玩家创建事件
func (p *ActorGame) onPlayerCreateEvent(e cfacade.IEventData) {
	evt, ok := e.(*event2.PlayerCreate)
	if !ok {
		return
	}

	clog.Infof("[PlayerCreateEvent] [%+v]", evt)
}
