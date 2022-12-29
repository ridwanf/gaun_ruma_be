package repository

import (
	"gaunRumaRestApi/config/db"
	"gaunRumaRestApi/entity"
)

type ProductTypeRepository interface {
	GetAllProductType() (*[]entity.ProductType, error)
	CreateProductType(name string, code string) (*entity.ProductType, error)
	UpdateProductType(id uint, name string, code string) (*entity.ProductType, error)
	DeleteProductType(id uint) (bool, error)
}

type ProductTypeRepositoryImpl struct {
	dbHandler *db.Handler
}

// CreateProductType implements ProductTypeRepository
func (p *ProductTypeRepositoryImpl) CreateProductType(name string, code string) (*entity.ProductType, error) {
	productType := entity.ProductType{TypeName: name, TypeCode: code}
	result := p.dbHandler.DB.Create(&productType)

	if result.Error != nil {
		return &entity.ProductType{}, result.Error
	}

	return &productType, nil
}

// DeleteProductType implements ProductTypeRepository
func (p *ProductTypeRepositoryImpl) DeleteProductType(id uint) (bool, error) {
	var productType entity.ProductType
	resp := p.dbHandler.DB.Delete(&productType, id)
	if resp.Error != nil {
		return false, resp.Error
	}

	return true, nil
}

// GetAllProductType implements ProductTypeRepository
func (p *ProductTypeRepositoryImpl) GetAllProductType() (*[]entity.ProductType, error) {
	var arrProductType []entity.ProductType
	var productType entity.ProductType
	resp := p.dbHandler.DB.Find(&productType)

	if resp.Error != nil {
		return &arrProductType, resp.Error
	}
	rows, err := resp.Rows()
	if err != nil {
		return &arrProductType, err
	}
	for rows.Next() {
		err = rows.Scan(&productType.ID, &productType.CreatedAt, &productType.UpdatedAt, &productType.TypeName, &productType.TypeCode)
		if err != nil {
			return &arrProductType, err
		}
		arrProductType = append(arrProductType, productType)
	}
	return &arrProductType, nil
}

// UpdateProductType implements ProductTypeRepository
func (p *ProductTypeRepositoryImpl) UpdateProductType(id uint, name string, code string) (*entity.ProductType, error) {
	productType := entity.ProductType{ID: id, TypeName: name, TypeCode: code}
	result := p.dbHandler.DB.Save(&productType)

	if result.Error != nil {
		return &entity.ProductType{}, result.Error
	}

	return &productType, nil
}

func NewProductTypeRepository(dbHandler *db.Handler) ProductTypeRepository {
	return &ProductTypeRepositoryImpl{
		dbHandler: dbHandler,
	}
}
