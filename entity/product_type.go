package entity

import "time"

type ProductType struct {
	ID        uint   `gorm:"primaryKey;autoIncrement:true"`
	TypeName  string `gorm:"not null"`
	TypeCode  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ProductType) TableName() string {
	return "t_productType"
}
