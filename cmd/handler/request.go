package handler

import "github.com/golang-jwt/jwt/v4"

type CreateProductRequest struct {
	ProductName       string  `json:"product_name"`
	ProductCode       string  `json:"product_code"`
	ProductPrice      float64 `json:"product_price"`
	ProductQuantity   uint64  `json:"product_quantity"`
	ProductColor      string  `json:"product_color"`
	ProductIsRejected bool    `json:"product_is_rejected"`
	ProductSize       string  `json:"product_size"`
	ProductType       uint    `json:"product_type"`
}

type UpdateProductRequest struct {
	ProductId         uint    `json:"product_id"`
	ProductName       string  `json:"product_name"`
	ProductCode       string  `json:"product_code"`
	ProductPrice      float64 `json:"product_price"`
	ProductQuantity   uint64  `json:"product_quantity"`
	ProductColor      string  `json:"product_color"`
	ProductIsRejected bool    `json:"product_is_rejected"`
	ProductSize       string  `json:"product_size"`
	ProductType       uint    `json:"product_type"`
}

type DeleteProductRequest struct {
	ProductId int64 `json:"product_id"`
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
