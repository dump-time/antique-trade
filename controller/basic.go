package controller

import (
	"github.com/dump-time/antique-trade/log"
	"github.com/dump-time/antique-trade/services"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-gonic/gin"
)

type RegisterData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginController(context *gin.Context) {

}

func RegisterController(context *gin.Context) {
	var registerData RegisterData
	if err := context.ShouldBindJSON(&registerData); err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}

	if err := services.RegisterUser(registerData.Username, registerData.Password, registerData.Role); err != nil {
		log.Error(err)
		util.InternalErrResp(context)
		return
	}

	util.SuccessResp(context, nil)
}
