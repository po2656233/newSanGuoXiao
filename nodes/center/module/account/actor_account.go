package account

import "C"
import (
	superConst "github.com/po2656233/superplace/const"
	cactor "github.com/po2656233/superplace/net/actor"
	"strings"
	. "superman/internal/constant"
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
	return strings.Trim(AccActor, superConst.DOT)
}

// OnInit center为后端节点，不直接与客户端通信，所以了一些remote函数，供RPC调用
func (p *ActorAccount) OnInit() {
	p.Remote().Register(p.Register)
	p.Remote().Register(p.Login)
	p.Remote().Register(p.GetUserID)
}

// Register 注册开发者帐号
func (p *ActorAccount) Register(req *pb.RegisterReq) (*pb.RegisterResp, int32) {
	accountName := req.Name
	password := req.Password

	if strings.TrimSpace(accountName) == "" || strings.TrimSpace(password) == "" {
		return nil, Register02
	}

	if len(accountName) < 3 || len(accountName) > 18 {
		return nil, Register02
	}

	if len(password) < 3 || len(password) > 18 {
		return nil, Register02
	}

	//db2.DevAccountRegister(accountName, password, req.Address)
	data, errCode := rpc.SendDataToDB(p.App(), req)
	resp, _ := data.(*pb.RegisterResp)
	return resp, errCode
}

// Login 根据帐号名获取开发者帐号表
func (p *ActorAccount) Login(req *pb.LoginReq) (*pb.LoginResp, int32) {
	data, errCode := rpc.SendDataToDB(p.App(), req)
	resp, _ := data.(*pb.LoginResp)
	return resp, errCode

	//accountName := req.Account
	//password := req.Password
	//devAccount, _ := db2.DevAccountWithName(accountName)
	//if devAccount == nil {
	//	//db2.DevAccountRegister(accountName, resp.MainInfo.UserInfo.Password,resp.MainInfo.UserInfo.ClientAddr)
	//} else if devAccount.Password != password {
	//	return nil, Login07
	//}
	//return &pb.LoginResp{
	//}, code.OK
}

// GetUserID 获取uid
func (p *ActorAccount) GetUserID(req *pb.GetUserIDReq) (*pb.GetUserIDResp, int32) {
	uid, ok := db2.BindUID(req.SdkId, req.Pid, req.OpenId)
	if uid == 0 || ok == false {
		return nil, Login07
	}

	return &pb.GetUserIDResp{Uid: uid}, SUCCESS
}
