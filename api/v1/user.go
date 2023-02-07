package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/cache"
	"go-blog/global/settings"
	"go-blog/message"
	"go-blog/message/errMsg"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/service"
	"go-blog/utils"
	"go-blog/utils/pretreat"
	"go-blog/utils/result"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 注册
func Register(c *gin.Context) {
	AddUser(c)
}

// 登录
func Login(c *gin.Context) {
	params := &request.LoginBody{}

	err := c.ShouldBindJSON(params)
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	params.Password, err = utils.DecodeBase64(params.Password)
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR, err)
		return
	}

	params.Password, err = utils.EncryptSha256(params.Password)
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR, err)
		return
	}

	if !service.IsCheckCodeValid(params.CheckCode) {
		result.ArgErrJson(c, errMsg.ERROT_CHECKCODE_WRONG, nil)
		return
	}

	user, token, err := service.Login(params)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR_USERNAME_PASSWORD_WRONG, nil, err)
		return
	}

	if err := cache.PushToken(user.ID, token); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	c.SetCookie("token", token, 3600*settings.TokenExp, "/", "*", false, true)
	result.SucJson(c, gin.H{"token": token})
}

// 查找当前用户详细数据
func FindSelfInfo(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR_NO_TOKEN, err)
		return
	}

	u, err := service.FindSelfInfo(token)
	if err != nil {
		result.ArgErrJson(c, errMsg.ERROR, err)
		return
	}

	result.SucJson(c, u)
}

// 添加用户
func AddUser(c *gin.Context) {
	var params request.CreateUserBody
	var user model.User
	var code int = message.SUCCESS
	var err error

	if !pretreat.ExtractAndTransParams(c, "json", &params, &user) {
		return
	}

	if service.CheckUserExist(&user) {
		result.ArgErrJson(c, errMsg.ERROR_USERNAME_USED, nil)
		return
	}

	err = service.CreateUser(&params)
	if err != nil {
		code = errMsg.ERROR
		result.ErrJson(c, code, nil, err)
		return
	}
	result.SucJson(c, nil)
}

// 查询单个用户
func GetUserInfo(c *gin.Context) {
	var params request.FindUserBody
	var user model.User
	if !pretreat.ExtractAndTransParams(c, "query", &params, &user) {
		return
	}

	data, err := service.FindUser(&user)
	if err != nil {

	}
	result.SucJson(c, data)
}

// 查询用户分页
func FindUserPaging(c *gin.Context) {
	var params request.FindUserPagingBody
	var count int64
	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	data, err := service.GetUserPaging(&params.FindUserBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, data)
}

// 查询用户列表
func FindUserList(c *gin.Context) {
	var params request.FindUserBody

	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	data, err := service.GetUserPaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}
	result.SucJson(c, data)
}

// 编辑用户
func UpdateUser(c *gin.Context) {
	var params request.UpdateUserBody
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.UpdateUser(&params); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	var params request.IDS
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.DeleteUser(params.IDS); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

func PatchAddRole(c *gin.Context) {
	var params request.PatchUserRoleBody
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.PatchAddRole(params.UserIDs, params.RoleIDs); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

func PatchRemoveRole(c *gin.Context) {
	var params request.PatchUserRoleBody
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.PatchRemoveRole(params.UserIDs, params.RoleIDs); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
