package mysql

import (
	"errors"
	"fmt"
	"github.com/po2656233/goleaf/log"
	"strconv"
	"strings"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/manger"
	"superman/nodes/leaf/jettengame/sql/model"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////

// CheckName 玩家昵称
func (sqlSelf *SqlOrm) CheckName(userID int64) string {
	name := ""
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select("name").Where("id = ?", userID).Find(&name).Error
	CheckError(err)
	return name
}

// CheckGold 获取玩家账户金币
func (sqlSelf *SqlOrm) CheckGold(userID int64) int64 {
	userMoney := int64(0)
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select("gold").Where("id = ?", userID).Find(&userMoney).Error
	CheckError(err)
	return userMoney
}

// CheckMoney 获取玩家账户Money
func (sqlSelf *SqlOrm) CheckMoney(userID int64) int64 {
	userMoney := int64(0)
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select("money").Where("id = ?", userID).Find(&userMoney).Error
	CheckError(err)
	return userMoney
}

// updateGold 更新玩家金币
func (sqlSelf *SqlOrm) updateGold(userID int64, gold int64) error {
	user := model.User{}
	now := time.Now()
	data := map[string]interface{}{"gold": gold, "updated_at": &now}
	err := sqlSelf.db.Table(user.TableName()).Where("id=?", userID).Updates(data).Error
	CheckError(err)
	return err
}

// updateMoney 更新玩家钱
func (sqlSelf *SqlOrm) updateMoney(userID int64, money int64) error {
	user := model.User{}
	now := time.Now()
	data := map[string]interface{}{"money": money, "updated_at": &now}
	err := sqlSelf.db.Table(user.TableName()).Where("id=?", userID).Updates(data).Error
	CheckError(err)
	return err
}

// CheckGameRecordLast 获取玩家最后一次游戏记录
func (sqlSelf *SqlOrm) CheckGameRecordLast(gid, uid int64) (model.GameRecord, error) {
	record := model.GameRecord{}
	err := sqlSelf.db.Table(record.TableName()).Select("id,father_id,rounds,inning,is_free,remain_free").Where("gid=? AND uid=?", gid, uid).
		Order("time DESC").Limit(1).Find(&record).Error
	CheckError(err)
	return record, err
}

// CheckUser 获取玩家ID
func (sqlSelf *SqlOrm) CheckUser(name string) (uid int64) {
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select("id").Where("name = ?", name).Find(&uid).Error
	if !CheckError(err) {
		return 0
	}
	return uid
}

// CheckPlatformInfo 获取平台信息
func (sqlSelf *SqlOrm) CheckPlatformInfo(uid int64) (platformID int64) {
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select("platformid").Where("id = ?", uid).Find(&platformID).Error
	CheckError(err)
	return
}

// CheckLogin 登录查询，改为redis
func (sqlSelf *SqlOrm) CheckLogin(name, password string) (uid int64, isSuccessful bool) {
	isSuccessful = false
	uid = 0
	user := model.User{}
	field := "id"
	query := "`account` = ? and `password`=?"
	err := sqlSelf.db.Table(user.TableName()).Select(field).Where(query, name, password).Find(&uid).Error
	if CheckError(err) && uid != 0 {
		isSuccessful = true
	}
	return uid, isSuccessful
}

// CheckLoginTime 登录时间
func (sqlSelf *SqlOrm) CheckLoginTime(uid int64) (timestamp int64) {
	user := model.User{}
	field := "`logintime`"
	query := "`id` = ? "
	err := sqlSelf.db.Table(user.TableName()).Select(field).Where(query, uid).Find(&timestamp).Error
	CheckError(err)
	return timestamp
}

// CheckLeaveTime 登出时间
func (sqlSelf *SqlOrm) CheckLeaveTime(uid int64) (timestamp int64) {
	user := model.User{}
	field := "`leavetime`"
	query := "`id` = ? "
	err := sqlSelf.db.Table(user.TableName()).Select(field).Where(query, uid).Find(&uid).Error
	CheckError(err)
	return timestamp
}

// CheckUserInfo 获取玩家信息
func (sqlSelf *SqlOrm) CheckUserInfo(userID int64) *protoMsg.UserInfo {
	fields := "`id`,`name`,`account`,`face`,`gender`, `age`,`vip`,`level`,`money`,`passport`,`realname`,`phone`,`email`, `address`,`identity`,`agentid`,`referralcode`,`clientaddr`,`serveraddr`,`machinecode`"
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select(fields).Where("id = ?", userID).Find(&user).Error
	CheckError(err)
	msg := &protoMsg.UserInfo{
		UserID:       userID,
		Name:         user.Name,
		Account:      user.Account,
		FaceID:       user.Face,
		Gender:       user.Gender,
		Age:          user.Age,
		VIP:          user.Vip,
		Level:        user.Level,
		Money:        user.Money,
		PassPortID:   user.Passport,
		RealName:     user.Realname,
		PhoneNum:     user.Phone,
		Email:        user.Email,
		Address:      user.Address,
		IDentity:     user.Identity,
		AgentID:      user.Agentid,
		ReferralCode: user.Referralcode,
		ClientAddr:   user.Clientaddr,
		ServerAddr:   user.Serveraddr,
		MachineCode:  user.Machinecode,
	}
	return msg
}

// GetRobotsInfo 获取机器人
func (sqlSelf *SqlOrm) GetRobotsInfo(count int) []*protoMsg.PlayerInfo {
	userList := make([]*model.User, 0)
	userM := model.User{}
	fields := "`id`,`name`,`account`,`face`,`gender`, `platformid`,`age`,`vip`,`level`,`money`,`passport`,`realname`,`phone`,`email`, `address`,`identity`,`agentid`,`referralcode`,`clientaddr`,`serveraddr`,`machinecode`"
	query := "`gender`=? AND `leavetime`>=`logintime`"
	err := sqlSelf.db.Table(userM.TableName()).Select(fields).Where(query, 0x0F).Limit(count).Find(&userList).Error
	if !CheckError(err) {
		return nil
	}

	robotList := make([]*protoMsg.PlayerInfo, 0)
	ids := make([]int64, 0)
	for _, user := range userList {
		msg := &protoMsg.PlayerInfo{
			UserID:     user.ID,
			Name:       user.Name,
			Account:    user.Account,
			FaceID:     user.Face,
			Sex:        user.Gender,
			Age:        user.Age,
			Level:      user.Level,
			Money:      user.Money,
			PlatformID: int64(user.Platformid),
		}
		ids = append(ids, msg.UserID)
		robotList = append(robotList, msg)
	}
	sqlSelf.UpdateLoginTime(ids)
	return robotList
}

/////////////////////////////////////////////////////////////////////////////////

// CheckGoods 查看商品信息
func (sqlSelf *SqlOrm) CheckGoods(goodsId int64) *protoMsg.GoodsInfo {
	goods := model.Goods{}
	fields := "`name`,`kind`,`level`,`price`,`store`,`sold`"
	err := sqlSelf.db.Table(goods.TableName()).Select(fields).Where("id = ?", goodsId).Find(&goods).Error
	if !CheckError(err) {
		return nil
	}
	return &protoMsg.GoodsInfo{
		Name:  goods.Name,
		ID:    goods.ID,
		Kind:  goods.Kind,
		Level: goods.Level,
		Price: goods.Price,
		Store: goods.Store,
		Sold:  goods.Sold,
	}
}

// CheckAllGoods 查看当前所有商品信息
func (sqlSelf *SqlOrm) CheckAllGoods() (goodsInfoList map[int64]*protoMsg.GoodsInfo) {
	goodsList := make([]*model.Goods, 0)
	goods := model.Goods{}
	fields := "*"
	err := sqlSelf.db.Table(goods.TableName()).Select(fields).Find(&goodsList).Error
	if !CheckError(err) {
		return nil
	}
	goodsInfoList = make(map[int64]*protoMsg.GoodsInfo, 0)
	for _, m := range goodsList {
		goodsInfoList[m.ID] = &protoMsg.GoodsInfo{
			Name:  m.Name,
			ID:    m.ID,
			Kind:  m.Kind,
			Level: m.Level,
			Price: m.Price,
			Store: m.Store,
			Sold:  m.Sold,
		}
	}
	return goodsInfoList
}

// CheckAssets 获取玩家资产信息 [放入背包]
func (sqlSelf *SqlOrm) CheckAssets(userID int64) *protoMsg.KnapsackInfo {
	assetList := make([]model.Asset, 0)
	asset := model.Asset{}
	fields := "`goodsid`, `amount`,`spending`"
	err := sqlSelf.db.Table(asset.TableName()).Select(fields).Where("uid=?", userID).Find(&assetList).Error
	if !CheckError(err) {
		return nil
	}
	allGoods := sqlSelf.CheckAllGoods()
	knapsack := &protoMsg.KnapsackInfo{
		ID:      userID,
		MyGoods: make([]*protoMsg.GoodsInfo, 0),
	}
	for _, m := range assetList {
		if item, ok := allGoods[m.Goodsid]; ok {
			item.Store = int64(m.Amount)
			item.Sold = int64(m.Spending)
			knapsack.MyGoods = append(knapsack.MyGoods, item)
		}
	}

	return knapsack
}

// CheckBaseInfo 游戏信息
func (sqlSelf *SqlOrm) CheckBaseInfo(gameID int64) (*protoMsg.GameInfo, *protoMsg.TableInfo) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	game := model.Game{}
	fields := "`name`,`type`,`level`,`kindid`,`enterscore`,`lessscore`,`state`,`maxonline`,`commission`,`maxchair`,`amount`,`playscore`,`hostid`,`robot_count`"
	err := sqlSelf.db.Table(game.TableName()).Select(fields).Where("id=?", gameID).Find(&game).Error
	if !CheckError(err) {
		return nil, nil
	}
	gInfo := &protoMsg.GameInfo{
		Name:   game.Name,
		Level:  game.Level,
		KindID: game.Kindid,
		Type:   protoMsg.GameType(game.Type),
		Scene:  protoMsg.GameScene(0),
	}
	tInfo := &protoMsg.TableInfo{
		HostID:     game.Hostid,
		Password:   game.Password,
		LessScore:  int32(game.Lessscore),
		EnterScore: int32(game.Enterscore),
		PlayScore:  game.Playscore,
		Amount:     game.Amount,
		MaxChair:   game.Maxchair,
		State:      protoMsg.TableState(game.State),
		Commission: game.Commission,
	}
	return gInfo, tInfo
}

