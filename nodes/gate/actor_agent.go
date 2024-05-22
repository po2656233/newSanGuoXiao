package gate

import (
	"github.com/po2656233/superplace/const/code"
	cstring "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	"github.com/po2656233/superplace/net/parser/pomelo"
	cproto "github.com/po2656233/superplace/net/proto"
	"sanguoxiao/internal/data"
	"sanguoxiao/internal/hints"
	pb2 "sanguoxiao/internal/protocol/gofile"
	"sanguoxiao/internal/rpc"
	"sanguoxiao/internal/session_key"
	"sanguoxiao/internal/token"
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
		Value: hints.Login12,
	})

	p.Local().Register(p.login)
	p.Remote().Register(p.setSession)
}

func (p *ActorAgent) setSession(req *pb2.StringKeyValue) {
	if req.Key == "" {
		return
	}

	if agent, ok := pomelo.GetAgent(p.ActorID()); ok {
		agent.Session().Set(req.Key, req.Value)
	}
}

// login 用户登录，验证帐号 (*pb.LoginResponse, int32)
func (p *ActorAgent) login(session *cproto.Session, req *pb2.LoginRequest) {
	agent, found := pomelo.GetAgent(p.ActorID())
	if !found {
		return
	}

	// 验证token
	userToken, errCode := p.validateToken(req.Token)
	if code.IsFail(errCode) {
		agent.Response(session, errCode)
		return
	}

	// 验证pid是否配置
	sdkRow := data.SdkConfig.Get(userToken.PID)
	if sdkRow == nil {
		agent.ResponseCode(session, hints.Login15, true)
		return
	}
	// 根据token带来的sdk参数，从中心节点获取uid
	resp, errCode := rpc.SendData(p.App(), rpc.SourcePath, rpc.AccountActor, rpc.CenterType, &pb2.GetUserIDReq{
		SdkId:  sdkRow.SdkId,
		Pid:    userToken.PID,
		OpenId: userToken.OpenID,
	})

	rsp := resp.(*pb2.GetUserIDResp)
	uid := rsp.Uid
	if uid == 0 || code.IsFail(errCode) {
		agent.ResponseCode(session, hints.Login07, true)
		return
	}

	p.checkGateSession(uid)

	if err := agent.Bind(uid); err != nil {
		clog.Warn(err)
		agent.ResponseCode(session, hints.Login07, true)
		return
	}

	agent.Session().Set(sessionKey.ServerID, cstring.ToString(req.ServerId))
	agent.Session().Set(sessionKey.PID, cstring.ToString(userToken.PID))
	agent.Session().Set(sessionKey.OpenID, userToken.OpenID)

	response := &pb2.LoginResponse{
		Uid:    uid,
		Pid:    userToken.PID,
		OpenId: userToken.OpenID,
	}

	agent.Response(session, response)
}

func (p *ActorAgent) validateToken(base64Token string) (*token.Token, int32) {
	userToken, ok := token.DecodeToken(base64Token)
	if ok == false {
		return nil, hints.Login15
	}

	platformRow := data.SdkConfig.Get(userToken.PID)
	if platformRow == nil {
		return nil, hints.Login15
	}

	statusCode, ok := token.Validate(userToken, platformRow.Salt)
	if ok == false {
		return nil, statusCode
	}

	return userToken, code.OK
}

func (p *ActorAgent) checkGateSession(uid cfacade.UID) {
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
		actorPath := cfacade.NewPath(member.GetNodeId(), "user")
		p.Call(actorPath, pomelo.KickFuncName, rsp)
	}
}

// onSessionClose  当agent断开时，关闭对应的ActorAgent
func (p *ActorAgent) onSessionClose(agent *pomelo.Agent) {
	session := agent.Session()
	serverId := session.GetString(sessionKey.ServerID)
	if serverId == "" {
		return
	}

	// 通知game节点关闭session
	childId := cstring.ToString(session.Uid)
	if childId != "" {
		targetPath := cfacade.NewChildPath(serverId, "player", childId)
		p.Call(targetPath, "sessionClose", nil)
	}

	// 自己退出
	p.Exit()
	clog.Infof("sessionClose path = %s", p.Path())
}
