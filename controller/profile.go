package controller

import (
	"strconv"

	"github.com/dump-time/antique-trade/log"
	"github.com/dump-time/antique-trade/services"
	"github.com/dump-time/antique-trade/util"
	"github.com/gin-gonic/gin"
)

func ProfileDetailController(context *gin.Context) {
	user_id_str := context.Param("user_id")
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		util.ParamsErrResp(context)
		log.Error(err)
		return
	}

	user, result := services.FetchUserProfile(uint(user_id))
	if result.Error != nil {
		log.Error(result.Error)
		if result.RowsAffected == 0 {
			util.NotFoundResp(context)
		} else {
			util.InternalErrResp(context)
		}
		return
	}

	util.SuccessResp(context, gin.H{
		"username":    user.Username,
		"name":        user.Name,
		"tel":         user.Tel,
		"shortTitle":  user.ShortTitle,
		"jobTitle":    user.JobTitle,
		"point":       user.Point,
		"sex":         user.Sex,
		"avatarUrl":   user.AvatarUrl,
		"role":        user.Role,
		"description": user.Description,
	})
}

func ProfileListController(context *gin.Context) {

}
