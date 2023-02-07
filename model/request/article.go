package request

type ArticleID struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateArticleBody struct {
	Title       string   `json:"title" form:"title"`
	Description string   `json:"description" form:"description"`
	Tags        []string `json:"tags" form:"tags"`
	CategoryID  string   `json:"categoryID" form:"categoryID"`
	AuthorID    string   `json:"authorID" form:"authorID"`
	Content     string   `json:"content"`
}

type UpdateArticleBody struct {
	CreateArticleBody
	ArticleID
}

type FindArticleBody struct {
	CreateArticleBody
	ID string `json:"id" form:"id"`
}

type FindArticlePagingBody struct {
	FindArticleBody
	Pagination
}

type DeleteArticleBody struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}
