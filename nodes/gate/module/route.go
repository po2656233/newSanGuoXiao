package module

import (
	cslice "github.com/po2656233/superplace/extend/slice"
	cstring "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/pomelo"
	pmessage "github.com/po2656233/superplace/net/parser/pomelo/message"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	. "superman/internal/constant"
	"superman/internal/protocol/go_file/common"
	gateMsg "superman/internal/protocol/go_file/gate"
	"superman/internal/rpc"
)

var (
	// 客户端连接后，必需先执行第一条协议，进行token验证后，才能进行后续的逻辑
	firstRouteName = Join(NodeTypeGate, ActIdGate, FuncLogin)
	// 角色进入游戏时的前三个协议
	beforeLoginRoutes = []string{
		//"game.player.select",  //查询玩家角色
		//"game.player.create",  //玩家创建角色
		//"game.player.enter",   //玩家角色进入游戏
		//Join(NodeTypeLeaf, ActIdGame, FuncEnter),
		//"leaf.game.enter", //玩家角色进入游戏
	}

	notLoginRsp = &pb.Int32{
		Value: Login07,
	}
)

// ///////////////////////////////pomelo////////////////////////////////////////////////////
// onDataRoute 数据路由规则

// OnPomeloDataRoute
// 登录逻辑:
// 1.(建立连接)客户端建立连接，服务端对应创建一个agent用于处理玩家消息,actorID == sid
// 2.(用户登录)客户端进行帐号登录验证，通过uid绑定当前sid
// 3.(角色登录)客户端通过'beforeLoginRoutes'中的协议完成角色登录
func OnPomeloDataRoute(agent *pomelo.Agent, route *pmessage.Route, msg *pmessage.Message) {
	session := pomelo.BuildSession(agent, msg)

	// agent没有"用户登录",且请求不是第一条协议，则踢掉agent，断开连接
	if !session.IsBind() && msg.Route != firstRouteName {
		agent.Kick(notLoginRsp, true)
		return
	}

	if agent.NodeType() == route.NodeType() {
		targetPath := cfacade.NewChildPath(agent.NodeId(), route.HandleName(), session.Sid)
		pomelo.LocalDataRoute(agent, session, route, msg, targetPath)
	} else {
		gameNodeRoute(agent, session, route, msg)
	}
}

// gameNodeRoute 实现agent路由消息到游戏节点
func gameNodeRoute(agent *pomelo.Agent, session *cproto.Session, route *pmessage.Route, msg *pmessage.Message) {
	if !session.IsBind() {
		return
	}

	// 如果agent没有完成"角色登录",则禁止转发到game节点
	if !session.Contains(PlayerID) {
		// 如果不是角色登录协议则踢掉agent
		if found := cslice.StringInSlice(msg.Route, beforeLoginRoutes); !found {
			agent.Kick(notLoginRsp, true)
			return
		}
	}

	serverId := session.GetString(ServerID)
	if serverId == "" {
		return
	}

	childId := cstring.ToString(session.Uid)
	targetPath := cfacade.NewChildPath(serverId, route.HandleName(), childId)

	clusterPacket := cproto.GetClusterPacket()
	clusterPacket.SourcePath = session.AgentPath
	clusterPacket.TargetPath = targetPath
	clusterPacket.FuncName = route.Method()
	clusterPacket.Session = session   // agent session
	clusterPacket.ArgBytes = msg.Data // packet -> message -> data

	if err := agent.Cluster().PublishLocal(serverId, clusterPacket); err != nil {
		clog.Errorf("[route][gameNodeRoute] ClusterLocalDataRoute err:%v", err)
	}

}

// ///////////////////////////////simple////////////////////////////////////////////////////

func OnSimpleDataRoute(agent *simple.Agent, msg *simple.Message, route *simple.NodeRoute) {
	session := agent.Session()
	session.Mid = msg.MID

	if session.Mid == MIDPing {
		data, _ := rpc.GetProtoData(&gateMsg.PongResp{})
		agent.SendRaw(data)
		return
	}
	// current node
	if agent.NodeType() == route.NodeType {
		targetPath := cfacade.NewChildPath(agent.NodeId(), route.ActorID, session.Sid)
		simple.LocalDataRoute(agent, session, msg, route, targetPath)
		return
	}

	serverId := Empty
	if !session.IsBind() {
		clog.Warnf("[sid = %s,uid = %d] Session is not bind with UID. failed to forward message.[route = %+v]",
			agent.SID(),
			agent.UID(),
			route,
		)
		//return
	} else {
		serverId = session.GetString(ServerID)
	}

	// 聊天服,走随机节点
	if serverId == Empty || route.NodeType == NodeTypeChat {
		member, found := agent.Discovery().Random(route.NodeType)
		if !found {
			return
		}
		serverId = member.GetNodeId()
		session.Set(ServerID, serverId)
		session.Set(OpenID, cstring.ToString(session.Uid))
		//targetPath = cfacade.NewPath(member.GetNodeId(), route.ActorID)
	}
	childId := cstring.ToString(session.Uid)
	targetPath := cfacade.NewChildPath(serverId, route.ActorID, childId)
	clusterPacket := cproto.GetClusterPacket()
	clusterPacket.SourcePath = session.AgentPath
	clusterPacket.TargetPath = targetPath
	clusterPacket.FuncName = route.FuncName
	clusterPacket.Session = session   // agent session
	clusterPacket.ArgBytes = msg.Data // packet -> message -> data
	if clusterPacket.FuncName == "" {
		clog.Errorf("[route][gameNodeRoute] There is no function name in the parameter. ")
		return
		//clusterPacket.FuncName = FuncRequest
	}
	if err := agent.Cluster().PublishLocal(serverId, clusterPacket); err != nil {
		clog.Errorf("[route][gameNodeRoute] ClusterLocalDataRoute err:%v", err)
		agent.SendMsg(&gateMsg.ResultResp{
			State: NetworkErr02,
			Hints: StatusText[NetworkErr02],
		})
	}
}
