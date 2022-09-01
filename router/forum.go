package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initForumRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	forumGroup := apiGroup.Group("/forum")

	forumGroup.Use(middleware.IsLogined)
	{
		forumGroup.POST("/add", controller.PostAddController)
		forumGroup.GET("/list", controller.PostListController)
		forumGroup.POST("/comment/:post_id", controller.CommentAddController)
	}

	return forumGroup
}
