package message

import (
	"fmt"
	"go-blog/message/errMsg"
)

const (
	SUCCESS = 200
)

var codeMsg = map[int]string{}

func init() {
	errMsg.InitErrMsg(codeMsg)
	InitMessage(codeMsg)
}

func InitMessage(m map[int]string) {
	m[SUCCESS] = "OK"
}

func GetMsg(code int) string {
	msg, ok := codeMsg[code]
	if !ok {
		fmt.Printf("message codeMsg[%s] get nil val\n")
	}

	return msg
}
