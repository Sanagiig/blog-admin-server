package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitArticleRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("article")
	pubr := pub.Group("article")

	pubr.GET("/find", v1.FindArticle)
	pubr.POST("/findPaging", v1.FindArticlePaging)
	pubr.POST("/findList", v1.FindArticleList)

	prir.POST("/add", v1.AddArticle)
	prir.POST("/update", v1.UpdateArticle)
	prir.POST("/delete", v1.DeleteArticle)

}
