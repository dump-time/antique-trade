package router

import (
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-gonic/gin"
)

func initBasicRouter(apiGroup *gin.RouterGroup) *gin.RouterGroup {
	apiGroup.GET("/ping", func(ctx *gin.Context) {
		util.SuccessResp(ctx, gin.H{
			"ping": "pong",
		})
	})
	return apiGroup
}
