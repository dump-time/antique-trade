package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initForumRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	forumGroup := apiGroup.Group("/forum")

	{
		forumGroup.POST("/add", middleware.IsLogined, controller.PostAddController)
		forumGroup.GET("/list", controller.PostListController)
		forumGroup.POST("/comment/:post_id", middleware.IsLogined, controller.CommentAddController)
	}

	return forumGroup
}
