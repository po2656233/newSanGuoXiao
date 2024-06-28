package mysql

import (
	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"strings"
	protoMsg "superman/internal/protocol/gofile"
	. "superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/conf"
	"superman/nodes/leaf/jettengame/manger"
	"superman/nodes/leaf/jettengame/sql/model"
	"superman/nodes/leaf/jettengame/sql/redis"
	"sync"
	"time"
)

// ConnectGorm 连接数据库
func ConnectGorm(user, password, address, port, dbName string) *gorm.DB {
	dataSourceName := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + dbName + "?charset=utf8"
	//注意 注意 此处为移除gorm的日志自定义了相关结构。正式使用时 请放开
	db, err1 := gorm.Open(gMysql.Open(dataSourceName), &gorm.Config{
		Logger:      DiscardLogger{}.LogMode(logger.Info),
		PrepareStmt: false,
	})
	if err1 != nil {
		panic(err1)
	}

	sqlDB, err2 := db.DB()
	if err2 != nil {
		panic(err2)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(0)

	return db.Debug()
}

var once sync.Once
var sqlObj *SqlOrm

func SqlHandle() *SqlOrm {
	once.Do(func() {
		sqlObj = &SqlOrm{
			user:     conf.Server.DBUser,
			password: conf.Server.DBPassword,
			address:  conf.Server.DBAddress,
			port:     conf.Server.DBPort,
			dbName:   conf.Server.DBName,
			db:       nil,
		}
		sqlObj.Init()
	})
	return sqlObj
}

type SqlOrm struct {
	sync.RWMutex
	user     string
	password string
	address  string
	port     string
	dbName   string
	db       *gorm.DB
}

func (sqlSelf *SqlOrm) Init() {
	if sqlSelf.db == nil {
		sqlSelf.db = ConnectGorm(sqlObj.user, sqlObj.password, sqlObj.address, sqlObj.port, sqlObj.dbName)
		//设置连接复用时间为60s 避免 closing bad idle connection: EOF
		//sqlSelf.db.SetConnMaxLifetime(Minute * time.Second)
	}
	//获取平台信息
	sqlSelf.initPlatformInfo()

	//获取游戏信息
	sqlSelf.initGamesInfo()

	//重置机器人
	//sqlSelf.ResetRobot()
}

// initPlatformInfo 初始化平台信息
func (sqlSelf *SqlOrm) initPlatformInfo() bool {
	platformList := make([]*model.Platform, 0)
	err := sqlSelf.db.Table("platform").Select("`id`,`name`,`rooms`,`servers`").Find(&platformList).Error
	if !CheckError(err) {
		return false
	}
	for _, platform := range platformList {
		var roomNums []int64
		var addrList []string

		//提取 房间 数据
		allRooms := strings.Split(platform.Rooms, ",")

		for _, room := range allRooms {
			if num, error1 := strconv.Atoi(room); error1 == nil {
				roomNums = append(roomNums, int64(num))
			}
		}

		//提取服务器地址数据
		allServers := strings.Split(platform.Servers, ";")
		for _, service := range allServers {
			addrList = append(addrList, service)
		}

		manger.GetPlatformManger().Append(&manger.PlatformInfo{
			ID:         int64(platform.ID),
			Name:       platform.Name,
			ClassIDs:   roomNums,
			ServerList: addrList,
		})
	}
	return true
}

// initGamesInfo 初始化游戏信息
func (sqlSelf *SqlOrm) initGamesInfo() bool {
	gameList := make([]model.Game, 0)
	field := "`id`,`hostid`,`name`, `password`,`kindid`, `type`,`lessscore`,`enterscore`,`playscore`,`amount`,`maxchair`, `state`,`level`, `commission`,`robot_count`"
	query := "`state`=?"
	game := model.Game{}
	err := sqlSelf.db.Table(game.TableName()).Select(field).Where(query, 1).Find(&gameList).Error
	if !CheckError(err) {
		return false
	}

	categoriesList := make([]*model.Categories, 0)
	field = "`type`,`kind`,`level`"
	categoriesM := model.Categories{}
	err = sqlSelf.db.Table(categoriesM.TableName()).Select(field).Find(&categoriesList).Error
	if !CheckError(err) {
		return false
	}
	for _, m := range gameList {
		gInfo := &protoMsg.GameInfo{
			Name:   m.Name,
			Level:  m.Level,
			KindID: m.Kindid,
			Type:   protoMsg.GameType(m.Type),
		}
		tInfo := &protoMsg.TableInfo{
			HostID:     m.Hostid,
			Password:   m.Password,
			LessScore:  int32(m.Lessscore),
			EnterScore: int32(m.Enterscore),
			PlayScore:  m.Playscore,
			Amount:     m.Amount,
			MaxChair:   m.Maxchair,
			State:      protoMsg.TableState(m.State),
			Commission: m.Commission,
			RobotCount: m.RobotCount,
		}
		have := false
		redis.RedisHandle().Set(GetGameKindKey(m.ID), m.Kindid, 0)
		manger.GetGamesManger().CreateGame(m.ID, gInfo, tInfo)
		for _, v := range categoriesList {
			if v.Type == int32(gInfo.Type) && v.Kind == gInfo.KindID && v.Level == gInfo.Level {
				have = true
				break
			}
		}
		if !have {
			categories := model.Categories{}
			categories.Name = m.Name
			categories.Kind = m.Kindid
			categories.Level = m.Level
			categories.Type = m.Type
			categories.CreatedAt = time.Now()
			err = sqlSelf.db.Table(categories.TableName()).Create(&categories).Error
			//记录到categories表
			//statement := "INSERT IGNORE INTO `categories`(`name`,`type`, `kind`,`level`, `remark`)VALUES(?,?,?,?,?)"
			if !CheckError(err) {
				return false
			}
		}
	}
	return true
}

// ResetRobot 重置机器人
func (sqlSelf *SqlOrm) ResetRobot() {
	user := model.User{}
	err := sqlSelf.db.Table(user.TableName()).Where("`gender`in(?)", 0x0F).Update("logintime", 0).Error
	CheckError(err)
}

// CloseMysql 关闭mysql
func (sqlSelf *SqlOrm) CloseMysql() {
	db, err := sqlSelf.db.DB()
	CheckError(err)
	err = db.Close()
	CheckError(err)
}
