package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddArticle(title string, content string, primaryImageUrl string, writerID uint) *gorm.DB {
	result := global.DB.Create(&model.Article{
		Title:          title,
		Content:        content,
		PimaryImageURL: primaryImageUrl,
		WriterID:       writerID,
	})

	return result
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
	user := model.User{Model: gorm.Model{ID: userID}}
	var favoriteArticles []model.Article
	err := global.DB.Model(&user).Association("FavoriteArticles").Find(&favoriteArticles)
	var favoriteArticleData []gin.H

	for _, article := range favoriteArticles {
		favoriteArticleData = append(favoriteArticleData, gin.H{
			"id":              article.Model.ID,
			"title":           article.Title,
			"content":         article.Content,
			"primaryImageURL": article.PimaryImageURL,
		})
	}

	return favoriteArticleData, err
}
