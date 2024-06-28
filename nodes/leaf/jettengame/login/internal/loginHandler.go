package internal

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/base"
	. "superman/nodes/leaf/jettengame/manger"
	"superman/nodes/leaf/jettengame/sql/model"
	"superman/nodes/leaf/jettengame/sql/mysql"
	"superman/nodes/leaf/jettengame/sql/redis"
	"time"

	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
	uuid "github.com/satori/go.uuid"
	"superman/nodes/leaf/jettengame/gamedata/goclib/util"
)

var sqlHandle = mysql.SqlHandle()    //数据库管理
var personManger = GetPlayerManger() //玩家管理
var GlobalSender = GetClientManger()
var redisClient = redis.RedisHandle()

func init() {

	// 向当前模块（login 模块）注册 Login 消息的消息处理函数 handleTest
	handleMsg(&protoMsg.RegisterReq{}, handleRegister)   //反馈--->用户信息
	handleMsg(&protoMsg.LoginReq{}, handleLogin)         //反馈--->主页信息
	handleMsg(&protoMsg.ReconnectReq{}, handleReconnect) //反馈--->主页信息

	handleMsg(&protoMsg.ChooseClassReq{}, handleChooseClass)   //反馈--->游戏分类列表
	handleMsg(&protoMsg.ChooseGameReq{}, handleChooseGame)     //反馈--->牌桌列表
	handleMsg(&protoMsg.SettingTableReq{}, handleSettingTable) //反馈--->牌桌信息
	handleMsg(&protoMsg.UpdateMoneyReq{}, handleUpdateMoney)   //反馈--->玩家余额

	handleMsg(&protoMsg.CheckInReq{}, handleCheckIn)       //签到
	handleMsg(&protoMsg.GetCheckInReq{}, handleGetCheckIn) //所有签到记录
	handleMsg(&protoMsg.EmailReq{}, handleEmail)           //获取邮件
	handleMsg(&protoMsg.EmailReadReq{}, handleEmailRead)   //删除邮件
	handleMsg(&protoMsg.EmailDelReq{}, handleEmailDel)     //删除邮件
	handleMsg(&protoMsg.ClaimReq{}, handleClaim)           //领取奖励
	handleMsg(&protoMsg.BarterReq{}, handleBarter)         //置换房卡

	handleMsg(&protoMsg.GetGoodsReq{}, handleGetGoods)           //查看商品
	handleMsg(&protoMsg.GetAllGoodsReq{}, handleGetAllGoods)     //获取所有商品信息
	handleMsg(&protoMsg.BuyGoodsReq{}, handleBuyGoods)           //购买房卡
	handleMsg(&protoMsg.RechargeReq{}, handleRecharge)           //钱包充值
	handleMsg(&protoMsg.GetRechargesReq{}, handleGetRecharges)   //获取充值记录
	handleMsg(&protoMsg.CheckKnapsackReq{}, handleCheckKnapsack) //查看玩家背包

	//心跳包
	handleMsg(&protoMsg.PingReq{}, handlePing) //反馈--->玩家余额
	//handleMsg(&jsonMsg.UserLogin{}, handleLoginJson)
}

