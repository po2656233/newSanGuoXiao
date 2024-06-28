package robot

//对外暴露接口
import (
	"superman/nodes/leaf/jettengame/robot/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
