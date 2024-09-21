package web

import (
	"github.com/gin-gonic/gin"
	"github.com/po2656233/superplace"
	"github.com/po2656233/superplace/components/cron"
	"github.com/po2656233/superplace/components/gin"
	"github.com/po2656233/superplace/extend/file"
	clog "github.com/po2656233/superplace/logger"
	"superman/internal/component/check_center"
	"superman/internal/conf"
	"superman/nodes/web/controller"
	"superman/nodes/web/middleware"
	"superman/nodes/web/sdk"
)

func Run(profileFilePath, nodeId string) {
	// 配置sgx引擎,加载profile配置文件
	app := superplace.Configure(profileFilePath, nodeId, false, superplace.Cluster)

	// 注册调度组件
	app.Register(superCron.New())

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

func httpServerComponent(addr string) *superGin.Component {
	gin.SetMode(gin.DebugMode)

	// new http server
	httpServer := superGin.NewHttp("http_server", addr)
	httpServer.Use(superGin.Cors())
	httpServer.Use(middleware.AuthRequired())
	// http server使用gin组件搭建，这里增加一个RecoveryWithZap中间件
	httpServer.Use(superGin.RecoveryWithZap(true))

	// 映射h5客户端静态文件到static目录，例如：http://127.0.0.1:8089/static/protocol.js
	httpServer.Static("/static", "./static/")

	// 加载./view目录的html模板文件(详情查看gin文档)
	viewFiles := exFile.WalkFiles("./view/", ".html")
	if len(viewFiles) < 1 {
		clog.Errorf("view files not found.")
	} else {
		httpServer.LoadHTMLFiles(viewFiles...)
	}

	//注册 controller
	httpServer.Register(new(controller.Controller))

	return httpServer
}
