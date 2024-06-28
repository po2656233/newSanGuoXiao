package internal

import (
	"github.com/po2656233/goleaf/gate"
	"superman/nodes/leaf/jettengame/conf"
	"superman/nodes/leaf/jettengame/game"
	"superman/nodes/leaf/jettengame/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		KCPAddr:         conf.Server.KCPAddr,
		WSAddr:          conf.Server.WSAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.ProcessorProto, //消息处理器对象(proto|json)
		AgentChanRPC:    game.ChanRPC,       // 极其重要
	}
}

// Run 启动后,注册服务至网关(中心服)
func (m *Module) Run(closeSig chan bool) {
	//log.Release("[gate] 正在注册服务...")
	//_ = conf.ToCenterServer(conf.RegisterPath, conf.MethodWeb, &conf.Server.WSAddr)
	//_ = conf.ToCenterServer(conf.RegisterPath, conf.MethodTcp, &conf.Server.TCPAddr)
	// 建议 kcp直连,可以不必注册到中心服
	//_ = conf.ToCenterServer(conf.RegisterPath, conf.MethodKcp, &conf.Server.KCPAddr)
	m.Gate.Run(closeSig)
}

func (m *Module) OnDestroy() {
	//log.Release("[gate] 正在移除服务")
	//_ = conf.ToCenterServer(conf.RemovePath, conf.MethodWeb, &conf.Server.WSAddr)
	//_ = conf.ToCenterServer(conf.RemovePath, conf.MethodTcp, &conf.Server.TCPAddr)
	//_ = conf.ToCenterServer(conf.RemovePath, conf.MethodKcp, &conf.Server.KCPAddr)
	m.Gate.OnDestroy()
}
