package player

import (
	"fmt"
	superConst "github.com/po2656233/superplace/const"
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	"gorm.io/gorm"
	"strings"
	cst "superman/internal/constant"
	event2 "superman/internal/event"
	commMsg "superman/internal/protocol/go_file/common"
	gameMsg "superman/internal/protocol/go_file/game"
	gateMsg "superman/internal/protocol/go_file/gate"
	sqlmodel "superman/internal/sql_model/minigame"
	"superman/internal/utils"
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

	p.Remote().Register(p.GetClassList)
	p.Remote().Register(p.GetRoomList)
	p.Remote().Register(p.GetTableList)
	p.Remote().Register(p.GetGameList)
	p.Remote().Register(p.GetTable)

	p.Remote().Register(p.CreateRoom)
	p.Remote().Register(p.CreateTable)
	p.Remote().Register(p.DeleteTable)

	p.Remote().Register(p.Recharge)
	p.Remote().Register(p.AddRecord)
	p.Remote().Register(p.DecreaseGameRun)
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

	player := &ActorPlayer{
		isOnline: false,
	}
	player.SetSession(msg.Session)
	childActor, err := p.Child().Create(childID, player)

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
	resp, err := p.GetGameList(&gateMsg.GetGameListReq{
		Kid: cst.Unlimited,
	})
	if err != nil {
		log.Errorf("[ActorGame] checkGameList err:%v ", err)
		return
	}
	if resp != nil && resp.Items != nil {
		p.Timer().Remove(p.checkGamesTimeId)
		mgr.GetGameInfoMgr().AddGames(resp.Items.Items)
		return
	}

	p.checkGamesTimeId = p.Timer().AddOnce(p.getGamesTime, p.checkGameList)
}

// cron
func (p *ActorGame) checkTableList(rid int64) {
	req := &gateMsg.GetTableListReq{
		Rid: rid,
	}
	resp, err := p.GetTableList(req)
	if err != nil {
		log.Errorf("[ActorGame] checkTableList err:%v", err)
		return
	}
	if resp == nil || resp.Items == nil {
		return
	}
	for _, item := range resp.Items.Items {
		room := mgr.GetRoomMgr().GetRoom(item.Rid)
		if room == nil {
			continue
		}
		if _, err = room.AddTable(item, NewGame); err != nil {
			log.Errorf("%v", err)
		}
	}

	//p.Timer().AddOnce(p.getGamesTime, func() {
	//	p.checkTableList(rid)
	//})
}

