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

func CreateDictionary(c *gin.Context) {
	var params request.CreateDictionaryBody
	var dictionary model.Dictionary

	if !pretreat.ExtractAndTransParams(c, "json", &params, &dictionary) {
		return
	}

	if err := service.CreateDictionary(&dictionary); err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	result.SucJson(c, nil)
}

func UpdateDictionary(c *gin.Context) {
	var err error
	var params request.UpdateDictionaryBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.UpdateDictionary(params.ID, &params.CreateDictionaryBody); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

func FindDictionary(c *gin.Context) {
	var params request.DictionaryID
	dictionary := model.Dictionary{}

	if !pretreat.ExtractAndTransParams(c, "query", &params, &dictionary) {
		return
	}

	if err := service.FindDictionary(&dictionary); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &dictionary)
}

func FindDictionaryList(c *gin.Context) {
	var params request.FindDictionaryBody
	var dictionaries []model.Dictionary
	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	dictionaries, err := service.FindDictionaryPaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &dictionaries)
}

func FindDictionaryPaging(c *gin.Context) {
	var params request.FindDictionaryPagingBody
	var count int64
	var dictionaries []model.Dictionary
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	dictionaries, err = service.FindDictionaryPaging(&params.FindDictionaryBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, &dictionaries)
}

func FindAllDictionaryTree(c *gin.Context) {
	dictionaries, err := service.FindAllDictionaryTree()
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, dictionaries)
}

func FindDictionaryAllParent(c *gin.Context) {
	var params request.DictionaryID
	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	dictionaries, err := service.FindDictionaryAllParent(params.ID)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, dictionaries)
}

func FindDictionaryChildren(c *gin.Context) {
	var params request.FindDictionaryBody
	var dictionary model.Dictionary

	if !pretreat.ExtractAndTransParams(c, "query", &params, &dictionary) {
		return
	}

	dictionaries, err := service.FindDictionaryChildren(&dictionary)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, dictionaries)
}

func FindDictionaryAllChildren(c *gin.Context) {
	var params request.FindDictionaryBody
	var dictionary model.Dictionary

	if !pretreat.ExtractAndTransParams(c, "query", &params, &dictionary) {
		return
	}

	dictionaries, err := service.FindDictionaryAllChildren(&dictionary)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, dictionaries)
}

func FindDictionaryTree(c *gin.Context) {
	var params request.DictionaryID
	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	dictionaries, err := service.FindDictionaryTree(params.ID)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, dictionaries)
}

func DeleteDictionary(c *gin.Context) {
	var params request.DeleteDictionaryBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.IDs) == 0 {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil)
		return
	}

	if err := service.DeleteDictionary(params.IDs); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
