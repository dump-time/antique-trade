package model

import "gorm.io/gorm"

type Forum struct {
	gorm.Model

	Content  string // 内容
	ImageURL string // 图片（1）链接
	FileURL  string // 文件（1）链接

	WriterID uint // 发布人 ID

	Comments []*Comment `gorm:"foreignkey:ForumID"`
}
