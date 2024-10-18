package db

import (
	"errors"
	"fmt"
	exSnowflake "github.com/po2656233/superplace/extend/snowflake"
	"gorm.io/gorm"
	. "superman/internal/constant"
	gateMsg "superman/internal/protocol/go_file/gate"
	"superman/internal/rpc"
	"superman/internal/sql_model/center"
	sqlmodel "superman/internal/sql_model/minigame"
	. "superman/internal/utils"

	"time"
)

// GetUserInfo 获取玩家信息
func (self *Component) GetUserInfo(uid int64) (*gateMsg.GetUserInfoResp, error) {
	resp, code := rpc.SendDataToAcc(self.App(), &gateMsg.GetUserInfoReq{
		Uid: uid,
	})
	if code != SUCCESS {
		return nil, errors.New(fmt.Sprintf("addRecharge get userinfo is err:%v", code))
	}
	userInfo, ok := resp.(*gateMsg.GetUserInfoResp)
	if !ok {
		return nil, errors.New("addRecharge get userinfo resp data broken")
	}
	return userInfo, nil
}

// AddRoom 新增房间
func (self *Component) AddRoom(room sqlmodel.Room) (int64, error) {

	rid := self.CheckRoom(room.Hostid, room.Name)
	if 0 < rid {
		return 0, errors.New("房间已经存在")
	}
	room.CreatedAt = time.Now()
	self.Lock()
	defer self.Unlock()
	err := self.db.Table(room.TableName()).Create(&room).Error
	if !CheckError(err) {
		return 0, err
	}
	return room.ID, nil
}

// AddTable 新增桌牌
func (self *Component) AddTable(table sqlmodel.Table) (id int64, maxSit int32, err error) {
	count := self.CheckTableCount(table.Rid)
	max := self.CheckRoomMaxTable(table.Rid)
	if max <= count {
		return 0, 0, fmt.Errorf("%s:%d", StatusText[Room15], max)
	}
	maxSit = self.CheckGameMaxPlayer(table.Gid)
	table.MaxSitter = maxSit
	table.CreatedAt = time.Now()
	self.Lock()
	defer self.Unlock()
	err = self.db.Table(table.TableName()).Create(&table).Error
	if !CheckError(err) {
		return 0, maxSit, err
	}
	room := sqlmodel.Room{}
	err = self.db.Table(room.TableName()).Where("id=?", table.Rid).UpdateColumn("table_count", count+1).Error
	id = table.ID
	return

	//// 使用事务来确保操作的原子性
	//err := self.db.Transaction(func(tx *gorm.DB) error {
	//	// 查询当前gid和rid组合下的最大num值
	//	var currentMaxNum int32
	//	err := tx.Model(&table).Where("gid = ? AND rid = ?", table.Gid, table.Rid).Select("max(num)").Scan(&currentMaxNum).Error
	//	if err != nil {
	//		return err
	//	}
	//
	//	// 计算新的num值
	//	table.Num = currentMaxNum + 1
	//
	//	// 创建一个新的记录
	//	// 插入新的记录
	//	if err := tx.Create(&table).Error; err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})
	//return table.ID, table.Num, err
}

