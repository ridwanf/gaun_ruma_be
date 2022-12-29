package handler

import "github.com/golang-jwt/jwt/v4"

type CreateStockRequest struct {
	StockName       string  `json:"stock_name"`
	StockCode       string  `json:"stock_code"`
	StockPrice      float64 `json:"stock_price"`
	StockQuantity   uint64  `json:"stock_quantity"`
	StockColor      string  `json:"stock_color"`
	StockIsRejected bool    `json:"stock_is_rejected"`
	StockSize       string  `json:"stock_size"`
	StockType       uint    `json:"stock_type"`
}

type UpdateStockRequest struct {
	StockId         uint    `json:"stock_id"`
	StockName       string  `json:"stock_name"`
	StockCode       string  `json:"stock_code"`
	StockPrice      float64 `json:"stock_price"`
	StockQuantity   uint64  `json:"stock_quantity"`
	StockColor      string  `json:"stock_color"`
	StockIsRejected bool    `json:"stock_is_rejected"`
	StockSize       string  `json:"stock_size"`
	StockType       uint    `json:"stock_type"`
}

type DeleteStockRequest struct {
	StockId int64 `json:"stock_id"`
}

type UserLoginRequest struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}

type JwtCustomClaims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

type CreateProductTypeRequest struct {
	TypeName string `json:"type_name"`
	TypeCode string `json:"type_code"`
}

type UpdateProductTypeRequest struct {
	TypeName string `json:"type_name"`
	TypeCode string `json:"type_code"`
	TypeID   uint   `json:"type_id"`
}
type DeleteProductTypeRequest struct {
	TypeId int64 `json:"type_id"`
}