// CheckClassInfo 获取分类信息
func (sqlSelf *SqlOrm) CheckClassInfo(classID int64) (item *protoMsg.ClassItem) {
	room := model.Room{}
	fields := "`name`,`roomkey`"
	err := sqlSelf.db.Table(room.TableName()).Select(fields).Where("num=?", classID).Find(&room).Error
	if !CheckError(err) {
		return nil
	}
	return &protoMsg.ClassItem{
		ID:   classID,
		Name: room.Name,
		Key:  room.Roomkey,
	}
}

// CheckRoomInfo 获取房间信息
func (sqlSelf *SqlOrm) CheckRoomInfo(roomNum int64) *model.Room {
	room := model.Room{}
	fields := "`type`,`num`,`name`,`games`,`roomkey`,`mark`"
	err := sqlSelf.db.Table(room.TableName()).Select(fields).Where("num=?", roomNum).Find(&room).Error
	if !CheckError(err) {
		return nil
	}
	return &room
}

// CheckGameKind 获取游戏种类
func (sqlSelf *SqlOrm) CheckGameKind(gid int64) (kindId int32, err error) {
	game := model.Game{}
	fields := "`kindid`"
	err = sqlSelf.db.Table(game.TableName()).Select(fields).Where("id=?", gid).First(&kindId).Error
	CheckError(err)
	return
}

