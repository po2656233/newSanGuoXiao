package base

import (
	"github.com/po2656233/goleaf/chanrpc"
	"github.com/po2656233/goleaf/module"
	"superman/nodes/leaf/jettengame/conf"
)

// 骨架
func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,                         //go的可使用数目
		TimerDispatcherLen: conf.TimerDispatcherLen,            //定时器分配数
		AsynCallLen:        conf.AsynCallLen,                   //异步调用的数目
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen), //模块间使用 ChanRPC 通讯，消息路由也不例外
	}
	skeleton.Init()
	return skeleton
}
