// Package database contains all database related methods
// Author : readyGo "JitenP@Outlook.Com"
// This code is generated by readyGo. You are free to make amendments as and where required
package database

import (
	"grpcsample/models"
	"gorm.io/gorm"
	"strconv"
)

// ProductDB is to maintain database related methods
type ProductDB struct {
	DB *Database
}

// CreateProduct is to insert a record in to the database. 
func (p *ProductDB) CreateProduct(product *models.Product) (result string, err error) {
	err = p.DB.Client.(*gorm.DB).AutoMigrate(product)
	if err != nil {
		return "", err
	}
	res := p.DB.Client.(*gorm.DB).Create(product)
	if res.Error != nil {
		return "", res.Error
	}

	result = strconv.Itoa(product.ID)
	return result, nil
}

// UpdateProductByID is to update a record in the database.The first param is to identify the record and the second is the list of the fields to update
func (p *ProductDB)  UpdateProductByID(id string, data map[string]interface{}) (result string, err error) {
	res := p.DB.Client.(*gorm.DB).Model(&models.Product{}).Where("id=?", id).Updates(data)
	if res.Error != nil {
		return "", res.Error
	}
	result = strconv.FormatInt(res.RowsAffected, 10)
	return result, nil
}

// DeleteProductByID is to hard delete a record from the database provided by id
func (p *ProductDB)  DeleteProductByID(id string) (result string, err error) {
	res := p.DB.Client.(*gorm.DB).Delete(&models.Product{}, &id)
	if res.Error != nil {
		return "", res.Error
	}
	result = strconv.FormatInt(res.RowsAffected, 10)
	return result, nil
}

// GetProductByID is to fetch a record from database provided by id
func  (p *ProductDB) GetProductByID(id string) (product *models.Product, err error) {
	product = &models.Product{}
	res := p.DB.Client.(*gorm.DB).First(product, &id)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

// GetAllProducts fetches all records from a table. skip and limit params are used to fetch specific number of records. selector is a additional param to fetch based on
func  (p *ProductDB) GetAllProducts(skip int64, limit int64, selector interface{}) ([]models.Product, error) {
products := []models.Product{}
	res := p.DB.Client.(*gorm.DB).Limit(int(skip)).Offset(int(limit)).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products , nil
}

// GetAllProductsBy fetches all records from a table. skip and limit params are used to fetch specific number of records.key and value are addition params to fetch based on a key with a value
func (p *ProductDB) GetAllProductsBy(key string, value interface{}, skip int64, limit int64) ([]models.Product, error) {
products := []models.Product{}
	res:=p.DB.Client.(*gorm.DB).Where(key+"=?", value).Limit(int(skip)).Offset(int(limit)).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products , nil
}
