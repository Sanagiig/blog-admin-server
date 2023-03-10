package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//var domainRe = regexp.MustCompile("https.//\\(" + settings.Domain + "\\)(/|$)")

func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	// 必须，设置服务器支持的所有跨域请求的方法
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	c.Header("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	c.Header("Access-Control-Allow-Credentials", "true")
	//放行所有OPTIONS方法
	if c.Request.Method == "OPTIONS" {
		c.JSON(http.StatusOK, "Options Request!")
	}
	// 处理请求
	c.Next() //  处理请求
}
