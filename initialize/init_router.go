package initialize

import (
	"github.com/gin-gonic/gin"
	"go-blog/global/settings"
	"go-blog/middlewares"
	"go-blog/routes"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORS)
	router.Static("/static", settings.StaticDir)
	pub := router.Group("")
	private := router.Group("")

	pub.Use(middlewares.Pub)
	private.Use(middlewares.Private)
	routes.InitRouter(pub, private)

	return router
}
