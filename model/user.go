package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username    string
	Password    string
	Role        string
	Name        string // 真实姓名
	Tel         string // 电话
	ShortTitle  string // 称谓
	JobTitle    string // 称谓
	Point       float64
	Sex         string
	AvatarUrl   string // 头像
	Description string // 简介

	// Has many relations
	Articles []*Article

	// Many to many relations
	FavoriteArticles []*Article `gorm:"many2many:user_article"`
	FollowedUsers    []*User    `gorm:"many2many:user_followed"`
	CartProducts     []*Product `gorm:"many2many:user_cart"`
}
