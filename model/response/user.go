package response

type UserBaseInfo struct {
	Username string `json:"username,omitempty"`
	Nickname string ` json:"nickname"`
	Sex      string ` json:"sex,omitempty"`
	Email    string ` json:"email,omitempty"`
	Phone    string ` json:"phone,omitempty"`
	Password string ` json:"-"`
}
