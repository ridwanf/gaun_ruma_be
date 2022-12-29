package repository

import (
	"gaunRumaRestApi/config/db"
	"gaunRumaRestApi/entity"
)

type ProductRepository interface {
	GetAllProduct() (*[]entity.Product, error)
	CreateProduct(name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*entity.Product, error)
	UpdateProduct(id uint, name string, code string, price float64, stock uint64, color string, isRejected bool, size string, stockType uint) (*entity.Product, error)
	DeleteProduct(id uint) (bool, error)
}

type ProductRepositoryImpl struct {
	dbHandler *db.Handler
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
func (s *ProductRepositoryImpl) GetAllProduct() (*[]entity.Product, error) {
	var arrProduct []entity.Product
	var product entity.Product
	resp := s.dbHandler.DB.Preload("ProductType").Find(&product)

	if resp.Error != nil {
		return &arrProduct, resp.Error
	}
	rows, err := resp.Rows()
	if err != nil {
		return &arrProduct, err
	}
	for rows.Next() {
		err = rows.Scan(&product.ID,
			&product.ProductName,
			&product.ProductCode,
			&product.TypeID,
			&product.ProductReject,
			&product.ProductColor,
			&product.ProductSize,
			&product.ProductStock,
			&product.ProductPrice,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return &arrProduct, err
		}
		arrProduct = append(arrProduct, product)
	}
	return &arrProduct, nil
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
