package account

import "C"
import (
	"fmt"
	"github.com/google/uuid"
	superConst "github.com/po2656233/superplace/const"
	log "github.com/po2656233/superplace/logger"
	cactor "github.com/po2656233/superplace/net/actor"
	"strconv"
	"strings"
	. "superman/internal/constant"
	commMsg "superman/internal/protocol/go_file/common"
	gateMsg "superman/internal/protocol/go_file/gate"
	sqlmodel "superman/internal/sql_model/minigame"
	"superman/internal/utils"
	db2 "superman/nodes/center/db"
	"time"
)

type (
	ActorAccount struct {
		cactor.Base
		dbComponent *db2.Component
	}
)

func (p *ActorAccount) AliasID() string {
	return strings.Trim(AccActor, superConst.DOT)
}

// OnInit center为后端节点，不直接与客户端通信，所以了一些remote函数，供RPC调用
func (p *ActorAccount) OnInit() {
	p.Remote().Register(p.Register)
	p.Remote().Register(p.Login)
	p.Remote().Register(p.Logout)
	p.Remote().Register(p.GetUserID)
	p.Remote().Register(p.GetUserInfo)
	p.dbComponent = p.App().Find(CenterDb).(*db2.Component)
	if p.dbComponent == nil {
		log.Fatal("db component not found")
	}
}

// Register 注册开发者账号
func (p *ActorAccount) Register(req *gateMsg.RegisterReq) (*gateMsg.RegisterResp, int32) {
	accountName := req.Name
	password := req.Password

	if strings.TrimSpace(accountName) == Empty || strings.TrimSpace(password) == Empty {
		return nil, Register09
	}

	if len(accountName) < NameLenMin || len(accountName) > NameLenMax {
		return nil, Register07
	}

	if len(password) < NameLenMin || len(password) > NameLenMax {
		return nil, Register08
	}

	identity := uuid.New().String()
	user := sqlmodel.User{
		Name:     req.Name,
		Account:  req.Name,
		Password: utils.Md5Sum(req.Password + strconv.FormatInt(int64(len(req.Password)), 10)),
		Passport: req.PassPortID,
		Realname: req.RealName,
		Phone:    req.PhoneNum,
		Address:  req.Address,
		Email:    req.Email,
		Identity: identity,
		//Clientaddr:   clientAddr,
		Machinecode:  req.MachineCode,
		Referralcode: req.InvitationCode,
		//Serveraddr:   serverAddr,
		Face:   req.FaceID,
		Gender: req.Gender,
		Age:    req.Age,
	}
	log.Warnf("actorDB register req:%v", req)
	uid, err := p.dbComponent.AddUser(user)
	if err != nil {
		uid = -1
		log.Warnf("[actor_gorm] Register uid:%v err:%v", uid, err)
	}
	db2.DevAccountRegister(uid, accountName, password, req.Address)
	resp := &gateMsg.RegisterResp{
		OpenId: fmt.Sprintf("%d", uid),
	}
	return resp, SUCCESS
}

// Login 根据帐号名获取开发者账号表
func (p *ActorAccount) Login(req *gateMsg.LoginReq) (*gateMsg.LoginResp, int32) {
	psw := utils.Md5Sum(req.Password + strconv.FormatInt(int64(len(req.Password)), 10))
	userInfo, err := p.dbComponent.CheckUserInfo(req.Account, psw)
	if err != nil {
		log.Error("[ActorAccount] login CheckUserInfo req:%v error:%v", req, err)
		return nil, Login08
	}
	if err = p.dbComponent.UpdateLoginTime(userInfo.ID); err != nil {
		log.Error("[ActorAccount] login UpdateLoginTime req:%v error:%v", req, err)
		return nil, Login03
	}
	resp := &gateMsg.LoginResp{}
	resp.MainInfo = &commMsg.MasterInfo{
		UserInfo: &commMsg.UserInfo{
			UserID:  userInfo.ID,
			Name:    userInfo.Name,
			Account: userInfo.Account,
			//Password:     userInfo.Password,
			Money:   userInfo.Money,
			Coin:    userInfo.Coin,
			YuanBao: userInfo.Yuanbao,
			FaceID:  userInfo.Face,
			Gender:  userInfo.Gender,
			Age:     userInfo.Age,
			Vip:     userInfo.Vip,
			Level:   userInfo.Level,
			//PassPortID:   userInfo.Passport,
			//Address:      userInfo.Address,
			//AgentID:      userInfo.Agentid,
			//ServerAddr:   userInfo.Serveraddr,
			//MachineCode:  userInfo.Machinecode,
			//RealName:     userInfo.Realname,
			//PhoneNum:     userInfo.Phone,
			//Email:        userInfo.Email,
			//ReferralCode: userInfo.Referralcode,
			//IDentity:     userInfo.Identity,
			//ClientAddr:   userInfo.Clientaddr,
		},
	}

	accountName := req.Account
	password := req.Password
	devAccount, _ := db2.DevAccountWithName(accountName)
	if devAccount == nil {
		db2.DevAccountRegister(userInfo.ID, accountName, password, req.MachineCode)
	} else if devAccount.Password != password {
		return nil, Login07
	}
	return resp, SUCCESS
}

func (p *ActorAccount) Logout(req *gateMsg.LogoutReq) (*gateMsg.LogoutResp, int32) {
	now := time.Now().Unix()
	if err := p.dbComponent.UpdateLeaveTime(req.Uid, now); err != nil {
		log.Error("[ActorAccount] Logout UpdateLeaveTime req:%v error:%v", req, err)
		return nil, Login16
	}
	return &gateMsg.LogoutResp{
		Uid:       req.Uid,
		Timestamp: now,
	}, SUCCESS
}

// GetUserID 获取uid
func (p *ActorAccount) GetUserID(req *commMsg.GetUserIDReq) (*commMsg.GetUserIDResp, int32) {
	uid, ok := db2.BindUID(req.SdkId, req.Pid, req.OpenId)
	if uid == 0 || ok == false {
		return nil, Login07
	}

	return &commMsg.GetUserIDResp{Uid: uid}, SUCCESS
}

func (p *ActorAccount) FixNickName(req *gateMsg.FixNickNameReq) (*gateMsg.FixNickNameResp, int32) {
	resp := &gateMsg.FixNickNameResp{}
	err := p.dbComponent.UpdateNickName(req.Uid, req.Name)
	if err != nil {
		log.Error("[ActorAccount] FixNickName UpdateNickName req:%v error:%v", req, err)
		return resp, Setting20
	}
	resp.Uid = req.Uid
	resp.Name = req.Name
	return resp, SUCCESS
}

func (p *ActorAccount) GetUserInfo(req *gateMsg.GetUserInfoReq) (*gateMsg.GetUserInfoResp, error) {
	resp := &gateMsg.GetUserInfoResp{}
	info, err := p.dbComponent.CheckUserSimpInfo(req.Uid)
	if err != nil {
		return resp, err
	}
	resp.Info = &commMsg.UserInfo{
		UserID:   info.ID,
		Name:     info.Name,
		Account:  info.Account,
		FaceID:   info.Face,
		Gender:   info.Gender,
		Age:      info.Age,
		Empirice: info.Empiric,
		Vip:      info.Vip,
		Level:    info.Level,
		YuanBao:  info.Yuanbao,
		Coin:     info.Coin,
		Money:    info.Money,
	}
	return resp, nil
}
