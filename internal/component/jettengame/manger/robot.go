package manger

import (
	"github.com/po2656233/goleaf/log"
	"golang.org/x/net/websocket"
	"net"
	"net/url"
	. "sanguoxiao/internal/component/jettengame/base"
	"sanguoxiao/internal/component/jettengame/msg"
	protoMsg "sanguoxiao/internal/protocol/gofile"
	"sync"
)

// Robot 机器人
type Robot struct {
	*Player
}

// 接收游戏创建成功的通知，根据配置文件，向游戏中投放机器人。
// 机器人根据游戏逻辑完成下注过程
// ------------------------管理机器人--------------------------------------------------//

// RobotManger 管理机器人
type RobotManger struct {
	sync.Map
	sync.Mutex
} // == persons map[int64]*Player

var mangerR *RobotManger = nil
var onceR sync.Once

// GetRobotManger 机器人管理对象(单例模式)//manger.persons = make(map[int64]*Player)
func GetRobotManger() *RobotManger {
	onceR.Do(func() {
		mangerR = &RobotManger{
			Map: sync.Map{},
		}
	})
	return mangerR
}

// Append 添加机器人
func (robotSelf *RobotManger) Append(play *Robot) (*Robot, bool) {
	v, ok := robotSelf.Load(play.UserID)
	if !ok || v == nil {
		//		log.Debug("[新增一个玩家]\tID:%v", play.UserID)
		robotSelf.Store(play.UserID, play)
		return play, true
	}
	//	log.Debug("[已經存在玩家]\tID:%v ", play.UserID)
	return v.(*Robot), false
}

// Get 获取指定机器人[根据索引,即userID]
func (robotSelf *RobotManger) Get(userID int64) *Robot {
	value, ok := robotSelf.Load(userID)
	if ok {
		return value.(*Robot)
	}
	return nil
}

// DeleteIndex 按索引删除机器人
func (robotSelf *RobotManger) DeleteIndex(i int64) {
	robotSelf.Delete(i)
}

// DeletePlayer 删除机器人
func (robotSelf *RobotManger) DeletePlayer(play *Robot) {
	index := int64(0)
	robotSelf.Range(func(key, value interface{}) bool {
		if key.(int64) == play.UserID {
			//log.Debug("[删除]\t找到要删除的机器人:%v ", play.UserID)
			index = key.(int64)
			value = nil
			return false
		}
		return true
	})
	robotSelf.Delete(index)
	play = nil
}

func enterGame(robot *Robot, room *Room, table *Table, game *Game) bool {
	//添加座位成功 再进入
	if !game.AddChair(robot.Player) {
		return false
	}
	robot.Player.State = protoMsg.PlayerState_PlayerSitDown
	robot.Player.Level = game.G.Level
	robot.Player.GameID = game.ID
	robot.Player.PtrTable = table
	robot.Player.PtrRoom = room
	robot.Player.RoomNum = room.Num

	agent := &ClientAgent{
		userData: robot.Player,
	}

	//添加 虚拟的客户端与用户ID 绑定
	GetClientManger().Append(robot.UserID, agent)

	// 进入游戏
	var args []interface{}
	args = append(args, game, agent)
	robot.Enter(args)

	// 更新玩家信息,并通知其他玩家在线
	args = make([]interface{}, 0)
	args = append(args, int32(GameUpdateEnter), robot.UserID)
	table.Instance.UpdateInfo(args)

	return true
}

// EnterOne 删除机器人
func (robotSelf *RobotManger) EnterOne(gid int64) *Robot {
	robotSelf.Lock()
	defer robotSelf.Unlock()
	var rob *Robot = nil
	play, tb, ret := GetPlayerManger().CheckTable(gid)
	if !ret {
		return rob
	}

	game, ok1 := GetGamesManger().GetGame(gid)
	if !ok1 {
		return rob
	}
	robotSelf.Range(func(key, value interface{}) bool {
		robot, ok := value.(*Robot)
		if !ok {
			return true
		}
		if robot.GameID == INVALID &&
			(protoMsg.PlayerState_PlayerLookOn == robot.State || protoMsg.PlayerState_PlayerStandUp == robot.State ||
				protoMsg.PlayerState_PlayerGiveUp == robot.State) {
			if enterGame(robot, play.PtrRoom, tb, game) {
				rob = robot
				return false
			}
		}
		return true
	})
	return rob
}

// ExitOne 指定机器人退出游戏
func (robotSelf *RobotManger) ExitOne(uid int64) {
	robotSelf.Range(func(key, value interface{}) bool {
		if robot, ok := value.(*Robot); ok && robot.UserID == uid {
			robot.GameID = INVALID
			robot.State = protoMsg.PlayerState_PlayerStandUp
			robotSelf.Store(uid, robot)
			return false
		}
		return true
	})
}

func (robotSelf *RobotManger) RegisterOne() {

	rob := &Robot{
		Player: &Player{
			PlayerInfo: &protoMsg.PlayerInfo{
				Name:       GetFullName(),
				Sex:        0x0F,
				PlatformID: 1,
			},
		},
	}
	msgReg := &protoMsg.RegisterReq{
		Name:       rob.Name,
		Gender:     rob.Sex,
		Password:   "rob",
		PlatformID: 1,
	}
	client := &ClientAgent{
		userData: rob,
	}
	_ = msg.ProcessorProto.Route(msgReg, client)
}

