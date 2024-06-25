package gate

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	superConst "github.com/po2656233/superplace/const"
	exReflect "github.com/po2656233/superplace/extend/reflect"
	cslice "github.com/po2656233/superplace/extend/slice"
	cstring "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/pomelo"
	pmessage "github.com/po2656233/superplace/net/parser/pomelo/message"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	"os"
	. "sanguoxiao/internal/constant"
	"sanguoxiao/internal/hints"
	"sanguoxiao/internal/protocol/gofile"
	pb2 "sanguoxiao/internal/protocol/gofile"
	"sanguoxiao/internal/session_key"
	"strconv"
)

var (
	// 客户端连接后，必需先执行第一条协议，进行token验证后，才能进行后续的逻辑
	firstRouteName = Join(NodeTypeGate, ActorGate, FuncLogin)
	// 角色进入游戏时的前三个协议
	beforeLoginRoutes = []string{
		//"game.player.select",  //查询玩家角色
		//"game.player.create",  //玩家创建角色
		//"game.player.enter",   //玩家角色进入游戏
		Join(NodeTypeLeaf, ActorGame, FuncEnter),
		//"leaf.game.enter", //玩家角色进入游戏
	}

	notLoginRsp = &pb.Int32{
		Value: hints.Login07,
	}
	mapFuncs = make(map[string]string)
	endian   = binary.BigEndian
)

func GetMsgFunc(fileName string) map[string]string {
	data, _ := os.ReadFile(fileName)
	//创建map，用于接收解码好的数据
	//创建文件的解码器
	_ = json.Unmarshal(data, &mapFuncs)
	//解码文件中的数据，丢入dataMap所在的内存
	//fmt.Println("解码成功", mapFuncs)
	return mapFuncs
}

func GetProtoData(msg proto.Message) ([]byte, error) {
	strId := ""
	name := exReflect.GetStructName(msg)
	for id, protoName := range mapFuncs {
		if protoName == name {
			strId = id
			break
		}
	}
	if strId == "" {
		return nil, errors.New("no data")
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, err
	}
	pkg := bytes.NewBuffer([]byte{})
	err = binary.Write(pkg, endian, uint32(id))
	if err != nil {
		return nil, err
	}

	data, err := proto.Marshal(msg)
	data = append(pkg.Bytes(), data...)
	return data, nil
}

func onSimpleDataRoute(agent *simple.Agent, msg *simple.Message, route *simple.NodeRoute) {
	session := agent.Session()
	session.Mid = msg.MID

	// 所有置空函数
	//if route.FuncName == "" {
	//	route.FuncName = exReflect.GetStructName(msg)
	//	//bytesReader := bytes.NewReader(msg.Data[:4])
	//	//var msgId uint32
	//	//err := binary.Read(bytesReader, endian, &msgId)
	//	//if err != nil {
	//	//	clog.Warnf("[sid = %s,uid = %d] Session is not msgId err:%v.[route = %+v]", agent.SID(), agent.UID(), err, route)
	//	//	return
	//	//}
	//	//strId := cstring.ToString(msgId)
	//	//if funcName, ok := mapFuncs[strId]; ok {
	//	//	route.FuncName = funcName
	//	//}
	//}
	if session.Mid == MIDPing {
		data, _ := GetProtoData(&pb2.PongResp{})
		agent.Response(MIDPing, data)
		return
	}
	// current node
	if agent.NodeType() == route.NodeType {
		targetPath := cfacade.NewChildPath(agent.NodeId(), route.ActorID, session.Sid)
		simple.LocalDataRoute(agent, session, msg, route, targetPath)
		return
	}

	if !session.IsBind() {
		clog.Warnf("[sid = %s,uid = %d] Session is not bind with UID. failed to forward message.[route = %+v]",
			agent.SID(),
			agent.UID(),
			route,
		)
		return
	}

	member, found := agent.Discovery().Random(route.NodeType)
	if !found {
		clog.Warnf("[sid = %s,uid = %d] discovery not found.[route = %+v]",
			agent.SID(),
			agent.UID(),
			route,
		)
		return
	}

	targetPath := cfacade.NewPath(member.GetNodeId(), route.ActorID)
	if route.FuncName != "" {
		targetPath += superConst.DOT + route.FuncName
	}
	if err := simple.ClusterLocalDataRoute(agent, session, msg, route, member.GetNodeId(), targetPath); err != nil {
		clog.Errorf("[sid = %s,uid = %d] Session post message err:%v.[route = %+v]",
			agent.SID(),
			agent.UID(),
			err,
			route,
		)
	}
}

// onDataRoute 数据路由规则
//
// 登录逻辑:
// 1.(建立连接)客户端建立连接，服务端对应创建一个agent用于处理玩家消息,actorID == sid
// 2.(用户登录)客户端进行帐号登录验证，通过uid绑定当前sid
// 3.(角色登录)客户端通过'beforeLoginRoutes'中的协议完成角色登录
func onPomeloDataRoute(agent *pomelo.Agent, route *pmessage.Route, msg *pmessage.Message) {
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
	if !session.Contains(sessionKey.PlayerID) {
		// 如果不是角色登录协议则踢掉agent
		if found := cslice.StringInSlice(msg.Route, beforeLoginRoutes); !found {
			agent.Kick(notLoginRsp, true)
			return
		}
	}

	serverId := session.GetString(sessionKey.ServerID)
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
