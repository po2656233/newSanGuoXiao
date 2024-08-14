package queue

import (
	superConst "github.com/po2656233/superplace/const"
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	"strings"
	cst "superman/internal/constant"
)

type (
	// ActorMatch 每位登录的玩家对应一个子actor
	ActorMatch struct {
		//pomelo.ActorBase
		simple.ActorBase
	}
)

func (p *ActorMatch) AliasID() string {
	return strings.Trim(cst.MatchActor, superConst.DOT)
}
func (p *ActorMatch) OnInit() {
	log.Debugf("[ActorMatch] path = %s init!", p.PathString())

}

func (p *ActorMatch) OnStop() {
	log.Debugf("[ActorMatch] path = %s exit!", p.PathString())
}
func (p *ActorMatch) OnFindChild(msg *cfacade.Message) (cfacade.IActor, bool) {
	// 动态创建 player child actor
	childID := msg.TargetPath().ChildID
	childActor, err := p.Child().Create(childID, &ActorQueue{
		isOnline: false,
	})

	if err != nil {
		return nil, false
	}

	return childActor, true
}
