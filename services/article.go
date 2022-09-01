package services

import (
	"database/sql"

	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddArticle(title string, content string, primaryImageUrl string, writerID uint) error {
	article := model.Article{
		Title:          title,
		Content:        content,
		PimaryImageURL: primaryImageUrl,
		UserID:         writerID,
	}
	if err := global.DB.Create(&article).Error; err != nil {
		return err
	}

	return nil
}

func MarkArticle(userID uint, articleID uint) error {
	err := global.DB.
		Model(&model.User{Model: gorm.Model{ID: userID}}).
		Association("FavoriteArticles").
		Append(&model.Article{Model: gorm.Model{ID: articleID}})
	return err
}

func UnMarkArticle(userID uint, articleID uint) error {
	err := global.DB.
		Model(&model.User{Model: gorm.Model{ID: userID}}).
		Association("FavoriteArticles").
		Delete(&model.Article{Model: gorm.Model{ID: articleID}})
	return err
}

func ListFavoriteArticle(userID uint) ([]gin.H, error) {
	type resultData struct {
		model.Article
		model.User
	}
	var result []resultData

	err := global.DB.Model(&model.User{}).
		Select("articles.*, users.*").
		Joins("inner join articles on articles.user_id = users.id inner join user_article on user_article.article_id = articles.id and user_article.user_id = ?", userID).
		Find(&result).Error
	var favoriteArticleData []gin.H

	for _, data := range result {
		favoriteArticleData = append(favoriteArticleData, gin.H{
			"id":              data.Article.Model.ID,
			"title":           data.Article.Title,
			"content":         data.Article.Content,
			"primaryImageURL": data.Article.PimaryImageURL,
			"writer": gin.H{
				"id":          data.User.Model.ID,
				"username":    data.User.Username,
				"name":        data.User.Name,
				"tel":         data.User.Tel,
				"shortTitle":  data.User.ShortTitle,
				"jobTitle":    data.User.JobTitle,
				"point":       data.User.Point,
				"sex":         data.User.Sex,
				"avatarUrl":   data.User.AvatarUrl,
				"role":        data.User.Role,
				"description": data.User.Description,
			},
		})
	}

	return favoriteArticleData, err
}

func ListArticles(userID sql.NullInt64) ([]gin.H, error) {
	type resultData struct {
		model.Article
		model.User
		UserID sql.NullInt64
	}
	var result []resultData

	err := global.DB.Select("articles.*, users.*, user_article.user_id").
		Model(&model.Article{}).
		Joins("left join user_article on articles.id = user_article.article_id AND user_article.user_id = ? inner join users on articles.user_id = users.id ", userID).
		Find(&result).Error

	var articleDatas []gin.H
	var articleData gin.H
	for _, data := range result {
		articleData = gin.H{
			"id":              data.Article.Model.ID,
			"title":           data.Article.Title,
			"content":         data.Article.Content,
			"primaryImageURL": data.Article.PimaryImageURL,
			"writer": gin.H{
				"id":          data.User.Model.ID,
				"username":    data.User.Username,
				"name":        data.User.Name,
				"tel":         data.User.Tel,
				"shortTitle":  data.User.ShortTitle,
				"jobTitle":    data.User.JobTitle,
				"point":       data.User.Point,
				"sex":         data.User.Sex,
				"avatarUrl":   data.User.AvatarUrl,
				"role":        data.User.Role,
				"description": data.User.Description,
			},
		}
		if data.UserID.Valid {
			articleData["isMarked"] = true
		} else {
			articleData["isMarked"] = false
		}
		articleDatas = append(articleDatas, articleData)
	}

	return articleDatas, err
}

func GetArticleDetail(userID sql.NullInt64, articleID int) (gin.H, error) {
	var resultArticle model.Article
	result := global.DB.Model(&model.Article{}).Preload("Products").Where("id = ?", articleID).Take(&resultArticle)
	if result.Error != nil {
		return nil, result.Error
	}
	var user model.User
	if err := global.DB.Model(&resultArticle).Association("User").Find(&user); err != nil {
		return nil, err
	}

	// Extract products data
	var products []gin.H
	for _, product := range resultArticle.Products {
		products = append(products, gin.H{
			"id":              product.ID,
			"title":           product.Title,
			"price":           product.Price,
			"description":     product.Description,
			"primaryImageURL": product.PrimaryImageURL,
			"category":        product.Category,
			"userID":          product.UserID,
		})
	}

	// Generate result data
	var resultData = gin.H{
		"title":           resultArticle.Title,
		"content":         resultArticle.Content,
		"primaryImageURL": resultArticle.PimaryImageURL,
		"writer": gin.H{
			"id":          user.Model.ID,
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
		},
	}
	resultData["products"] = products

	count := global.DB.
		Model(&model.User{Model: gorm.Model{ID: uint(userID.Int64)}}).
		Where("user_article.user_id = ? and user_article.article_id = ?", userID, articleID).
		Association("FavoriteArticles").Count()
	if count == 0 {
		resultData["isMarked"] = false
	} else {
		resultData["isMarked"] = true
	}

	return resultData, result.Error
}
