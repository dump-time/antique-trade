package services

import (
	"log"

	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
)

func CreatePost(uid uint, content string, imageURL *string, fileURL *string) error {
	result := global.DB.Create(
		&model.Post{
			Content:  content,
			WriterID: uid,
			ImageURL: imageURL,
			FileURL:  fileURL,
		},
	)

	return result.Error
}

func ListPost() ([]model.Post, error) {
	var posts []model.Post

	result := global.DB.Model(&model.Post{}).Preload("Comments").Find(&posts)

	log.Println(posts)

	return posts, result.Error
}

func CreateComment(uid uint, content string, post_id uint) error {
	result := global.DB.Create(
		&model.Comment{
			Content:  content,
			WriterID: uid,
			PostID:   post_id,
		},
	)

	return result.Error
}
