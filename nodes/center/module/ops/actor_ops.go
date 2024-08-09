package ops

import (
	superConst "github.com/po2656233/superplace/const"
	cactor "github.com/po2656233/superplace/net/actor"
	"strings"
	. "superman/internal/constant"
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
	return strings.Trim(OpsActor, superConst.DOT)
}

// OnInit 注册remote函数
func (p *ActorOps) OnInit() {
	p.Remote().Register(p.Ping)
}

// Ping 请求center是否响应
func (p *ActorOps) Ping() (*pb.Bool, int32) {
	return pingReturn, SUCCESS
}
