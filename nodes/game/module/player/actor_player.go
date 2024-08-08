package player

import (
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	"golang.org/x/net/websocket"
	"net"
	"net/url"
	"reflect"
	. "superman/internal/constant"
	event2 "superman/internal/event"
	protoMsg "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	"superman/internal/utils"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/online"
	"superman/nodes/game/msg"
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

// 绑定协议与其处理函数
func init() {
	//玩家行为
	handlerMsg(&protoMsg.EnterGameReq{}, enter)
	handlerMsg(&protoMsg.JoinGameReadyQueueReq{}, join)
	handlerMsg(&protoMsg.ExitGameReq{}, exit)
	handlerMsg(&protoMsg.ChineseChessReadyReq{}, ready)
	handlerMsg(&protoMsg.ChineseChessSetTimeReq{}, setTime)
	//handlerMsg(&protoMsg.DisbandedGameReq{}, disbandedGame)
	//handlerMsg(&protoMsg.TrusteeReq{}, trustee)
	//handlerMsg(&protoMsg.ChangeTableReq{}, changeTable)
	//handlerMsg(&protoMsg.GetInningsInfoReq{}, getInnings)
	//handlerMsg(&protoMsg.GameRecord{}, getGameRecords)
	//handlerMsg(&protoMsg.GetRecordReq{}, getRecords)
	//handlerMsg(&protoMsg.GetBackPasswordReq{}, getBackPassword)
	//
	////系统
	//handlerMsg(&protoMsg.UpdateGoldReq{}, updateGold)
	//handlerMsg(&protoMsg.SuggestReq{}, handleSuggest)           //意见反馈
	//handlerMsg(&protoMsg.NotifyNoticeReq{}, handleNotifyNotice) //意见反馈
	////slotGame
	//// zhaocaimiao
	//handlerMsg(&protoMsg.ZhaocaimiaoBetReq{}, playing)

	//more people Game
	handlerMsg(&protoMsg.BaccaratBetReq{}, playing)
	handlerMsg(&protoMsg.BrcowcowBetReq{}, playing)
	handlerMsg(&protoMsg.TigerXdragonBetReq{}, playing)
	handlerMsg(&protoMsg.BrtoubaoBetReq{}, playing)
	handlerMsg(&protoMsg.BrTuitongziBetReq{}, playing)

	go func() {
		for {
			msg.ServerChanRPC.Exec(<-msg.ServerChanRPC.ChanCall)
		}
	}()
}

// 注册传输消息
func handlerMsg(m interface{}, h interface{}) {
	msg.ServerChanRPC.Register(reflect.TypeOf(m), h)
}
func (p *ActorPlayer) OnInit() {
	// 注册 session关闭的remote函数(网关触发连接断开后，会调用RPC发送该消息)
	p.Remote().Register(p.sessionClose)
	p.Local().Register(p.request)

}

// sessionClose 接收角色session关闭处理
func (p *ActorPlayer) sessionClose() {
	online.UnBindPlayer(p.uid)
	p.isOnline = false
	p.Exit()

	logoutEvent := event2.NewPlayerLogout(p.ActorID(), p.playerId)
	p.PostEvent(&logoutEvent)

}
func (p *ActorPlayer) request(session *cproto.Session, req *protoMsg.Request) {
	// 派发给各个模块
	msgData, err := msg.ProcessorProto.Unmarshal(req.Data)
	if err != nil {
		clog.Debugf("unmarshal message error: %v", err)
		p.SendResult(FAILED, StatusText[NetworkErr09])
		return
	}
	uid := session.GetUid()
	p.Session = session
	p.playerId = uid
	person := mgr.GetPlayerMgr().Get(uid)
	if person == nil {
		// 获取玩家信息
		data, errCode := rpc.SendData(p.App(), SourcePath, DBActor, NodeTypeCenter, &protoMsg.GetUserInfoReq{Uid: uid})
		if errCode == 0 && data != nil {
			resp, ok := data.(*protoMsg.GetUserInfoResp)
			if ok && resp.Info != nil {
				person = mgr.ToPlayer(resp.Info)
				online.BindPlayer(uid, person.UserID, p.PathString())
				mgr.GetPlayerMgr().Append(person)
			}
		}
	}
	if person == nil {
		userData := p.UserData()
		if userData != nil {
			person = userData.(*mgr.Player)
			mgr.GetPlayerMgr().Append(person)
		}
	} else {
		p.SetUserData(person)
		mgr.GetClientMgr().Append(person.UserID, p)
	}
	if !checkArgs([]interface{}{msgData, p}) {
		return
	}
	err = msg.ProcessorProto.Route(msgData, p)
	utils.CheckError(err)

	//p.Call(session.AgentPath, "Response", &protoMsg.Response{})
}

func (p *ActorPlayer) Send(funName string, resp interface{}) {
	p.Call(p.Session.AgentPath, funName, resp)
}

func (p *ActorPlayer) SendResult(state int32, hints string) {
	p.WriteMsg(&protoMsg.ResultResp{
		State: state,
		Hints: hints,
	})
}

func (p *ActorPlayer) SendResultPop(state int32, title, hints string) {
	p.WriteMsg(&protoMsg.ResultPopResp{
		Flag:  state,
		Title: title,
		Hints: hints,
	})
}

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
	if person != nil && person.Exit() {
		mgr.GetPlayerMgr().Delete(person.UserID)
		person = nil
	}
}

