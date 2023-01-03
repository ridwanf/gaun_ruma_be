package models

import "time"

type ProductType struct {
	TypeID    uint      `json:"type_id"`
	TypeName  string    `json:"type_name"`
	TypeCode  string    `json:"type_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
