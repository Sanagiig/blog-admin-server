package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitSysRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("sys")
	//pubr := pub.Group("tag")

	prir.POST("/initPermission", v1.InitPermision)
	prir.POST("/initDictionary", v1.InitDictionary)
}
