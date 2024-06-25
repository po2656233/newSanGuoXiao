package robot

//对外暴露接口
import (
	"sanguoxiao/internal/component/jettengame/robot/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
