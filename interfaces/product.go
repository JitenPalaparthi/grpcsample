// Package interfaces is to define all interface methods for each model type
// Author : readyGo "JitenP@Outlook.Com"
// This code is generated by readyGo. You are free to make amendments as and where required
package interfaces

import (
	"grpcsample/models"
)


// ProductInterface is to define all product related methods
type ProductInterface interface {
    CreateProduct(product *models.Product)(string, error)
	UpdateProductByID(id string, data map[string]interface{}) (string, error)
	DeleteProductByID(id string) (string, error)
	GetProductByID(id string) (*models.Product, error)
	GetAllProducts(skip int64, limit int64, selector interface{}) ([]models.Product, error)
	GetAllProductsBy(key string, value interface{}, skip int64, limit int64) ([]models.Product, error)
    // Write any additional methods and implement them in database package
}
