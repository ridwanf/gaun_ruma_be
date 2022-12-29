package handler

import (
	"gaunRumaRestApi/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductTypeHandler struct {
	repository repository.ProductTypeRepository
}

func NewProductTypeHandler(repository repository.ProductTypeRepository) *ProductTypeHandler {
	return &ProductTypeHandler{
		repository: repository,
	}
}

func (s *ProductTypeHandler) GetAllProductType(c echo.Context) error {
	result, err := s.repository.GetAllProductType()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductTypeHandler) CreateProductType(c echo.Context) error {
	req := new(CreateProductTypeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.repository.CreateProductType(req.TypeName, req.TypeCode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductTypeHandler) UpdateProductType(c echo.Context) error {
	req := new(UpdateProductTypeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.repository.UpdateProductType(req.TypeID, req.TypeName, req.TypeCode)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductTypeHandler) DeleteProductType(c echo.Context) error {
	req := new(DeleteProductTypeRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.repository.DeleteProductType(uint(req.TypeId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
