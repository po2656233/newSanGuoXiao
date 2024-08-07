package player

import (
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	cst "superman/internal/constant"
	event2 "superman/internal/event"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/category"
	"superman/nodes/game/module/online"
	"time"
)

type (
	// ActorGame 每位登录的玩家对应一个子actor
	ActorGame struct {
		//pomelo.ActorBase
		simple.ActorBase
		checkGamesTimeId uint64
		childExitTime    time.Duration
		getGamesTime     time.Duration
		getRoomsTime     time.Duration
	}
)

func (p *ActorGame) AliasID() string {
	return "game"
}
func (p *ActorGame) OnInit() {
	log.Debugf("[ActorGame] path = %s init!", p.PathString())
	p.childExitTime = time.Minute * 30
	p.getGamesTime = time.Second * 5
	p.getRoomsTime = time.Minute * 1
	category.InitConfig()

	// 注册角色登陆事件
	p.Event().Register(p.onLoginEvent)
	p.Event().Register(p.onLogoutEvent)
	p.Event().Register(p.onPlayerCreateEvent)
	p.Remote().Register(p.checkChild)
	// 接收来自中心服的数据修改
	p.Remote().Register(p.CreateRoomResp)
	p.Remote().Register(p.CreateTableResp)
	p.Remote().Register(p.DeleteTableResp)

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
	log.Debugf("[ActorGame] path = %s exit!", p.PathString())
}

// cron
func (p *ActorGame) checkGameList() {
	data, errCode := rpc.SendData(p.App(), cst.SourcePath, cst.DBActor, cst.NodeTypeCenter, &pb.GetGameListReq{
		Kid: cst.Unlimited,
	})
	if errCode == 0 {
		resp, ok := data.(*pb.GetGameListResp)
		if ok && resp != nil && resp.Items != nil {
			p.Timer().Remove(p.checkGamesTimeId)
			mgr.GetGameInfoMgr().AddGames(resp.Items.Items)
			return
		}
	}
	p.checkGamesTimeId = p.Timer().AddOnce(p.getGamesTime, p.checkGameList)
}

// cron
func (p *ActorGame) checkTableList(rid int64) {
	req := &pb.GetTableListReq{
		Rid: rid,
	}

	data, errCode := rpc.SendData(p.App(), cst.SourcePath, cst.DBActor, cst.NodeTypeCenter, req)
	if errCode == 0 {
		resp, ok := data.(*pb.GetTableListResp)
		if !ok || resp == nil || resp.Items == nil {
			return
		}
		for _, item := range resp.Items.Items {
			room := mgr.GetRoomMgr().GetRoom(item.Rid)
			if room == nil {
				continue
			}
			if _, err := room.AddTable(item, NewGame); err != nil {
				log.Errorf("rid:%v add table:%v is err:%v !!!!", room.Id, item.Id, err)
			}

		}
		return
	}
	p.Timer().AddOnce(p.getGamesTime, func() {
		p.checkTableList(rid)
	})
}

// cron
func (p *ActorGame) checkRoomList() {
	req := &pb.GetRoomListReq{
		Uid: cst.Unlimited,
	}

	data, errCode := rpc.SendData(p.App(), cst.SourcePath, cst.DBActor, cst.NodeTypeCenter, req)
	if errCode == 0 {
		resp, ok := data.(*pb.GetRoomListResp)
		if ok && resp != nil && resp.Items != nil {
			for _, item := range resp.Items.Items {
				rm := mgr.GetRoomMgr().AddRoom(item)
				p.checkTableList(rm.Id)
			}
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

func (p *ActorGame) CreateRoomResp(resp *pb.CreateRoomResp) {
	if _, ok := mgr.GetRoomMgr().Create(resp.Info); !ok {
		log.Errorf("[ActorGame][CreateRoomResp] [Create] no good!")
	}
}
func (p *ActorGame) CreateTableResp(resp *pb.CreateTableResp) {
	if resp.Table == nil {
		log.Errorf("[ActorGame][CreateTableResp] no resp!")
		return
	}
	if rm := mgr.GetRoomMgr().GetRoom(resp.Table.Rid); rm != nil {
		if _, err := rm.AddTable(resp.Table, NewGame); err != nil {
			log.Errorf("[ActorGame][CreateTableResp] [AddTable] no resp!")
		}
	}
}
func (p *ActorGame) DeleteTableResp(resp *pb.DeleteTableResp) {
	if rm := mgr.GetRoomMgr().GetRoom(resp.Rid); rm != nil {
		rm.DelTable(resp.Tid)
	}
}

/////////////////////////////////////////////////////////////////

// onLoginEvent 玩家登陆事件处理
func (p *ActorGame) onLoginEvent(e cfacade.IEventData) {
	evt, ok := e.(*event2.PlayerLogin)
	if ok == false {
		return
	}

	log.Infof("[PlayerLoginEvent] [playerId = %d, onlineCount = %d]",
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

	log.Infof("[PlayerLogoutEvent] [playerId = %d, onlineCount = %d]",
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

	log.Infof("[PlayerCreateEvent] [%+v]", evt)
}
