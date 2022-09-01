package router

import (
	"github.com/dump-time/antique-trade/controller"
	"github.com/dump-time/antique-trade/middleware"
	"github.com/gin-gonic/gin"
)

func initCartRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	cartGroup := apiGroup.Group("/cart")

	cartGroup.Use(middleware.IsLogined)
	{
		cartGroup.GET("/add/:product_id", controller.CartAddProductController)
		cartGroup.GET("/remove/:product_id", controller.CartDeleteProductController)
		cartGroup.GET("/list", controller.CartListController)
	}

	return cartGroup
}
