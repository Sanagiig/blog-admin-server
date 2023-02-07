package request

type TagID struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateTagBody struct {
	Name        string `json:"name" form:"name"`
	Type        string `json:"type" form:"type"`
	Description string `json:"description" form:"description"`
}

type UpdateTagBody struct {
	CreateTagBody
	TagID
}

type FindTagBody struct {
	CreateTagBody
	ID   string   `json:"id" form:"id"`
	Type []string `json:"type" form:"type"`
}

type FindTagPagingBody struct {
	FindTagBody
	Pagination
}

type DeleteTagBody struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}
