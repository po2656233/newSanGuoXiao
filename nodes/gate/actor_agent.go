package gate

import (
	"github.com/po2656233/superplace/const/code"
	cstring "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	"github.com/po2656233/superplace/net/parser/pomelo"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	"superman/internal/conf"
	. "superman/internal/constant"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	"superman/internal/token"
)

var (
	duplicateLoginCode []byte
)

type (
	// ActorAgent 每个网络连接对应一个ActorAgent
	ActorAgent struct {
		cactor.Base
	}
)

func (p *ActorAgent) OnInit() {
	duplicateLoginCode, _ = p.App().Serializer().Marshal(&cproto.I32{
		Value: Login12,
	})
	// pomelo
	p.Remote().Register(p.setSession)
	p.Local().Register(p.login)

	// simple
	p.Remote().Register(p.setSimpleSession)
	p.Local().Register(p.Login)
}

////////////////////////pomelo//////////////////////////////////////

func (p *ActorAgent) setSession(req *pb.StringKeyValue) {
	if req.Key == "" {
		return
	}

	if agent, ok := pomelo.GetAgent(p.ActorID()); ok {
		agent.Session().Set(req.Key, req.Value)
	}
}

// login 用户登录，验证帐号 (*pb.LoginResponse, int32)
func (p *ActorAgent) login(session *cproto.Session, req *pb.LoginRequest) {
	agent, found := pomelo.GetAgent(p.ActorID())
	if !found {
		return
	}

	// 验证token
	userToken, errCode := token.ValidateBase64(req.Token)
	if code.IsFail(errCode) {
		agent.Response(session, errCode)
		return
	}

	// 验证pid是否配置
	sdkRow := conf.SdkConfig.Get(userToken.PID)
	if sdkRow == nil {
		agent.ResponseCode(session, Login15, true)
		return
	}
	// 根据token带来的sdk参数，从中心节点获取uid
	resp, errCode := rpc.SendDataToAcc(p.App(), &pb.GetUserIDReq{
		SdkId:  sdkRow.SdkId,
		Pid:    userToken.PID,
		OpenId: userToken.OpenID,
	})

	rsp := resp.(*pb.GetUserIDResp)
	uid := rsp.Uid
	if uid == 0 || code.IsFail(errCode) {
		agent.ResponseCode(session, Login07, true)
		return
	}

	p.checkSession(uid)

	if err := agent.Bind(uid); err != nil {
		clog.Warn(err)
		agent.ResponseCode(session, Login07, true)
		return
	}

	agent.Session().Set(ServerID, cstring.ToString(req.ServerId))
	agent.Session().Set(PID, cstring.ToString(userToken.PID))
	agent.Session().Set(OpenID, userToken.OpenID)
	agent.Response(session, &pb.LoginResponse{
		Uid:    uid,
		Pid:    userToken.PID,
		OpenId: userToken.OpenID,
	})
}

func (p *ActorAgent) checkSession(uid cfacade.UID) {
	if agent, found := pomelo.GetAgentWithUID(uid); found {
		agent.Kick(duplicateLoginCode, true)
	}

	rsp := &cproto.PomeloKick{
		Uid:    uid,
		Reason: duplicateLoginCode,
	}

	// 遍历所有网关节点，踢除旧的session
	members := p.App().Discovery().ListByType(p.App().NodeType(), p.App().NodeId())
	for _, member := range members {
		// user是gate.go里自定义的agentActorID
		actorPath := cfacade.NewPath(member.GetNodeId(), ActIdGate)
		p.Call(actorPath, pomelo.KickFuncName, rsp)
	}
}

// onSessionClose  当agent断开时，关闭对应的ActorAgent
func (p *ActorAgent) onPomeloSessionClose(agent *pomelo.Agent) {
	session := agent.Session()
	serverId := session.GetString(ServerID)
	if serverId == "" {
		return
	}

	// 通知game节点关闭session
	childId := cstring.ToString(session.Uid)
	if childId != "" {
		targetPath := cfacade.NewChildPath(serverId, ActIdGame, childId)
		p.Call(targetPath, "sessionClose", nil)
	}

	// 自己退出
	p.Exit()
	clog.Infof("sessionClose path = %s", p.Path())
}

// ///////////////////////////simple/////////////////////////////////////////////
func (p *ActorAgent) setSimpleSession(req *pb.StringKeyValue) {
	if req.Key == "" {
		return
	}

	if agent, ok := simple.GetAgent(p.ActorID()); ok {
		agent.Session().Set(req.Key, req.Value)
	}
}

// Login 登录
func (p *ActorAgent) Login(session *cproto.Session, req *pb.LoginRequest) {
	agent, found := simple.GetAgent(p.ActorID())
	if !found {
		return
	}

	// 验证token
	uid, _ := token.GetIDForToken(req.Token)
	if uid == 0 {
		agent.Response(session.Mid, &pb.ResultResp{
			State: Login07,
			Hints: StatusText[Login07],
		})
		return
	}

	p.checkSimpleSession(uid)

	if err := agent.Bind(uid); err != nil {
		clog.Warn(err)
		agent.Response(session.Mid, &pb.ResultResp{
			State: Flag0003,
			Hints: StatusText[Flag0003],
		})
		return
	}

	agent.Session().Set(ServerID, cstring.ToString(req.ServerId))
	agent.Session().Set(OpenID, cstring.ToString(uid))
	res := &pb.LoginResponse{
		Uid:    uid,
		OpenId: cstring.ToString(uid),
	}
	data, err := rpc.GetProtoData(res)
	agent.SendRaw(data)
	clog.Infof("Login mid:%v  err:%v", data[:4], err)
}

func (p *ActorAgent) checkSimpleSession(uid cfacade.UID) {
	if agent, found := simple.GetAgentWithUID(uid); found {
		session := agent.Session()
		agent.Response(session.Mid, &pb.ResultResp{
			State: Login12,
			Hints: StatusText[Login12],
		})
	}

	rsp := &cproto.PomeloKick{
		Uid:    uid,
		Reason: duplicateLoginCode,
	}

	// 遍历所有网关节点，踢除旧的session
	members := p.App().Discovery().ListByType(p.App().NodeType(), p.App().NodeId())
	for _, member := range members {
		// user是gate.go里自定义的agentActorID
		actorPath := cfacade.NewPath(member.GetNodeId(), ActIdGate)
		p.Call(actorPath, pomelo.KickFuncName, rsp)
	}
}

// onSessionClose  当agent断开时，关闭对应的ActorAgent
func (p *ActorAgent) onSimpleSessionClose(agent *simple.Agent) {
	session := agent.Session()
	serverId := session.GetString(ServerID)
	if serverId == "" {
		return
	}

	// 通知game节点关闭session
	childId := cstring.ToString(session.Uid)
	if childId != "" {
		targetPath := cfacade.NewChildPath(serverId, ActIdGame, childId)
		p.Call(targetPath, FuncSessionClose, nil)
	}

	// 自己退出
	p.Exit()
	clog.Infof("sessionClose path = %s", p.Path())
}
