package internal

import (
	"github.com/po2656233/goleaf/module"
	"sanguoxiao/internal/component/jettengame/base"
	"sanguoxiao/internal/component/jettengame/manger"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer //模块之间可以通过这个进行交互
	//AsyncChan = chanrpc.NewServer(10)  // 测试异步chan用
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	manger.GetGamesManger().UpCaches()
}

func (m *Module) OnDestroy() {
	//PlayerManager.Close()
	//mgodb.Close()
	//log.Release("closed")
	manger.GetGamesManger().DownCaches()
	//mysql.SqlHandle().CloseMysql()

}
