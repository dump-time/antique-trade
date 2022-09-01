package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model

	Title          string // 标题
	Content        string // 内容
	PimaryImageURL string // 主图链接内容
	WriterID       uint   // 发布人 ID
}
