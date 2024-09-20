package gate

import (
	"github.com/po2656233/superplace"
	sgxGops "github.com/po2656233/superplace/components/gops"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	cconnector "github.com/po2656233/superplace/net/connector"
	"github.com/po2656233/superplace/net/parser/pomelo"
	"github.com/po2656233/superplace/net/parser/simple"
	"strconv"
	"strings"
	"superman/internal/component/check_center"
	"superman/internal/conf"
	. "superman/internal/constant"
	"superman/internal/rpc"
	"time"
)

// !!! 注意 ---->节点名称包含pomelo则表示启动该协议，否则使用simple协议(即: mid+数据长度[不含mid和长度本身]+数据)

// Run 运行gate节点
// gate 主要用于对外提供网络连接、管理用户连接、消息转发
func Run(profileFilePath, nodeId string) {
	// 创建一个sgx实例
	app := superplace.Configure(
		profileFilePath,
		nodeId,
		true,
		superplace.Cluster,
	)

	// 设置网络数据包解析器
	if strings.Contains(nodeId, "pomelo") {
		netParser := buildPomeloParser(app)
		app.SetNetParser(netParser)
	} else {
		netParser := buildSimpleParser(app)
		app.SetNetParser(netParser)
	}

	app.Register(sgxGops.New())
	// 注册检则中心服组件，用于检则中心服是否先启动
	app.Register(checkCenter.New())
	// 注册数据配表组件，具体详见data-config的使用方法和参数配置
	app.Register(conf.New())

	//启动sgx引擎
	app.Startup()
}

func buildPomeloParser(app *superplace.AppBuilder) cfacade.INetParser {
	// 使用pomelo网络数据包解析器
	agentActor := pomelo.NewActor(ActIdGate)
	//创建一个tcp监听，用于client/robot压测机器人连接网关tcp
	agentActor.AddConnector(cconnector.NewTCP(TcpAddr))
	agentActor.AddConnector(cconnector.NewKCP(KcpAddr))
	//再创建一个websocket监听，用于h5客户端建立连接
	agentActor.AddConnector(cconnector.NewWS(app.Address()))
	//当有新连接创建Agent时，启动一个自定义(ActorAgent)的子actor
	agentActor.SetOnNewAgent(func(newAgent *pomelo.Agent) {
		childActor := &ActorAgent{}
		newAgent.AddOnClose(childActor.onPomeloSessionClose)
		_, _ = agentActor.Child().Create(newAgent.SID(), childActor) // actorID == sid
	})

	// 设置数据路由函数
	agentActor.SetOnDataRoute(onPomeloDataRoute)

	return agentActor
}

// 构建简单的网络数据包解析器
func buildSimpleParser(app *superplace.AppBuilder) cfacade.INetParser {
	agentActor := simple.NewActor(ActIdGate)
	agentActor.AddConnector(cconnector.NewTCP(TcpAddr))
	agentActor.AddConnector(cconnector.NewKCP(KcpAddr))
	agentActor.AddConnector(cconnector.NewWS(app.Address()))

	agentActor.SetOnNewAgent(func(newAgent *simple.Agent) {
		childActor := &ActorAgent{}
		newAgent.AddOnClose(childActor.onSimpleSessionClose)
		_, _ = agentActor.Child().Create(newAgent.SID(), childActor)
		clog.Infof("新增客户端 sid:%v uid:%v", newAgent.SID(), newAgent.UID())
	})

	// 设置大头&小头
	agentActor.SetEndian(rpc.Endian)
	// 设置心跳时间
	agentActor.SetHeartbeatTime(60 * time.Second)
	// 设置积压消息数量
	agentActor.SetWriteBacklog(64)

	// 设置数据路由函数 (leaf)
	agentActor.SetOnDataRoute(onSimpleDataRoute)

	// 设置消息节点路由(建议配合data-config组件进行使用)
	// mid = 1 的消息路由到  gate节点.user的Actor.login函数上
	//agentActor.AddNodeRoute(MIDGate, &simple.NodeRoute{
	//	NodeType: NodeTypeGate,
	//	ActorID:  ActIdGate,
	//	FuncName: FuncSimpLogin,
	//})
	//
	//agentActor.AddNodeRoute(MIDPing, &simple.NodeRoute{
	//	NodeType: NodeTypeGate,
	//	ActorID:  ActIdGate,
	//})
	//
	//agentActor.AddNodeRoute(MIDLeaf, &simple.NodeRoute{
	//	NodeType: NodeTypeLeaf,
	//	ActorID:  ActIdGame,
	//	FuncName: FuncRequest,
	//})
	//
	//agentActor.AddNodeRoute(MIDGame, &simple.NodeRoute{
	//	NodeType: NodeTypeGame,
	//	ActorID:  ActIdGame,
	//	FuncName: FuncRequest,
	//})
	//
	//agentActor.AddNodeRoute(MIDMatch, &simple.NodeRoute{
	//	NodeType: NodeTypeMatch,
	//	ActorID:  ActIdMatch,
	//	FuncName: FuncMatch,
	//})

	mapIDs := rpc.LoadMsgInfos()
	gateStart := int64(236)
	gateEnd := int64(244)
	registerProto := func(id int64, field, key string) bool {
		size := len(field)
		n := len(key)
		if size < n+1 || field[size-n:] != key {
			return false
		}
		nodeType := NodeTypeGame
		actorID := ActIdGame
		if gateStart <= id && id <= gateEnd {
			nodeType = NodeTypeGate
			actorID = ActIdGate
		}
		agentActor.AddNodeRoute(uint32(id), &simple.NodeRoute{
			NodeType: nodeType,
			ActorID:  actorID,
			FuncName: field[:size-n],
		})
		clog.Infof("路由消息注册--节点类型:%v\t 实例:%v\t 消息ID:%v 消息类型:%v", nodeType, actorID, id, field)

		return true
	}
	for id, field := range mapIDs {
		nID, _ := strconv.ParseInt(id, 10, 32)
		if !registerProto(nID, field, "Req") {
			registerProto(nID, field, "Request")
		}

	}
	return agentActor
}
