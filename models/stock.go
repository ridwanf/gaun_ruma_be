package models

type Stock struct {
	StockId       int32   `json:"stock_id"`
	StockName     string  `json:"stock_name"`
	StockCode     string  `json:"stock_code"`
	StockPrice    float64 `json:"stock_price"`
	StockQuantity int64   `json:"stock_quantity"`
}
