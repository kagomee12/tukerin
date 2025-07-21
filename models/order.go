package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	User_id    int     `gorm:"not null;" json:"user"`
	User       User    `gorm:"foreignKey:User_id"`
	Status     string  `gorm:"not null;" json:"status"`
	TotalPrice float64 `gorm:"not null;" json:"total_price"`
}