package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initArticleRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	articleRouter := apiGroup.Group("/article")
	{
		articleRouter.POST("/add", middleware.IsLogined, controller.AddArticleController)
		articleRouter.GET("/mark/:id", middleware.IsLogined, controller.MarkArticleController)
		articleRouter.GET("/unmark/:id", middleware.IsLogined, controller.UnMarkArticleController)
		articleRouter.GET("/list-favorite", middleware.IsLogined, controller.ListFavoriteArticleController)
	}
	return articleRouter
}
