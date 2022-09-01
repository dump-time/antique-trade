package router

import (
	"net/http"

	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initImageRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	imageRouter := apiGroup.Group("/image")
	{
		imageRouter.POST("/upload", middleware.IsLogined, controller.UploadImageController)
		imageRouter.StaticFS("/show", http.Dir("./data/image"))
	}
	return imageRouter
}
