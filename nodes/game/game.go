package game

import (
	"github.com/po2656233/superplace"
	sgxCron "github.com/po2656233/superplace/components/cron"
	sgxGops "github.com/po2656233/superplace/components/gops"
	sgxSnowflake "github.com/po2656233/superplace/extend/snowflake"
	cstring "github.com/po2656233/superplace/extend/string"
	sgxUtils "github.com/po2656233/superplace/extend/utils"
	"sanguoxiao/internal/component/check_center"
	"sanguoxiao/internal/data"
	"sanguoxiao/nodes/game/db"
	"sanguoxiao/nodes/game/module/player"
	"sanguoxiao/nodes/game/module/sanguoxiao"
)

func Run(profileFilePath, nodeId string) {
	if sgxUtils.IsNumeric(nodeId) == false {
		panic("node parameter must is number.")
	}

	// snowflake global id
	serverId, _ := cstring.ToInt64(nodeId)
	sgxSnowflake.SetDefaultNode(serverId)

	// 配置sgx引擎
	app := superplace.Configure(profileFilePath, nodeId, false, superplace.Cluster)

	// diagnose
	app.Register(sgxGops.New())
	// 注册调度组件
	app.Register(sgxCron.New())
	// 注册数据配置组件
	app.Register(data.New())
	// 注册检测中心节点组件，确认中心节点启动后，再启动当前节点
	app.Register(checkCenter.New())
	// 注册db组件
	app.Register(db.New())

	app.AddActors(
		&player.ActorPlayers{},
		&sanguoxiao.ActorSanGuoXiao{},
	)

	app.Startup()
}
