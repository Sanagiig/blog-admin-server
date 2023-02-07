package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitUserRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("user")
	pubr := pub.Group("user")

	pubr.POST("/register", v1.Register)
	pubr.POST("/login", v1.Login)

	prir.GET("/findSelfInfo", v1.FindSelfInfo)
	prir.POST("/add", v1.AddUser)
	prir.POST("/update", v1.UpdateUser)
	prir.POST("/delete", v1.DeleteUser)
	prir.GET("/findDetail", v1.GetUserInfo)
	prir.GET("/findList", v1.FindUserList)
	prir.GET("/findPaging", v1.FindUserPaging)
	prir.GET("/patchAddRole", v1.PatchAddRole)
	prir.GET("/patchRemoveRole", v1.PatchRemoveRole)
}
