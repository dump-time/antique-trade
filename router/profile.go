package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/gin-gonic/gin"
)

func initProfileRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	profileRouter := apiGroup.Group("/profile")
	{
		profileRouter.GET("/detail/:user_id", controller.ProfileDetailController)
	}
	return apiGroup
}
