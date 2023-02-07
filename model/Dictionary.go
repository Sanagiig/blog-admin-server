package model

type Dictionary struct {
	Base
	ParentID    string `gorm:"type:varchar(36);default:null" json:"parentId"`
	Code        string `gorm:"type:varchar(20);not null;uniqueIndex" json:"code"`
	Name        string `gorm:"type:varchar(30);not null" json:"name"`
	Description string `gorm:"type:varchar(200);" json:"description"`
}
