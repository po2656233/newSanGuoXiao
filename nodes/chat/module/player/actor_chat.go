package player

import (
	"strings"
	cst "superman/internal/constant"
	"time"

	superConst "github.com/po2656233/superplace/const"
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
)

type (
	// ActorChat 每位登录的玩家对应一个子actor
	ActorChat struct {
		//pomelo.ActorBase
		simple.ActorBase
		childExitTime time.Duration
	}
)

func (p *ActorChat) AliasID() string {
	return strings.Trim(cst.ChatActor, superConst.DOT)
}
func (p *ActorChat) OnInit() {
	// 注册协议
	//p.registerLocalMsg()  // 注册(与客户端通信)的协议
	p.registerRemoteMsg() // 注意服务间交互的协议

	// 检查游戏基础信息(房间列表、牌桌列表)
	p.checkBaseInfo()

}

// registerRemoteMsg 注意服务间交互的协议
func (p *ActorChat) registerRemoteMsg() {
	p.Remote().Register(p.checkChild) // 与子节点交互
}

// checkBaseInfo 检查基础信息
func (p *ActorChat) checkBaseInfo() {
	log.Debugf("[ActorChat] path = %s init!", p.PathString())
	p.childExitTime = time.Minute * 30

	p.Timer().RemoveAll()
	p.Call(p.PathString(), "checkChild", nil)

}

func (p *ActorChat) OnFindChild(msg *cfacade.Message) (cfacade.IActor, bool) {
	// 动态创建 player child actor
	childID := msg.TargetPath().ChildID
	childActor, err := p.Child().Create(childID, &ActorPlayer{
		session:  msg.Session,
		isOnline: false,
	})

	if err != nil {
		return nil, false
	}

	return childActor, true
}

func (p *ActorChat) OnStop() {
	log.Debugf("[ActorChat] path = %s exit!", p.PathString())
}

func (p *ActorChat) checkChild() {
	// 扫描所有玩家actor
	p.Child().Each(func(iActor cfacade.IActor) {
		child, ok := iActor.(*ActorPlayer)
		if !ok || child.isOnline {
			return
		}

		// 玩家下线，并且超过childExitTime时间没有收发消息，则退出actor
		deadline := time.Now().Add(-p.childExitTime).Unix()
		if child.LastAt() < deadline {
			child.Exit() //actor退出
		}
	})
}

/////////////////////////////////////////////////////////////////