// 注册模块间的通信
func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// -----------------消息处理-----------------
// 注册
func handleRegister(args []interface{}) {
	m := args[0].(*protoMsg.RegisterReq)
	a := args[1].(gate.Agent)
	log.Debug("msg: %v psw:%v", m.GetName(), m.GetPassword())
	if m.Name == "" || m.Password == "" {
		a.WriteMsg(&protoMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Register02],
		})
		return
	}

	clientAddr := a.RemoteAddr().String()
	//serverAddr从平台中获取
	serverAddr := ""
	list := GetPlatformManger().Get(m.PlatformID).ServerList
	if 0 < len(list) {
		serverAddr = list[0]
	}

	// AgentID应当由解析邀请码所得
	infos := strings.Split(m.InvitationCode, "_")
	agentID := int64(0)
	if 2 < len(infos) {
		if uid, err := strconv.ParseInt(infos[0], 10, 64); err == nil {
			agentID = uid
		}
	}

	//由平台生成唯一标识
	identity := uuid.NewV4().String()
	rand.Seed(time.Now().Unix())
	strconv.FormatInt(int64(len(m.Password)), 10)
	//数据库新增用户信息
	user := model.User{
		Name:         m.Name,
		Account:      m.Name,
		Password:     util.Md5Sum(m.Password + strconv.FormatInt(int64(len(m.Password)), 10)),
		Passport:     m.PassPortID,
		Realname:     m.RealName,
		Phone:        m.PhoneNum,
		Address:      m.Address,
		Email:        m.Email,
		Identity:     identity,
		Clientaddr:   clientAddr,
		Machinecode:  m.MachineCode,
		Referralcode: m.InvitationCode,
		Serveraddr:   serverAddr,
		Face:         int32(m.FaceID),
		Gender:       int32(m.Gender),
		Age:          int32(m.Age),
		Vip:          LessVIP,
		Level:        LessLevel,
		Platformid:   int32(m.PlatformID),
		Agentid:      agentID,
		Withdraw:     INVALID,
		Deposit:      LessMoney,
		Money:        LessMoney,
	}
	if uid, err := sqlHandle.AddUser(user); err == nil {
		msg := &protoMsg.RegisterResp{
			//Info: sqlHandle.CheckUserInfo(uid),
		}
		a.WriteMsg(msg)

		//注册奖励
		if award, goods := sqlHandle.CheckEmail(IndexStart); award != nil {
			_, err = sqlHandle.AddEmail(INVALID, uid, award.Cc, award.Topic, award.Content, goods, "用户注册")
			CheckError(err)
			a.WriteMsg(award)
		}
	} else {
		log.Release("注册失败 uid:%v  err:%v !", uid, err.Error())
		a.WriteMsg(&protoMsg.ResultResp{
			State: FAILED,
			Hints: err.Error(),
		})
	}

}
func login(account, password, address string) (int64, *protoMsg.MasterInfo, int) {
	//数据库校验玩家数据
	if uid, ok := sqlHandle.CheckLogin(account, password); ok {
		pid := sqlHandle.CheckPlatformInfo(uid)
		if plat := GetPlatformManger().Get(pid); plat != nil {

			log.Debug("[Login] 游戏类别列表:%v!", plat.ClassIDs)
			//踢掉之前登录的账号
			loginTime := sqlHandle.CheckLoginTime(uid)
			nowTime := time.Now().Unix()
			// 不允许异地登录,超1小时 可以被解封//超过24小时 可以被解封
			if preAgent, ok := GlobalSender.Get(uid); ok && nowTime-loginTime < Hour {
				preAddr := preAgent.RemoteAddr().String()
				log.Release("[Login] 失败!!!!! curAddr:%v preAddr:%v", address, preAddr)
				if strings.Split(address, ":")[0] != strings.Split(preAddr, ":")[0] { // 不允许异地登录
					GlobalSender.SendData(preAgent, &protoMsg.AllopatricResp{
						UserID: uid,
					})
					return pid, nil, Login12
				}
				GlobalSender.SendPopResult(preAgent, FAILED, StatusText[Login05], StatusText[Login10])
				GlobalSender.DeleteClient(uid)
				preAgent.SetUserData(nil)
				preAgent.Close()
			}
			//限制频繁登录
			if nowTime-loginTime < 10 {
				return pid, nil, Login13
			}

			//房间列表
			mainInfo := &protoMsg.MasterInfo{}
			mainInfo.UserInfo = sqlHandle.CheckUserInfo(uid)
			plat.ClassIDs = sqlHandle.CheckRooms(pid)
			mainInfo.Classes = sqlHandle.CheckClassList(uid, plat.ClassIDs)
			sqlHandle.UpdateLoginTime([]int64{uid})
			// 默认签到
			if _, err := sqlHandle.CheckIn(uid, StatusText[Login14]); err != nil {
				log.Release("[Login] 默认签到 失败!!!!! err:%v", err)
			}
			return pid, mainInfo, SUCCESS
		} else {
			return pid, nil, Login07
		}
	}

	return 0, nil, Login08
}

func doLogin(platID int64, account, password string, agent gate.Agent) int {
	//检测用户是否存在
	password = password + strconv.FormatInt(int64(len(password)), 10)
	pid, msg, code := login(account, util.Md5Sum(password), agent.RemoteAddr().String())
	if code != SUCCESS {
		return code
	}

	person := &Player{
		PlayerInfo: &protoMsg.PlayerInfo{},
	}
	person.UserID = msg.UserInfo.UserID
	person.Name = msg.UserInfo.Name
	person.Age = msg.UserInfo.Age
	person.Sex = msg.UserInfo.Gender
	person.Account = msg.UserInfo.Account
	person.Money = msg.UserInfo.Money
	person.State = INVALID
	person.PlatformID = pid
	person.RoomNum = INVALID
	person.GameID = INVALID
	person.ChairID = INVALID
	person.PtrTable = nil
	person.PtrRoom = &Room{
		Num:     INVALID,
		Kind:    INVALID,
		Level:   INVALID,
		PageNum: INVALID,
	}

	// 获取token
	token, ok1 := person.GetToken()
	if !ok1 {
		return Login11
	}

	// 保存玩家数据
	person, _ = personManger.Append(person)
	// 往agent里添加数据
	if agent.UserData() == nil {
		agent.SetUserData(person)
		GetClientManger().Append(person.UserID, agent)
	}

	//反馈登录信息
	sqlHandle.UpdateLoginTime([]int64{person.UserID})
	GetClientManger().SendData(agent, &protoMsg.LoginResp{
		MainInfo: msg,
		InGameID: person.GameID,
		Token:    token,
	})
	GetClientManger().SendResult(agent, SUCCESS, StatusText[Login01])

	//获取平台公告
	hallNotices := sqlHandle.CheckNotice(platID)
	for _, notice := range hallNotices {
		notice.UserID = person.UserID
		GetClientManger().SendData(agent, notice)
	}

	//公告通知[一般是某玩家赢取金币公告]
	if notice := redisClient.Get(KeyNormalNotify); notice != nil {
		dataS, err1 := notice.Bytes()
		if err1 != nil {
			log.Release("[Notice]  user:[%v] platId:[%v] But notice Bytes was err:%v ", account, platID, err1)
			return SUCCESS
		}
		noticeData, err2 := GetClientManger().Unmarshal(dataS)
		if err2 != nil {
			log.Release("[Notice]  user:[%v] platId:[%v]  But notice Unmarshal was err:%v ", account, platID, err2)
			return SUCCESS
		}
		if noticeR, ok1 := noticeData.(*protoMsg.NotifyNoticeResp); ok1 {
			GetClientManger().SendData(agent, noticeR)
			log.Release("[Notice]  %v[platId:%v]", account, platID)
		}
	}
	redisClient.Set(GetAddressKey(agent.RemoteAddr().String()), fmt.Sprintf("%v|%v", platID, account), 0)

	log.Release("[Login] 登录成功! %v[platId:%v]", account, platID)
	return SUCCESS

}