func (self *Component) AddRecharge(table *sqlmodel.Recharge) error {
	// 使用事务来确保操作的原子性
	userInfo, err := self.GetUserInfo(table.UID)
	if err != nil {
		return err
	}
	err = self.centerDB.Transaction(func(tx *gorm.DB) error {
		// 获取充值前的金额
		money := userInfo.Info.Money
		// 插入新的记录
		user := center.User{
			ID: table.UID,
		}
		switch table.Switch {
		case 1:
			yuanbao := table.Payment * 100
			err = self.db.Model(user).Select("yuanbao").UpdateColumn("yuanbao", gorm.Expr("yuanbao+?", yuanbao)).First(&user).Error
		case 2:
			coin := table.Payment * 100 * 100
			err = self.db.Model(user).Select("coin").UpdateColumn("coin", gorm.Expr("coin+?", coin)).First(&user).Error
		default:
			err = self.db.Model(user).Select("money").UpdateColumn("money", gorm.Expr("money+?", table.Payment)).First(&user).Error
		}
		table.Premoney = money
		if user.Money < 0 {
			user.Money = 0
		}
		table.Money = user.Money
		node, err := exSnowflake.NewNode(1)
		if err != nil {
			return err
		}
		table.Order = fmt.Sprintf("%v%v%v", table.Timestamp, RandomStrLetter(4), node.Generate())
		// 创建一个新的 充值记录
		if err = tx.Create(&table).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

func (self *Component) AddRecord(table *sqlmodel.Record) error {
	userInfo, err := self.GetUserInfo(table.UID)
	if err != nil {
		return err
	}
	err = self.centerDB.Transaction(func(tx *gorm.DB) error {
		// 获取充值前的金额
		coin := userInfo.Info.Coin
		// 获取充值前的金额
		if err != nil {
			return err
		}
		table.Pergold = coin
		// 插入新的记录
		table.Gold = table.Pergold + table.Payment
		// 创建一个新的 充值记录
		if err = self.db.Create(&table).Error; err != nil {
			return err
		}
		user := center.User{
			ID: table.UID,
		}
		if err = self.db.Model(user).UpdateColumn("coin", gorm.Expr("coin + ?", table.Payment)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// EraseRemain
func (self *Component) EraseRemain(tid int64, amount int32) (int32, error) {
	tb := sqlmodel.Table{
		ID: tid,
	}
	self.Lock()
	defer self.Unlock()
	err := self.db.Model(tb).Select("remain").Where("0 < maxround AND 0 < remain").
		UpdateColumn("remain", gorm.Expr("remain - ?", amount)).First(&tb).Error
	return tb.Remain, err
}

// DelTable 新增桌牌
func (self *Component) DelTable(rid, tid int64) error {
	tb := sqlmodel.Table{
		ID: tid,
	}
	self.Lock()
	defer self.Unlock()
	err := self.db.Model(tb).Delete(tb).Error
	if !CheckError(err) {
		return err
	}
	room := sqlmodel.Room{}
	return self.db.Table(room.TableName()).Where("id=?", rid).UpdateColumn("table_count", gorm.Expr("table_count - ?", 1)).Error
}

///////////////////////Check//////////////////////////////////////////////

// CheckRoom 检测房间是否包含
func (self *Component) CheckRoom(hostid int64, name string) (rid int64) {
	room := sqlmodel.Room{}
	err := self.db.Table(room.TableName()).Select("id").Where("hostid = ? AND name = ?", hostid, name).Find(&rid).Error
	CheckError(err)
	return
}

// checkRoomCount查看房间数目
func (self *Component) CheckRoomExist(hostid, rid int64) (bool, error) {
	room := sqlmodel.Room{}
	count := int64(0)
	err := self.db.Table(room.TableName()).Where("hostid = ? AND id = ? ", hostid, rid).Count(&count).Error
	CheckError(err)
	return 0 < count, err
}

////////////////////////////////////////////////////////////

// CheckRoomMaxTable 检测桌牌存在数量
func (self *Component) CheckRoomMaxTable(rid int64) (max int64) {
	room := sqlmodel.Room{}
	err := self.db.Table(room.TableName()).Select("max_table").Where("id = ? ", rid).Find(&max).Error
	CheckError(err)
	return
}

// CheckTableCount 检测桌牌存在数量
func (self *Component) CheckTableCount(rid int64) (count int64) {
	tb := sqlmodel.Table{}
	err := self.db.Table(tb.TableName()).Select("id").Where("rid = ? ", rid).Count(&count).Error
	CheckError(err)
	return
}

// CheckTableRid 检测桌牌存在数量
func (self *Component) CheckTableRid(tid int64) (rid int64) {
	room := sqlmodel.Table{}
	err := self.db.Table(room.TableName()).Select("rid").Where("id = ? ", tid).Find(&rid).Error
	CheckError(err)
	return
}

// CheckGameMaxPlayer 检测桌牌存在数量
func (self *Component) CheckGameMaxPlayer(gid int64) (max int32) {
	gm := sqlmodel.Game{}
	err := self.db.Table(gm.TableName()).Select("max_player").Where("id = ? ", gid).Find(&max).Error
	CheckError(err)
	return
}

// CheckRoomInfo 获取玩家信息
func (self *Component) CheckRoomInfo(rid int64) (*sqlmodel.Room, error) {
	rm := &sqlmodel.Room{}
	err := self.db.Table(rm.TableName()).Select("*").Where("id=?", rid).Find(rm).Error
	count := int64(0)
	self.db.Model(&sqlmodel.Table{}).Where("rid = ?", rid).Count(&count)
	rm.TableCount = int32(count)
	CheckError(err)
	return rm, err
}

// CheckGameName 获取游戏名称
func (self *Component) CheckGameName(gid int64) (name string, err error) {
	game := sqlmodel.Game{}
	err = self.db.Table(game.TableName()).Select("name").Where("id = ?", gid).Find(&name).Error
	CheckError(err)
	return
}

// CheckGameEnName 获取游戏英文名
func (self *Component) CheckGameEnName(gid int64) (name string, err error) {
	game := sqlmodel.Game{}
	err = self.db.Table(game.TableName()).Select("en_name").Where("id = ?", gid).Find(&name).Error
	CheckError(err)
	return
}

//////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////

// CheckClassify 获取游戏种类
func (self *Component) CheckClassify() (user []*sqlmodel.Kindinfo, err error) {
	user = make([]*sqlmodel.Kindinfo, 0)
	kind := sqlmodel.Kindinfo{}
	err = self.db.Table(kind.TableName()).Select("*").Find(&user).Error
	CheckError(err)
	return
}

// CheckRooms 获取房间列表 校验TableCount
func (self *Component) CheckRooms(hostId, startTime int64, pageSize, pageNumber int) (rooms []*sqlmodel.Room, err error) {
	rooms = make([]*sqlmodel.Room, 0)
	room := sqlmodel.Room{}
	if pageSize < 0 {
		pageSize = -1
	}
	if pageNumber < 1 {
		pageNumber = 1
	}
	if 0 < startTime && 0 < hostId {
		start := time.Unix(startTime, 0)
		err = self.db.Table(room.TableName()).Select("*").Where("`created_at` >= ? AND `hostid`=?", start, hostId).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&rooms).Error
	} else if 0 < startTime {
		start := time.Unix(startTime, 0)
		err = self.db.Table(room.TableName()).Select("*").Where("`created_at` >= ?", start).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&rooms).Error
	} else if 0 < hostId {
		err = self.db.Table(room.TableName()).Select("*").Where("`hostid`=?", hostId).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&rooms).Error
	} else {
		err = self.db.Table(room.TableName()).Select("*").Find(&rooms).Error
	}
	count := int64(0)
	for _, s := range rooms {
		self.db.Model(&sqlmodel.Table{}).Where("rid = ?", s.ID).Count(&count)
		s.TableCount = int32(count)
	}
	CheckError(err)
	return
}

// CheckTables 获取牌桌列表
func (self *Component) CheckTables(rid int64, pageSize, pageNumber int) (tables []*sqlmodel.Table, err error) {
	tables = make([]*sqlmodel.Table, 0)
	table := sqlmodel.Table{}
	if 0 < rid {
		err = self.db.Table(table.TableName()).Select("*").Where("rid=?", rid).Find(&tables).Error
	} else {
		if pageSize < 0 {
			pageSize = -1
		}
		if pageNumber < 1 {
			pageNumber = 1
		}
		err = self.db.Table(table.TableName()).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&tables).Error
	}
	CheckError(err)
	return
}

// CheckTable 获取牌桌列表
func (self *Component) CheckTable(tid int64) (*sqlmodel.Table, error) {
	table := &sqlmodel.Table{}
	err := self.db.Table(table.TableName()).Select("*").Where("id=?", tid).Find(table).Error
	CheckError(err)
	return table, err
}

// CheckGames 获取游戏列表
func (self *Component) CheckGames(kid int64, pageSize, pageNumber int) (games []*sqlmodel.Game, err error) {
	games = make([]*sqlmodel.Game, 0)
	game := sqlmodel.Game{}
	if kid == -1 {
		if pageSize < 0 {
			pageSize = -1
		}
		if pageNumber < 1 {
			pageNumber = 1
		}
		err = self.db.Table(game.TableName()).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&games).Error
	} else {
		err = self.db.Table(game.TableName()).Select("*").Where("kid=?", kid).Find(&games).Error
	}

	CheckError(err)
	return
}

// CheckGame 获取游戏
func (self *Component) CheckGame(gid int64) (*sqlmodel.Game, error) {
	game := &sqlmodel.Game{}
	err := self.db.Table(game.TableName()).Select("*").Where("id=?", gid).Find(game).Error
	CheckError(err)
	return game, err
}

// CheckAllGames 获取所有游戏列表
func (self *Component) CheckAllGames(pageSize, pageNumber int) (games []*sqlmodel.Game, err error) {
	games = make([]*sqlmodel.Game, 0)
	game := sqlmodel.Game{}
	if pageSize < 0 {
		pageSize = -1
	}
	if pageNumber < 1 {
		pageNumber = 1
	}
	err = self.db.Table(game.TableName()).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Find(&games).Error
	CheckError(err)
	return
}
