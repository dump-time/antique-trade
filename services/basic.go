package services

import (
	"errors"

	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
)

func RegisterUser(username string, password string, role string, sex string, tel string) error {
	user := model.User{
		Username: username,
		Password: password,
		Role:     role,
		Sex:      sex,
		Tel:      tel,
		Name:     username,
	}
	result := global.DB.Create(&user)
	return result.Error
}

func Login(username string, password string) (model.User, error) {
	var user model.User
	global.DB.Where("username = ?", username).Take(&user)
	if user.Password != password {
		return model.User{}, errors.New("password error")
	}

	return user, nil
}
