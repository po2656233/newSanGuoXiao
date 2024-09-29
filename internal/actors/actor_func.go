package actors

import (
	"fmt"
	"github.com/google/uuid"
	clog "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	"strconv"
	commMsg "superman/internal/protocol/go_file/common"
	gameMsg "superman/internal/protocol/go_file/game"

	gateMsg "superman/internal/protocol/go_file/gate"
	"superman/internal/rpc"
	sqlmodel "superman/internal/sql_model/minigame"
	. "superman/internal/utils"
	"time"
)

// Register /////////////////////////////////////////////////////////
func (self *ActorDB) Register(req *gateMsg.RegisterReq) (*gateMsg.RegisterResp, error) {
	m := req
	identity := uuid.New().String()
	user := sqlmodel.User{
		Name:     m.Name,
		Account:  m.Name,
		Password: Md5Sum(m.Password + strconv.FormatInt(int64(len(m.Password)), 10)),
		Passport: m.PassPortID,
		Realname: m.RealName,
		Phone:    m.PhoneNum,
		Address:  m.Address,
		Email:    m.Email,
		Identity: identity,
		//Clientaddr:   clientAddr,
		Machinecode:  m.MachineCode,
		Referralcode: m.InvitationCode,
		//Serveraddr:   serverAddr,
		Face:   m.FaceID,
		Gender: m.Gender,
		Age:    m.Age,
	}
	clog.Warnf("actorDB register req:%v", req)
	uid, err := self.addUser(user)
	if err != nil {
		uid = -1
		clog.Errorf("[actor_gorm] Register uid:%v err:%v", uid, err)
	}
	resp := &gateMsg.RegisterResp{
		OpenId: fmt.Sprintf("%d", uid),
	}

	return resp, err
}

func (self *ActorDB) Login(req *gateMsg.LoginReq) (*gateMsg.LoginResp, error) {
	psw := Md5Sum(req.Password + strconv.FormatInt(int64(len(req.Password)), 10))
	userInfo, err := self.checkUserInfo(req.Account, psw)
	if err != nil {
		return nil, err
	}
	if err := self.updateLoginTime(userInfo.ID); err != nil {
		return nil, err
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
	return resp, err
}

func (self *ActorDB) Logout(req *gateMsg.LogoutReq) (*gateMsg.LogoutResp, error) {
	now := time.Now().Unix()
	if err := self.updateLeaveTime(req.Uid, now); err != nil {
		return nil, err
	}
	return &gateMsg.LogoutResp{
		Uid:       req.Uid,
		Timestamp: now,
	}, nil
}

// GetClassList 取分类列表
func (self *ActorDB) GetClassList() (*gateMsg.GetClassListResp, error) {
	resp := &gateMsg.GetClassListResp{
		Items: &commMsg.ClassList{
			Classify: make([]*commMsg.ClassItem, 0),
		},
	}
	ret, err := self.checkClassify()
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
func (self *ActorDB) GetRoomList(req *gateMsg.GetRoomListReq) (*gateMsg.GetRoomListResp, error) {
	resp := &gateMsg.GetRoomListResp{
		Items: &commMsg.RoomList{
			Items: make([]*commMsg.RoomInfo, 0),
		},
	}
	ret, err := self.checkRooms(req.Uid, 0, -1, 0)
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
func (self *ActorDB) GetTableList(req *gateMsg.GetTableListReq) (*gateMsg.GetTableListResp, error) {
	resp := &gateMsg.GetTableListResp{
		Items: &commMsg.TableList{
			Items: make([]*commMsg.TableInfo, 0),
		},
	}
	ret, err := self.checkTables(req.Rid, -1, 0)
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
func (self *ActorDB) GetGameList(req *gateMsg.GetGameListReq) (*gateMsg.GetGameListResp, error) {
	resp := &gateMsg.GetGameListResp{
		Items: &commMsg.GameList{
			Items: make([]*commMsg.GameInfo, 0),
		},
	}
	ret, err := self.checkGames(req.Kid, -1, 0)
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

func (self *ActorDB) CreateRoom(req *gateMsg.CreateRoomReq) (*gateMsg.CreateRoomResp, error) {
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
	req.RoomKey = Md5Sum(req.RoomKey)
	roomId, err := self.addRoom(sqlmodel.Room{
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
	rpc.SendDataToGame(self.App(), resp)
	return resp, nil
}

func (self *ActorDB) CreateTable(req *gateMsg.CreateTableReq) (*gateMsg.CreateTableResp, error) {
	resp := &gateMsg.CreateTableResp{}
	if req.Name == "" {
		// 则获取游戏名称
		req.Name, _ = self.checkGameName(req.Gid)
	}
	tid, maxSit, err := self.addTable(sqlmodel.Table{
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
	rpc.SendDataToGame(self.App(), resp)
	return resp, err
}

func (self *ActorDB) DeleteTable(req *gateMsg.DeleteTableReq) (*gateMsg.DeleteTableResp, error) {
	resp := &gateMsg.DeleteTableResp{}
	rid := self.checkTableRid(req.Tid)
	if rid == 0 {
		return resp, fmt.Errorf("tid:%d no have rid", req.Tid)
	}
	ok, err := self.checkRoomExist(req.HostId, rid)
	if !ok || err != nil {
		return resp, fmt.Errorf("hostid:%d rid:%d no exist err:%v", req.HostId, rid, err)
	}
	// 检测游戏是否处于关闭状态，才能
	err = self.delTable(req.HostId, req.Tid)
	resp.Tid = req.Tid
	resp.Rid = rid
	rpc.SendDataToGame(self.App(), resp)
	return resp, err
}

// /////////////////////////////////////////////////////////////////////////////////////

func (self *ActorDB) GetTable(req *gateMsg.GetTableReq) (*gateMsg.GetTableResp, error) {
	resp := &gateMsg.GetTableResp{}
	info, err := self.checkTable(req.Tid)
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
func (self *ActorDB) GetUserInfo(req *gateMsg.GetUserInfoReq) (*gateMsg.GetUserInfoResp, error) {
	resp := &gateMsg.GetUserInfoResp{}
	info, err := self.checkUserSimpInfo(req.Uid)
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

func (self *ActorDB) FixNickName(req *gateMsg.FixNickNameReq) (*gateMsg.FixNickNameResp, error) {
	resp := &gateMsg.FixNickNameResp{}
	err := self.updateNickName(req.Uid, req.Name)
	if err != nil {
		return resp, err
	}
	resp.Uid = req.Uid
	resp.Name = req.Name
	return resp, nil
}

func (self *ActorDB) Recharge(req *gateMsg.RechargeReq) (*gateMsg.RechargeResp, error) {
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
	err := self.addRecharge(info)
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

func (self *ActorDB) AddRecord(req *gameMsg.AddRecordReq) (*gameMsg.AddRecordResp, error) {
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
	err := self.addRecord(record)
	if err != nil {
		return resp, err
	}
	resp.PerGold = record.Pergold
	resp.Gold = record.Gold
	return resp, nil
}
func (self *ActorDB) DecreaseGameRun(req *gameMsg.DecreaseGameRunReq) (*gameMsg.DecreaseGameRunResp, error) {
	resp := &gameMsg.DecreaseGameRunResp{}
	remain, err := self.eraseRemain(req.Tid, req.Amount)
	resp.Tid = req.Tid
	resp.Remain = remain
	return resp, err
}