// CheckGameList 服务列表查询
func (sqlSelf *SqlOrm) CheckGameList(classNum int32) (name, key string, games *protoMsg.GameList) {
	//从房间中找到kindID和Level
	room := model.Room{}
	fields := "`id`,`name`,`roomkey`,`games`,`type`"
	err := sqlSelf.db.Table(room.TableName()).Select(fields).Where("num=?", classNum).Find(&room).Error
	if !CheckError(err) {
		return "", "", nil
	}
	games = &protoMsg.GameList{}
	roomGames := strings.Split(room.Games, ",")
	if room.Games == "" || 0 == len(roomGames) {
		gameList := make([]model.Game, 0)
		fields = "`id`,`name`,`kindid`,`level`,`hostid`,`robot_count`"
		err := sqlSelf.db.Table((&model.Game{}).TableName()).Select(fields).Where("type=?", room.Type).Find(&gameList).Error
		if !CheckError(err) {
			return "", "", nil
		}
		for _, game := range gameList {
			games.Items = append(games.Items, &protoMsg.GameItem{
				ID: game.ID,
				Info: &protoMsg.GameInfo{
					Type:   protoMsg.GameType(room.Type),
					KindID: game.Kindid,
					Level:  game.Level,
					Scene:  protoMsg.GameScene(0),
					Name:   game.Name,
				},
			})
		}

	}
	for _, strKid := range roomGames {
		id, _ := strconv.Atoi(strKid)
		if gameInfo, _ := sqlSelf.CheckBaseInfo(int64(id)); gameInfo != nil {
			games.Items = append(games.Items, &protoMsg.GameItem{
				ID:   int64(id),
				Info: gameInfo,
			})
		}
	}
	return room.Name, room.Roomkey, games
}

func (sqlSelf *SqlOrm) CheckRooms(pid int64) []int64 {
	roomList := make([]int64, 0)
	platform := &model.Platform{}
	err := sqlSelf.db.Table(platform.TableName()).Select("`rooms`").Where("id = ?", pid).Find(&platform).Error
	if !CheckError(err) {
		return roomList
	}

	rooms := strings.Split(platform.Rooms, ",")
	for _, room := range rooms {
		if num, err1 := strconv.Atoi(room); num != 0 && err1 == nil {
			roomList = append(roomList, int64(num))
		}
	}
	return roomList
}

