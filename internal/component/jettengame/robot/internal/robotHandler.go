package internal

import (
	"crypto/rand"
	"github.com/po2656233/goleaf/gate"
	"github.com/po2656233/goleaf/log"
	"math/big"
	"reflect"
	. "sanguoxiao/internal/component/jettengame/base"
	"sanguoxiao/internal/component/jettengame/manger"
	protoMsg "sanguoxiao/internal/protocol/gofile"
	"time"
)

func init() {

	//准备

	// 游戏退出(重新进入,营造进进出出的氛围)
	handlerMsg(&protoMsg.ExitGameResp{}, handleExitGameResp) //反馈--->主页信息
}

// 注册模块间的通信
func handlerMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// ////////////////////////////////////////////
func handleExitGameResp(args []interface{}) {
	m := args[0].(*protoMsg.ExitGameResp)
	a := args[1].(gate.Agent)
	if a.UserData() != nil {
		person := a.UserData().(*manger.Player)
		if person.UserID == m.UserID {
			if person.Level == Limit {
				log.Release("机器人:%v 账号:%v 已被限制进入游戏:[%v]\n\n", person.UserID, person.Account, m.GameID)
				return
			}
			manger.GetRobotManger().ExitOne(person.UserID)
			r1, _ := rand.Int(rand.Reader, big.NewInt(100))
			wait := time.Duration(r1.Int64()%6 + 3)
			time.AfterFunc(wait*time.Second, func() {
				rob := manger.GetRobotManger().EnterOne(m.GameID)
				if rob != nil {
					//log.Release("玩家:%v 账号:%v <<---进入游戏:[%v]\n\n", rob.UserID, rob.Account, m.GameID)
				} else {
					log.Release("机器人无法进入 <<---进入游戏:[%v]\n\n", m.GameID)
				}
			})
		}

	}
}
