package entity

import "time"

type Product struct {
	ID            uint   `gorm:"primaryKey;autoIncrement:true"`
	ProductName   string `gorm:"not null"`
	ProductCode   string `gorm:"not null"`
	TypeID        uint
	ProductType   ProductType `gorm:"foreignKey:TypeID;references:ID"`
	ProductReject bool
	ProductColor  string
	ProductSize   string
	ProductStock  uint64
	ProductPrice  float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (Product) TableName() string {
	return "t_product"
}
