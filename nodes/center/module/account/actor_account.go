package account

import (
	"github.com/po2656233/superplace/const/code"
	cactor "github.com/po2656233/superplace/net/actor"
	"strings"
	"superman/internal/hints"
	"superman/internal/protocol/gofile"
	"superman/internal/rpc"
	db2 "superman/nodes/center/db"
)

type (
	ActorAccount struct {
		cactor.Base
	}
)

func (p *ActorAccount) AliasID() string {
	return "account"
}

// OnInit center为后端节点，不直接与客户端通信，所以了一些remote函数，供RPC调用
func (p *ActorAccount) OnInit() {
	p.Remote().Register(p.RegisterReq)
	p.Remote().Register(p.LoginReq)
	p.Remote().Register(p.GetUserIDReq)
}

// RegisterReq 注册开发者帐号
func (p *ActorAccount) RegisterReq(req *pb.RegisterReq) (*pb.RegisterResp, int32) {
	accountName := req.Name
	password := req.Password

	if strings.TrimSpace(accountName) == "" || strings.TrimSpace(password) == "" {
		return nil, hints.Register02
	}

	if len(accountName) < 3 || len(accountName) > 18 {
		return nil, hints.Register02
	}

	if len(password) < 3 || len(password) > 18 {
		return nil, hints.Register02
	}

	db2.DevAccountRegister(accountName, password, req.Address)
	app := p.App()
	target := rpc.GetTargetPath(app, ".db", rpc.CenterType)
	resp := &pb.RegisterResp{}
	errCode := p.App().ActorSystem().CallWait(rpc.SourcePath, target, "Register", req, resp)
	_ = errCode
	return resp, code.OK
}

// LoginReq 根据帐号名获取开发者帐号表
func (p *ActorAccount) LoginReq(req *pb.LoginReq) (*pb.LoginResp, int32) {
	accountName := req.Account
	password := req.Password

	devAccount, _ := db2.DevAccountWithName(accountName)
	if devAccount == nil || devAccount.Password != password {
		return nil, hints.Login07
	}

	return &pb.LoginResp{
		MainInfo: &pb.MasterInfo{
			UserInfo: &pb.UserInfo{UserID: devAccount.AccountId},
		},
	}, code.OK
}

// GetUserIDReq 获取uid
func (p *ActorAccount) GetUserIDReq(req *pb.GetUserIDReq) (*pb.GetUserIDResp, int32) {
	uid, ok := db2.BindUID(req.SdkId, req.Pid, req.OpenId)
	if uid == 0 || ok == false {
		return nil, hints.Login07
	}

	return &pb.GetUserIDResp{Uid: uid}, code.OK
}
