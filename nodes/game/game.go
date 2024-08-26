package game

import (
	"github.com/po2656233/superplace"
	"github.com/po2656233/superplace/components/cron"
	"github.com/po2656233/superplace/components/gops"
	"github.com/po2656233/superplace/extend/snowflake"
	cstring "github.com/po2656233/superplace/extend/string"
	sgxUtils "github.com/po2656233/superplace/extend/utils"
	"superman/internal/component/check_center"
	"superman/internal/conf"
	"superman/nodes/game/module/player"
)

func Run(profileFilePath, nodeId string) {
	if sgxUtils.IsNumeric(nodeId) == false {
		panic("node parameter must is number.")
	}

	// snowflake global id
	serverId, _ := cstring.ToInt64(nodeId)
	exSnowflake.SetDefaultNode(serverId)

	// 配置sgx引擎
	app := superplace.Configure(profileFilePath, nodeId, false, superplace.Cluster)
	// diagnose
	app.Register(superGops.New())
	// 注册调度组件
	app.Register(superCron.New())
	// 注册数据配置组件
	app.Register(conf.New())
	// 注册检测中心节点组件，确认中心节点启动后，再启动当前节点
	app.Register(checkCenter.New())
	// 注册db组件[废弃]:方便扩展部署,移除游戏服直接操作db数据库
	// app.Register(db.New())

	app.AddActors(
		&player.ActorGame{},
	)

	app.Startup()
}
