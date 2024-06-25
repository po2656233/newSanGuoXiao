package internal

import (
	. "sanguoxiao/internal/component/jettengame/base"
	. "sanguoxiao/internal/component/jettengame/manger"
	"sanguoxiao/internal/component/jettengame/sql/mysql"
	"sync/atomic"

	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
)

//广播消息
//这里是对所有玩家进行通知，通知单个游戏的所有玩家，请在单个游戏里实现

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("OffLine", rpcOfflineAgent)
	//清场
	//AsyncChan.Register("clearUp", func(args []interface{}) {
	//
	//	log.Debug("clearUp:%v", args)
	//	//table := args[0].(*Table)
	//
	//}) // 广播消息 调用参考:game.ChanRPC.Go("Broadcast", agent, args)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent) //【模块间通信】跟路由之间的通信
	//GetClientManger().Append(INVALID, a)
	if nil != a.UserData() {
		atomic.AddInt32(&OnlineCount, 1)
		log.Release("总在线人数: %v (新增) 信息：%+v", atomic.LoadInt32(&OnlineCount), a.UserData())
	}
	_ = a

}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a

	//断线通知
	userData := a.UserData()
	if nil != userData {
		atomic.AddInt32(&OnlineCount, -1)
		if person, ok := userData.(*Player); ok && nil != person {
			log.Release("[玩家离开]玩家:%v 总在线人数:%v  IP:%v", person.UserID, atomic.LoadInt32(&OnlineCount), a.RemoteAddr())
			// 平台维度
			platform := GetPlatformManger().Get(person.PlatformID)
			// 房间维度
			if nil != platform && nil != person.PtrTable && nil != person.PtrTable.Instance {
				var updateArgs []interface{}
				updateArgs = append(updateArgs, int32(GameUpdateOut), person.UserID, a)
				person.PtrTable.Instance.UpdateInfo(updateArgs)
			}

			// 玩家维度
			log.Debug("[玩家离开]\t->清理 玩家:%v(IP:%v) 数据", person.UserID, a.RemoteAddr())
			mysql.SqlHandle().UpdateLeaveTime([]int64{person.UserID})
			if person.GameID == INVALID {
				GetPlayerManger().DeletePlayer(person)
				log.Release("[玩家离开][移除玩家]\t玩家ID:%v 不在任何游戏中,且掉线.", person.UserID)
			}

			GetClientManger().DeleteClient(person.UserID)
		} else {
			log.Error("[玩家离开][error]\t找不到对应的玩家数据:%v", a.RemoteAddr())
		}
		return
	}
	a.SetUserData(nil)
	a.Destroy()
	//GetClientManger().DeleteClient(INVALID)

}

func rpcOfflineAgent(args []interface{}) {
	//断线通知
	log.Error("[args] 玩家数据:%v", args)
}
