package web

import (
	"github.com/gin-gonic/gin"
	"github.com/po2656233/superplace"
	sgxCron "github.com/po2656233/superplace/components/cron"
	sgxGin "github.com/po2656233/superplace/components/gin"
	sgxFile "github.com/po2656233/superplace/extend/file"
	"sanguoxiao/internal/component/check_center"
	"sanguoxiao/internal/conf"
	"sanguoxiao/nodes/web/controller"
	"sanguoxiao/nodes/web/sdk"
)

func Run(profileFilePath, nodeId string) {
	// 配置sgx引擎,加载profile配置文件
	app := superplace.Configure(profileFilePath, nodeId, false, superplace.Cluster)

	// 注册调度组件
	app.Register(sgxCron.New())

	// 注册检查中心服是否启动组件
	app.Register(checkCenter.New())

	// 注册数据配表组件
	app.Register(conf.New())

	// 加载http server组件
	app.Register(httpServerComponent(app.Address()))

	// 加载sdk逻辑
	sdk.Init(app)

	// 启动sgx引擎
	app.Startup()
}

func httpServerComponent(addr string) *sgxGin.Component {
	gin.SetMode(gin.DebugMode)

	// new http server
	httpServer := sgxGin.NewHttp("http_server", addr)
	httpServer.Use(sgxGin.Cors())

	// http server使用gin组件搭建，这里增加一个RecoveryWithZap中间件
	httpServer.Use(sgxGin.RecoveryWithZap(true))

	// 映射h5客户端静态文件到static目录，例如：http://127.0.0.1/static/protocol.js
	httpServer.Static("/static", "./static/")

	// 加载./view目录的html模板文件(详情查看gin文档)
	viewFiles := sgxFile.WalkFiles("./view/", ".html")
	if len(viewFiles) < 1 {
		panic("view files not found.")
	}
	httpServer.LoadHTMLFiles(viewFiles...)

	//注册 controller
	httpServer.Register(new(controller.Controller))

	return httpServer
}
