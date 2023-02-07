package service

import (
	"fmt"
	"go-blog/global"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/utils"
	"strings"
)

func CreatePermission(p *model.Permission) error {
	return global.DB.Model(&model.Permission{}).Create(p).Error
}

func FindPermission(p *model.Permission) error {
	return global.DB.First(p).Error
}

func FindPermissionPaging(params *request.FindPermissionBody, pageInfo *request.Pagination, count *int64) ([]model.Permission, error) {
	var permissions []model.Permission
	var err error

	db := global.DB
	sql := strings.Builder{}

	sql.WriteString(fmt.Sprintf(
		"name like '%%%s%%' AND Code like '%%%s%%' AND value like '%%%s%%' AND description like '%%%s%%'",
		params.Name, params.Code, params.Value, params.Description,
	))
	if params.ID != "" {
		sql.WriteString(fmt.Sprintf("AND id = %s", params.ID))
	}

	if params.Type != "" {
		sql.WriteString(fmt.Sprintf("AND type = %s", params.Type))
	}

	if len(params.Roles) > 0 {
		db = db.Model(&model.Permission{}).Where("id in (?)",
			db.Table("role_permission").Select("permission_id").Where("role_id in (?)", strings.Join(params.Roles, ",")),
		).Where(sql.String())
	} else {
		db = db.Model(&model.Permission{}).Where(sql.String())
	}

	if pageInfo != nil {
		db.Count(count)
		db = db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).Limit(pageInfo.PageSize)

		// 此处用 scopes 会报错，有可能是调整了 sql 拼接的顺序
		//err = db.Scopes(scopes.Pager(pageInfo.Page, pageInfo.PageSize, count)).Find(&permissions).Error
	}

	err = db.Find(&permissions).Error
	return permissions, err
}

func UpdatePermission(id string, params *request.CreatePermissionBody) error {
	var err error

	permission := model.Permission{}
	if err = utils.FieldValTransfer(&permission, params); err != nil {
		return err
	}

	err = global.DB.Where("id = ?", id).Updates(&permission).Error
	return err
}

func DeletePermission(params *request.DeletePermissionBody) error {
	permissions := []model.Permission{}
	// 此处不能用 where 拼接，会报 主键不存在 ，必须建实例（蛋疼）
	for i := 0; i < len(params.IDs); i++ {
		permissions = append(permissions, model.Permission{Base: model.Base{ID: params.IDs[i]}})
	}

	db := global.DB.Model(&permissions)
	if err := db.Association("Roles").Clear(); err != nil {
		return err
	}

	if err := global.DB.Delete(&permissions).Error; err != nil {
		return err
	}

	return nil
}
