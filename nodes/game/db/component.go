package db

import (
	superGorm "github.com/po2656233/superplace/components/gorm"
	"github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	. "superman/internal/constant"
	sqlmodel "superman/internal/sql_model"
	. "superman/internal/utils"
)

type Component struct {
	facade.Component
	db    *gorm.DB
	curDB string
}

func New() *Component {
	return &Component{}
}

func (c *Component) Name() string {
	return GameDb
}

// OnInit Actor初始化前触发该函数
func (self *Component) OnInit() {
	self.changeDB(GameDb)
	//// 每秒查询一次db
	//p.Timer().Add(5*time.Second, p.selectDB)
	//// 1秒后进行一次分页查询
	//p.Timer().AddOnce(1*time.Second, p.selectPagination)
}
func (self *Component) changeDB(dbNode string) {
	if self.curDB == dbNode {
		return
	}
	self.curDB = dbNode
	// db配置的注解
	// 打开profile-dev.json，找到"game-1"和"db"配置
	// 当前示例启动的节点id为 game-1
	// db_id_list参数配置了center_db_1，表示当前节点可以连接该数据库
	// 当前节点启时注册了gorm组件  app.Register(cherryGORM.NewComponent())
	// 通过gorm组件可以获取对应的gorm.DB对象
	// 后续操作请参考gorm的用法

	// 获取 db_id = "center_db_1" 的配置
	dbID := self.App().Settings().GetConfig(DbList).GetString(dbNode)
	if self.db != nil {
		dbObj, _ := self.db.DB()
		_ = dbObj.Close()
		self.db = nil
	}
	self.db = superGorm.NewComponent().GetDb(dbID)
	if self.db == nil {
		clog.Panic(dbID, " not found")
	}
}

/////////////////////////////////////////////////////////////////////////////////

// //////////////////////////////以下有关写操作 强烈要求加锁////////////////////////////////////////////////////////////////////

// AddRecord 添加游戏记录
func (self *Component) AddRecord(table *sqlmodel.Record) error {
	err := self.db.Transaction(func(tx *gorm.DB) error {
		// 获取充值前的金额
		coin, err := self.checkUserCoin(table.UID)
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
		user := sqlmodel.User{
			ID: table.UID,
		}
		if err = self.db.Model(user).UpdateColumn("coin", gorm.Expr("coin + ?", table.Payment)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// EraseRemain 减少剩余次数
func (self *Component) EraseRemain(tid int64, amount int32) (int32, error) {
	tb := sqlmodel.Table{
		ID: tid,
	}
	err := self.db.Model(tb).Select("remain").Where("0 < maxround AND 0 < remain").
		UpdateColumn("remain", gorm.Expr("remain - ?", amount)).First(&tb).Error
	return tb.Remain, err
}

func (self *Component) checkUserCoin(uid int64) (coin int64, err error) {
	user := &sqlmodel.User{}
	err = self.db.Table(user.TableName()).Select("coin").Where("id=?", uid).Find(&coin).Error
	CheckError(err)
	return
}
