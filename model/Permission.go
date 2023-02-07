package model

type Permission struct {
	Base
	Name        string `gorm:"type:varchar(30);not null;uniqueIndex" json:"name"`
	Code        string `gorm:"type:varchar(20);not null;uniqueIndex" json:"code"`
	Type        string `gorm:"type:varchar(30);not null;index;default:common" json:"type"`
	Value       string `gorm:"type:varchar(200);not null" json:"value"`
	Description string `gorm:"type:varchar(200);" json:"description"`
	Roles       []Role `gorm:"many2many:role_permission" json:"roles,omitempty"`
}
