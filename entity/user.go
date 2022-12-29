package entity

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement:true"`
	UserName     string `gorm:"not null"`
	UserPassword string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (User) TableName() string {
	return "t_user"
}
