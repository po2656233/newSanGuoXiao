package player

import (
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	cst "superman/internal/constant"
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
		checkGamesTimeId uint64
		haveRooms        bool
		childExitTime    time.Duration
		getGamesTime     time.Duration
		getRoomsTime     time.Duration
	}
)

func (p *ActorGame) AliasID() string {
	return "game"
}
func (p *ActorGame) OnInit() {
	clog.Debugf("[ActorGame] path = %s init!", p.PathString())
	p.childExitTime = time.Minute * 30
	p.getGamesTime = time.Second * 5
	p.getRoomsTime = time.Minute * 1
	p.haveRooms = false

	// 注册角色登陆事件
	p.Event().Register(p.onLoginEvent)
	p.Event().Register(p.onLogoutEvent)
	p.Event().Register(p.onPlayerCreateEvent)
	p.Remote().Register(p.checkChild)

	p.Timer().RemoveAll()
	p.Timer().AddOnce(time.Second, p.checkGameList)
	p.Timer().AddOnce(time.Second, p.checkRoomList)
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
func (p *ActorGame) checkGameList() {
	data, errCode := rpc.SendData(p.App(), rpc.SourcePath, rpc.DBActor, rpc.CenterType, &pb.GetGameListReq{
		Kid: cst.Unlimited,
	})
	if errCode == 0 {
		resp, ok := data.(*pb.GetGameListResp)
		if ok && resp != nil && resp.Items != nil {
			p.Timer().Remove(p.checkGamesTimeId)
			manger.GetGameInfoMgr().AddGames(resp.Items.Items)
			return
		}
	}
	p.checkGamesTimeId = p.Timer().AddOnce(p.getGamesTime, p.checkGameList)
}

// cron
func (p *ActorGame) checkRoomList() {
	req := &pb.GetRoomListReq{
		Uid: cst.Unlimited,
	}
	if p.haveRooms {
		// 获取最新的房间列表
		req.StartTime = time.Now().Add(-p.getRoomsTime).Unix()
	}

	data, errCode := rpc.SendData(p.App(), rpc.SourcePath, rpc.DBActor, rpc.CenterType, req)
	if errCode == 0 {
		resp, ok := data.(*pb.GetRoomListResp)
		if ok && resp != nil && resp.Items != nil {
			manger.GetRoomMgr().AddRooms(resp.Items.Items)
			p.haveRooms = true
			p.Timer().AddOnce(p.getRoomsTime, p.checkRoomList)
			return
		}
	}

	p.Timer().AddOnce(time.Second, p.checkRoomList)
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

/////////////////////////////////////////////////////////////////

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
