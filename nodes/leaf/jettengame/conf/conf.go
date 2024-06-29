package conf

import (
	"log"
	"time"
)

var (
	// LogFlag log conf
	LogFlag = log.LstdFlags | log.Lshortfile | log.Lmsgprefix //Llongfile Lshortfile:输出文件名和行号 LstdFlags：标准输出

	// PendingWriteNum gate conf
	PendingWriteNum        = 10000   //挂起数目
	MaxMsgLen       uint32 = 1 << 16 //包体最大长度64k
	MinMsgLen              = 1
	HTTPTimeout            = 15 * time.Second //HTTP网络延迟15秒
	LenMsgLen              = 2                //所占字节长度
	LittleEndian           = false            //小端模式
	//UseMsgJsonProcessor        = true

	// GoLen skeleton conf
	GoLen              = 111000 //go在骨架（skeleton）中可使用的数目
	TimerDispatcherLen = 1000   //定时器分配数
	AsynCallLen        = 10000  //异步调用的数目
	ChanRPCLen         = 100000 //信道RPC的数目
	RobotMax           = 2000   // RobotMax

	ServerName       = "jetten" // 当前服务名
	ProtoPackageName = "pb"     // proto的包名
	ServerJsonPath   = "config/leafconf/server.json"
	NacosJsonPath    = "config/leafconf/nacos.json"
	GameYamlPath     = "config/leafconf/game.yaml"
)
