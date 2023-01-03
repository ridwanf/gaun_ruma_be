package models

import (
	"gaunRumaRestApi/entity"
)

func ProductTypeModelToEntity(model *ProductType) *entity.ProductType {
	productType := &entity.ProductType{
		ID:        model.TypeID,
		TypeName:  model.TypeName,
		TypeCode:  model.TypeCode,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
	return productType
}

func ProductTypeEntityToModel(entity *entity.ProductType) *ProductType {
	productType := &ProductType{
		TypeID:    entity.ID,
		TypeName:  entity.TypeName,
		TypeCode:  entity.TypeCode,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return productType
}

func ProductTypeEntityToListProductTypeModel(entityList []*entity.ProductType) []*ProductType {
	var modelList []*ProductType

	for _, data := range entityList {
		modelList = append(modelList, ProductTypeEntityToModel(data))
	}
	return modelList
}

func ProductTypeModeloListProductTypeEntity(modelList []*ProductType) []*entity.ProductType {
	var entityList []*entity.ProductType

	for _, data := range modelList {
		entityList = append(entityList, ProductTypeModelToEntity(data))
	}
	return entityList
}
