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

func CreateCategory(c *gin.Context) {
	var params request.CreateCategoryBody
	var category model.Category

	if !pretreat.ExtractAndTransParams(c, "json", &params, &category) {
		return
	}

	if err := service.CreateCategory(&category); err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	result.SucJson(c, nil)
}

func UpdateCategory(c *gin.Context) {
	var err error
	var params request.UpdateCategoryBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.UpdateCategory(params.ID, &params.CreateCategoryBody); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

func FindCategory(c *gin.Context) {
	var params request.CategoryID
	category := model.Category{}

	if !pretreat.ExtractAndTransParams(c, "query", &params, &category) {
		return
	}

	if err := service.FindCategory(&category); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &category)
}

func FindCategoryList(c *gin.Context) {
	var params request.FindCategoryBody
	var categories []model.Category
	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	categories, err := service.FindCategoryPaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &categories)
}

func FindCategoryPaging(c *gin.Context) {
	var params request.FindCategoryPagingBody
	var count int64
	var categories []model.Category
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	categories, err = service.FindCategoryPaging(&params.FindCategoryBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, &categories)
}

func FindAllCategoryTree(c *gin.Context) {
	categories, err := service.FindAllCategoryTree()
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, categories)
}

func FindCategoryAllParent(c *gin.Context) {
	var params request.CategoryID
	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	categories, err := service.FindCategoryAllParent(params.ID)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, categories)
}

func FindCategoryChildren(c *gin.Context) {
	var params request.FindCategoryBody
	var category model.Category

	if !pretreat.ExtractAndTransParams(c, "query", &params, &category) {
		return
	}

	dictionaries, err := service.FindCategoryChildren(&category)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, dictionaries)
}

func FindCategoryAllChildren(c *gin.Context) {
	var params request.CategoryID
	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	categories, err := service.FindCategoryAllChildren(params.ID)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, categories)
}

func FindCategoryTree(c *gin.Context) {
	var params request.CategoryID
	if !pretreat.ExtractParams(c, "query", &params) {
		return
	}

	categories, err := service.FindCategoryTree(params.ID)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, categories)
}

func DeleteCategory(c *gin.Context) {
	var params request.DeleteCategoryBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.IDs) == 0 {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil)
		return
	}

	if err := service.DeleteCategory(params.IDs); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
