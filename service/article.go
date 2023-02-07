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

func CreateArticle(p *request.CreateArticleBody) error {
	article := &model.Article{}
	tags := []model.Tag{}
	if err := utils.FieldValTransfer(article, p); err != nil {
		return err
	}

	if len(p.Tags) > 0 {
		for i := 0; i < len(p.Tags); i++ {
			tags = append(tags, model.Tag{Base: model.Base{ID: p.Tags[i]}})
		}
		article.Tags = tags
	}
	return global.DB.Model(&model.Article{}).Create(article).Error
}

func FindArticle(p *model.Article) error {
	return global.DB.Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, nickname")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).First(p).Error
}

func FindArticlePaging(params *request.FindArticleBody, pageInfo *request.Pagination, count *int64) ([]model.Article, error) {
	var Articles []model.Article
	var err error

	db := global.DB.Omit("content")
	sql := strings.Builder{}

	sql.WriteString(
		fmt.Sprintf("title like '%%%s%%' AND description like '%%%s%%'",
			params.Title, params.Description,
		))

	if params.ID != "" {
		sql.WriteString(fmt.Sprintf("AND id = %s", params.ID))
	}

	if len(params.Tags) > 0 {
		db = db.Model(&model.Article{}).Where("id in (?)",
			db.Table("article_tag").Select("article_id").Where("tag_id in (?)", strings.Join(params.Tags, ",")),
		).Where(sql.String())
	} else {
		db = db.Model(&model.Article{}).Where(sql.String())
	}

	if pageInfo != nil {
		db.Count(count)
		db = db.Offset((pageInfo.Page - 1) * pageInfo.PageSize).Limit(pageInfo.PageSize)

		// 此处用 scopes 会报错，有可能是调整了 sql 拼接的顺序
		//err = db.Scopes(scopes.Pager(pageInfo.Page, pageInfo.PageSize, count)).Find(&Articles).Error
	}

	err = db.Preload("Author", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, nickname")
	}).Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Preload("Tags", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Find(&Articles).Error

	return Articles, err
}

func UpdateArticle(id string, params *request.CreateArticleBody) error {
	var err error
	article := model.Article{}

	if err = utils.FieldValTransfer(&article, params); err != nil {
		return err
	}

	db := global.DB.Model(&model.Article{Base: model.Base{ID: id}})
	if len(params.Tags) > 0 {
		tags := []model.Tag{}
		for i := 0; i < len(params.Tags); i++ {
			tags = append(tags, model.Tag{Base: model.Base{ID: params.Tags[i]}})
		}
		err = db.Association("Tags").Replace(tags)
	} else {
		err = db.Association("Tags").Clear()
	}

	if err != nil {
		return err
	}

	return db.Updates(&article).Error
}

func DeleteArticle(ids []string) error {
	articles := []model.Article{}
	// 此处不能用 where 拼接，会报 主键不存在 ，必须建实例（蛋疼）
	for i := 0; i < len(ids); i++ {
		articles = append(articles, model.Article{Base: model.Base{ID: ids[i]}})
	}

	db := global.DB.Model(&articles)
	if err := db.Association("Tags").Clear(); err != nil {
		return err
	}

	if err := global.DB.Delete(&articles).Error; err != nil {
		return err
	}

	return nil
}
