package mysql

import (
	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "superman/internal/utils"
	"sync"
)

// ConnectGorm 连接数据库
func ConnectGorm(user, password, address, port, dbName string) *gorm.DB {
	dataSourceName := user + ":" + password + "@tcp(" + address + ":" + port + ")/" + dbName + "?charset=utf8"
	//注意 注意 此处为移除gorm的日志自定义了相关结构。正式使用时 请放开
	db, err1 := gorm.Open(gMysql.Open(dataSourceName), &gorm.Config{
		//Logger:      DiscardLogger{}.LogMode(logger.Info),
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
			//user:     conf.Server.DBUser,
			//password: conf.Server.DBPassword,
			//address:  conf.Server.DBAddress,
			//port:     conf.Server.DBPort,
			//dbName:   conf.Server.DBName,
			db: nil,
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

	//重置机器人
	//sqlSelf.ResetRobot()
}

// CloseMysql 关闭mysql
func (sqlSelf *SqlOrm) CloseMysql() {
	db, err := sqlSelf.db.DB()
	CheckError(err)
	err = db.Close()
	CheckError(err)
}
