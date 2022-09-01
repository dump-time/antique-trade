package controller

import (
	"strconv"

	"github.com/dump-time/antique-trade/log"
	"github.com/dump-time/antique-trade/services"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ArticleData struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	PrimaryImageURL string `json:"primaryImageURL"`
}

func AddArticleController(context *gin.Context) {
	var articleData ArticleData
	if err := context.ShouldBindJSON(&articleData); err != nil {
		util.ParamsErrResp(context)
		log.Error(err)
		return
	}
	userID := sessions.Default(context).Get("id").(uint)

	if result := services.AddArticle(
		articleData.Title,
		articleData.Content,
		articleData.PrimaryImageURL,
		userID,
	); result.Error != nil {
		log.Error(result.Error)
		util.InternalErrResp(context)
		return
	}

	util.SuccessResp(context, nil)
}

func MarkArticleController(context *gin.Context) {
	articleID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	userID := sessions.Default(context).Get("id").(uint)

	if err := services.MarkArticle(userID, uint(articleID)); err != nil {
		util.InternalErrResp(context)
		log.Error(err)
		return
	}

	util.SuccessResp(context, nil)
}

func UnMarkArticleController(context *gin.Context) {

}
