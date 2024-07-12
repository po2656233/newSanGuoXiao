package ops

import (
	"github.com/po2656233/superplace/const/code"
	cactor "github.com/po2656233/superplace/net/actor"
	"superman/internal/protocol/gofile"
)

var (
	pingReturn = &pb.Bool{Value: true}
)

type (
	ActorOps struct {
		cactor.Base
	}
)

func (p *ActorOps) AliasID() string {
	return "ops"
}

// OnInit 注册remote函数
func (p *ActorOps) OnInit() {
	p.Remote().Register(p.Ping)
}

// Ping 请求center是否响应
func (p *ActorOps) Ping() (*pb.Bool, int32) {
	return pingReturn, code.OK
}