//////////////////////////////实现Agent/////////////////////////////////////////////////////////

func (p *ActorPlayer) WriteMsg(v interface{}) {
	msgData, err := msg.ProcessorProto.Marshal(v)
	if err != nil {
		clog.Debug("unmarshal message error: %v", err)
		return
	}
	data := make([]byte, 0)
	data = append(data, msgData[0]...)
	data = append(data, msgData[1]...)
	p.Response(p.Session, data)
}
func (p *ActorPlayer) LocalAddr() net.Addr {
	return nil
}
func (p *ActorPlayer) RemoteAddr() net.Addr {
	return &websocket.Addr{
		URL: &url.URL{
			Host: p.Session.GetIp(),
		},
	}
}
func (p *ActorPlayer) Close() {
	p.sessionClose()
}
func (p *ActorPlayer) Destroy() {

}
func (p *ActorPlayer) UserData() interface{} {
	return p.userData
}
func (p *ActorPlayer) SetUserData(data interface{}) {
	p.userData = data
}

////////////////////////////////////////////////////////////////////////////

func checkArgs(args []interface{}) bool {
	m := args[0]
	agent := args[1].(*ActorPlayer)
	if agent == nil || m == nil {
		return false
	}
	userData := agent.UserData()
	if userData == nil {
		clog.Warnf("[checkArgs] params:%+v  err:%s", m, StatusText[Login06])
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return false
	}
	person := userData.(*mgr.Player)
	if person == nil {
		clog.Warnf("[checkArgs] params:%+v  err:%s", m, StatusText[Login06])
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[Login06])
		return false
	}

	if play := mgr.GetPlayerMgr().Get(person.UserID); play == nil {
		clog.Warnf("[checkArgs] params:%+v  uid:%+v err:%s", m, person.UserID, StatusText[User03])
		agent.SendResultPop(FAILED, StatusText[Title004], StatusText[User03])
		return false
	}
	return true
}

////////////////////////////////////////////////////////////////////////////

