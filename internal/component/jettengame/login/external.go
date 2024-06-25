package login

//对外暴露接口
import (
	"sanguoxiao/internal/component/jettengame/login/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)
