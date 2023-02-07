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

// 添加标签
func AddTag(c *gin.Context) {
	var params request.CreateTagBody
	var Tag model.Tag

	if !pretreat.ExtractAndTransParams(c, "json", &params, &Tag) {
		return
	}
	if err := service.CreateTag(&Tag); err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	result.SucJson(c, nil)
}

// FindTagList
//
//	@Description: 获取标签列表
//	@param c
func FindTag(c *gin.Context) {
	var params request.FindTagBody
	var Tag model.Tag

	if !pretreat.ExtractAndTransParams(c, "query", &params, &Tag) {
		return
	}

	if err := service.FindTag(&Tag); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &Tag)
}

// 查询标签列表
func FindTagList(c *gin.Context) {
	var params request.FindTagBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}
	Tags, err := service.FindTagPaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, Tags)
}

// 分页查询标签列表
func FindTagPaging(c *gin.Context) {
	var params request.FindTagPagingBody
	var count int64

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	Tags, err := service.FindTagPaging(&params.FindTagBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, Tags)
}

// 编辑标签
func UpdateTag(c *gin.Context) {
	var params request.UpdateTagBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err := service.UpdateTag(params.ID, &params.CreateTagBody); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

// 删除标签
func DeleteTag(c *gin.Context) {
	var params request.DeleteTagBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.IDs) == 0 {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil)
		return
	}

	if err := service.DeleteTag(&params); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
