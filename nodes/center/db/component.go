package db

import (
	"github.com/po2656233/superplace/extend/utils"
	"github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
)

var (
	onLoadFuncList []func() // db初始化时加载函数列表
)

type Component struct {
	facade.Component
}

func (c *Component) Name() string {
	return "db_center_component"
}

// Init 组件初始化函数
// 为了简化部署的复杂性，本示例取消了数据库连接相关的逻辑
func (c *Component) Init() {
}

func (c *Component) OnAfterInit() {
	//addOnload(loadDevAccount)

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
