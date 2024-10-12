package chat

import (
	superGORM "github.com/po2656233/superplace/components/gorm"
	"github.com/po2656233/superplace/net/parser/simple"
	checkCenter "superman/internal/component/check_center"
	"superman/internal/conf"
	"superman/internal/rpc"
	db "superman/nodes/chat/db"
	"superman/nodes/chat/module/player"

	"github.com/po2656233/superplace"
	superCron "github.com/po2656233/superplace/components/cron"
	superGops "github.com/po2656233/superplace/components/gops"
)

func Run(profileFilePath, nodeId string) {
	//if sgxUtils.IsNumeric(nodeId) == false {
	//	panic("node parameter must is number.")
	//}
	//
	//// snowflake global id
	//serverId, _ := cstring.ToInt64(nodeId)
	//exSnowflake.SetDefaultNode(serverId)

	//
	simple.SetParseProtoFunc(rpc.ParseProto)

	// 配置sgx引擎
	app := superplace.Configure(profileFilePath, nodeId, false, superplace.Cluster)
	// 注册gorm组件来处理数据库数据
	app.Register(superGORM.NewComponent())
	// diagnose
	app.Register(superGops.New())
	// 注册调度组件
	app.Register(superCron.New())
	// 注册数据配置组件
	app.Register(conf.New())
	// 注册检测中心节点组件，确认中心节点启动后，再启动当前节点
	app.Register(checkCenter.New())
	// 注册db组件[废弃]:方便扩展部署,移除游戏服直接操作db数据库
	app.Register(db.New())

	app.AddActors(
		&player.ActorChat{},
	)

	app.Startup()
}
