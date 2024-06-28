package superman

import (
	clog "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/pomelo"
)

type (
	// ActorSanGuoXiao 三国消
	ActorSanGuoXiao struct {
		pomelo.ActorBase
	}
)

func (p *ActorSanGuoXiao) OnInit() {
	clog.Debugf("[actorPlayer] path = %s init!", p.PathString())

	//p.Local().Register(p.playerSelect) // 注册 查看角色
	//p.Local().Register(p.playerCreate) // 注册 创建角色
	//p.Local().Register(p.playerEnter)  // 注册 进入角色
}

func (p *ActorSanGuoXiao) OnStop() {
	clog.Debugf("[actorPlayer] path = %s exit!", p.PathString())
}
