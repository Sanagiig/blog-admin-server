package utils

import (
	"fmt"
	"runtime"
)

func GetErrorStack(preStr string, err error) string {
	pc, file, line, ok := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	if !ok {
		return "GetErrorStack 方法获取堆栈失败，返回错误原信息：" + err.Error()
	}
	if err == nil {
		return ""
	} else {
		errMsg := fmt.Sprintf("%s \n\tat %s:%d (Method %s)\nCause by: %s\n", preStr, file, line, f.Name(), err.Error())
		return errMsg
	}
}
