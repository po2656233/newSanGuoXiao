package jettengame

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/po2656233/goleaf/cluster"
	lconf "github.com/po2656233/goleaf/conf" //该包的init函数会被调用
	"github.com/po2656233/goleaf/console"
	"github.com/po2656233/goleaf/log"
	"github.com/po2656233/goleaf/module"
	sgxFacade "github.com/po2656233/superplace/facade"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/trace"
	"superman/nodes/leaf/jettengame/conf"
	"superman/nodes/leaf/jettengame/game"
	"superman/nodes/leaf/jettengame/gate"
	"superman/nodes/leaf/jettengame/login"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

const (
	pprofAddr string = ":7890"
	version          = "1.1.4"
)

func StartHTTPDebug() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: pprofAddr, Handler: pprofHandler}
	go server.ListenAndServe()
}

// Component 启动时,检查center节点是否存活
type Component struct {
	sgxFacade.Component
}

func New() *Component {
	return &Component{}
}

func (c *Component) Name() string {
	return "leaf"
}

func (c *Component) OnAfterInit() {
	//确保并发执行
	runtime.GOMAXPROCS(runtime.NumCPU())
	//go func() {
	//	http.ListenAndServe("localhost:9000", nil)
	//}()
	////性能测试
	////cpu
	//cpuProfile, _ := os.Create("cpu_profile")
	//pprof.StartCPUProfile(cpuProfile)
	//defer pprof.StopCPUProfile()
	////内存
	//memProfile, _ := os.Create("mem_profile")
	//pprof.WriteHeapProfile(memProfile)

	//StartHTTPDebug()
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	// 集群[弃用]
	//lconf.ListenAddr = "127.0.0.1:10010"
	//lconf.ConnAddrs = []string{"127.0.0.1:10010"}
	//lconf.PendingWriteNum = 1

	// logger
	if lconf.LogLevel != "" {
		logger, err := log.New(lconf.LogLevel, lconf.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
		defer logger.Close()
	}
	log.Release("Leaf :%v  starting up", version)

	mods := make([]module.Module, 0)
	mods = append(mods, gate.Module, login.Module, game.Module)
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	//对应的模块实现
	//操作步骤：
	//1、msg 构建传输协议（消息体）
	//2、gate 的router路由分发数据
	//3、game 的handle下进行消息处理

	//---------骨架---------
	//gate 模块，负责游戏客户端的接入。
	//login 模块，负责登录流程。
	//game 模块，负责游戏主逻辑。

	//---------平台（大厅）---------
	//task 模块，负责任务活动。领取与完成
	//notice 模块，负责游戏通告。平台通告、房间通告、邮箱通告。
	//spread 模块，负责游戏推广(分享)。
	//statistics 模块，统计面板。玩家输赢后的平台掉星或加星 亦或加减分

	//---------玩家---------
	//stageBag 模块， 玩家道具背包。
	//chat 模块，负责用户聊天。大厅谈论，弹幕，私聊。
	//Recharge 模块，与在线充值，负责在线充值 |customer service 在线客服   充值成功接口,通知金币变化
	//live 模块

	//--------本地信息---------
	//log 模块，负责日志，记录节点信息。
	//sql || database 模块，负责读写数据库。

	//一: 断线重连。 游戏时，保持长连接状态; 在游戏没结束，玩家都可以重新连接进来，并且处于托管状态。
	//二: 定时器。
	//三：事件派发。
	//配置信息是通过server.json配置的
}

func (c *Component) OnStop() {
	log.Release("Leaf closing down!")
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
