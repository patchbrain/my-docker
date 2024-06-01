package tool

import (
	"runtime"
)
// GetCurrentFuncName 返回当前执行函数的名称
func GetCurrentFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}

	// 获取函数详细信息
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}

	return fn.Name()
}