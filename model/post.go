package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title          string // 标题
	Content        string // 内容
	PimaryImageURL string // 主图链接内容
	ImageURLs      string // 其他图片链接
	PosterID       uint   // 发布人 ID
}
