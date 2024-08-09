package actors

import (
	"fmt"
	"github.com/google/uuid"
	clog "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	"strconv"
	. "superman/internal/constant"
	pb "superman/internal/protocol/gofile"
	"superman/internal/rpc"
	sqlmodel "superman/internal/sql_model"
	. "superman/internal/utils"
	"time"
)

// Register /////////////////////////////////////////////////////////
func (self *ActorDB) Register(req *pb.RegisterReq) (*pb.RegisterResp, error) {
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
	resp := &pb.RegisterResp{
		OpenId: fmt.Sprintf("%d", uid),
	}

	return resp, err
}

func (self *ActorDB) Login(req *pb.LoginReq) (*pb.LoginResp, error) {
	psw := Md5Sum(req.Password + strconv.FormatInt(int64(len(req.Password)), 10))
	userInfo, err := self.checkUserInfo(req.Account, psw)
	if err != nil {
		return nil, err
	}
	resp := &pb.LoginResp{}
	resp.MainInfo = &pb.MasterInfo{
		UserInfo: &pb.UserInfo{
			UserID:       userInfo.ID,
			Name:         userInfo.Name,
			Account:      userInfo.Account,
			Password:     userInfo.Password,
			Money:        userInfo.Money,
			Coin:         userInfo.Coin,
			YuanBao:      userInfo.Yuanbao,
			FaceID:       userInfo.Face,
			Gender:       userInfo.Gender,
			Age:          userInfo.Age,
			Vip:          userInfo.Vip,
			Level:        userInfo.Level,
			PassPortID:   userInfo.Passport,
			Address:      userInfo.Address,
			AgentID:      userInfo.Agentid,
			ServerAddr:   userInfo.Serveraddr,
			MachineCode:  userInfo.Machinecode,
			RealName:     userInfo.Realname,
			PhoneNum:     userInfo.Phone,
			Email:        userInfo.Email,
			ReferralCode: userInfo.Referralcode,
			IDentity:     userInfo.Identity,
			ClientAddr:   userInfo.Clientaddr,
		},
	}
	return resp, err
}

// GetClassList 取分类列表
func (self *ActorDB) GetClassList() (*pb.GetClassListResp, error) {
	resp := &pb.GetClassListResp{
		Items: &pb.ClassList{
			Classify: make([]*pb.ClassItem, 0),
		},
	}
	ret, err := self.checkClassify()
	for _, kind := range ret {
		resp.Items.Classify = append(resp.Items.Classify, &pb.ClassItem{
			Id:     kind.ID,
			Name:   kind.Name,
			EnName: kind.EnName,
		})
	}
	return resp, err
}

// GetRoomList 取分类列表
func (self *ActorDB) GetRoomList(req *pb.GetRoomListReq) (*pb.GetRoomListResp, error) {
	resp := &pb.GetRoomListResp{
		Items: &pb.RoomList{
			Items: make([]*pb.RoomInfo, 0),
		},
	}
	ret, err := self.checkRooms(req.Uid, req.StartTime, -1, 0)
	for _, room := range ret {
		resp.Items.Items = append(resp.Items.Items, &pb.RoomInfo{
			Id:         room.ID,
			HostId:     room.Hostid,
			Level:      pb.RoomLevel(room.Level),
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
func (self *ActorDB) GetTableList(req *pb.GetTableListReq) (*pb.GetTableListResp, error) {
	resp := &pb.GetTableListResp{
		Items: &pb.TableList{
			Items: make([]*pb.TableInfo, 0),
		},
	}
	ret, err := self.checkTables(req.Rid, -1, 0)
	for _, table := range ret {
		resp.Items.Items = append(resp.Items.Items, &pb.TableInfo{
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
func (self *ActorDB) GetGameList(req *pb.GetGameListReq) (*pb.GetGameListResp, error) {
	resp := &pb.GetGameListResp{
		Items: &pb.GameList{
			Items: make([]*pb.GameInfo, 0),
		},
	}
	ret, err := self.checkGames(req.Kid, -1, 0)
	for _, game := range ret {
		resp.Items.Items = append(resp.Items.Items, &pb.GameInfo{
			Id:        game.ID,
			Name:      game.Name,
			Kid:       game.Kid,
			Lessscore: game.Lessscore,
			State:     pb.GameState(game.State),
			MaxPlayer: game.MaxPlayer,
			HowToPlay: game.HowToPlay,
		})
	}
	return resp, err
}

///////////////////////create////////////////////////////////////

func (self *ActorDB) CreateRoom(req *pb.CreateRoomReq) (*pb.CreateRoomResp, error) {
	resp := &pb.CreateRoomResp{}
	maxCount := int32(0)
	maxTable := int32(0)
	switch req.Level {
	case pb.RoomLevel_GeneralRoom:
		maxCount = 20
		maxTable = 5
	case pb.RoomLevel_MiddleRoom:
		maxCount = 50
		maxTable = 10
	case pb.RoomLevel_HighRoom:
		maxCount = 200
		maxTable = 20
	case pb.RoomLevel_TopRoom:
		maxCount = 1000
		maxTable = 50
	case pb.RoomLevel_SuperRoom:
		maxCount = 10000
		maxTable = 100
	case pb.RoomLevel_SystemRoom:
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
	resp.Info = &pb.RoomInfo{
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

func (self *ActorDB) CreateTable(req *pb.CreateTableReq) (*pb.CreateTableResp, error) {
	resp := &pb.CreateTableResp{}
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
	resp.Table = &pb.TableInfo{
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

func (self *ActorDB) DeleteTable(req *pb.DeleteTableReq) (*pb.DeleteTableResp, error) {
	resp := &pb.DeleteTableResp{}
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

func (self *ActorDB) GetTable(req *pb.GetTableReq) (*pb.GetTableResp, error) {
	resp := &pb.GetTableResp{}
	info, err := self.checkTable(req.Tid)
	if err != nil {
		return nil, err
	}
	resp.Info = &pb.TableInfo{
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
func (self *ActorDB) GetUserInfo(req *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	resp := &pb.GetUserInfoResp{}
	info, err := self.checkUserSimpInfo(req.Uid)
	if err != nil {
		return resp, err
	}
	resp.Info = &pb.UserInfo{
		UserID:  info.ID,
		Name:    info.Name,
		Account: info.Account,
		FaceID:  info.Face,
		Gender:  info.Gender,
		Age:     info.Age,
		Empiric: info.Empiric,
		Vip:     info.Vip,
		Level:   info.Level,
		YuanBao: info.Yuanbao,
		Coin:    info.Coin,
		Money:   info.Money,
	}
	return resp, nil
}

func (self *ActorDB) FixNickName(req *pb.FixNickNameReq) (*pb.FixNickNameResp, error) {
	resp := &pb.FixNickNameResp{}
	err := self.updateNickName(req.Uid, req.Name)
	if err != nil {
		return resp, err
	}
	resp.Uid = req.Uid
	resp.Name = req.Name
	return resp, nil
}

func (self *ActorDB) Recharge(req *pb.RechargeReq) (*pb.RechargeResp, error) {
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
	return &pb.RechargeResp{
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

func (self *ActorDB) AddRecord(req *pb.AddRecordReq) (*pb.AddRecordResp, error) {
	resp := &pb.AddRecordResp{}
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
func (self *ActorDB) DecreaseGameRun(req *pb.DecreaseGameRunReq) (*pb.DecreaseGameRunResp, error) {
	resp := &pb.DecreaseGameRunResp{}
	remain, err := self.eraseRemain(req.Tid, req.Amount)
	resp.Tid = req.Tid
	resp.Remain = remain
	return resp, err
}
