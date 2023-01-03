package repository

import (
	"fmt"
	"gaunRumaRestApi/config/db"
	"gaunRumaRestApi/entity"
)

type ProductRepository interface {
	GetAllProduct() ([]*entity.Product, error)
	GetProductById(id uint) (*entity.Product, error)
	CreateProduct(name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*entity.Product, error)
	UpdateProduct(id uint, name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*entity.Product, error)
	DeleteProduct(id uint) (bool, error)
}

type ProductRepositoryImpl struct {
	dbHandler *db.Handler
}

// GetProductById implements ProductRepository
func (s *ProductRepositoryImpl) GetProductById(id uint) (*entity.Product, error) {
	var product entity.Product
	result := s.dbHandler.DB.Preload("ProductType").First(&product, id)

	if result.Error != nil {
		return &product, result.Error
	}
	return &product, nil
}

// CreateProduct implements ProductRepository
func (s *ProductRepositoryImpl) CreateProduct(name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*entity.Product, error) {
	product := entity.Product{ProductName: name, ProductCode: code, ProductReject: isRejected, ProductColor: color, ProductSize: size, ProductStock: stock, ProductPrice: price, TypeID: stockType}
	result := s.dbHandler.DB.Create(&product)

	if result.Error != nil {
		return &entity.Product{}, result.Error
	}

	return &product, nil
}

// DeleteProduct implements ProductRepository
func (s *ProductRepositoryImpl) DeleteProduct(id uint) (bool, error) {
	var product entity.Product
	resp := s.dbHandler.DB.Delete(&product, id)
	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

// GetAllProduct implements ProductRepository
func (s *ProductRepositoryImpl) GetAllProduct() ([]*entity.Product, error) {
	var arrProduct []*entity.Product
	resp := s.dbHandler.DB.Preload("ProductType").Find(&arrProduct)
	fmt.Println(resp.RowsAffected)

	if resp.Error != nil {
		return arrProduct, resp.Error
	}
	return arrProduct, nil
}

// UpdateProduct implements ProductRepository
func (s *ProductRepositoryImpl) UpdateProduct(id uint, name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*entity.Product, error) {
	product := entity.Product{ID: id, ProductName: name, ProductCode: code, ProductReject: isRejected, ProductColor: color, ProductSize: size, ProductStock: stock, ProductPrice: price, TypeID: stockType}
	result := s.dbHandler.DB.Save(&product)

	if result.Error != nil {
		return &entity.Product{}, result.Error
	}

	return &product, nil
}

func NewProductRepository(dbHandler *db.Handler) ProductRepository {
	return &ProductRepositoryImpl{
		dbHandler: dbHandler,
	}
}