// CheckClassList 获取房间列表(应当先查redis 后查mysql)
func (sqlSelf *SqlOrm) CheckClassList(userID int64, platRooms []int64) (roomsInfo *protoMsg.ClassList) {
	strNums := ""
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Select("roomnums").Where("id = ?", userID).Find(&strNums).Error
	if !CheckError(err) {
		return nil
	}

	log.Debug("读取数据库数据:->...房间列表 rooms[%v]\n", strNums)

	//提取数据
	allRooms := strings.Split(strNums, ",")
	roomsInfo = &protoMsg.ClassList{}
	roomsInfo.Classify = make([]*protoMsg.ClassItem, 0)
	if strNums == "" || 0 == len(allRooms) {
		for _, v := range platRooms {
			info := sqlSelf.CheckClassInfo(v)
			roomsInfo.Classify = CopyInsert(roomsInfo.Classify, len(roomsInfo.Classify), info).([]*protoMsg.ClassItem)
		}
	} else {
		for _, room := range allRooms {
			if num, error1 := strconv.Atoi(room); error1 == nil {
				info := sqlSelf.CheckClassInfo(int64(num))
				roomsInfo.Classify = CopyInsert(roomsInfo.Classify, len(roomsInfo.Classify), info).([]*protoMsg.ClassItem)
			}
		}
	}
	roomsInfo.Classify = SliceRemoveDuplicate(roomsInfo.Classify).([]*protoMsg.ClassItem)
	return roomsInfo
}

func (sqlSelf *SqlOrm) CheckTaskList(userID int64) (tasks []int32) {
	_ = userID
	return tasks
}

/////////////////////////////////////////////////////////////////

// CheckInnings 查找牌局记录
func (sqlSelf *SqlOrm) CheckInnings(uid, gid int64) []*protoMsg.InningInfo {
	recordList := make([]*model.Record, 0)
	record := model.Record{}
	fields := "`order`,`payment`,`time`,`code`,`success`,`remark`"
	err := sqlSelf.db.Table(record.TableName()).Select(fields).Where("`uid`=? AND `byid`=? AND `gid`=?", uid, SYSTEMID, gid).Find(&recordList).Error
	if !CheckError(err) {
		return nil
	}
	infos := make([]*protoMsg.InningInfo, 0)
	for _, m := range recordList {
		if m.Success == SUCCESS && m.Code == CodeSettle {
			info := &protoMsg.InningInfo{
				GameID:    gid,
				GameName:  m.Remark,
				Number:    m.Order,
				Payoff:    m.Gold,
				TimeStamp: m.Time,
			}
			infos = CopyInsert(infos, len(infos), info).([]*protoMsg.InningInfo)
		}
	}
	return infos
}

// CheckRecharges 查看充值详情
func (sqlSelf *SqlOrm) CheckRecharges(uid int64) []*protoMsg.RechargeResp {
	rechargeList := make([]*model.Recharge, 0)
	recharge := model.Recharge{}
	fields := "`byid`,`premoney`,`payment`,`money`,`code`,`success`,`timestamp`,`order`,`remark`"
	err := sqlSelf.db.Table(recharge.TableName()).Select(fields).Where("`uid`=?", uid).Find(&rechargeList).Error
	if !CheckError(err) {
		return nil
	}
	infos := make([]*protoMsg.RechargeResp, 0)
	for _, m := range rechargeList {
		infos = CopyInsert(infos, len(infos), &protoMsg.RechargeResp{
			UserID:    uid,
			ByiD:      m.Byid,
			PreMoney:  m.Premoney,
			Payment:   m.Payment,
			Money:     m.Money,
			Reason:    m.Remark,
			Method:    m.Code,
			Order:     m.Order,
			TimeStamp: m.Timestamp,
			IsSuccess: m.Success == 0,
		}).([]*protoMsg.RechargeResp)
	}
	return infos
}

// CheckRecords 查看游戏记录
func (sqlSelf *SqlOrm) CheckRecords(uid, kindID int64, level int32, start, end int64) []*protoMsg.InningInfo {
	if end == 0 {
		end = time.Now().Unix()
	}
	recordList := make([]*model.Record, 0)
	record := model.Record{}
	fields := "`gid`,`order`,`payment`,`time`,`code`,`success`,`remark`"
	//query := "`uid`=? AND `byid`=? AND `time` between ? AND ?"
	query := "`uid`=?  AND `time` between ? AND ?"

	var err error
	if kindID == INVALID && level == INVALID {
		err = sqlSelf.db.Table(record.TableName()).Select(fields).Where(query, uid, start, end).Find(&recordList).Error
	} else if Limit <= level {
		query += " AND `kid`=? "
		err = sqlSelf.db.Table(record.TableName()).Select(fields).Where(query, uid, start, end, kindID).Find(&recordList).Error
	} else {
		query += " AND `kid`=? AND `level`=? "
		err = sqlSelf.db.Table(record.TableName()).Select(fields).Where(query, uid, start, end, kindID, level).Find(&recordList).Error
	}
	if !CheckError(err) {
		return nil
	}
	infos := make([]*protoMsg.InningInfo, 0)
	for _, m := range recordList {
		infos = CopyInsert(infos, len(infos), &protoMsg.InningInfo{
			GameID:    m.Gid,
			GameName:  m.Remark,
			Number:    m.Order,
			Payoff:    m.Payment,
			TimeStamp: m.Time,
		}).([]*protoMsg.InningInfo)
	}
	return infos
}