// 重连
func handleReconnect(args []interface{}) {
	m := args[0].(*protoMsg.ReconnectReq)
	log.Debug("[reconnect]LoginInfo:->%v", m)

	//
	agent := args[1].(gate.Agent)
	if code := doLogin(m.PlatformID, m.Account, m.Password, agent); code != SUCCESS {
		agent.WriteMsg(&protoMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[code],
		})
		//日志打印
		log.Error("Reconnect Failed! m:%v", m)
	}

}

// 登录(房间列表)
func handleLogin(args []interface{}) {
	m := args[0].(*protoMsg.LoginReq)
	log.Debug("[receive]LoginInfo:->%v", m)

	//
	agent := args[1].(gate.Agent)
	if code := doLogin(m.PlatformID, m.Account, m.Password, agent); code != SUCCESS {
		agent.WriteMsg(&protoMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[code],
		})
		//日志打印
		log.Error("Reconnect Failed! m:%v", m)
	}

}

// 选择分类--->反馈游戏列表
func handleChooseClass(args []interface{}) {
	m := args[0].(*protoMsg.ChooseClassReq)
	agent := args[1].(gate.Agent)
	if INVALID == m.ID {
		GetClientManger().SendResult(agent, FAILED, StatusText[ClassInfo01])
		return
	}

	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		person.RoomNum = int64(m.ID)
		person.PtrRoom.Num = person.RoomNum

		_, key, games := sqlHandle.CheckGameList(m.ID)
		if key != m.TableKey {
			GetClientManger().SendResult(agent, FAILED, StatusText[ClassInfo01])
			return
		}
		data, _ := PBToBytes(person.ToMsg())
		ret, err1 := redisClient.Set(GetToken(person.UserID), data, Date*time.Second).Result()
		log.Debug("%v login %v %v", person.UserID, ret, err1)

		msg := &protoMsg.ChooseClassResp{
			ID:      m.ID,
			Games:   &protoMsg.GameList{},
			PageNum: INVALID,
		}
		if games != nil && MaxPage < len(games.Items) {
			//
			msg.PageNum = IndexStart
			items := make([]*protoMsg.GameItem, 0)
			for i := 0; i < len(games.Items); i++ {
				if 0 == i%MaxPage {
					if 0 < len(items) {
						msg.Games.Items = items
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					items = make([]*protoMsg.GameItem, 0)
				}
				items = CopyInsert(items, len(items), games.Items[i]).([]*protoMsg.GameItem)
			}
			if 0 < len(items) {
				msg.Games.Items = items
				GetClientManger().SendData(agent, msg)
			} else {
				GetClientManger().SendResult(agent, FAILED, StatusText[ClassInfo03])
			}
		} else {
			msg.Games = games
			GetClientManger().SendData(agent, msg)
		}
		return
	}
	agent.WriteMsg(&protoMsg.ResultResp{
		State: FAILED,
		Hints: StatusText[Login06],
	})

}

// 选择游戏-->反馈牌桌列表
func handleChooseGame(args []interface{}) {
	m := args[0].(*protoMsg.ChooseGameReq)
	agent := args[1].(gate.Agent)
	//data, _ := PBToBytes(person.ToMsg())
	//ret, err1 := redisClient.Set(token, data, Date*time.Second).Result()

	if userData := agent.UserData(); userData == nil {
		agent.WriteMsg(&protoMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Login06],
		})
		return
	}

	person := agent.UserData().(*Player)
	person.PtrRoom.Kind = int32(m.Info.KindID)
	person.PtrRoom.Level = m.Info.Level
	person.PtrRoom.PageNum = m.PageNum
	msg := &protoMsg.ChooseGameResp{
		Info:    m.Info,
		Tables:  &protoMsg.TableList{},
		PageNum: IndexStart,
	}
	// 检测当前玩家是否在游戏当中 [废弃]不再主动拉进房间，由客户端决定
	//player := GetPlayerManger().Get(person.UserID)
	//if player.GameID != INVALID {
	//game, ok := GetGamesManger().GetGame(player.GameID)
	//if ok {
	//	if game.G.KindID != m.Info.KindID || game.G.Level != m.Info.Level {
	//		return
	//	}
	//	var enterArgs []interface{}
	//	enterArgs = append(enterArgs, game, agent)
	//	person.Enter(enterArgs)
	//	return
	//}
	//}

	tables := GetGamesManger().GetTables(m.Info.KindID, m.Info.Level)
	if tables == nil {
		//GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo01])
		return
	}
	//			log.Debug("[牌桌列表]\t%v", len(tables.Items))
	if tables == nil || len(tables.Items) > MaxPage {
		msg.Tables = tables
		GetClientManger().SendData(agent, msg)
		return
	}

	// 游戏分页
	items := make([]*protoMsg.TableItem, 0)
	msg.PageNum = IndexStart
	for i := 0; i < len(tables.Items); i++ {
		if 0 == i%MaxPage {
			if 0 < len(items) {
				msg.Tables.Items = items
				if m.PageNum == INVALID {
					GetClientManger().SendData(agent, msg)
				}
				if m.PageNum == msg.PageNum {
					break
				}
				msg.PageNum++

			}
			items = make([]*protoMsg.TableItem, 0)
		}
		items = CopyInsert(items, len(items), tables.Items[i]).([]*protoMsg.TableItem)
	}

	if 0 == len(items) {
		GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo07])
		return
	}
	msg.Tables.Items = items
	GetClientManger().SendData(agent, msg)
}

