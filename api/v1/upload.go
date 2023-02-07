package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/message/errMsg"
	"go-blog/model/request"
	"go-blog/service"
	"go-blog/utils"
	"go-blog/utils/pretreat"
	"go-blog/utils/result"
)

func GetUploadAuth(c *gin.Context) {
	var params request.UploadAuthBody

	err := c.BindQuery(&params)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	if params.UploadType == "" {
		result.ErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil, nil)
		return
	} else if params.Username == "" {
		token, _ := c.Cookie("token")
		userTokenInfo, err := utils.ParseToken(token)
		if err != nil {
			result.ErrJson(c, errMsg.ERROR, nil, err)
			return
		}
		params.Username = userTokenInfo.Username
	}

	uploadAuthData, err := service.GetUploadAuth(&params)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	res, err := service.RefreshCdn([]string{uploadAuthData.RefreshUrl})
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}
	fmt.Printf("%v", res)
	result.SucJson(c, uploadAuthData)
}

func RefreshCDN(c *gin.Context) {
	var params request.UploadRefreshCDNBody

	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	res, err := service.RefreshCdn(params.URLs)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, res)
}

func UploadFile(c *gin.Context) {
	typ := c.Param("type")
	username, _ := c.Get("username")
	file, err := c.FormFile("file")
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	url, err := service.UploadFile(typ, username.(string), file)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, url)
}