// CheckGetCheckIn 获取签到信息
func (sqlSelf *SqlOrm) CheckGetCheckIn(userID int64) []*protoMsg.CheckInResp {
	checkinList := make([]*model.Checkin, 0)
	checkin := model.Checkin{}
	fields := "`remark`,`timestamp`"
	query := "`uid`=? "
	err := sqlSelf.db.Table(checkin.TableName()).Select(fields).Where(query, userID).Find(&checkinList).Error
	if !CheckError(err) {
		return nil
	}

	infos := make([]*protoMsg.CheckInResp, 0)
	for _, m := range checkinList {
		infos = CopyInsert(infos, len(infos), &protoMsg.CheckInResp{
			UserID:    userID,
			Remark:    m.Remark,
			Timestamp: m.Timestamp,
		}).([]*protoMsg.CheckInResp)
	}
	return infos
}

// CheckEmail 获取某条Email
func (sqlSelf *SqlOrm) CheckEmail(eid int64) (*protoMsg.EmailInfo, string) {
	email := model.Email{}
	fields := "`sender`,`accepter`,`carboncopy`,`topic`,`content`,`goods`,`timestamp`,`isread`"
	query := "`id`=? "
	err := sqlSelf.db.Table(email.TableName()).Select(fields).Where(query, eid).Find(&email).Error
	if !CheckError(err) {
		return nil, ""
	}

	info := &protoMsg.EmailInfo{
		EmailID:    email.ID,
		AcceptName: email.Accepter,
		Sender:     email.Sender,
		Cc:         email.Carboncopy,
		Topic:      email.Topic,
		Content:    email.Content,
		IsRead:     email.Isread == 1,
		TimeStamp:  email.Timestamp,
	}
	if email.Goods != "" {
		all := strings.Split(email.Goods, " ")
		data := make([]byte, 0)
		for _, item := range all {
			if i, err2 := strconv.Atoi(item); err2 == nil {
				data = append(data, byte(i))
			}
		}
		info.AwardList = &protoMsg.GoodsList{}
		err = BytesToPB(data, info.AwardList)
		CheckError(err)
	}

	return info, email.Goods
}

func (sqlSelf *SqlOrm) CheckEmails(userID int64) []*protoMsg.EmailInfo {
	accepter := sqlSelf.CheckName(userID)

	emailList := make([]*model.Email, 0)
	emailM := model.Email{}
	fields := "`id`,`sender`,`carboncopy`,`topic`,`content`,`goods`,`timestamp`,`isread`"
	query := "`accepter`=? "
	err := sqlSelf.db.Table(emailM.TableName()).Select(fields).Where(query, accepter).Find(&emailList).Error
	if !CheckError(err) {
		return nil
	}
	infos := make([]*protoMsg.EmailInfo, 0)
	for _, email := range emailList {
		info := &protoMsg.EmailInfo{
			EmailID:    email.ID,
			AcceptName: email.Accepter,
			Sender:     email.Sender,
			Cc:         email.Carboncopy,
			Topic:      email.Topic,
			Content:    email.Content,
			IsRead:     email.Isread == 1,
			TimeStamp:  email.Timestamp,
		}
		if email.Goods != "" {
			all := strings.Split(email.Goods, " ")
			data := make([]byte, 0)
			for _, item := range all {
				if i, err2 := strconv.Atoi(item); err2 == nil {
					data = append(data, byte(i))
				}
			}
			info.AwardList = &protoMsg.GoodsList{}
			err = BytesToPB(data, info.AwardList)
			CheckError(err)
		}
		infos = CopyInsert(infos, len(infos), info).([]*protoMsg.EmailInfo)
	}
	return infos
}

// CheckNotice 【大厅公告】库内的公告没有userID字段 且暂不支持子游戏内公告 需外部赋值
func (sqlSelf *SqlOrm) CheckNotice(platID int64) []*protoMsg.NotifyNoticeResp {
	noticeList := make([]*model.Notice, 0)
	noticeM := model.Notice{}
	fields := "`kindid`,`level`,`type`,`start`,`end`,`title`,`content`"
	query := "`platid`=? "
	err := sqlSelf.db.Table(noticeM.TableName()).Select(fields).Where(query, platID).Find(&noticeList).Error
	if !CheckError(err) {
		return nil
	}
	timestamp := time.Now().Unix()
	infos := make([]*protoMsg.NotifyNoticeResp, 0)
	for _, m := range noticeList {
		notice := &protoMsg.NotifyNoticeResp{
			UserID:  INVALID,
			GameID:  INVALID,
			Level:   protoMsg.NTFLevel(m.Type),
			Title:   m.Title,
			Content: m.Content,
			TimeInfo: &protoMsg.TimeInfo{
				TimeStamp: m.Start,
				TotalTime: int32(m.End - m.Start),
				OutTime:   int32(timestamp - m.Start),
				WaitTime:  int32(m.End - timestamp),
			},
		}
		infos = CopyInsert(infos, len(infos), notice).([]*protoMsg.NotifyNoticeResp)
	}
	return infos
}

