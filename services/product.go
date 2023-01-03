package services

import (
	"gaunRumaRestApi/models"
	"gaunRumaRestApi/repository"
)

type ProductService interface {
	GetAllProduct() ([]*models.Product, error)
	GetProductById(id uint) (*models.Product, error)
	CreateProduct(name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*models.Product, error)
	UpdateProduct(id uint, name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*models.Product, error)
	DeleteProduct(id uint) (bool, error)
}

type ProductServiceImpl struct {
	repository repository.ProductRepository
}

// CreateProduct implements ProductService
func (p *ProductServiceImpl) CreateProduct(name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*models.Product, error) {
	result, err := p.repository.CreateProduct(name, code, price, stock, color, isRejected, size, stockType)
	if err != nil {
		return nil, err
	}
	return models.ProductEntityToModel(result), nil
}

// DeleteProduct implements ProductService
func (p *ProductServiceImpl) DeleteProduct(id uint) (bool, error) {
	result, err := p.repository.DeleteProduct(id)
	if err != nil {
		return false, err
	}

	return result, nil
}

// GetAllProduct implements ProductService
func (p *ProductServiceImpl) GetAllProduct() ([]*models.Product, error) {
	result, err := p.repository.GetAllProduct()
	if err != nil {
		return nil, err
	}

	return models.ProductEntityToListProductModel(result), nil
}

// GetProductById implements ProductService
func (p *ProductServiceImpl) GetProductById(id uint) (*models.Product, error) {
	result, err := p.repository.GetProductById(id)
	if err != nil {
		return nil, err
	}

	return models.ProductEntityToModel(result), nil
}

// UpdateProduct implements ProductService
func (p *ProductServiceImpl) UpdateProduct(id uint, name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*models.Product, error) {
	result, err := p.repository.UpdateProduct(id, name, code, price, stock, color, isRejected, size, stockType)
	if err != nil {
		return nil, err
	}
	return models.ProductEntityToModel(result), nil
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		repository: repository,
	}
}
