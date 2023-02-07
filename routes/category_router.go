package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-blog/api/v1"
)

func InitCategoryRouter(pub *gin.RouterGroup, private *gin.RouterGroup) {
	prir := private.Group("category")
	//pubr := pub.Group("category")

	prir.GET("/findTreeData", v1.FindCategoryTree)
	prir.POST("/add", v1.CreateCategory)
	prir.POST("/update", v1.UpdateCategory)
	prir.POST("/delete", v1.DeleteCategory)
	prir.GET("/find", v1.FindCategory)
	prir.POST("/findList", v1.FindCategoryList)
	prir.POST("/findPaging", v1.FindCategoryPaging)
	prir.GET("/findAllTree", v1.FindAllCategoryTree)
	prir.GET("/findAllParent", v1.FindCategoryAllParent)
	prir.GET("/findChildren", v1.FindCategoryChildren)
	prir.GET("/findAllChildren", v1.FindCategoryAllChildren)
}
