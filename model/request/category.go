package request

type CategoryID struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateCategoryBody struct {
	ParentID    string `json:"ParentID" form:"ParentID"`
	Name        string `json:"name" form:"name"`
	Type        string `json:"type" form:"type"`
	Description string `json:"description" form:"description"`
}

type UpdateCategoryBody struct {
	CreateCategoryBody
	CategoryID
}

type FindCategoryBody struct {
	CreateCategoryBody
	ID   string   `json:"id" form:"id"`
	Type []string `json:"type" form:"type"`
}

type FindCategoryPagingBody struct {
	FindCategoryBody
	Pagination
}

type DeleteCategoryBody struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}
