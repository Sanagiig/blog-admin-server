package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitPermissionRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("permission")
	//pubr := pub.Group("permission")

	prir.POST("/add", v1.AddPemission)
	prir.POST("/update", v1.UpdatePemission)
	prir.POST("/delete", v1.DeletePemission)
	prir.GET("/find", v1.FindPemission)
	prir.POST("/findPaging", v1.FindPemissionPaging)
	prir.POST("/findList", v1.FindPemissionList)
}
