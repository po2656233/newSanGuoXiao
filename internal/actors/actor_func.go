package actors

import (
	"fmt"
	"github.com/google/uuid"
	clog "github.com/po2656233/superplace/logger"
	"strconv"
	pb "superman/internal/protocol/gofile"
	sqlmodel "superman/internal/sql_model"
	. "superman/internal/utils"
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
	if userInfo != nil {
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
	}

	return resp, err
}

// ClassList 取分类列表
func (self *ActorDB) ClassList() (*pb.GetClassListResp, error) {
	resp := &pb.GetClassListResp{
		Items: &pb.ClassList{
			Classify: make([]*pb.ClassItem, 0),
		},
	}
	ret, err := self.GetClassify()
	if ret != nil {
		for _, kind := range ret {
			resp.Items.Classify = append(resp.Items.Classify, &pb.ClassItem{
				Id:     kind.ID,
				Name:   kind.Name,
				EnName: kind.EnName,
			})
		}
	}
	return resp, err
}

// RoomList 取分类列表
func (self *ActorDB) RoomList(req *pb.GetRoomListReq) (*pb.GetRoomListResp, error) {
	resp := &pb.GetRoomListResp{
		Items: &pb.RoomList{
			Items: make([]*pb.RoomInfo, 0),
		},
	}
	ret, err := self.GetRooms(req.Uid)
	if ret != nil {
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
	}
	return resp, err
}
