package result

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-blog/global"
	"go-blog/message"
	"net/http"
)

func SucJson(c *gin.Context, data any) {
	code := message.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": message.GetMsg(code),
	})
}

func QueryTableJson(c *gin.Context, total int64, data any) {
	code := message.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": gin.H{
			"total": total,
			"list":  data,
		},
		"message": message.GetMsg(code),
	})
}

func ErrJson(c *gin.Context, code int, data any, err any) {
	var errStr string
	switch err.(type) {
	case error:
		if err != nil {
			errStr = err.(error).Error()
			global.Log.WithFields(logrus.Fields{
				"Remote": c.RemoteIP(),
				"Method": c.Request.Method,
				"Path":   c.FullPath(),
			}).Error(errStr)
		}
	case string:
		errStr = err.(string)
	default:
		fmt.Println("========================ErrJson nil")
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"errMsg":  errStr,
		"message": message.GetMsg(code),
	})
}

func ArgErrJson(c *gin.Context, code int, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
		global.Log.Error(errMsg)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    nil,
		"errMsg":  errMsg,
		"message": message.GetMsg(code),
	})
}
