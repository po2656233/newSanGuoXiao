package login

//对外暴露接口
import (
	"superman/nodes/leaf/jettengame/login/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
