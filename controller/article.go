package controller

import (
	"database/sql"
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

	if err := services.AddArticle(
		articleData.Title,
		articleData.Content,
		articleData.PrimaryImageURL,
		userID,
	); err != nil {
		log.Error(err)
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
	articleID, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		log.Error(err)
		util.ParamsErrResp(context)
		return
	}
	userID := sessions.Default(context).Get("id").(uint)

	if err := services.UnMarkArticle(userID, uint(articleID)); err != nil {
		util.InternalErrResp(context)
		log.Error(err)
		return
	}

	util.SuccessResp(context, nil)
}

func ListFavoriteArticleController(context *gin.Context) {
	userID := sessions.Default(context).Get("id").(uint)
	articles, err := services.ListFavoriteArticle(userID)
	if err != nil {
		log.Error(err.Error)
		util.InternalErrResp(context)
		return
	}

	util.SuccessResp(context, articles)
}

func ListArticlesController(context *gin.Context) {
	var userID sql.NullInt64

	userIDData := sessions.Default(context).Get("id")
	if userIDData == nil {
		userID.Valid = false
	} else {
		userID.Int64 = int64(userIDData.(uint))
		userID.Valid = true
	}

	articles, err := services.ListArticles(userID)
	if err != nil {
		log.Error(err.Error)
		util.InternalErrResp(context)
		return
	}

	util.SuccessResp(context, articles)
}