// 配置牌桌
func handleSettingTable(args []interface{}) {
	m := args[0].(*protoMsg.SettingTableReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if "" == m.TInfo.Name {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting12])
			return
		}
		if strings.Contains(m.TInfo.Name, " ") {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting13])
			return
		}
		if NameLen < len(m.TInfo.Name) {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting19])
			return
		}

		if m.TInfo.PlayScore <= INVALID {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting14])
			return
		}

		if m.TInfo.PlayScore <= int64(m.TInfo.LessScore) {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting15])
			return
		}
		if m.TInfo.Amount <= INVALID {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting17])
			return
		}
		if MaxRounds < m.TInfo.Amount {
			GetClientManger().SendResult(agent, FAILED, StatusText[Setting18])
			return
		}
		//1、该游戏是否开售
		allGoods := sqlHandle.CheckAllGoods()
		if len(allGoods) <= 0 {
			GetClientManger().SendPopResult(agent, FAILED, StatusText[Title001], StatusText[Setting16])
			return
		}
		goodsId := int64(INVALID)
		for id, goods := range allGoods {
			if goods.Kind == m.GInfo.KindID && goods.Level == m.GInfo.Level {
				goodsId = id
				break
			}
		}
		if goodsId == INVALID {
			GetClientManger().SendPopResult(agent, FAILED, StatusText[Title001], StatusText[Setting16])
			return
		}

		//2、查看玩家资产里是否有该类型房卡 且数量需满足 ||( goods.ID == goodsid && INVALID < goods.Store && m.TInfo.Amount <= int32(goods.Store))
		asset := sqlHandle.CheckAssets(person.UserID)
		normal := int64(INVALID)
		fixGoods := &protoMsg.GoodsInfo{}
		generalGoods := &protoMsg.GoodsInfo{}
		for _, goods := range asset.MyGoods {
			if goods.ID == goodsId {
				fixGoods = goods
				normal += goods.Store
			}
		}
		//通用房卡
		countGeneral := int64(INVALID)
		if normal == INVALID || normal < int64(m.TInfo.Amount) {
			for _, goods := range asset.MyGoods {
				if goods.ID == GeneralRoom {
					generalGoods = goods
					countGeneral += goods.Store
				}
			}
			if normal+countGeneral < int64(m.TInfo.Amount) {
				GetClientManger().SendData(agent, &protoMsg.ToShoppingResp{
					ID:     goodsId,
					Count:  int32(m.TInfo.Amount),
					Reason: StatusText[User23],
				})
				//GetClientManger().SendPopResult(agent, FAILED, StatusText[Title001], StatusText[User23])
				return
			}
			countGeneral = int64(m.TInfo.Amount) - normal
		} else {
			normal = int64(m.TInfo.Amount)
		}

		//游戏的最小人数
		getMinCount := func(kid int32) int32 {
			switch kid {
			case Zhaocaimiao:
				return MaxChairMore
			}
			return MaxChairMore
		}
		m.TInfo.MaxChair = getMinCount(m.GInfo.KindID)

		//3、先从数据库创建
		gid, err := sqlHandle.AddGame(m.GInfo, m.TInfo)
		if err != nil {
			log.Error("玩家%v 配置游戏:%v", person.UserID, err.Error())
			GetClientManger().SendPopResult(agent, Fault, StatusText[Title001], StatusText[Mysql12])
			return
		}

		//4、正式创建
		if g, ok := GetGamesManger().CreateGame(gid, m.GInfo, m.TInfo); ok { //[1
			//5、扣除玩家的房卡
			if err1 := sqlHandle.UpdateAsset(person.UserID, goodsId, INVALID, -int32(normal), CodePayRoomCard, "消费房卡"); err1 != nil {
				err = sqlHandle.DelGame(gid, INVALID)
				CheckError(err)

				log.Error("玩家%v 配置游戏:%v 消费房卡:%v", person.UserID, err1.Error(), normal)
				GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo02])
				return
			} else {
				goodMsg := &protoMsg.BuyGoodsResp{
					UserID: person.UserID,
					Info:   fixGoods,
				}
				goodMsg.Info.Amount = fixGoods.Amount - int32(normal)
				goodMsg.Info.Sold = normal
				GetClientManger().SendData(agent, goodMsg)
			}

			if countGeneral != INVALID {
				if err1 := sqlHandle.UpdateAsset(person.UserID, GeneralRoom, INVALID, -int32(countGeneral), CodePayRoomCard, "消费通用房卡"); err1 != nil {
					err = sqlHandle.DelGame(gid, INVALID)
					CheckError(err)

					log.Error("玩家%v 配置游戏:%v %v", person.UserID, countGeneral, normal)
					GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo02])
					return
				} else {
					goodMsg := &protoMsg.BuyGoodsResp{
						UserID: person.UserID,
						Info:   generalGoods,
					}
					goodMsg.Info.Amount = generalGoods.Amount - int32(countGeneral)
					goodMsg.Info.Sold = countGeneral
					GetClientManger().SendData(agent, goodMsg)
				}
			}

			//6、反馈配置信息
			num := GetGamesManger().GetTableNum(g.ID, g.G.KindID, g.G.Level)
			msg := &protoMsg.SettingTableResp{
				Item: &protoMsg.TableItem{
					Num:    num,
					GameID: g.ID,
					Info:   m.TInfo,
				},
			}
			msg.Item.Info = &protoMsg.TableInfo{
				HostID:     m.TInfo.HostID,
				Name:       m.TInfo.Name,
				Password:   "",
				State:      m.TInfo.State,
				EnterScore: m.TInfo.EnterScore,
				LessScore:  m.TInfo.LessScore,
				PlayScore:  m.TInfo.PlayScore,
				Commission: m.TInfo.Commission,
				MaxChair:   m.TInfo.MaxChair,
				Amount:     m.TInfo.Amount,
				MaxOnline:  m.TInfo.MaxOnline,
			}
			GetClientManger().SendData(agent, msg)

			//7、记录玩家房间状态
			person.PtrRoom.Kind = int32(m.GInfo.KindID)
			person.PtrRoom.Level = m.GInfo.Level
			log.Debug("成功添加桌子%v:%v", g.ID, m.TInfo)

			//8、更新页码,并通知同页码下的玩家
			person.PtrRoom.PageNum = GetGamesManger().UpdateTableList(m.GInfo, person.UserID, INVALID)
		} else {
			log.Error("玩家%v 配置游戏失败", person.UserID)
			err = sqlHandle.DelGame(gid, INVALID)
			CheckError(err)
			GetClientManger().SendResult(agent, FAILED, StatusText[TableInfo02])
		} //0]
		return
	}
	agent.WriteMsg(&protoMsg.ResultResp{
		State: FAILED,
		Hints: StatusText[Login06],
	})
}

