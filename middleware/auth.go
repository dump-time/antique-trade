package middleware

import (
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogined(context *gin.Context) {
	userID := sessions.Default(context).Get("id")
	if userID == nil {
		util.NotLoginResp(context)
		context.Abort()
	}
}
