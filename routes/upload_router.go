package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitUploadRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("upload")
	pubr := pub.Group("upload")

	//pubr.POST("/register", v1.Register)
	//pubr.POST("/login", v1.Login)

	pubr.GET("/getAuth", v1.GetUploadAuth)
	pubr.GET("/refreshCDN", v1.RefreshCDN)
	prir.POST("/file/:type", v1.UploadFile)
}
