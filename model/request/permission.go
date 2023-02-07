package request

type PermissionID struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreatePermissionBody struct {
	Name        string `json:"name" form:"name"`
	Code        string `json:"code" form:"code"`
	Type        string `json:"type" form:"type"`
	Value       string `json:"value" form:"value"`
	Description string `json:"description" form:"description"`
}

type UpdatePermissionBody struct {
	CreatePermissionBody
	PermissionID
}

type FindPermissionBody struct {
	CreatePermissionBody
	ID    string   `json:"id" form:"id"`
	Roles []string `json:"roles" form:"roles"`
}

type FindPermissionPagingBody struct {
	FindPermissionBody
	Pagination
	Type []string `json:"type" form:"type"`
}

type DeletePermissionBody struct {
	IDs []string `json:"ids" form:"ids" binding:"required"`
}

type PatchRolePermissionBody struct {
	RoleIds       []string `json:"roleIds"  binding:"required"`
	PermissionIds []string `json:"permissionIds"  binding:"required"`
}
