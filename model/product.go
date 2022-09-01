package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Title           string  // 标题
	Price           float64 // 价格
	Description     string  // 描述
	PrimaryImageURL string  // 主图链接

	UserID uint // 发布人 ID
}
