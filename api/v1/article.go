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

// 添加文章
func AddArticle(c *gin.Context) {
	var params request.CreateArticleBody
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	// 作者
	if params.AuthorID == "" {
		params.AuthorID = c.GetString("id")
	}

	if err = service.CreateArticle(&params); err != nil {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, err)
		return
	}

	result.SucJson(c, nil)
}

// 编辑文章
func UpdateArticle(c *gin.Context) {
	var err error
	var params request.UpdateArticleBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if err = service.UpdateArticle(params.ID, &params.CreateArticleBody); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}

// 查询单个文章
func FindArticle(c *gin.Context) {
	var params request.ArticleID
	article := model.Article{}

	if !pretreat.ExtractAndTransParams(c, "query", &params, &article) {
		return
	}

	if err := service.FindArticle(&article); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &article)
}

// 查询文章列表
func FindArticleList(c *gin.Context) {
	var params request.FindArticleBody
	var articles []model.Article
	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	articles, err := service.FindArticlePaging(&params, nil, nil)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, &articles)
}

func FindArticlePaging(c *gin.Context) {
	var params request.FindArticlePagingBody
	var count int64
	var articles []model.Article
	var err error

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	articles, err = service.FindArticlePaging(&params.FindArticleBody, &params.Pagination, &count)
	if err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.QueryTableJson(c, count, &articles)
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	var params request.DeleteArticleBody

	if !pretreat.ExtractParams(c, "json", &params) {
		return
	}

	if len(params.IDs) == 0 {
		result.ArgErrJson(c, errMsg.ERROR_ARGUMENTS_WRONG, nil)
		return
	}

	if err := service.DeleteArticle(params.IDs); err != nil {
		result.ErrJson(c, errMsg.ERROR, nil, err)
		return
	}

	result.SucJson(c, nil)
}
