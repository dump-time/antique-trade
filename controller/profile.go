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

type EditProfileData struct {
	AvatarURL   string `json:"avatarUrl"`
	Description string `json:"description"`
	JobTitle    string `json:"jobTitle"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Sex         string `json:"sex"`
	ShortTitle  string `json:"shortTitle"`
	Tel         string `json:"tel"`
	Username    string `json:"username"`
}

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
	var userIDNullable sql.NullInt64

	userID := sessions.Default(context).Get("id")
	role := context.Param("role")
	if userID == nil {
		userIDNullable.Valid = false
	} else {
		userIDNullable.Valid = true
		userIDNullable.Int64 = int64(userID.(uint))
	}

	log.Info(userIDNullable) // debug

	profileData, result := services.FetchProfileList(userIDNullable, role)
	if result.Error != nil {
		log.Error(result.Error)
		if result.RowsAffected == 0 {
			util.NotFoundResp(context)
		} else {
			util.InternalErrResp(context)
		}
		return
	}

	// Generate
	util.SuccessResp(context, profileData)
}

func EditProfileController(context *gin.Context) {
	userID := sessions.Default(context).Get("id").(uint)
	var editProfileData EditProfileData
	if err := context.ShouldBindJSON(&editProfileData); err != nil {
		util.ParamsErrResp(context)
		log.Error(err)
		return
	}

	profileData := map[string]any{
		"username":    editProfileData.Username,
		"name":        editProfileData.Name,
		"tel":         editProfileData.Tel,
		"short_title": editProfileData.ShortTitle,
		"job_title":   editProfileData.JobTitle,
		"sex":         editProfileData.Sex,
		"avatar_url":  editProfileData.AvatarURL,
		"role":        editProfileData.Role,
		"description": editProfileData.Description,
	}

	if result := services.EditProfile(userID, profileData); result.Error != nil {
		log.Error(result.Error)
		if result.RowsAffected == 0 {
			util.NotFoundResp(context)
		} else {
			util.InternalErrResp(context)
		}
		return
	}

	util.SuccessResp(context, nil)
}

func ListFavoritePeopleController(context *gin.Context) {
	userID := sessions.Default(context).Get("id").(uint)
	favoritePeopleData, result := services.ListFavoritePeople(userID)
	if result.Error != nil {
		log.Error(result.Error)
		if result.RowsAffected == 0 {
			util.NotFoundResp(context)
		} else {
			util.InternalErrResp(context)
		}
		return
	}

	util.SuccessResp(context, favoritePeopleData)
}
