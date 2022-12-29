package handler

import (
	"gaunRumaRestApi/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StockHandler struct {
	repository repository.ProductRepository
}

func NewStockHandler(repository repository.ProductRepository) *StockHandler {
	return &StockHandler{
		repository: repository,
	}
}
func (s *StockHandler) GetAllStock(c echo.Context) error {
	result, err := s.repository.GetAllProduct()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *StockHandler) CreateStock(c echo.Context) error {
	req := new(CreateStockRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.repository.CreateProduct(req.StockName, req.StockCode, req.StockPrice, req.StockQuantity, req.StockColor, req.StockIsRejected, req.StockSize, req.StockType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *StockHandler) UpdateStock(c echo.Context) error {
	req := new(UpdateStockRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.repository.UpdateProduct(req.StockId, req.StockName, req.StockCode, req.StockPrice, req.StockQuantity, req.StockColor, req.StockIsRejected, req.StockSize, req.StockType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (s *StockHandler) DeleteStock(c echo.Context) error {
	req := new(DeleteStockRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	result, err := s.repository.DeleteProduct(uint(req.StockId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}
