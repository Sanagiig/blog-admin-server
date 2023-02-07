package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitDictionaryRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("dictionary")
	pubr := pub.Group("dictionary")

	pubr.GET("/findTreeData", v1.FindDictionaryTree)

	prir.POST("/add", v1.CreateDictionary)
	prir.POST("/update", v1.UpdateDictionary)
	prir.POST("/delete", v1.DeleteDictionary)
	prir.GET("/find", v1.FindDictionary)
	prir.POST("/findList", v1.FindDictionaryList)
	prir.POST("/findPaging", v1.FindDictionaryPaging)
	prir.GET("/findAllTree", v1.FindAllDictionaryTree)
	prir.GET("/findAllParent", v1.FindDictionaryAllParent)
	prir.GET("/findChildren", v1.FindDictionaryChildren)
	prir.GET("/findAllChildren", v1.FindDictionaryAllChildren)

}
