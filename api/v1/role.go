package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-blog/message/errMsg"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/service"
	"go-blog/utils/pretreat"
	"go-blog/utils/result"
)

// 添加角色
func AddRole(c *gin.Context) {
	var params request.CreateRoleBody
	var Role model.Role

	if !pretreat.ExtractAndTransParams(c, "json", &params, &Role) {
		return
	}
	if err := service.CreateRole(&Role); err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	result.SucJson(c, nil)
}

// FindRoleList
//
//	@Description: 获取角色列表
//	@param c
func FindRole(c *gin.Context) {
	var params request.FindRoleBody
	var Role model.Role

	if !pretreat.ExtractAndTransParams(c, "query", &params, &Role) {
		return
	}

	if err := service.FindRole(&Role); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &Role)
}

// 查询角色列表
func GetRoleList(c *gin.Context) {
	var params request.FindRoleBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}
	Roles, err := service.FindRolePaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, Roles)
}

// 分页查询角色列表
func GetRolePaging(c *gin.Context) {
	var params request.FindRolePagingBody
	var count int64

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	Roles, err := service.FindRolePaging(&params.FindRoleBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, Roles)
}

// 编辑角色
func UpdateRole(c *gin.Context) {
	var params request.UpdateRoleBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err := service.UpdateRole(params.ID, &params.CreateRoleBody); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

// 删除角色
func DeleteRole(c *gin.Context) {
	var params request.DeleteRoleBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.IDs) == 0 {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil)
		return
	}

	if err := service.DeleteRole(&params); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

func PatchAddPermission(c *gin.Context) {
	var params request.PatchRolePermissionBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.RoleIds) == 0 || len(params.PermissionIds) == 0 {
		result.ErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil, errors.New("角色和权限不能为空"))
		return
	}

	if err := service.PatchAddPermission(params.RoleIds, params.PermissionIds); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

func PatchRemovePermission(c *gin.Context) {
	var params request.PatchRolePermissionBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.RoleIds) == 0 || len(params.PermissionIds) == 0 {
		result.ErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil, errors.New("角色和权限不能为空"))
		return
	}

	if err := service.PatchRemovePermission(params.RoleIds, params.PermissionIds); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
