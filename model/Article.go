package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title      string     `gorm:"type:varchar(100);not null" json:"title"`
	Desc       string     `gorm:"type:varchar(255)" json:"desc"`
	Content    string     `gorm:"type:longtext;not null" json:"content"`
	Img        string     `gorm:"type:varchar(100)" json:"img"`
	Categories []Category `gorm:"many2many:article_categories;" json:"categories"`
}
