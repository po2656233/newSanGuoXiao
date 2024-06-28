package internal

import (
	"github.com/po2656233/goleaf/module"
	"superman/nodes/leaf/jettengame/base"
	"superman/nodes/leaf/jettengame/conf"
	"superman/nodes/leaf/jettengame/manger"
	"superman/nodes/leaf/jettengame/sql/mysql"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	// 重置机器人 登出时间
	mysql.SqlHandle().ResetRobot()
	// 获取可用机器人
	robotList := mysql.SqlHandle().GetRobotsInfo(conf.RobotMax)
	num := conf.RobotMax - len(robotList)
	for i := 0; i < num; i++ {
		//注册机器人
		manger.GetRobotManger().RegisterOne()
	}

	for _, info := range robotList {
		item := &manger.Robot{
			Player: manger.ToPlayer(info),
		}
		item.Gold = item.Money
		// 添加至 用户队列
		manger.GetPlayerManger().Append(item.Player)
		// 添加至 机器人队列
		manger.GetRobotManger().Append(item)

	}
}

func (m *Module) OnDestroy() {

}