////////////////////////////////////////////////////////

func (robotSelf *RobotManger) Enter(room *Room, table *Table, game *Game) (int32, bool) {
	// 如果已经有机器人则不再添加 game.HaveRobot ||
	if game.T.RobotCount == 0 || game.T.HostID != SYSTEMID {
		return 0, false
	}
	robotSelf.Lock()
	defer robotSelf.Unlock()
	// 先尝试退出游戏
	//robotSelf.ExitGame(game.ID)

	// 修正机器人数量 不超过座椅数-1
	robotCount := game.T.RobotCount
	if 0 < game.T.MaxChair && game.T.MaxChair <= robotCount {
		robotCount = game.T.MaxChair - 1
	}

	total := int32(0)
	// 取符合条件的机器人进入游戏
	robotSelf.Range(func(key, value interface{}) bool {
		robot, ok := value.(*Robot)
		if !ok {
			return true
		}
		if robot.GameID == game.ID && robot.Player.Level != Limit {
			if robotCount <= total {
				// 超出指定数量,则机器人站起 等待离开
				robot.Player.Level = Limit
				log.Release("机器人:%v 即将离开 游戏(ID):%v 当前机器人数量:%v", robot.Name, game.ID, total)
				return true
			}
			total++
			log.Release("已经 机器人:%v 已在 游戏(ID):%v 当前机器人数量:%v", robot.Name, game.ID, total)

		}
		return true
	})

	if robotCount <= total {
		return robotCount, true
	}

	// 还差多少
	robotCount = robotCount - total
	newTotal := int32(0)
	robotSelf.Range(func(key, value interface{}) bool {
		robot, ok := value.(*Robot)
		if !ok {
			return true
		}
		if robotCount <= newTotal {
			return false
		}
		if robot.Level == Limit && robot.GameID == game.ID && enterGame(robot, room, table, game) {
			newTotal++
			log.Release("重返游戏 机器人:%v 进入 游戏(ID):%v 当前机器人数量:%v", robot.Name, game.ID, total+newTotal)
		}
		return true
	})

	if robotCount <= newTotal {
		return total + newTotal, true
	}

	// 还差多少
	total += newTotal
	robotCount = robotCount - newTotal
	newTotal = int32(0)
	robotSelf.Range(func(key, value interface{}) bool {
		robot, ok := value.(*Robot)
		if !ok {
			return true
		}
		if robotCount <= newTotal {
			return false
		}
		if robot.GameID == INVALID && enterGame(robot, room, table, game) {
			newTotal++
			log.Release("第一次 机器人:%v 进入 游戏(ID):%v 当前机器人数量:%v", robot.Name, game.ID, total+newTotal)
		}
		return true
	})

	//total = robotCount
	log.Release("当前 进入 游戏人数：%v", total+newTotal)
	return total + newTotal, robotCount <= newTotal
}

// ExitGame 退出指定游戏
func (robotSelf *RobotManger) ExitGame(gameId int64) {
	count := 0
	robotSelf.Range(func(key, value interface{}) bool {
		if robot, ok := value.(*Robot); ok && robot.GameID == gameId {
			robot.Player.Level = Limit
			robot.State = protoMsg.PlayerState_PlayerStandUp
			robot.GameID = INVALID
			count++
			//message := &protoMsg.ExitGameReq{
			//	GameID: gameId,
			//}
			//// 发给游戏服
			//if agent, ok1 := GetClientManger().Get(robot.UserID); ok1 {
			//	_ = msg.ProcessorProto.Route(message, agent)
			//	count++
			//	//玩家退出游戏
			//}
		}
		return true
	})
	log.Release("当前 退出 游戏人数：%v", count)
}

//////////////////////////构造虚拟agent//////////////////////////////////////

type ClientAgent struct {
	userData interface{}
}

func (a *ClientAgent) Run() {

}

func (a *ClientAgent) OnClose() {

}

func (a *ClientAgent) WriteMsg(data interface{}) {
	//
	//log.Debug("机器人[%v] 接收信息:%v", a.userData, msg)
	err := msg.ProcessorProto.Route(data, a)
	if err != nil {
		log.Error("机器人[%v] 接收信息:%v err:%v", a.userData, data, err)
	} else {
		//log.Debug("机器人[%v] 接收信息 ok", a.userData)
	}
}

func (a *ClientAgent) LocalAddr() net.Addr {
	return &websocket.Addr{
		URL: &url.URL{
			Host: "127.0.0.1",
		},
	}
}

func (a *ClientAgent) RemoteAddr() net.Addr {
	return &websocket.Addr{
		URL: &url.URL{
			Host: "127.0.0.1",
		},
	}
}

func (a *ClientAgent) Close() {
}

func (a *ClientAgent) Destroy() {
}

func (a *ClientAgent) UserData() interface{} {
	return a.userData
}

func (a *ClientAgent) SetUserData(data interface{}) {
	a.userData = data
}
