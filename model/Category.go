package model

type Category struct {
	Base
	ParentID    string `gorm:"type:varchar(36);default:null" json:"parentId,omitempty"`
	Name        string `gorm:"type:varchar(30);not null" json:"name,omitempty"`
	Type        string `gorm:"type:varchar(36);not null;index" json:"type,omitempty"`
	Description string `gorm:"type:varchar(200);" json:"description,omitempty"`
}
