package service

import (
	"fmt"
	"go-blog/global"
	"go-blog/model"
	"go-blog/model/request"
	"go-blog/model/response"
	"go-blog/model/scopes"
	"go-blog/utils"
	"strings"
)

func CreateCategory(c *model.Category) error {
	return global.DB.Create(&c).Error
}

func UpdateCategory(id string, params *request.CreateCategoryBody) error {
	var err error
	category := model.Category{}

	if err = utils.FieldValTransfer(&category, params); err != nil {
		return err
	}

	return global.DB.Where("id = ?", id).Updates(&category).Error
}

func FindCategory(c *model.Category) error {
	return global.DB.First(c).Error
}

func FindCategoryPaging(params *request.FindCategoryBody, pageInfo *request.Pagination, count *int64) ([]model.Category, error) {
	var categories []model.Category
	var err error

	db := global.DB.Model(&model.Category{}).Where(
		fmt.Sprintf("name like '%%%s%%' AND  description like '%%%s%%'",
			params.Name, params.Description),
	)

	if params.ParentID == "" {
		db = db.Where("parent_id IS NULL")
	} else {
		db = db.Where("parent_id = ?", params.ParentID)
	}

	if len(params.Type) > 0 {
		db = db.Where("type in (?)", params.Type)
	}

	if pageInfo != nil {
		//db = db.Model(&model.Category{}).Scopes(scopes.Pager(pageInfo.Page, pageInfo.PageSize, count))
		db.Count(count)
		db = db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).Limit(pageInfo.PageSize)
	}

	err = db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	// 类型转中文
	if len(categories) > 0 {
		typeList, err := FindDictionaryChildren(&model.Dictionary{Code: "categoryType"})
		if err != nil {
			return nil, err
		}

		utils.TransStrucField2Another(typeList, categories, "Type")
	}

	return categories, err
}

func FindCategoryAllParent(id string) ([]model.Category, error) {
	var parentIDs string
	categories := []model.Category{}

	err := global.DB.Scopes(scopes.AllParentID("category", id)).Scan(&parentIDs).Error
	if err != nil {
		return categories, err
	}

	err = global.DB.Find(&categories, strings.Split(parentIDs, ",")).Error
	return categories, err
}

func FindCategoryChildren(params *model.Category) ([]model.Category, error) {
	res := []model.Category{}
	if params.ID == "" {
		err := FindCategory(params)
		if err != nil {
			return nil, err
		} else if params.ID == "" {
			return nil, fmt.Errorf("查找不到相关字典信息")
		}
	}

	err := global.DB.Where("parent_id = ?", params.ID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

func FindCategoryAllChildren(id string) ([]model.Category, error) {
	var childrenIDs string
	categories := []model.Category{}

	err := global.DB.Scopes(scopes.AllChildrenID("category", id)).Scan(&childrenIDs).Error
	if err != nil {
		return categories, err
	}

	err = global.DB.Find(&categories, strings.Split(childrenIDs, ",")).Error
	return categories, err
}

func FindCategoryTree(id string) ([]*response.CategoryTreeData, error) {
	var childrenIDs string
	var res []*response.CategoryTreeData
	categories := []model.Category{}
	categoryNodes := []utils.TreeNode{}

	err := global.DB.Scopes(scopes.AllChildrenID("category", id)).Scan(&childrenIDs).Error
	if err != nil {
		return nil, err
	}

	err = global.DB.Model(&model.Category{}).Find(&categories, strings.Split(childrenIDs, ",")).Error
	for i := 0; i < len(categories); i++ {
		categoryNodes = append(categoryNodes, &response.CategoryTreeData{Category: &categories[i]})
	}

	// 类型转换
	treeList := (utils.List2TreeMap(categoryNodes))
	for _, node := range treeList {
		res = append(res, node.(*response.CategoryTreeData))
	}
	return res, nil
}

func FindAllCategoryTree() ([]*response.CategoryTreeData, error) {
	var res []*response.CategoryTreeData
	categories := []model.Category{}
	categoryNodes := []utils.TreeNode{}

	err := global.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(categories); i++ {
		categoryNodes = append(categoryNodes, &response.CategoryTreeData{Category: &categories[i]})
	}

	// 类型转换
	treeList := (utils.List2TreeMap(categoryNodes))
	for _, node := range treeList {
		if node.GetParentID() == "" {
			res = append(res, node.(*response.CategoryTreeData))
		}
	}
	return res, nil
}

func DeleteCategory(ids []string) error {
	return global.DB.Delete(&model.Category{}, ids).Error
}
