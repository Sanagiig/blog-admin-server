package request

type IDS struct {
	IDS []string `json:"ids" binding:"required"`
}

type CreateRange struct {
	StartCreateTime string `json:"startCreateTime"`
	EndCreateTime   string `json:"endCreateTime"`
}

type UpdateRange struct {
	StartUpdateTime string `json:"startUpdateTime"`
	EndUpdateTime   string `json:"endUpdateTime"`
}

type Pagination struct {
	Page     int `json:"page" form:"page" binding:"required,min=1"`
	PageSize int `json:"pageSize" form:"pageSize"  binding:"required,min=1"`
}
