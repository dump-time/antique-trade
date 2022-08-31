package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
)

func RegisterUser(username string, password string, role string) error {
	user := model.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	result := global.DB.Create(&user)
	return result.Error
}