// 更新金钱
func handleUpdateMoney(args []interface{}) {
	_ = args[0].(*protoMsg.UpdateMoneyReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		person.Money = sqlHandle.CheckMoney(person.UserID)
		msg := &protoMsg.UpdateMoneyResp{
			UserID: person.UserID,
			Money:  person.Money,
		}
		GetClientManger().SendData(agent, msg)
	}
}

// 签到
func handleCheckIn(args []interface{}) {
	m := args[0].(*protoMsg.CheckInReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		msg := &protoMsg.CheckInResp{
			UserID: person.UserID,
			Remark: m.Remark,
		}
		if timestamp, err := sqlHandle.CheckIn(person.UserID, m.Remark); err == nil {
			msg.Timestamp = timestamp
			list := &protoMsg.GoodsList{}
			list.AllGoods = make([]*protoMsg.GoodsInfo, 0)
			info := &protoMsg.GoodsInfo{
				ID:     1001,
				Kind:   1,
				Name:   StatusText[Room14],
				Amount: 10,
			}
			list.AllGoods = append(list.AllGoods, info)
			msg.AwardList = list
			GetClientManger().SendData(agent, msg)
			//发送奖励

		} else if err.Error() == StatusText[Flag0001] {
			GetClientManger().SendResult(agent, FAILED, StatusText[User28])
		}
		return
	}
	GetClientManger().SendError(agent)
}

// 查看签到记录
func handleGetCheckIn(args []interface{}) {
	_ = args[0].(*protoMsg.GetCheckInReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)

		msg := &protoMsg.GetCheckInResp{
			UserID:     person.UserID,
			AllCheckin: make([]*protoMsg.CheckInResp, 0),
			PageNum:    IndexStart,
		}
		msg.AllCheckin = sqlHandle.CheckGetCheckIn(person.UserID)

		sort.Slice(msg.AllCheckin, func(i, j int) bool {
			return msg.AllCheckin[i].Timestamp < msg.AllCheckin[j].Timestamp
		})
		//分页发送
		//仅记录最后一页数据
		if MaxPage < len(msg.AllCheckin) {
			items := make([]*protoMsg.CheckInResp, 0)
			msg.PageNum = IndexStart
			for i := 0; i < len(msg.AllCheckin); i++ {
				if 0 == i%MaxPage {
					if 0 < len(items) {
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					items = make([]*protoMsg.CheckInResp, 0)
				}
				items = CopyInsert(items, len(items), msg.AllCheckin[i]).([]*protoMsg.CheckInResp)

			}
			if 0 < len(items) {
				msg.AllCheckin = items
			}
			GetClientManger().SendData(agent, msg)
		} else {
			GetClientManger().SendData(agent, msg)
		}
	}
}