// CheckRobotCount 获取机器人数量
func (sqlSelf *SqlOrm) CheckRobotCount(gameID int64) (count int32, err error) {
	gameM := model.Game{}
	fields := "`robot_count`"
	query := "`id`=? "
	err = sqlSelf.db.Table(gameM.TableName()).Select(fields).Where(query, gameID).Find(&count).Error
	CheckError(err)
	return
}

// CheckRoomNum 获取房间号
func (sqlSelf *SqlOrm) CheckRoomNum(gid int64) (num int64, err error) {
	room := model.Room{}
	gameM := model.Game{}
	fields := "a.num"
	query := "b.id=? "
	rTable := fmt.Sprintf("%v a", room.TableName())
	strJoin := fmt.Sprintf("LEFT join %v b on a.type = b.type", gameM.TableName())
	err = sqlSelf.db.Table(rTable).Select(fields).Joins(strJoin).Where(query, gid).Find(&num).Error
	CheckError(err)
	return
}

// //////////////////////////////以下有关写操作 强烈要求加锁////////////////////////////////////////////////////////////////////

// AddUser 完成注册 【新增一个玩家】
func (sqlSelf *SqlOrm) AddUser(user model.User) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	if 0 < sqlSelf.CheckUser(user.Account) {
		return 0, errors.New("用户已经存在")
	}
	// SignInTime
	now := time.Now()
	user.Signintime = now.Unix()
	user.CreatedAt = now
	err := sqlSelf.db.Table(user.TableName()).Create(&user).Error
	if !CheckError(err) {
		return 0, err
	}
	return user.ID, nil
}

// AddGame 新增游戏 返回游戏ID [废弃返回值,有程序决定游戏ID]
func (sqlSelf *SqlOrm) AddGame(info *protoMsg.GameInfo, tInfo *protoMsg.TableInfo) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	game := model.Game{
		Hostid:     tInfo.HostID,
		Name:       info.Name,
		Password:   tInfo.Password,
		Kindid:     info.KindID,
		Type:       int32(info.Type),
		Lessscore:  int64(tInfo.LessScore),
		Enterscore: int64(tInfo.EnterScore),
		Playscore:  tInfo.PlayScore,
		Amount:     tInfo.Amount,
		Maxchair:   tInfo.MaxChair,
		State:      int32(tInfo.State),
		Level:      info.Level,
		Commission: tInfo.Commission,
		Maxonline:  tInfo.MaxOnline,
		RobotCount: tInfo.RobotCount,
		CreatedAt:  time.Now(),
	}
	err := sqlSelf.db.Table(game.TableName()).Create(&game).Error
	if !CheckError(err) {
		return INVALID, err
	}
	if game.ID == INVALID {
		return game.ID, errors.New("游戏ID不合法")
	}
	return game.ID, nil
}

// AddGameRecord 添加游戏记录
func (sqlSelf *SqlOrm) AddGameRecord(record model.GameRecord) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	record.Time = time.Now().Unix()
	record.CreatedAt = time.Now()
	err := sqlSelf.db.Table(record.TableName()).Create(&record).Error
	if !CheckError(err) {
		return INVALID, err
	}
	return record.ID, nil

}

// AddEmail 邮件
func (sqlSelf *SqlOrm) AddEmail(userID, toID int64, carboncopy, topic, content, goods, remark string) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	sender := StatusText[Title008]
	if userID != SYSTEMID {
		sender = sqlSelf.CheckName(userID)
	}
	accepter := sqlSelf.CheckName(toID)
	now := time.Now()
	email := model.Email{
		Sender:     sender,
		Accepter:   accepter,
		Carboncopy: carboncopy,
		Topic:      topic,
		Content:    content,
		Goods:      goods,
		Remark:     remark,
		Timestamp:  now.Unix(),
		CreatedAt:  now,
	}
	err := sqlSelf.db.Table(email.TableName()).Create(&email).Error
	if !CheckError(err) {
		return INVALID, err
	}
	return email.ID, nil
}

// AddNotice 公告
func (sqlSelf *SqlOrm) AddNotice(notice model.Notice) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	notice.CreatedAt = time.Now()
	err := sqlSelf.db.Table(notice.TableName()).Create(&notice).Error
	if !CheckError(err) {
		return INVALID, err
	}
	return int64(notice.ID), nil
}

