package player

import (
	superConst "github.com/po2656233/superplace/const"
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	"strings"
	cst "superman/internal/constant"
	event2 "superman/internal/event"
	gateProto "superman/internal/protocol/go_file/gate"
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
	return strings.Trim(cst.GameActor, superConst.DOT)
}
func (p *ActorGame) OnInit() {
	// 加载游戏配置
	category.InitConfig()
	// 加载协议文件
	rpc.LoadMsgInfos()
	// 客户端管理实例句柄
	mgr.GetClientMgr().SetApp(p.App())

	// 注册协议
	p.registerEvent() // 响应事件
	//p.registerLocalMsg()  // 注册(与客户端通信)的协议
	p.registerRemoteMsg() // 注意服务间交互的协议

	// 检查游戏基础信息(房间列表、牌桌列表)
	p.checkBaseInfo()

}

// registerEvent 响应事件
func (p *ActorGame) registerEvent() {
	// 注册角色登陆事件
	p.Event().Register(p.onLoginEvent)
	p.Event().Register(p.onLogoutEvent)
	p.Event().Register(p.onPlayerCreateEvent)

}

// registerRemoteMsg 注意服务间交互的协议
func (p *ActorGame) registerRemoteMsg() {
	p.Remote().Register(p.checkChild) // 与子节点交互
	// 接收来自中心服的数据修改
	p.Remote().Register(p.CreateRoomResp)
	p.Remote().Register(p.CreateTableResp)
	p.Remote().Register(p.DeleteTableResp)
}

// checkBaseInfo 检查基础信息
func (p *ActorGame) checkBaseInfo() {
	log.Debugf("[ActorGame] path = %s init!", p.PathString())
	p.childExitTime = time.Minute * 30
	p.getGamesTime = time.Second * 5
	p.getRoomsTime = time.Minute * 1

	p.Timer().RemoveAll()
	p.Timer().AddOnce(time.Second, p.checkGameList)
	p.Timer().AddOnce(time.Second, p.checkRoomList)
	//p.Call(p.PathString(), "checkChild", nil)

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
	data, errCode := rpc.SendDataToDB(p.App(), &gateProto.GetGameListReq{
		Kid: cst.Unlimited,
	})
	if errCode == 0 {
		resp, ok := data.(*gateProto.GetGameListResp)
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
	req := &gateProto.GetTableListReq{
		Rid: rid,
	}

	data, errCode := rpc.SendDataToDB(p.App(), req)
	if errCode == 0 {
		resp, ok := data.(*gateProto.GetTableListResp)
		if !ok || resp == nil || resp.Items == nil {
			return
		}
		for _, item := range resp.Items.Items {
			room := mgr.GetRoomMgr().GetRoom(item.Rid)
			if room == nil {
				continue
			}
			if _, err := room.AddTable(item, NewGame); err != nil {
				log.Errorf("%v", err)
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
	req := &gateProto.GetRoomListReq{
		Uid: cst.Unlimited,
	}

	data, errCode := rpc.SendDataToDB(p.App(), req)
	if errCode == 0 {
		resp, ok := data.(*gateProto.GetRoomListResp)
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
		if !ok || child.isOnline {
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

func (p *ActorGame) CreateRoomResp(resp *gateProto.CreateRoomResp) {
	if _, ok := mgr.GetRoomMgr().Create(resp.Info); !ok {
		log.Errorf("[ActorGame][CreateRoomResp] [Create] no good!")
	}
}
func (p *ActorGame) CreateTableResp(resp *gateProto.CreateTableResp) {
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
func (p *ActorGame) DeleteTableResp(resp *gateProto.DeleteTableResp) {
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
