package request

type RegisterUserBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginBody struct {
	RegisterUserBody
	CheckCode string `json:"checkCode" binding:"required"`
}

type BasicUserBody struct {
	Nickname string   `json:"nickname"`
	Phone    string   `json:"phone"`
	Sex      string   `json:"sex"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

type UserInfoBody struct {
	Description string `json:"description"`
	Profile     string `json:"profile"`
}

type CreateUserBody struct {
	RegisterUserBody
	BasicUserBody
	UserInfoBody
}

type UpdateUserBody struct {
	BasicUserBody
	UpdateUserInfoBody
	ID string `json:"id"`
}

type FindUserBody struct {
	ID       string   `json:"id"  form:"id"`
	Username string   `json:"username" form:"username"`
	Email    string   `json:"email" form:"email"`
	Phone    string   `json:"phone" form:"phone"`
	Roles    []string `json:"roles" form:"roles"`
}

type FindUserPagingBody struct {
	FindUserBody
	Pagination
}

type UpdateUserInfoBody struct {
	UserInfoBody
}

type PatchUserRoleBody struct {
	UserIDs []string `json:"userIds"`
	RoleIDs []string `json:"roleIds"`
}
