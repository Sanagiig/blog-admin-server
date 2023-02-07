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

func CreateDictionary(params *model.Dictionary) error {
	return global.DB.Create(&params).Error
}

func UpdateDictionary(id string, params *request.CreateDictionaryBody) error {
	var err error
	dictionary := model.Dictionary{}

	if err = utils.FieldValTransfer(&dictionary, params); err != nil {
		return err
	}

	return global.DB.Where("id = ?", id).Updates(&dictionary).Error
}

func FindDictionary(c *model.Dictionary) error {
	return global.DB.Where(c).First(c).Error
}

func FindDictionaryPaging(params *request.FindDictionaryBody, pageInfo *request.Pagination, count *int64) ([]model.Dictionary, error) {
	var dictionaries []model.Dictionary
	var err error

	db := global.DB.Model(&model.Dictionary{}).Where(
		fmt.Sprintf("name like '%%%s%%' AND code like '%%%s%%' AND  description like '%%%s%%'",
			params.Name, params.Code, params.Description),
	)

	if params.ParentID == "" {
		db = db.Where("parent_id IS NULL")
	} else {
		db = db.Where("parent_id = ?", params.ParentID)
	}

	if pageInfo != nil {
		//db = db.Model(&model.Dictionary{}).Scopes(scopes.Pager(pageInfo.Page, pageInfo.PageSize, count))
		db.Count(count)
		db = db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).Limit(pageInfo.PageSize)
	}

	err = db.Find(&dictionaries).Error
	return dictionaries, err
}

func FindDictionaryAllParent(id string) ([]model.Dictionary, error) {
	var parentIDs string
	dictionaries := []model.Dictionary{}

	err := global.DB.Scopes(scopes.AllParentID("dictionary", id)).Scan(&parentIDs).Error
	if err != nil {
		return dictionaries, err
	}

	err = global.DB.Find(&dictionaries, strings.Split(parentIDs, ",")).Error
	return dictionaries, err
}

func FindDictionaryChildren(params *model.Dictionary) ([]model.Dictionary, error) {
	res := []model.Dictionary{}
	if params.ID == "" {
		err := FindDictionary(params)
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

func FindDictionaryAllChildren(params *model.Dictionary) ([]model.Dictionary, error) {
	var childrenIDs string
	dictionaries := []model.Dictionary{}

	err := FindDictionary(params)
	if err != nil {
		return nil, err
	} else if params.ID == "" {
		return nil, fmt.Errorf("查找不到相关字典信息")
	}

	err = global.DB.Scopes(scopes.AllChildrenID("dictionary", params.ID)).Scan(&childrenIDs).Error
	if err != nil {
		return dictionaries, err
	}

	err = global.DB.Find(&dictionaries, strings.Split(childrenIDs, ",")).Error
	return dictionaries, err
}

func FindDictionaryTree(id string) ([]*response.DictionaryTreeData, error) {
	var childrenIDs string
	var res []*response.DictionaryTreeData
	dictionaries := []model.Dictionary{}
	dictionaryNodes := []utils.TreeNode{}

	err := global.DB.Scopes(scopes.AllChildrenID("dictionary", id)).Scan(&childrenIDs).Error
	if err != nil {
		return nil, err
	}

	err = global.DB.Model(&model.Dictionary{}).Find(&dictionaries, strings.Split(childrenIDs, ",")).Error
	for i := 0; i < len(dictionaries); i++ {
		dictionaryNodes = append(dictionaryNodes, &response.DictionaryTreeData{Dictionary: &dictionaries[i]})
	}

	// 类型转换
	treeList := (utils.List2TreeMap(dictionaryNodes))
	for _, node := range treeList {
		res = append(res, node.(*response.DictionaryTreeData))
	}
	return res, nil
}

func FindAllDictionaryTree() ([]*response.DictionaryTreeData, error) {
	var res []*response.DictionaryTreeData
	dictionaries := []model.Dictionary{}
	dictionaryNodes := []utils.TreeNode{}

	err := global.DB.Find(&dictionaries).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(dictionaries); i++ {
		dictionaryNodes = append(dictionaryNodes, &response.DictionaryTreeData{Dictionary: &dictionaries[i]})
	}

	// 类型转换
	treeList := (utils.List2TreeMap(dictionaryNodes))
	for _, node := range treeList {
		if node.GetParentID() == "" {
			res = append(res, node.(*response.DictionaryTreeData))
		}
	}
	return res, nil
}

func DeleteDictionary(ids []string) error {
	return global.DB.Delete(&model.Dictionary{}, ids).Error
}
