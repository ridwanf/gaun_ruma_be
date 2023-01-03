package models

import (
	"gaunRumaRestApi/entity"
)

func ProductModelToEntity(model *Product) *entity.Product {
	product := &entity.Product{
		ID:            uint(model.ProductId),
		ProductName:   model.ProductName,
		ProductCode:   model.ProductCode,
		TypeID:        model.TypeID,
		ProductType:   *ProductTypeModelToEntity(&model.ProductType),
		ProductReject: model.IsRejected,
		ProductColor:  model.ProductColor,
		ProductSize:   model.ProductSize,
		ProductStock:  model.ProductStock,
		ProductPrice:  model.ProductPrice,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
	return product
}

func ProductEntityToModel(entity *entity.Product) *Product {
	product := &Product{
		ProductId:    int32(entity.ID),
		ProductName:  entity.ProductName,
		ProductCode:  entity.ProductCode,
		TypeID:       entity.TypeID,
		ProductType:  *ProductTypeEntityToModel(&entity.ProductType),
		IsRejected:   entity.ProductReject,
		ProductColor: entity.ProductColor,
		ProductSize:  entity.ProductSize,
		ProductStock: entity.ProductStock,
		ProductPrice: entity.ProductPrice,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
	return product
}

func ProductEntityToListProductModel(entityList []*entity.Product) []*Product {
	var modelList []*Product

	for _, data := range entityList {
		modelList = append(modelList, ProductEntityToModel(data))
	}
	return modelList
}

func ProductModeloListProductEntity(modelList []*Product) []*entity.Product {
	var entityList []*entity.Product

	for _, data := range modelList {
		entityList = append(entityList, ProductModelToEntity(data))
	}
	return entityList
}
