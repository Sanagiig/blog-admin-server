package v1

import (
	"github.com/gin-gonic/gin"
	"go-blog/message/errMsg"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/service"
	"go-blog/utils/pretreat"
	"go-blog/utils/result"
)

// 添加权限
func AddPemission(c *gin.Context) {
	var params request.CreatePermissionBody
	var permission model.Permission

	if !pretreat.ExtractAndTransParams(c, "json", &params, &permission) {
		return
	}

	if err := service.CreatePermission(&permission); err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	result.SucJson(c, nil)
}

// FindPemissionList
//
//	@Description: 获取权限列表
//	@param c
func FindPemission(c *gin.Context) {
	var params request.FindPermissionBody
	var permission model.Permission

	if !pretreat.ExtractAndTransParams(c, "query", &params, &permission) {
		return
	}

	if err := service.FindPermission(&permission); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &permission)
}

// 查询权限列表
func FindPemissionList(c *gin.Context) {
	var params request.FindPermissionBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}
	permissions, err := service.FindPermissionPaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, permissions)
}

// 查询权限分页
func FindPemissionPaging(c *gin.Context) {
	var params request.FindPermissionPagingBody
	var count int64

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}
	permissions, err := service.FindPermissionPaging(&params.FindPermissionBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, permissions)
}

// 编辑权限
func UpdatePemission(c *gin.Context) {
	var params request.UpdatePermissionBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err := service.UpdatePermission(params.ID, &params.CreatePermissionBody); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

// 删除权限
func DeletePemission(c *gin.Context) {
	var params request.DeletePermissionBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.IDs) == 0 {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil)
		return
	}

	if err := service.DeletePermission(&params); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
