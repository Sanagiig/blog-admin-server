package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitTagRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("tag")
	//pubr := pub.Group("tag")

	prir.POST("/add", v1.AddTag)
	prir.POST("/update", v1.UpdateTag)
	prir.POST("/delete", v1.DeleteTag)
	prir.GET("/find", v1.FindTag)
	prir.POST("/findPaging", v1.FindTagPaging)
	prir.POST("/findList", v1.FindTagList)
}
