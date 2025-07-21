package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    int     `gorm:"not null;" json:"user_id"`
	TotalPrice float64 `gorm:"not null;" json:"total_price"`
	CartItems []CartItem `gorm:"foreignKey:CartId" json:"cart_items"`
}