package db

// [废弃]
import (
	superGorm "github.com/po2656233/superplace/components/gorm"
	"github.com/po2656233/superplace/facade"
	clog "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	"log"
	. "superman/internal/constant"
	"sync"
)

type Component struct {
	facade.Component
	db       *gorm.DB
	centerDB *gorm.DB
	sync.Mutex
}

func New() *Component {
	return &Component{}
}

func (c *Component) Name() string {
	return GameDb
}

func (self *Component) Init() {

}

// OnAfterInit 初始化前触发该函数
func (self *Component) OnAfterInit() {
	component := superGorm.NewComponent()
	name := component.Name()

	// 获取gorm组件
	gormCpt := self.App().Find(name).(*superGorm.Component)
	if gormCpt == nil {
		log.Panic("[Component = %s] not found.", name)
		return
	}
	// 获取 db_id = "center_db_1" 的配置
	dbID := self.App().Settings().GetConfig(DbList).GetString(self.Name())
	if self.db != nil {
		dbObj, _ := self.db.DB()
		_ = dbObj.Close()
		self.db = nil
	}
	self.db = gormCpt.GetDb(dbID)
	if self.db == nil {
		clog.Panic(dbID, " not found")
	}

	centerDbID := self.App().Settings().GetConfig(DbList).GetString(CenterDb)
	if self.centerDB != nil {
		dbObj, _ := self.centerDB.DB()
		_ = dbObj.Close()
		self.centerDB = nil
	}
	self.centerDB = gormCpt.GetDb(centerDbID)
	if self.centerDB == nil {
		clog.Panic(centerDbID, " not found CenterDb ")
	}

	//// 每秒查询一次db
	//p.Timer().Add(5*time.Second, p.selectDB)
	//// 1秒后进行一次分页查询
	//p.Timer().AddOnce(1*time.Second, p.selectPagination)
}
