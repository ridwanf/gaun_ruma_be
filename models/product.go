package models

import "time"

type Product struct {
	ProductId       int32       `json:"product_id"`
	ProductName     string      `json:"product_name"`
	ProductCode     string      `json:"product_code"`
	ProductPrice    float64     `json:"product_price"`
	TypeID          uint        `json:"product_type_id"`
	ProductQuantity int64       `json:"product_quantity"`
	ProductType     ProductType `json:"product_type"`
	IsRejected      bool        `json:"product_is_rejected"`
	ProductColor    string      `json:"product_color"`
	ProductSize     string      `json:"product_size"`
	ProductStock    uint64      `json:"product_stock"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
}