//// playerSelect 玩家查询角色列表
//func (p *ActorPlayer) playerSelect(session *cproto.Session, _ *protoMsg.None) {
//	response := &protoMsg.PlayerSelectResponse{}
//
//	playerId := db.GetPlayerIdWithUID(session.Uid)
//	if playerId > 0 {
//		// 游戏设定单服单角色，协议设计成可返回多角色
//		playerTable, found := db.GetPlayerTable(playerId)
//		if found {
//			playerInfo := buildPBPlayer(playerTable)
//			response.List = append(response.List, &playerInfo)
//		}
//	}
//
//	p.Response(session, response)
//}
//
//// playerCreate 玩家创角
//func (p *ActorPlayer) playerCreate(session *cproto.Session, req *protoMsg.PlayerCreateRequest) {
//	if req.Gender > 1 {
//		p.ResponseCode(session, code.PlayerCreateFail)
//		return
//	}
//
//	// 检查玩家昵称
//	if len(req.PlayerName) < 1 {
//		p.ResponseCode(session, code.PlayerCreateFail)
//		return
//	}
//
//	// 帐号是否已经在当前游戏服存在角色
//	if db.GetPlayerIdWithUID(session.Uid) > 0 {
//		p.ResponseCode(session, code.PlayerCreateFail)
//		return
//	}
//
//	// 获取创角初始化配置
//	playerInitRow, found := data.PlayerInitConfig.Get(req.Gender)
//	if found == false {
//		p.ResponseCode(session, code.PlayerCreateFail)
//		return
//	}
//
//	// 创建角色&添加角色初始的资产
//	serverId := session.GetInt32(sessionKey.ServerID)
//	newPlayerTable, errCode := db.CreatePlayer(session, req.PlayerName, serverId, playerInitRow)
//	if code.IsFail(errCode) {
//		p.ResponseCode(session, errCode)
//		return
//	}
//
//	// TODO 更新最后一次登陆的角色信息到中心节点
//
//	// 抛出角色创建事件
//	playerCreateEvent := event.NewPlayerCreate(newPlayerTable.PlayerId, req.PlayerName, req.Gender)
//	p.PostEvent(&playerCreateEvent)
//
//	playerInfo := buildPBPlayer(newPlayerTable)
//	response := &pb.PlayerCreateResponse{
//		Player: &playerInfo,
//	}
//
//	p.Response(session, response)
//}
//
//// playerEnter 玩家进入游戏
//func (p *ActorPlayer) playerEnter(session *cproto.Session, req *pb.Int64) {
//	playerId := req.Value
//	if playerId < 1 {
//		p.ResponseCode(session, code.PlayerIdError)
//		return
//	}
//
//	// 检查并查找该用户下的该角色
//	playerTable, found := db.GetPlayerTable(req.GetValue())
//	if found == false {
//		p.ResponseCode(session, code.PlayerIdError)
//		return
//	}
//
//	// 保存进入游戏的玩家对应的agentPath
//	online.BindPlayer(playerId, playerTable.UID, session.AgentPath)
//
//	// 设置网关节点session的PlayerID属性
//	p.Call(session.ActorPath(), "setSession", &pb.StringKeyValue{
//		Key:   sessionKey.PlayerID,
//		Value: cstring.ToString(playerId),
//	})
//
//	p.uid = playerTable.UID
//	p.playerId = playerTable.PlayerId
//	p.isOnline = true // 设置为在线状态
//
//	// 这里改为客户端主动请求更佳
//	// [01]推送角色 道具数据
//	//module.Item.ListPush(session, playerId)
//	// [02]推送角色 英雄数据
//	//module.Hero.ListPush(session, playerId)
//	// [03]推送角色 成就数据
//	//module.Achieve.CheckNewAndPush(playerId, true, true)
//	// [04]推送角色 邮件数据
//	//module.Mail.ListPush(session, playerId)
//
//	// [99]最后推送 角色进入游戏响应结果
//	response := &pb.PlayerEnterResponse{}
//	response.GuideMaps = map[int32]int32{}
//
//	p.Response(session, response)
//
//	// 角色登录事件
//	loginEvent := event.NewPlayerLogin(p.ActorID(), playerId)
//	p.PostEvent(&loginEvent)
//}
//
//func buildPBPlayer(playerTable *db.PlayerTable) pb.Player {
//	return pb.Player{
//		PlayerId:   playerTable.PlayerId,
//		PlayerName: playerTable.Name,
//		Level:      playerTable.Level,
//		CreateTime: playerTable.CreateTime,
//		Exp:        playerTable.Exp,
//		Gender:     playerTable.Gender,
//	}
//}
