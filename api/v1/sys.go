package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/message/errMsg"
	"go-blog/service"
	"go-blog/utils/result"
)

func InitPermision(c *gin.Context) {
	err := service.InitPermision()
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}
	result.SucJson(c, nil)
}

func InitDictionary(c *gin.Context) {
	err := service.InitDictionary()
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}
	result.SucJson(c, nil)
}
