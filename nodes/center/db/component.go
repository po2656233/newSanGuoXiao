package db

import (
	"github.com/po2656233/superplace/components/gorm"
	"github.com/po2656233/superplace/extend/utils"
	"github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"gorm.io/gorm"
	. "superman/internal/constant"
	"superman/internal/utils"
	"sync"
)

var (
	onLoadFuncList []func() // db初始化时加载函数列表
)

type Component struct {
	facade.Component
	db *gorm.DB
	sync.Mutex
}

func (c *Component) Name() string {
	return CenterDb
}

// Init 组件初始化函数
// 为了简化部署的复杂性，本示例取消了数据库连接相关的逻辑
func (c *Component) Init() {
}

func (c *Component) OnAfterInit() {
	//addOnload()
	
	// db配置的注解
	// 打开profile-dev.json，找到"game-1"和"db"配置
	// 当前示例启动的节点id为 game-1
	// db_id_list参数配置了center_db_1，表示当前节点可以连接该数据库
	// 当前节点启时注册了gorm组件  app.Register(cherryGORM.NewComponent())
	// 通过gorm组件可以获取对应的gorm.DB对象
	// 后续操作请参考gorm的用法
	component := superGORM.NewComponent()
	name := component.Name()
	// 获取gorm组件
	gormCpt := c.App().Find(name).(*superGORM.Component)
	if gormCpt == nil {
		log.Panic("[component = %s] not found.", name)
		return
	}
	// 获取 db_id = "center_db_1" 的配置
	dbID := c.App().Settings().GetConfig(DbList).GetString(c.Name())
	if c.db != nil {
		dbObj, err := c.db.DB()
		utils.CheckError(err)
		err = dbObj.Close()
		utils.CheckError(err)
		c.db = nil
	}
	c.db = gormCpt.GetDb(dbID)
	if c.db == nil {
		log.Panic(dbID, " not found")
	}
	for _, fn := range onLoadFuncList {
		exUtils.Try(fn, func(errString string) {
			log.Warnf(errString)
		})
	}
}

func (*Component) OnStop() {
	//组件停止时触发逻辑
}

func New() *Component {
	return &Component{}
}

func addOnload(fn func()) {
	//onLoadFuncList = append(onLoadFuncList, fn)
}
