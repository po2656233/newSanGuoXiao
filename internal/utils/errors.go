package utils

import clog "github.com/po2656233/superplace/logger"

// CheckError 异常处理
func CheckError(err error) bool {
	if err != nil {
		//fmt.Println(err.Error())
		clog.Error("数据异常:%v", err) // 正式开服的时候，这里一定是错误类型 错误内容 等异常信息输出到日志上面的
		return false
	}
	return true
}
