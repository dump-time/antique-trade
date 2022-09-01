package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/gin-gonic/gin"
)

func initForumRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	forumGroup := apiGroup.Group("/forum")

	{
		forumGroup.POST("/add", controller.PostAddController)
		forumGroup.POST("/list", controller.PostListController)
		forumGroup.POST("/comment/add/:post_id", controller.CommentAddController)
	}

	return forumGroup
}
