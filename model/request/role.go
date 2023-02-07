package request

type CreateRoleBody struct {
	Name        string   `json:"name" form:"name"`
	Code        string   `json:"code" form:"code"`
	Description string   `json:"description" form:"description"`
	Permissions []string `json:"permissions" form:"permissions"`
}

type FindRoleBody struct {
	CreateRoleBody
	ID string `json:"id" form:"id"`
}

type FindRolePagingBody struct {
	FindRoleBody
	Pagination
}

type UpdateRoleBody struct {
	ID string `json:"id"`
	CreateRoleBody
}

type DeleteRoleBody struct {
	IDs []string `json:"ids"`
}
