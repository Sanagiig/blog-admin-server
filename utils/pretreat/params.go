package pretreat

import (
	"github.com/gin-gonic/gin"
	"go-blog/message/errMsg"
	"go-blog/utils"
	"go-blog/utils/result"
)

// ExtractParams
//
//	@Description: 提取参数，如果提取失败则直接返回 false,并结束该次请求
//	@param c
//	@param paramsType
//	@param obj
func ExtractParams(c *gin.Context, paramsType string, obj any) (res bool) {
	var err error
	res = true

	switch paramsType {
	case "query":
		err = c.ShouldBindQuery(obj)
	case "json":
		err = c.ShouldBindJSON(obj)
	default:
		panic("参数类型错误")
	}

	if err != nil {
		res = false
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
	}

	return res
}

// ExtractAndTransParams
//
//	@Description: 提取参数并且将合适的字段赋值到model
//	@param c
//	@param paramsType
//	@param params
//	@param model
//	@return res
func ExtractAndTransParams(c *gin.Context, paramsType string, params any, model any) (res bool) {
	if !ExtractParams(c, paramsType, params) {
		return false
	}

	err := utils.FieldValTransfer(model, params)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return false
	}

	return true
}
