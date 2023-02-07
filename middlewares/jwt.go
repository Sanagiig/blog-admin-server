package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-blog/cache"
	"go-blog/global/settings"
	"go-blog/message/errMsg"
	"go-blog/utils"
	"go-blog/utils/result"
)

//type Claim struct {
//	Id string `json:"id"`
//	Username string `json:"username"`
//}
//
//var JWTKey = []byte(utils.JWTKey)
//func GenToken() {
//	jwt.
//}

func setToken2Ctx(c *gin.Context, tokenInfo *utils.TokenInfo) {
	c.Set("id", tokenInfo.ID)
	c.Set("username", tokenInfo.Username)
}

func isTokenValid(c *gin.Context) bool {
	token, err := c.Cookie("token")
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR_NO_TOKEN, err)
		return false
	}

	if token == "" {
		token = c.GetHeader("token")
		if token == "" {
			result.ArgErrJson(c, errMsg.ERROR_NO_TOKEN, nil)
			return false
		}
	}

	tokenInfo, err := utils.ParseToken(token)
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR_TOKEN_WRONG, err)
		return false
	}

	rdToken, err := cache.PullToken(tokenInfo.ID)
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR, nil)
		return false
	} else if token != "" && rdToken != token {
		result.ErrJson(c, errMsg.ERROR_LOGIN_ANOTHER, nil, nil)
		return false
	} else if tokenInfo.CreateAt.Hour() >= settings.TokenExp {
		result.ErrJson(c, errMsg.ERROR_TOKEN_RUNTIME, nil, nil)
		return false
	}

	setToken2Ctx(c, tokenInfo)
	return true
}

func Private(ctx *gin.Context) {
	if !isTokenValid(ctx) {
		ctx.Abort()
		return
	}
}

func Pub(ctx *gin.Context) {
	ctx.Next()
}
