package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/gin-gonic/gin"
)

func initBasicRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	// login & register
	apiGroup.GET("/login", controller.LoginController)
	apiGroup.POST("/register", controller.RegisterController)
	return apiGroup
}
