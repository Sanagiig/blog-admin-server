package request

type DictionaryID struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateDictionaryBody struct {
	ParentID    string `json:"ParentID" form:"ParentID"`
	Name        string `json:"name" form:"name"`
	Code        string `json:"code" form:"code"`
	Description string `json:"description" form:"description"`
}

type UpdateDictionaryBody struct {
	CreateDictionaryBody
	DictionaryID
}

type FindDictionaryBody struct {
	CreateDictionaryBody
	ID string `json:"id" form:"id"`
}

type FindDictionaryPagingBody struct {
	FindDictionaryBody
	Pagination
}

type DeleteDictionaryBody struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}
