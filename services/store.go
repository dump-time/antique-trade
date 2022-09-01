package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
)

func CreateProduct(uid uint, title string, price float64, description string, primaryImageURL string) error {
	result := global.DB.Create(&model.Product{
		Title:           title,
		Price:           price,
		Description:     description,
		PrimaryImageURL: primaryImageURL,
		UserID:          uid,
	})

	return result.Error
}

type ListAllProductData struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	PrimaryImageURL string  `json:"primary_image_url"`
	UserID          uint    `json:"user_id"`
}

func ListAllProducts() ([]ListAllProductData, error) {
	var products []ListAllProductData
	result := global.DB.Model(&model.Product{}).Find(&products)

	return products, result.Error
}

type ListByUIDProductData struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	PrimaryImageURL string  `json:"primary_image_url"`
}

func ListProductsByUID(uid uint) ([]ListByUIDProductData, error) {
	var products []ListByUIDProductData
	result := global.DB.Model(&model.Product{}).Where("user_id = ?", uid).Find(&products)

	return products, result.Error
}
