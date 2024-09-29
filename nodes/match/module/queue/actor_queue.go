package queue

import (
	"context"
	"encoding/json"
	"github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	log "github.com/po2656233/superplace/logger"
	"github.com/po2656233/superplace/net/parser/simple"
	cproto "github.com/po2656233/superplace/net/proto"
	"google.golang.org/protobuf/proto"
	"math/rand"
	. "superman/internal/constant"
	protoMsg "superman/internal/protocol/go_file/common"
	gameMsg "superman/internal/protocol/go_file/game"
	gateMsg "superman/internal/protocol/go_file/gate"
	"superman/internal/redis_cluster"
	"superman/internal/rpc"
	"superman/nodes/game/module/online"
	//"superman/nodes/game/msg"
)

type (
	// ActorQueue 每位登录的玩家对应一个子actor
	ActorQueue struct {
		//pomelo.ActorBase
		simple.ActorBase
		isOnline bool
		Session  *cproto.Session
		userData interface{}
	}
)

func (p *ActorQueue) OnInit() {
	// 注册 session关闭的remote函数(网关触发连接断开后，会调用RPC发送该消息)
	p.Remote().Register(p.sessionClose)
	p.Local().Register(p.joinAllReadyQueue)
}

// sessionClose 接收角色session关闭处理
func (p *ActorQueue) sessionClose() {
	p.isOnline = false
	p.Exit()
}
func (p *ActorQueue) OnStop() {
	log.Infof("OnStop onlineCount = %d", online.Count())
	if p.Session == nil {
		log.Warnf("OnStop ")
		return
	}
}

func (p *ActorQueue) joinAllReadyQueue(session *cproto.Session, req *gameMsg.JoinAllReadyQueueReq) {
	// 从redis中查找评分最高分的房间作为可选项 一般房间仅差一个人满座的，则评分最高
	list, err := redis_cluster.SingleRedis().GetTopCountMembers(context.Background(), GetMatchKey(req.GameID, req.RoomID), INVALID)
	if err != nil {
		log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] 加入牌桌失败! err:%v", req.GameID, req.RoomID, err)
		p.Response(session, &gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Game51],
		})
		return
	}
	size := len(list)
	if 0 == size {
		log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] 加入牌桌失败! no enough", req.GameID, req.RoomID)
		p.Response(session, &gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Game51],
		})
		return
	}
	// 取第一个检测是否允许无限座次
	val, ok := list[0].(string)
	if !ok {
		log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] 加入牌桌失败! no value", req.GameID, req.RoomID)
		p.Response(session, &gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Game51],
		})
		return
	}
	info := redis_cluster.RedisMatchInfo{}
	if err = json.Unmarshal([]byte(val), &info); err != nil {
		log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] info:%+v err:%v", req.GameID, req.RoomID, info, err)
		p.Response(session, &gateMsg.ResultResp{
			State: FAILED,
			Hints: StatusText[Game51],
		})
		return
	}
	// 如果牌桌数不止一个
	if 1 < size {
		// 不受限制,则随机获取
		if info.MaxSit == Unlimited {
			index := rand.Int() % size
			val, ok = list[index].(string)
			if !ok {
				log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] 加入牌桌失败! no value", req.GameID, req.RoomID)
				p.Response(session, &gateMsg.ResultResp{
					State: FAILED,
					Hints: StatusText[Game51],
				})
				return
			}
			err = json.Unmarshal([]byte(val), &info)
			if err != nil {
				log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] err:%v", req.GameID, req.RoomID, err)
				p.Response(session, &gateMsg.ResultResp{
					State: FAILED,
					Hints: StatusText[Game51],
				})
				return
			}
		} else {
			// 找出少于 第一个已座人数 的牌桌
			for i := 1; i < size; i++ {
				val1, ok1 := list[i].(string)
				if !ok1 {
					continue
				}
				info1 := redis_cluster.RedisMatchInfo{}
				err = json.Unmarshal([]byte(val1), &info1)
				if err != nil {
					log.Errorf("[ActorQueue] 游戏[%v] 房间[%v] err:%v", req.GameID, req.RoomID, err)
					continue
				}
				// 已经坐满人的排除
				if info1.MaxSit == info1.SitCount {
					continue
				}
				// 找到最小的为止
				if info.SitCount < info1.SitCount {
					val = val1
					info = info1
				}
			}
		}

	}

	log.Infof("[ActorQueue] 游戏[%v] 房间[%v] 牌桌信息:%+v", req.GameID, req.RoomID, info)
	quest := &protoMsg.Request{
		//Sid:   session.Sid,
		//Route: session.ActorPath(),
		Data: make([]byte, 0),
	}
	msgData, _ := rpc.GetProtoData(req)
	quest.Data = append(quest.Data, msgData...)
	childId := exString.ToString(session.Uid)
	data, _ := proto.Marshal(quest)
	clusterPacket := cproto.GetClusterPacket()
	clusterPacket.SourcePath = session.AgentPath
	clusterPacket.TargetPath = cfacade.NewChildPath(info.Sid, info.Route, childId)
	clusterPacket.FuncName = FuncRequest
	clusterPacket.Session = session // agent session
	clusterPacket.ArgBytes = data   // packet -> message -> data

	if err = p.App().Cluster().PublishLocal(info.Sid, clusterPacket); err != nil {
		log.Errorf("[ActorQueue][gameNodeRoute] ClusterLocalDataRoute err:%v", err)
	}
}
