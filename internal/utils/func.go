package utils

// 对于函数的相关操作
import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetFuncName 获取函数名
func GetFuncName(layer int) string {
	//参数 layer 函数所在的层数
	if pc, _, _, ok := runtime.Caller(layer + 1); ok {
		funcName := runtime.FuncForPC(pc).Name()
		funcs := strings.Split(funcName, ".")
		size := len(funcs)
		if size-1 >= 0 {
			return funcs[size-1]
		}
	}
	return ""
}

// GoID 获取goroutine的id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

// DumpStacks 堆栈信息输出
func DumpStacks() []byte {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	//fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
	return buf
}
