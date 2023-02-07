package model

type Article struct {
	Base
	AuthorID   string   `json:"authorID"`
	Author     User     `gorm:"foreignkey:AuthorID;references:ID" json:"author,omitempty"`
	CategoryID string   `json:"-"`
	Category   Category `json:"category"`

	Tags        []Tag  `gorm:"many2many:article_tag" json:"tags"`
	Title       string `gorm:"type:varchar(30);not null" json:"title"`
	Description string `gorm:"type:varchar(100);" json:"description,omitempty"`
	Content     string `gorm:"type:mediumtext;not null;" json:"content,omitempty"`
}