// 查看邮件
func handleEmail(args []interface{}) {
	_ = args[0].(*protoMsg.EmailReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		msg := &protoMsg.EmailResp{
			UserID:  person.UserID,
			PageNum: IndexStart,
		}
		msg.Infos = sqlHandle.CheckEmails(person.UserID)
		if msg.Infos == nil || 0 == len(msg.Infos) {
			GetClientManger().SendResult(agent, FAILED, StatusText[User27])
			return
		}
		sort.Slice(msg.Infos, func(i, j int) bool {
			return msg.Infos[i].TimeStamp < msg.Infos[j].TimeStamp
		})
		//分页发送
		//仅记录最后一页数据
		if MaxPage < len(msg.Infos) {
			items := make([]*protoMsg.EmailInfo, 0)
			msg.PageNum = IndexStart
			for i := 0; i < len(msg.Infos); i++ {
				if 0 == i%MaxPage {
					if 0 < len(items) {
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					items = make([]*protoMsg.EmailInfo, 0)
				}
				items = CopyInsert(items, len(items), msg.Infos[i]).([]*protoMsg.EmailInfo)

			}
			if 0 < len(items) {
				msg.Infos = items
			}
			GetClientManger().SendData(agent, msg)
		} else {
			GetClientManger().SendData(agent, msg)
		}
		return
	}
	GetClientManger().SendResult(agent, FAILED, StatusText[User27])
}

// 查看邮件
func handleEmailRead(args []interface{}) {
	m := args[0].(*protoMsg.EmailReadReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if err := sqlHandle.ReadEmail(m.EmailID); err != nil {
			GetClientManger().SendResult(agent, FAILED, StatusText[Mysql15])
			return
		}
		msg := &protoMsg.EmailReadResp{
			UserID:  person.UserID,
			EmailID: m.EmailID,
		}
		GetClientManger().SendData(agent, msg)
	}
}

// 查看邮件
func handleEmailDel(args []interface{}) {
	m := args[0].(*protoMsg.EmailDelReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		if err := sqlHandle.DelEmail(m.EmailID); err != nil {
			GetClientManger().SendResult(agent, FAILED, StatusText[Mysql15])
			return
		}
		msg := &protoMsg.EmailDelResp{
			UserID:  person.UserID,
			EmailID: m.EmailID,
		}
		GetClientManger().SendData(agent, msg)
	}
}

// 补偿
func handleClaim(args []interface{}) {
	m := args[0].(*protoMsg.ClaimReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)

		if info, _ := sqlHandle.CheckEmail(m.EmailID); info != nil && info.AwardList != nil {
			for _, goods := range info.AwardList.AllGoods {
				if err1 := sqlHandle.UpdateAsset(person.UserID, goods.ID, INVALID, goods.Amount, CodeAward, "奖励房卡"); err1 != nil {
					GetClientManger().SendResult(agent, FAILED, StatusText[Mysql13])
					continue
				}
			}
		} else {
			GetClientManger().SendResult(agent, FAILED, StatusText[Mysql13])
			return
		}

		if err := sqlHandle.ClaimEmail(m.EmailID); err != nil {
			GetClientManger().SendResult(agent, FAILED, StatusText[Mysql13])
			return
		}
		msg := &protoMsg.ClaimResp{
			UserID:  person.UserID,
			EmailID: m.EmailID,
		}
		GetClientManger().SendData(agent, msg)
		person.Knapsack = sqlHandle.CheckAssets(person.UserID)

		//更新包裹
		msgX := &protoMsg.CheckKnapsackResp{
			UserID: person.UserID,
			Info:   person.Knapsack,
		}
		GetClientManger().SendData(agent, msgX)

	}
}

// 置换房卡
func handleBarter(args []interface{}) {
	m := args[0].(*protoMsg.BarterReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		allGoods := sqlHandle.CheckAllGoods()
		if _, ok := allGoods[m.ID]; ok {
			//置换房卡
			count := int32(INVALID)
			if m.ID == SuperRoom {
				count = m.Amount * TenTwice
			} else if m.ID == GeneralRoom {
				count = m.Amount
			}
			timestamp := time.Now().Unix()
			order := strconv.FormatInt(timestamp, 10)
			order = util.Md5Sum(order)
			if err := sqlHandle.UpdateAsset(person.UserID, m.ID, INVALID, -m.Amount, CodePayRoomCard, order); err != nil {
				log.Debug("玩家%v 消耗房卡失败:%v", person.UserID, m)
				GetClientManger().SendResult(agent, FAILED, StatusText[Mysql14])
				return
			}
			if err := sqlHandle.UpdateAsset(person.UserID, m.ToID, INVALID, count, CodeBarterCard, order); err != nil {
				log.Debug("玩家%v 置换房卡失败:%v", person.UserID, m)
				GetClientManger().SendResult(agent, FAILED, StatusText[Mysql14])
				return
			}
			person.Knapsack = sqlHandle.CheckAssets(person.UserID)
			msg := &protoMsg.BarterResp{
				UserID: person.UserID,
				Info:   person.Knapsack,
			}
			GetClientManger().SendData(agent, msg)
		}
	}
}

