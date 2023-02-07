package service

import (
	"fmt"
	"go-blog/global"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/utils"
	"gorm.io/gorm"
	"strings"
)

func CreateRole(p *model.Role) error {
	return global.DB.Model(&model.Role{}).Create(p).Error
}

func FindRole(p *model.Role) error {
	return global.DB.Preload("Permissions", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name")
	}).Where(p).First(p).Error
}

func FindRolePaging(params *request.FindRoleBody, pageInfo *request.Pagination, count *int64) ([]model.Role, error) {
	var Roles []model.Role
	var err error

	db := global.DB
	sql := strings.Builder{}

	sql.WriteString(fmt.Sprintf("name like '%%%s%%' AND Code like '%%%s%%'", params.Name, params.Code))
	if params.ID != "" {
		sql.WriteString(fmt.Sprintf("AND id = %s", params.ID))
	}

	if len(params.Permissions) > 0 {
		db = db.Model(&model.Role{}).Where("id in (?)",
			db.Table("role_permission").Select("role_id").Where("permission_id in (?)", strings.Join(params.Permissions, ",")),
		).Where(sql.String())
	} else {
		db = db.Model(&model.Role{}).Where(sql.String())
	}

	if pageInfo != nil {
		db.Count(count)
		db = db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).Limit(pageInfo.PageSize)

		// 此处用 scopes 会报错，有可能是调整了 sql 拼接的顺序
		//err = db.Scopes(scopes.Pager(pageInfo.Page, pageInfo.PageSize, count)).Find(&Roles).Error
	}

	err = db.Find(&Roles).Error
	return Roles, err
}

func UpdateRole(id string, params *request.CreateRoleBody) error {
	var err error

	role := model.Role{}
	if err = utils.FieldValTransfer(&role, params); err != nil {
		return err
	}

	if params.Permissions != nil {
		db := global.DB.Model(&model.Role{Base: model.Base{ID: id}})
		if len(params.Permissions) > 0 {
			permissions := []model.Permission{}
			for i := 0; i < len(params.Permissions); i++ {
				permissions = append(permissions, model.Permission{Base: model.Base{ID: params.Permissions[i]}})
			}

			if err = db.Preload("Permissions").Association("Permissions").Replace(permissions); err != nil {
				return err
			}

		} else {
			if err = db.Association("Permissions").Clear(); err != nil {
				return err
			}
		}
	}

	return global.DB.Where("id = ?", id).Updates(&role).Error
}

func DeleteRole(params *request.DeleteRoleBody) error {
	roles := []model.Role{}

	//此处不能用 where 拼接，会报 主键不存在 ，必须建实例（蛋疼）
	for i := 0; i < len(params.IDs); i++ {
		roles = append(roles, model.Role{Base: model.Base{ID: params.IDs[i]}})
	}

	if err := global.DB.Model(&roles).Association("Permissions").Clear(); err != nil {
		return err
	}

	if err := global.DB.Delete(&model.Role{}, params.IDs).Error; err != nil {
		return err
	}

	return nil
}

func PatchAddPermission(roleIds []string, permissionIds []string) error {
	roles := []model.Role{}
	permissions := []model.Permission{}

	for i := 0; i < len(roleIds); i++ {
		roles = append(roles, model.Role{Base: model.Base{ID: roleIds[i]}})
	}

	for i := 0; i < len(permissionIds); i++ {
		permissions = append(permissions, model.Permission{Base: model.Base{ID: permissionIds[i]}})
	}

	return global.DB.Model(&roles).Association("Permissions").Append(permissions)
}

func PatchRemovePermission(roleIds []string, permissionIds []string) error {
	roles := []model.Role{}

	for i := 0; i < len(roleIds); i++ {
		roles = append(roles, model.Role{Base: model.Base{ID: roleIds[i]}})
	}

	return global.DB.Model(&roles).Where("permission_id in (?)", permissionIds).Association("Permissions").Clear()
}
