package handler

import (
	"gaunRumaRestApi/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}
func (s *ProductHandler) GetAllProduct(c echo.Context) error {
	result, err := s.service.GetAllProduct()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductHandler) CreateProduct(c echo.Context) error {
	req := new(CreateProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.service.CreateProduct(req.ProductName, req.ProductCode, req.ProductPrice, req.ProductQuantity, req.ProductColor, req.ProductIsRejected, req.ProductSize, req.ProductType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductHandler) UpdateProduct(c echo.Context) error {
	req := new(UpdateProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.service.UpdateProduct(req.ProductId, req.ProductName, req.ProductCode, req.ProductPrice, req.ProductQuantity, req.ProductColor, req.ProductIsRejected, req.ProductSize, req.ProductType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductHandler) DeleteProduct(c echo.Context) error {
	req := new(DeleteProductRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.service.DeleteProduct(uint(req.ProductId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *ProductHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	result, err := s.service.GetProductById(uint(uintId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