// 查看某商品
func handleGetGoods(args []interface{}) {
	m := args[0].(*protoMsg.GetGoodsReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		allGoods := sqlHandle.CheckAllGoods()
		msg := &protoMsg.GetGoodsResp{
			UserID: person.UserID,
			Info:   allGoods[m.ID],
		}
		GetClientManger().SendData(agent, msg)
	}
}

// 查看所有商品
func handleGetAllGoods(args []interface{}) {
	_ = args[0].(*protoMsg.GetAllGoodsReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		allGoods := sqlHandle.CheckAllGoods()
		msg := &protoMsg.GetAllGoodsResp{
			UserID:  person.UserID,
			Info:    make([]*protoMsg.GoodsInfo, 0),
			PageNum: INVALID,
		}
		for _, item := range allGoods {
			msg.Info = append(msg.Info, item)
		}

		sort.Slice(msg.Info, func(i, j int) bool {
			return msg.Info[i].ID < msg.Info[j].ID
		})
		//分页发送
		//仅记录最后一页数据
		if MaxPage < len(msg.Info) {
			items := make([]*protoMsg.GoodsInfo, 0)
			msg.PageNum = IndexStart
			for i := 0; i < len(msg.Info); i++ {
				if 0 == i%MaxPage {
					if 0 < len(items) {
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					items = make([]*protoMsg.GoodsInfo, 0)
				}
				items = CopyInsert(items, len(items), msg.Info[i]).([]*protoMsg.GoodsInfo)

			}
			if 0 < len(items) {
				msg.Info = items
			}
			GetClientManger().SendData(agent, msg)
		} else {
			GetClientManger().SendData(agent, msg)
		}
	}
}

// 充值
func handleRecharge(args []interface{}) {
	m := args[0].(*protoMsg.RechargeReq)
	agent := args[1].(gate.Agent)
	if m.Payment <= INVALID || false == GetPlayerManger().Exist(m.UserID) || false == GetPlayerManger().Exist(m.ByiD) || m.Method != CodeRecharge {
		GetClientManger().SendResult(agent, FAILED, StatusText[User02])
		log.Debug("玩家%v 充值失败:%v", m.UserID, m)
		return
	}
	person := GetPlayerManger().Get(m.UserID)
	timestamp := time.Now().Unix()
	order := strconv.FormatInt(timestamp, 10)
	msg := &protoMsg.RechargeResp{
		UserID:    m.UserID,
		ByiD:      m.ByiD,
		PreMoney:  person.Money,
		Payment:   m.Payment,
		Reason:    m.Reason,
		Method:    m.Method,
		Order:     util.Md5Sum(order),
		TimeStamp: timestamp,
		IsSuccess: false,
	}
	//扣除金币
	if nowM, err := sqlHandle.RechargeMoney(m.UserID, m.ByiD, m.Payment, CodeRecharge, msg.Order, m.Reason); err == nil {
		GetClientManger().SendData(agent, &protoMsg.UpdateMoneyResp{
			UserID: person.UserID,
			Money:  nowM,
		})
		msg.Money = nowM
		msg.IsSuccess = nowM == msg.PreMoney
		person.Money = nowM
		GetClientManger().SendData(agent, msg)
		GetClientManger().SendResult(agent, SUCCESS, StatusText[User24])
		log.Debug("[充值]玩家%v 充值成功:%v", person.UserID, m)
		return
	} else {
		//GetClientManger().SendData(agent, msg)
		GetClientManger().SendResult(agent, FAILED, StatusText[User12])
		log.Debug("玩家%v 充值失败:%v", person.UserID, m)
	}
}

// 购买房卡
func handleBuyGoods(args []interface{}) {
	m := args[0].(*protoMsg.BuyGoodsReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		//
		if m.Count <= INVALID {
			log.Debug("玩家%v 充值房卡失败:%v", person.UserID, m)
			GetClientManger().SendResult(agent, FAILED, StatusText[User26])
			return
		}

		allGoods := sqlHandle.CheckAllGoods()
		if goods, ok := allGoods[m.ID]; ok {
			total := goods.Price * m.Count
			if m.Payment < total {
				log.Debug("玩家%v 充值房卡失败:%v", person.UserID, m)
				GetClientManger().SendResult(agent, FAILED, StatusText[User12])
				return
			}
			//置换房卡
			timestamp := time.Now().Unix()
			order := strconv.FormatInt(timestamp, 10)
			order = util.Md5Sum(order)
			if err := sqlHandle.UpdateAsset(person.UserID, m.ID, m.Payment, int32(m.Count), CodeBuyRoomCard, order); err != nil {
				log.Debug("玩家%v 充值房卡失败:%v", person.UserID, m)
				GetClientManger().SendResult(agent, FAILED, StatusText[User25])
				return
			}
			//扣除金币
			info := CalculateInfo{
				UserID:   person.UserID,
				ByUID:    SYSTEMID,
				Gid:      INVALID,
				HostID:   INVALID,
				PreMoney: person.Money,
				Payment:  total,
				Code:     CodeBuyRoomCard,
				TypeID:   protoMsg.GameType_General,
				Kid:      goods.Kind,
				Level:    goods.Level,
				Order:    order,
				Remark:   "购买房卡",
			}
			if nowM, _, ok := sqlHandle.DeductMoney(info); ok {
				GetClientManger().SendData(agent, &protoMsg.UpdateMoneyResp{
					UserID: person.UserID,
					Money:  nowM,
				})
				person.Money = nowM
				msg := &protoMsg.BuyGoodsResp{
					UserID: person.UserID,
					Info:   goods,
				}
				msg.Info.Amount = goods.Amount + int32(m.Count)
				msg.Info.Sold = m.Count
				GetClientManger().SendData(agent, msg)
				//GetClientManger().SendResult(agent, SUCCESS, StatusText[User24])
				log.Debug("[充值]玩家%v 充值房卡成功:%v", person.UserID, m)
			} else {
				log.Debug("玩家%v 充值房卡失败:%v", person.UserID, m)
				GetClientManger().SendResult(agent, FAILED, StatusText[User12])
			}
			return
		}

	}
	agent.WriteMsg(&protoMsg.ResultResp{
		State: FAILED,
		Hints: StatusText[User25],
	})
}

// 查看充值记录
func handleGetRecharges(args []interface{}) {
	_ = args[0].(*protoMsg.GetRechargesReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		msg := &protoMsg.GetRechargesResp{
			UserID:  person.UserID,
			PageNum: INVALID,
		}
		msg.AllRecharges = sqlHandle.CheckRecharges(person.UserID)
		//仅记录最后一页数据
		if MaxPage < len(msg.AllRecharges) {
			items := make([]*protoMsg.RechargeResp, 0)
			msg.PageNum = IndexStart
			for i := 0; i < len(msg.AllRecharges); i++ {
				if 0 == i%MaxPage {
					if 0 < len(items) {
						GetClientManger().SendData(agent, msg)
						msg.PageNum++
					}
					items = make([]*protoMsg.RechargeResp, 0)
				}
				items = CopyInsert(items, len(items), msg.AllRecharges[i]).([]*protoMsg.RechargeResp)

			}
			if 0 < len(items) {
				msg.AllRecharges = items
			}
			GetClientManger().SendData(agent, msg)
		} else {
			GetClientManger().SendData(agent, msg)
		}
	}

}

// 查看背包
func handleCheckKnapsack(args []interface{}) {
	_ = args[0].(*protoMsg.CheckKnapsackReq)
	agent := args[1].(gate.Agent)
	if userData := agent.UserData(); userData != nil { //[0
		person := userData.(*Player)
		person.Knapsack = sqlHandle.CheckAssets(person.UserID)
		msg := &protoMsg.CheckKnapsackResp{
			UserID: person.UserID,
			Info:   person.Knapsack,
		}
		GetClientManger().SendData(agent, msg)
	}
}

// ping
func handlePing(args []interface{}) {
	_ = args[0].(*protoMsg.PingReq)
	agent := args[1].(gate.Agent)
	agent.WriteMsg(&protoMsg.PongResp{})
	log.Debug("心跳 %v", agent.RemoteAddr())
}

/////////////////json-->测试用/////////////////////////////
// 消息处理
//func handleLoginJson(args []interface{}) {
//	// 收到的 Test 消息
//	m := args[0].(jsonMsg.UserLogin)
//	// 消息的发送者
//	a := args[1].(gate.Agent)
//	// 1 查询数据库--判断用户是不是合法
//	// 2 如果数据库返回查询正确--保存到缓存或者内存
//	// 输出收到的消息的内容
//	log.Debug("Test login %v", m.LoginName)
//	// 给发送者回应一个 Test 消息
//	a.WriteMsg(&jsonMsg.UserLogin{
//		LoginName: "client",
//	})
//}

//func handleRequestRoomInfoJson(args []interface{})  {
//	m := args[0].(jsonMsg.RequestRoomInfo)
//	// 消息的发送者
//	//a := args[1].(gate.Agent)
//	// 1 查询数据库--判断用户是不是合法
//	// 2 如果数据库返回查询正确--保存到缓存或者内存
//	// 输出收到的消息的内容
//	log.Debug("Test handleRequestRoomInfoJson %v", m)
//	// 给发送者回应一个 Test 消息
//	//a.WriteMsg(&jsonMsg.UserLogin{
//	//	LoginName: "client",
//	//})
//}

//////////////////数据库查询////////////////////////////

//[测试用]
//a := args[1].(gate.Agent)
//a.WriteMsg(&protoT.TestPro{
//	Name:*proto.String("kaile"),
//	Password:*proto.String("doo"),
//})
//Processor.Unmarshal(args[0].([]byte))
//
//buf := make([]byte, 32)
//// 接收消息
//n:=len(args)
//m := &proto.TestPro{}
//proto.Unmarshal(buf[4:n], m)
//
//// 消息的发送者
//a := args[1].(gate.Agent)
//defer a.Close()
//
//// 输出收到的消息的内容
//log.Debug("name:%v password:%v", m.GetName(), m.GetPassword())
//
//
//// 给发送者回应一个 Hello 消息
//a.WriteMsg(proto.UserLogin{})
