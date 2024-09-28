package player

import (
	superConst "github.com/po2656233/superplace/const"
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	"strings"
	cst "superman/internal/constant"
	"superman/internal/rpc"
	mgr "superman/nodes/game/manger"
	"superman/nodes/game/module/category"
	"time"
)

type (
	// ActorGame 每位登录的玩家对应一个子actor
	ActorGame struct {
		//pomelo.ActorBase
		simple.ActorBase
		childExitTime time.Duration
	}
)

func (p *ActorGame) AliasID() string {
	return strings.Trim(cst.GameActor, superConst.DOT)
}
func (p *ActorGame) OnInit() {
	// 加载游戏配置
	category.InitConfig()
	// 加载协议文件
	rpc.LoadMsgInfos()
	// 客户端管理实例句柄
	mgr.GetClientMgr().SetApp(p.App())

	// 注册协议
	//p.registerLocalMsg()  // 注册(与客户端通信)的协议
	p.registerRemoteMsg() // 注意服务间交互的协议

	// 检查游戏基础信息(房间列表、牌桌列表)
	p.checkBaseInfo()

}

// registerRemoteMsg 注意服务间交互的协议
func (p *ActorGame) registerRemoteMsg() {
	p.Remote().Register(p.checkChild) // 与子节点交互
}

// checkBaseInfo 检查基础信息
func (p *ActorGame) checkBaseInfo() {
	log.Debugf("[ActorGame] path = %s init!", p.PathString())
	p.childExitTime = time.Minute * 30

	p.Timer().RemoveAll()
	//p.Call(p.PathString(), "checkChild", nil)

}

func (p *ActorGame) OnFindChild(msg *cfacade.Message) (cfacade.IActor, bool) {
	// 动态创建 player child actor
	childID := msg.TargetPath().ChildID
	childActor, err := p.Child().Create(childID, &ActorPlayer{
		isOnline: false,
	})

	if err != nil {
		return nil, false
	}

	return childActor, true
}

func (p *ActorGame) OnStop() {
	log.Debugf("[ActorGame] path = %s exit!", p.PathString())
}

func (p *ActorGame) checkChild() {
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
