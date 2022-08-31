package router

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/log"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

// init global router
func init() {
	R = gin.New()

	// Setup trusted reverse proxies
	if err := R.SetTrustedProxies(global.Config.Serv.TrustedProxies); err != nil {
		log.Fatal(err)
		return
	}

	// Log formatter
	R.Use(gin.LoggerWithFormatter(log.Formatter))

	// Panic auto recovery & return 500
	R.Use(gin.Recovery())

	// Setup routers
	v1 := R.Group("/api/v1")
	initBasicRouter(v1)
}
