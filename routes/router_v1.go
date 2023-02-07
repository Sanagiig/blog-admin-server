package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	pubV1 := pub.Group("api/")
	privateV1 := private.Group("api/")

	InitSysRouter(pubV1, privateV1)
	InitDictionaryRouter(pubV1, privateV1)
	InitTagRouter(pubV1, privateV1)
	InitUserRouter(pubV1, privateV1)
	InitPermissionRouter(pubV1, privateV1)
	InitRoleRouter(pubV1, privateV1)
	InitCategoryRouter(pubV1, privateV1)
	InitArticleRouter(pubV1, privateV1)
	InitUploadRouter(pubV1, privateV1)
}
