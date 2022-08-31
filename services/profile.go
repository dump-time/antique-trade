package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
	"gorm.io/gorm"
)

func FetchUserProfile(id uint) (*model.User, *gorm.DB) {
	var user model.User
	result := global.DB.Where("id = ?", id).Take(&user)
	return &user, result
}
