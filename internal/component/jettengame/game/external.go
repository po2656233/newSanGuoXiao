package game

//对外暴露接口
import (
	"sanguoxiao/internal/component/jettengame/game/internal"
)

var (
	// 实例化 game 模块
	Module = new(internal.Module)
	//暴露ChanRPC
	ChanRPC = internal.ChanRPC

	//AsyncChan = internal.AsyncChan
	GameActor = &internal.ActorGame{}
)
