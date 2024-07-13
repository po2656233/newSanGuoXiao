package actors

import (
	"fmt"
	"github.com/google/uuid"
	clog "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	"strconv"
	pb "superman/internal/protocol/gofile"
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
		//Vip:          LessVIP,
		//Level:        LessLevel,
		//Agentid: m.PlatformID,
		//Agentid:      agentID,
		//Withdraw:     INVALID,
		//Deposit:      LessMoney,
		//Money:        LessMoney,
	}
	clog.Warnf("actorDB register req:%v", req)
	uid, err := self.AddUser(user)
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
	userInfo, err := self.GetUserInfo(req.Account, psw)
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
			VIP:          userInfo.Vip,
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
	ret, err := self.GetClassify()
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
	ret, err := self.GetRooms(req.Uid)
	for _, room := range ret {
		resp.Items.Items = append(resp.Items.Items, &pb.RoomInfo{
			Id:          room.ID,
			HostId:      room.Hostid,
			Level:       room.Level,
			Name:        room.Name,
			RoomKey:     room.Roomkey,
			EnterScore:  room.Enterscore,
			OnlineCount: room.OnlineCount,
			RobotCount:  room.RobotCount,
			MaxCount:    room.MaxCount,
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
	ret, err := self.GetTables(req.Rid)
	for _, table := range ret {
		resp.Items.Items = append(resp.Items.Items, &pb.TableInfo{
			Id:          table.ID,
			Num:         table.Num,
			Rid:         table.Rid,
			Gid:         table.Gid,
			OpenTime:    table.Opentime,
			Taxation:    table.Taxation,
			Commission:  table.Commission,
			Amount:      table.Amount,
			PlayScore:   table.Playscore,
			SitterCount: table.SitterCount,
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
	ret, err := self.GetGames(req.Kid)
	for _, game := range ret {
		resp.Items.Items = append(resp.Items.Items, &pb.GameInfo{
			Id:        game.ID,
			Name:      game.Name,
			Kid:       game.Kid,
			Lessscore: game.Lessscore,
			State:     pb.GameState(game.State),
			MaxCount:  game.Maxcount,
			HowToPlay: game.HowToPlay,
		})
	}
	return resp, err
}

///////////////////////create////////////////////////////////////

func (self *ActorDB) CreateRoom(req *pb.CreateRoomReq) (interface{}, error) {
	resp := &pb.CreateRoomResp{}
	roomId, err := self.AddRoom(sqlmodel.Room{
		Hostid:      req.HostId,
		Level:       req.Level,
		Name:        req.Name,
		Roomkey:     req.RoomKey,
		Enterscore:  req.EnterScore,
		OnlineCount: req.OnlineCount,
		RobotCount:  req.RobotCount,
		MaxCount:    req.MaxCount,
		Remark:      req.Remark,
		CreatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
		UpdateBy:    0,
		CreateBy:    req.HostId,
	})
	if 0 == roomId {
		return nil, err
	}
	resp.HostId = req.HostId
	resp.RoomId = roomId
	resp.RoomKey = req.RoomKey
	resp.Name = req.Name
	return resp, nil
}

func (self *ActorDB) CreateTable(req *pb.CreateTableReq) (*pb.CreateTableResp, error) {
	resp := &pb.CreateTableResp{}
	tid, num, err := self.AddTable(sqlmodel.Table{
		Gid:        req.Gid,
		Rid:        req.Rid,
		Name:       req.Name,
		Playscore:  req.Playscore,
		Taxation:   req.Taxation,
		Commission: req.Commission,
	})
	if err != nil {
		return resp, err
	}
	resp.Table = &pb.TableInfo{
		Id:         tid,
		Num:        num,
		Gid:        req.Gid,
		Name:       req.Name,
		Commission: req.Commission,
		Taxation:   req.Taxation,
		PlayScore:  req.Playscore,
	}
	return resp, err
}
