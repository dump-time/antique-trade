package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initStoreRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	storeRouter := apiGroup.Group("/store")

	{
		storeRouter.POST("/add", middleware.IsLogined, controller.ProductAddController)
		storeRouter.GET("/list", controller.ProductsListAllController)
		storeRouter.GET("/list/:uid", controller.ProductListByUIDController)
		storeRouter.GET("/category/:category", controller.ProductListByCategoryController)
	}

	return storeRouter
}