// cron
func (p *ActorGame) checkRoomList() {
	req := &gateMsg.GetRoomListReq{
		Uid: cst.Unlimited,
	}

	resp, err := p.GetRoomList(req)
	if err != nil {
		log.Errorf("[ActorGame] checkRoomList err:%v", err)
	}
	if resp != nil && resp.Items != nil {
		for _, item := range resp.Items.Items {
			rm := mgr.GetRoomMgr().AddRoom(item)
			p.checkTableList(rm.Id)
		}
		return
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

//////////////////////////////////////////////////////////////////

// GetClassList 取分类列表
func (self *ActorGame) GetClassList() (*gateMsg.GetClassListResp, error) {
	resp := &gateMsg.GetClassListResp{
		Items: &commMsg.ClassList{
			Classify: make([]*commMsg.ClassItem, 0),
		},
	}
	ret, err := mgr.GetDBCmpt().CheckClassify()
	for _, kind := range ret {
		resp.Items.Classify = append(resp.Items.Classify, &commMsg.ClassItem{
			Id:     kind.ID,
			Name:   kind.Name,
			EnName: kind.EnName,
		})
	}
	return resp, err
}

// GetRoomList 取分类列表
func (self *ActorGame) GetRoomList(req *gateMsg.GetRoomListReq) (*gateMsg.GetRoomListResp, error) {
	resp := &gateMsg.GetRoomListResp{
		Items: &commMsg.RoomList{
			Items: make([]*commMsg.RoomInfo, 0),
		},
	}
	ret, err := mgr.GetDBCmpt().CheckRooms(req.Uid, 0, -1, 0)
	for _, room := range ret {
		resp.Items.Items = append(resp.Items.Items, &commMsg.RoomInfo{
			Id:         room.ID,
			HostId:     room.Hostid,
			Level:      commMsg.RoomLevel(room.Level),
			Name:       room.Name,
			RoomKey:    room.Roomkey,
			EnterScore: room.Enterscore,
			MaxPerson:  room.MaxPerson,
			TableCount: room.TableCount,
			MaxTable:   room.MaxTable,
		})
	}
	return resp, err
}

// GetTableList 取分类列表
func (self *ActorGame) GetTableList(req *gateMsg.GetTableListReq) (*gateMsg.GetTableListResp, error) {
	resp := &gateMsg.GetTableListResp{
		Items: &commMsg.TableList{
			Items: make([]*commMsg.TableInfo, 0),
		},
	}
	ret, err := mgr.GetDBCmpt().CheckTables(req.Rid, -1, 0)
	for _, table := range ret {
		resp.Items.Items = append(resp.Items.Items, &commMsg.TableInfo{
			Id:         table.ID,
			Rid:        table.Rid,
			Name:       table.Name,
			Gid:        table.Gid,
			OpenTime:   table.Opentime,
			Commission: table.Commission,
			Remain:     table.Remain,
			MaxRound:   table.Maxround,
			PlayScore:  table.Playscore,
			MaxSitter:  table.MaxSitter,
		})
	}
	return resp, err
}

// GetGameList 取分类列表
func (self *ActorGame) GetGameList(req *gateMsg.GetGameListReq) (*gateMsg.GetGameListResp, error) {
	resp := &gateMsg.GetGameListResp{
		Items: &commMsg.GameList{
			Items: make([]*commMsg.GameInfo, 0),
		},
	}
	ret, err := mgr.GetDBCmpt().CheckGames(req.Kid, -1, 0)
	for _, game := range ret {
		resp.Items.Items = append(resp.Items.Items, &commMsg.GameInfo{
			Id:        game.ID,
			Name:      game.Name,
			Kid:       game.Kid,
			Lessscore: game.Lessscore,
			State:     commMsg.GameState(game.State),
			MaxPlayer: game.MaxPlayer,
			HowToPlay: game.HowToPlay,
		})
	}
	return resp, err
}

///////////////////////create////////////////////////////////////

func (self *ActorGame) CreateRoom(req *gateMsg.CreateRoomReq) (*gateMsg.CreateRoomResp, error) {
	resp := &gateMsg.CreateRoomResp{}
	maxCount := int32(0)
	maxTable := int32(0)
	switch req.Level {
	case commMsg.RoomLevel_GeneralRoom:
		maxCount = 20
		maxTable = 5
	case commMsg.RoomLevel_MiddleRoom:
		maxCount = 50
		maxTable = 10
	case commMsg.RoomLevel_HighRoom:
		maxCount = 200
		maxTable = 20
	case commMsg.RoomLevel_TopRoom:
		maxCount = 1000
		maxTable = 50
	case commMsg.RoomLevel_SuperRoom:
		maxCount = 10000
		maxTable = 100
	case commMsg.RoomLevel_SystemRoom:
		maxCount = -1
		maxTable = -1

	}
	req.RoomKey = utils.Md5Sum(req.RoomKey)
	roomId, err := mgr.GetDBCmpt().AddRoom(sqlmodel.Room{
		Hostid:     req.HostId,
		Level:      int32(req.Level),
		Name:       req.Name,
		Roomkey:    req.RoomKey,
		Enterscore: req.EnterScore,
		Taxation:   req.Taxation,
		MaxTable:   maxTable,
		TableCount: 0,
		MaxPerson:  maxCount,
		Remark:     req.Remark,
		CreatedAt:  time.Now(),
		DeletedAt:  gorm.DeletedAt{},
		UpdateBy:   0,
		CreateBy:   req.HostId,
	})
	if 0 == roomId || err != nil {
		return nil, err
	}
	resp.Info = &commMsg.RoomInfo{
		Id:         roomId,
		HostId:     req.HostId,
		Level:      req.Level,
		Name:       req.Name,
		RoomKey:    req.RoomKey,
		EnterScore: req.EnterScore,
		MaxPerson:  maxCount,
		TableCount: 0,
		MaxTable:   maxTable,
	}
	return resp, nil
}

func (self *ActorGame) CreateTable(req *gateMsg.CreateTableReq) (*gateMsg.CreateTableResp, error) {
	resp := &gateMsg.CreateTableResp{}
	if req.Name == "" {
		// 则获取游戏名称
		req.Name, _ = mgr.GetDBCmpt().CheckGameName(req.Gid)
	}
	tid, maxSit, err := mgr.GetDBCmpt().AddTable(sqlmodel.Table{
		Gid:        req.Gid,
		Rid:        req.Rid,
		Name:       req.Name,
		Playscore:  req.PlayScore,
		Commission: req.Commission,
		Maxround:   req.MaxRound,
		Remain:     req.MaxRound,
		Opentime:   req.Opentime,
	})
	if err != nil {
		return resp, err
	}
	resp.Table = &commMsg.TableInfo{
		Id:         tid,
		Rid:        req.Rid,
		Gid:        req.Gid,
		Name:       req.Name,
		Commission: req.Commission,
		PlayScore:  req.PlayScore,
		MaxSitter:  maxSit,
		MaxRound:   req.MaxRound,
		Remain:     req.MaxRound, // 剩余以最大局数为准
		OpenTime:   req.Opentime,
	}
	return resp, err
}

func (self *ActorGame) DeleteTable(req *gateMsg.DeleteTableReq) (*gateMsg.DeleteTableResp, error) {
	resp := &gateMsg.DeleteTableResp{}
	rid := mgr.GetDBCmpt().CheckTableRid(req.Tid)
	if rid == 0 {
		return resp, fmt.Errorf("tid:%d no have rid", req.Tid)
	}
	ok, err := mgr.GetDBCmpt().CheckRoomExist(req.HostId, rid)
	if !ok || err != nil {
		return resp, fmt.Errorf("hostid:%d rid:%d no exist err:%v", req.HostId, rid, err)
	}
	// 检测游戏是否处于关闭状态，才能
	err = mgr.GetDBCmpt().DelTable(req.HostId, req.Tid)
	resp.Tid = req.Tid
	resp.Rid = rid
	return resp, err
}

// /////////////////////////////////////////////////////////////////////////////////////

func (self *ActorGame) GetTable(req *gateMsg.GetTableReq) (*gateMsg.GetTableResp, error) {
	resp := &gateMsg.GetTableResp{}
	info, err := mgr.GetDBCmpt().CheckTable(req.Tid)
	if err != nil {
		return nil, err
	}
	resp.Info = &commMsg.TableInfo{
		Id:         info.ID,
		Name:       info.Name,
		Rid:        info.Rid,
		Gid:        info.Gid,
		OpenTime:   info.Opentime,
		Commission: info.Commission,
		Remain:     info.Remain,
		MaxRound:   info.Maxround,
		PlayScore:  info.Playscore,
		MaxSitter:  info.MaxSitter,
	}
	return resp, err
}

func (self *ActorGame) Recharge(req *gateMsg.RechargeReq) (*gateMsg.RechargeResp, error) {
	info := &sqlmodel.Recharge{
		UID:       req.UserID,
		Byid:      req.ByiD,
		Payment:   req.Payment,
		Premoney:  0,
		Money:     0,
		Code:      req.Method,
		Order:     "",
		Timestamp: time.Now().Unix(),
		Remark:    req.Reason,
		Switch:    req.Switch,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		UpdateBy:  0,
		CreateBy:  0,
	}
	err := mgr.GetDBCmpt().AddRecharge(info)
	if err != nil {
		return nil, err
	}
	return &gateMsg.RechargeResp{
		UserID:    info.UID,
		ByiD:      info.Byid,
		PreMoney:  info.Premoney,
		Payment:   info.Payment,
		Money:     info.Money,
		YuanBao:   info.Payment * 100,
		Coin:      info.Payment * 10000,
		Method:    info.Switch,
		IsSuccess: true,
		Order:     info.Order,
		TimeStamp: info.Timestamp,
		Reason:    info.Remark,
	}, nil
}

func (self *ActorGame) AddRecord(req *gameMsg.AddRecordReq) (*gameMsg.AddRecordResp, error) {
	resp := &gameMsg.AddRecordResp{}
	record := &sqlmodel.Record{
		UID:       req.Uid,
		Tid:       req.Tid,
		Payment:   req.Payment,
		Code:      req.Code,
		Order:     req.Order,
		Result:    req.Result,
		Remark:    req.Remark,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
		UpdateBy:  0,
		CreateBy:  0,
	}
	err := mgr.GetDBCmpt().AddRecord(record)
	if err != nil {
		return resp, err
	}
	resp.PerGold = record.Pergold
	resp.Gold = record.Gold
	return resp, nil
}
func (self *ActorGame) DecreaseGameRun(req *gameMsg.DecreaseGameRunReq) (*gameMsg.DecreaseGameRunResp, error) {
	resp := &gameMsg.DecreaseGameRunResp{}
	remain, err := mgr.GetDBCmpt().EraseRemain(req.Tid, req.Amount)
	resp.Tid = req.Tid
	resp.Remain = remain
	return resp, err
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
