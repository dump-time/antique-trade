package services

import (
	"database/sql"

	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchUserProfile(id uint) (*model.User, *gorm.DB) {
	var user model.User
	result := global.DB.Where("id = ?", id).Take(&user)
	return &user, result
}

func FetchProfileList(userID sql.NullInt64, role string) ([]gin.H, *gorm.DB) {
	type profileResult struct {
		model.User
		UserID     sql.NullInt64 // Fans ID
		FollowedID sql.NullInt64 // Follwed people ID
	}
	var profileList []profileResult

	result := global.DB.
		Model(&model.User{}).
		Select("users.*, user_followed.*").
		Joins("left join user_followed on users.id = user_followed.followed_user_id and user_followed.user_id = ?", userID).
		Where("users.`role` = ?", role).Find(&profileList)

	var profileData []gin.H
	var data gin.H
	for _, profile := range profileList {
		data = gin.H{
			"id":          profile.User.Model.ID,
			"username":    profile.User.Username,
			"name":        profile.User.Name,
			"tel":         profile.User.Tel,
			"shortTitle":  profile.User.ShortTitle,
			"jobTitle":    profile.User.JobTitle,
			"point":       profile.User.Point,
			"sex":         profile.User.Sex,
			"avatarUrl":   profile.User.AvatarUrl,
			"role":        profile.User.Role,
			"description": profile.User.Description,
		}
		if profile.UserID.Valid {
			data["followed"] = true
		} else {
			data["followed"] = false
		}

		profileData = append(profileData, data)
	}
	return profileData, result
}
