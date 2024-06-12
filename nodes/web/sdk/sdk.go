package sdk

import (
	sgxGin "github.com/po2656233/superplace/components/gin"
	sgxString "github.com/po2656233/superplace/extend/string"
	cfacade "github.com/po2656233/superplace/facade"
	sgxError "github.com/po2656233/superplace/logger/error"
	"sanguoxiao/internal/conf"
)

// sdk平台类型
const (
	DevMode  int32 = 1 // 开发模式，注册开发帐号登陆(开发时使用)
	QuickSDK int32 = 2 // quick sdk
)

var (
	invokeMaps = make(map[int32]Invoke)
)

type (
	Invoke interface {
		SdkId() int32                                                // sdk id
		Login(config *conf.SdkRow, params Params, callback Callback) // Login 登录验证接口
		PayCallback(config *conf.SdkRow, c *sgxGin.Context)          // PayCallback 支付回调接口
	}

	Params map[string]string

	Callback func(code int32, result Params, error ...error)
)

func (p Params) GetInt(key string, defaultValue ...int) int {
	defVal := 0
	if len(defaultValue) > 0 {
		defVal = defaultValue[0]
	}

	val, found := p[key]
	if !found {
		return defVal
	}

	intVal, ok := sgxString.ToInt(val)
	if ok {
		return intVal
	}

	return defVal
}

func (p Params) GetString(key string) (string, bool) {
	v, ok := p[key]
	return v, ok
}

func register(invoke Invoke) {
	invokeMaps[invoke.SdkId()] = invoke
}

func GetInvoke(sdkId int32) (invoke Invoke, error error) {
	invoke, found := invokeMaps[sdkId]
	if found == false {
		return nil, sgxError.Errorf("[sdkId = %d] not found.", sdkId)
	}

	return invoke, nil
}

func Init(app cfacade.IApplication) {
	register(devSdk{app})
	register(quickSdk{})
}
