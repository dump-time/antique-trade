package controller

import (
	"log"
	"strconv"

	"github.com/dump-time/antique-trade/services"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CartAddProductController(ctx *gin.Context) {
	uid := sessions.Default(ctx).Get("id").(uint)

	productIDStr := ctx.Param("product_id")

	log.Default().Println(productIDStr)

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	if err := services.AddProductIntoUserCart(uid, uint(productID)); err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, nil)
}

func CartDeleteProductController(ctx *gin.Context) {
	uid := sessions.Default(ctx).Get("id").(uint)

	productIDStr := ctx.Param("product_id")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		util.ParamsErrResp(ctx)
		return
	}

	if err := services.DeleteProductFromUserCart(uid, uint(productID)); err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, nil)
}

func CartListController(ctx *gin.Context) {
	uid := sessions.Default(ctx).Get("id").(uint)

	products, err := services.ListUserCart(uid)
	if err != nil {
		util.InternalErrResp(ctx)
		return
	}

	util.SuccessResp(ctx, products)
}
