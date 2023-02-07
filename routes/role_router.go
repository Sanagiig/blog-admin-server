package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitRoleRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("role")
	//pubr := pub.Group("role")

	prir.POST("/add", v1.AddRole)
	prir.POST("/update", v1.UpdateRole)
	prir.POST("/delete", v1.DeleteRole)
	prir.GET("/find", v1.FindRole)
	prir.POST("/findPaging", v1.GetRolePaging)
	prir.POST("/findList", v1.GetRoleList)
	prir.POST("/patchAddPermission", v1.PatchAddPermission)
	prir.POST("/patchRemovePermission", v1.PatchRemovePermission)
}
