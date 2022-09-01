package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initProfileRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	profileRouter := apiGroup.Group("/profile")
	{
		profileRouter.GET("/detail/:user_id", controller.ProfileDetailController)
		profileRouter.GET("/list/:role", controller.ProfileListController)
		profileRouter.POST("/edit", middleware.IsLogined, controller.EditProfileController)
		profileRouter.GET("/favorite-people", middleware.IsLogined, controller.ListFavoritePeopleController)
		profileRouter.GET("/follow/:id", middleware.IsLogined, controller.FollowController)
		profileRouter.GET("/unfollow/:id", middleware.IsLogined, controller.UnfollowController)
	}
	return profileRouter
}
