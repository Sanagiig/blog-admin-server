package model

type User struct {
	Base
	Username string   `gorm:"type:varchar(20);not null;uniqueIndex" json:"username,omitempty"`
	Nickname string   `gorm:"type:varchar(20);not null;uniqueIndex" json:"nickname"`
	Sex      string   `gorm:"type:char(10);check:sex in ('man','woman','unknown');default:'unkonw';comment:('man','womanm','unkonw') => (男,女,保密)" json:"sex,omitempty"`
	Email    string   `gorm:"type:varchar(50);uniqueIndex;default:null" json:"email,omitempty"`
	Phone    string   `gorm:"type:varchar(15);uniqueIndex;default:null" json:"phone,omitempty"`
	Password string   `gorm:"type:char(64);not null" json:"-"`
	Roles    []Role   `gorm:"many2many:user_role" json:"roles"`
	UserInfo UserInfo `gorm:"foreignKey:UserID;references:id" json:"userInfo,omitempty,inline" `
}

type UserInfo struct {
	UserID      string `json:"-"`
	Description string `gorm:"type:varchar(100);" json:"description"`
	Profile     string `gorm:"type:varchar(200);" json:"profile"`
}
