package controller

import (
	"strconv"

	"github.com/dump-time/antique-trade/services"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type ProductAddRequest struct {
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	PrimaryImageURL string  `json:"primary_image_url"`
}

func ProductAddController(ctx *gin.Context) {
	uid := sessions.Default(ctx).Get("id").(uint)

	var req ProductAddRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	if err := services.CreateProduct(uid, req.Title, req.Price, req.Description, req.Category, req.PrimaryImageURL); err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, nil)
}

func ProductsListAllController(ctx *gin.Context) {
	products, err := services.ListAllProducts()
	if err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, products)
}

func ProductListByUIDController(ctx *gin.Context) {
	uidStr := ctx.Param("uid")

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	products, err := services.ListProductsByUID(uint(uid))
	if err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, products)
}

func ProductListByCategoryController(ctx *gin.Context) {
	category := ctx.Param("category")

	products, err := services.ListProductsByCategory(category)
	if err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, products)
}