// CheckIn 签到
func (sqlSelf *SqlOrm) CheckIn(userID int64, remark string) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	sTime, eTime := GetOneDay()

	checkin := model.Checkin{}
	fields := "`timestamp`"
	query := "(?) <=`timestamp` and `timestamp`<(?) and `uid`in(?) "
	err := sqlSelf.db.Table(checkin.TableName()).Select(fields).Where(query, sTime, eTime, userID).Find(&checkin).Error
	if !CheckError(err) {
		return INVALID, err
	}

	if INVALID < checkin.Timestamp {
		return checkin.Timestamp, errors.New(StatusText[Flag0001])
	}
	now := time.Now()
	checkin.UID = userID
	checkin.Timestamp = now.Unix()
	checkin.Remark = remark
	checkin.CreatedAt = now
	checkin.UpdatedAt = now
	err = sqlSelf.db.Table(checkin.TableName()).Create(&checkin).Error
	if !CheckError(err) {
		return INVALID, err
	}
	return checkin.Timestamp, nil
}

// RechargeMoney 充值接口
func (sqlSelf *SqlOrm) RechargeMoney(userID, byUID int64, payment int64, code int32, order, remark string) (int64, error) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	person := manger.GetPlayerManger().Get(userID)
	if person == nil {
		log.Error("无效的玩家数据%v  时间:%v", userID, time.Now().Unix())
		return INVALID, errors.New("无效的玩家数据")
	}
	premoney := person.Money
	success := int32(0)
	money := premoney + payment
	if payment == INVALID {
		success = 1
	}
	now := time.Now()
	recharge := model.Recharge{
		UID:       userID,
		Byid:      byUID,
		Premoney:  premoney,
		Payment:   payment,
		Money:     money,
		Code:      code,
		Order:     order,
		Timestamp: now.Unix(),
		Success:   success,
		Remark:    remark,
		CreatedAt: now,
	}
	err := sqlSelf.db.Table(recharge.TableName()).Create(&recharge).Error
	if !CheckError(err) {
		return INVALID, err
	}
	return money, err
}

// DeductMoney 扣除金币 注:返回类型最好使用error
// userID:接受充值的ID
// money：充值的额度
// code: 操作码
// byUID: 由谁操作
func (sqlSelf *SqlOrm) DeductMoney(info manger.CalculateInfo) (nowMoney, factDeduct int64, isOK bool) {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()

	userID, byUID, gid, hostID := info.UserID, info.ByUID, info.Gid, info.HostID
	preMoney, payment, code, order, remark := info.PreMoney, info.Payment, info.Code, info.Order, info.Remark
	typeID, kid, level := info.TypeID, info.Kid, info.Level

	now := time.Now()
	person := manger.GetPlayerManger().Get(userID)
	if person == nil {
		log.Error("[SQL] [DeductMoney] 失败! 无效的玩家(uid:%v)数据  时间:%v", userID, now.Unix())
		return 0, 0, false
	}

	//对非房卡类型的游戏进行扣款
	if typeID != protoMsg.GameType_TableCard {
		//获取现有金币
		preMoney = sqlSelf.CheckMoney(userID)
		//防止扣出负数
		nowMoney = preMoney - payment
		if nowMoney <= 0 {
			log.Debug("[SQL] [DeductMoney] 金额成负数了:%v  %v %v,则全部扣除完.", nowMoney/100, payment/100, userID)
			//将钱全部扣完
			payment = preMoney
			nowMoney = int64(0)
		}

		log.Debug("[SQL] [DeductMoney] 玩家:%v 之前:%v  当前:%v  扣除:%v", userID, preMoney, nowMoney, payment)
		err := sqlSelf.updateMoney(userID, nowMoney)
		if !CheckError(err) {
			return nowMoney, factDeduct, false
		}
		factDeduct = preMoney - nowMoney
		person.Money = nowMoney
	} else {
		// 房卡类型
		nowMoney = preMoney - payment
		factDeduct = payment
		err := sqlSelf.updateGold(userID, nowMoney)
		if !CheckError(err) {
			return nowMoney, factDeduct, false
		}
		code |= 1 << 8
	}

	// 纪录数据变化 	//过滤机器人
	if manger.GetPlayerManger().Get(userID).Sex != 0x0F {
		record := model.Record{
			UID:       userID,
			Byid:      byUID,
			HostID:    hostID,
			Gid:       gid,
			Kid:       int64(kid),
			Level:     level,
			Pergold:   preMoney,
			Payment:   -payment,
			Gold:      nowMoney,
			Code:      code,
			Remark:    remark,
			Time:      now.Unix(),
			Order:     order,
			Success:   SUCCESS,
			CreatedAt: now,
		}
		err := sqlSelf.db.Table(record.TableName()).Create(&record).Error
		if !CheckError(err) {
			return nowMoney, factDeduct, false
		}
	}

	return nowMoney, payment, true
}

/////////////////////////////////////////////////////////////////////////////

