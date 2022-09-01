package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
	"gorm.io/gorm"
)

func AddProductIntoUserCart(userID uint, productID uint) error {
	result := global.DB.
		Model(&model.User{Model: gorm.Model{ID: userID}}).
		Association("CartProducts").
		Append(&model.Product{Model: gorm.Model{ID: productID}})
	return result
}

func DeleteProductFromUserCart(userID uint, productID uint) error {
	result := global.DB.
		Model(&model.User{Model: gorm.Model{ID: userID}}).
		Association("CartProducts").
		Delete(&model.Product{Model: gorm.Model{ID: productID}})
	return result
}

type CartProduct struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	PrimaryImageURL string  `json:"primary_image_url"`
}

func ListUserCart(userID uint) ([]CartProduct, error) {
	var products []CartProduct
	result := global.DB.
		Model(&model.User{Model: gorm.Model{ID: userID}}).
		Association("CartProducts").
		Find(&products)
	return products, result
}
