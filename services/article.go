package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
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
		Association("FavoritePosts").
		Append(&model.Article{Model: gorm.Model{ID: articleID}})
	return err
}