func (sqlSelf *SqlOrm) DelGame(gameID int64, roomNum int32) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	gameM := model.Game{
		ID: gameID,
	}
	//err := sqlSelf.db.Table(gameM.TableName()).Delete(&gameM).Error
	//if !CheckError(err) {
	//	return err
	//}
	err := sqlSelf.db.Table(gameM.TableName()).UpdateColumn("state", GameStateClose).Error
	if !CheckError(err) {
		return err
	}
	if roomNum != INVALID {
		room := model.Room{
			Num: roomNum,
		}
		err = sqlSelf.db.Table(room.TableName()).Where("`num`=?", roomNum).Find(&room).Error
		if !CheckError(err) {
			return err
		}
		game := strconv.FormatInt(gameID, 10) + ","
		room.Games = strings.Trim(room.Games, game)
		err = sqlSelf.db.Table(room.TableName()).Save(&room).Error
		if !CheckError(err) {
			return err
		}
	}
	return nil
}

// DelEmail 删除邮件
func (sqlSelf *SqlOrm) DelEmail(eid int64) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	email := model.Email{
		ID: eid,
	}
	err := sqlSelf.db.Table(email.TableName()).Delete(&email).Error
	CheckError(err)
	return err
}

// ClaimEmail 领取奖励
func (sqlSelf *SqlOrm) ClaimEmail(eid int64) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	email := model.Email{
		ID:     eid,
		Goods:  "",
		Isread: Default,
	}
	err := sqlSelf.db.Table(email.TableName()).Updates(&email).Error
	CheckError(err)
	return err
}

// ReadEmail 读取邮件
func (sqlSelf *SqlOrm) ReadEmail(eid int64) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	email := model.Email{}
	err := sqlSelf.db.Table(email.TableName()).Where("`id`=?", eid).Update("`isread`", 1).Error
	CheckError(err)
	return err
}

// UpdateGameAmount 更新游戏剩余次数
func (sqlSelf *SqlOrm) UpdateGameAmount(gameID int64, amount int32) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	game := model.Game{}
	err := sqlSelf.db.Table(game.TableName()).Where("`id`=?", gameID).Update("amount", amount).Error
	CheckError(err)
	return err
}

// UpdateRobotCount 更新机器人数量
func (sqlSelf *SqlOrm) UpdateRobotCount(gameID, hostId int64, count int32) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	game := model.Game{}
	err := sqlSelf.db.Table(game.TableName()).Where("`id`=? AND `hostid`=?", gameID, hostId).Update("robot_count", count).Error
	CheckError(err)
	if err == nil {
		manger.GetRobotManger().ExitGame(gameID)
	}
	return err
}

// UpdateAsset 更新资产 todo 写用主库  读用从库
func (sqlSelf *SqlOrm) UpdateAsset(uid, goodsID int64, payment int64, buyCount, code int32, remark string) error {
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	assetList := make([]*model.Asset, 0)
	asset := model.Asset{}
	field := "`uid`,`goodsid`,`id`,`amount`,`spending`,`count`,`totalprice`"
	query := "`uid`in (?) and `goodsid`in(?)"
	err := sqlSelf.db.Table(asset.TableName()).Select(field).Where(query, uid, goodsID).Find(&assetList).Error
	if !CheckError(err) {
		return nil
	}
	num := INVALID
	now := time.Now()
	for _, m := range assetList {
		num++
		m.Amount += buyCount
		m.Count += int64(buyCount)
		if buyCount < INVALID {
			m.Spending += -buyCount
		}

		m.Totalprice += payment
		m.UpdatedAt = now
		err = sqlSelf.db.Table(asset.TableName()).Save(&m).Error
		if !CheckError(err) {
			log.Error("UpdateAsset err:%v", err)
			return err
		}
	}
	if num == INVALID {
		asset.UID = uid
		asset.Goodsid = goodsID
		asset.Amount = buyCount
		asset.Spending = INVALID
		asset.Count = int64(buyCount)
		asset.Totalprice = payment
		asset.Code = code
		asset.Remark = remark
		asset.Time = now.Unix()

		err = sqlSelf.db.Table(asset.TableName()).Create(&asset).Error
		CheckError(err)
	}

	return err
}

// UpdateLoginTime 更新登录时间
func (sqlSelf *SqlOrm) UpdateLoginTime(ids []int64) {
	if len(ids) == 0 {
		return
	}
	sqlSelf.Lock()
	defer sqlSelf.Unlock()
	user := model.User{}
	now := time.Now().Unix()
	err := sqlSelf.db.Table(user.TableName()).Where("`id`in(?)", ids).Update("logintime", now).Error
	CheckError(err)
}

// UpdateLeaveTime 更新离线时间
func (sqlSelf *SqlOrm) UpdateLeaveTime(ids []int64) {
	if len(ids) == 0 {
		return
	}
	sqlSelf.Lock()
	defer sqlSelf.Unlock()

	user := model.User{}
	now := time.Now().Unix()
	err := sqlSelf.db.Table(user.TableName()).Where("`id`in(?)", ids).Update("leavetime", now).Error
	CheckError(err)
}
