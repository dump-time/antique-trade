package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	Content string // 内容

	WriterID uint // 发布人 ID

	ForumID uint // 外键
}
