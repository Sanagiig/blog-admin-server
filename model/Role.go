package model

type Role struct {
	Base
	Name        string       `gorm:"type:varchar(30);not null;uniqueIndex" json:"name,omitempty"`
	Code        string       `gorm:"type:varchar(20);not null;uniqueIndex" json:"code,omitempty"`
	Description string       `gorm:"type:varchar(200);" json:"description,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permission" json:"permissions,omitempty"`
	Users       []User       `gorm:"many2many:user_role" json:"users,omitempty"`
}
