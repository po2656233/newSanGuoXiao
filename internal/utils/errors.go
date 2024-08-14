package utils

import (
	exUtils "github.com/po2656233/superplace/extend/utils"
	clog "github.com/po2656233/superplace/logger"
)

// CheckError 异常处理
func CheckError(err error) (isOk bool) {
	isOk = true
	if err != nil {
		isOk = false
	}
	exUtils.Try(func() {
		if !isOk {
			clog.Errorf("err:%v", err)
			panic(err)
		}
	}, func(errString string) {
		// 上层的上上层级 调试后可得7
		typeName, fnName := GetFuncName(7)
		clog.Errorf("[%s]%s err:%s", fnName, typeName, errString)
	})
	return isOk
}
