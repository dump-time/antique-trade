package controller

import (
	"strconv"

	"github.com/dump-time/antique-trade/services"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type PostAddRequest struct {
	Content  string  `json:"content" binding:"required"`
	ImageURL *string `json:"image_url"`
	FileURL  *string `json:"file_url"`
}

func PostAddController(ctx *gin.Context) {
	uid := sessions.Default(ctx).Get("id").(uint)

	var req PostAddRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	if err := services.CreatePost(uint(uid), req.Content, req.ImageURL, req.FileURL); err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, nil)
}

func PostListController(ctx *gin.Context) {
	posts, err := services.ListPost()
	if err != nil {
		util.InternalErrResp(ctx)
		return
	}

	var resp []map[string]interface{}

	for _, post := range posts {
		comments := post.Comments

		var commentResp []map[string]interface{}

		for _, comment := range comments {
			commentResp = append(commentResp, map[string]interface{}{
				"id":      comment.ID,
				"content": comment.Content,
				"writer":  comment.WriterID,
			})
		}

		resp = append(resp, map[string]interface{}{
			"id":        post.ID,
			"content":   post.Content,
			"image_url": post.ImageURL,
			"file_url":  post.FileURL,
			"writer":    post.WriterID,
			"comments":  commentResp,
		})
	}

	util.SuccessResp(ctx, resp)
}

type CommentAddRequest struct {
	Content string `json:"content" binding:"required"`
}

func CommentAddController(ctx *gin.Context) {
	uid := sessions.Default(ctx).Get("id").(uint)

	postIDStr := ctx.Param("post_id")
	postID, err := strconv.Atoi(postIDStr)

	if err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	var req CommentAddRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	if err := services.CreateComment(uint(uid), req.Content, uint(postID)); err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, nil)
}
