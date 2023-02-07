package model

type Tag struct {
	Base
	Name        string    `gorm:"type:varchar(30);not null;uniqueIndex" json:"name"`
	Type        string    `gorm:"type:varchar(36);not null" json:"type"`
	Description string    `gorm:"type:varchar(200);" json:"description"`
	Articles    []Article `gorm:"many2many:article_tag" json:"-"`
}
