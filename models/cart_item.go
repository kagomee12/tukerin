package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartId    int     `gorm:"not null;" json:"cart_id"`
	Cart      Cart    `gorm:"foreignKey:CartId"`
	ProductId int     `gorm:"not null;" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
	Quantity  int     `gorm:"not null;" json:"quantity"`
	UnitPrice float64 `gorm:"not null;" json:"unit_price"`
}