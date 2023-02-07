package service

import (
	"fmt"
	"go-blog/global"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/utils"
)

func CreateTag(p *model.Tag) error {
	return global.DB.Model(&model.Tag{}).Create(p).Error
}

func FindTag(p *model.Tag) error {
	return global.DB.Where(p).First(p).Error
}

func FindTagPaging(params *request.FindTagBody, pageInfo *request.Pagination, count *int64) ([]model.Tag, error) {
	var tags []model.Tag
	var err error

	db := global.DB.Model(&model.Tag{}).Where(
		fmt.Sprintf("name like '%%%s%%' AND description like '%%%s%%'",
			params.Name, params.Description,
		))

	if params.ID != "" {
		db = db.Where(fmt.Sprintf("AND id = %s", params.ID))
	}

	if len(params.Type) > 0 {
		db = db.Where("type in (?)", params.Type)
	}

	if pageInfo != nil {
		db.Count(count)
		db = db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).Limit(pageInfo.PageSize)
	}

	err = db.Find(&tags).Error
	if err != nil {
		return nil, err
	}

	// 类型转中文
	if len(tags) > 0 {
		typeList, err := FindDictionaryChildren(&model.Dictionary{Code: "tagType"})
		if err != nil {
			return nil, err
		}

		utils.TransStrucField2Another(typeList, tags, "Type")
	}

	return tags, nil
}

func UpdateTag(id string, params *request.CreateTagBody) error {
	var err error

	Tag := model.Tag{}
	if err = utils.FieldValTransfer(&Tag, params); err != nil {
		return err
	}

	return global.DB.Where("id = ?", id).Updates(&Tag).Error
}

func DeleteTag(params *request.DeleteTagBody) error {
	return global.DB.Delete(&model.Tag{}, params.IDs).Error
}
