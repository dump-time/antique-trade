package services

import (
	"github.com/dump-time/antique-trade/global"
	"github.com/dump-time/antique-trade/model"
)

func CreateProduct(uid uint, title string, price float64, description string, category string, primaryImageURL string) error {
	result := global.DB.Create(&model.Product{
		Title:           title,
		Price:           price,
		Description:     description,
		Category:        category,
		PrimaryImageURL: primaryImageURL,
		UserID:          uid,
	})

	return result.Error
}

type ListAllProduct struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	PrimaryImageURL string  `json:"primary_image_url"`
	UserID          uint    `json:"user_id"`
}

func ListAllProducts() ([]ListAllProduct, error) {
	var products []ListAllProduct
	result := global.DB.Model(&model.Product{}).Find(&products)

	return products, result.Error
}

type ListByUIDProduct struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	PrimaryImageURL string  `json:"primary_image_url"`
}

func ListProductsByUID(uid uint) ([]ListByUIDProduct, error) {
	var products []ListByUIDProduct
	result := global.DB.Model(&model.Product{}).Where("user_id = ?", uid).Find(&products)

	return products, result.Error
}

type ListByCategoryProduct struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	Price           float64 `json:"price"`
	Description     string  `json:"description"`
	PrimaryImageURL string  `json:"primary_image_url"`
	UserID          uint    `json:"user_id"`
}

func ListProductsByCategory(category string) ([]ListByCategoryProduct, error) {
	var products []ListByCategoryProduct
	result := global.DB.Model(&model.Product{}).Where("category = ?", category).Find(&products)

	return products, result.Error
}
