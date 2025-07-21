package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"size:100;not null" json:"name"`
	Description string `gorm:"size:100;not null;" json:"description"`
	Price float64 `gorm:"not null;" json:"price"`
	CategoryId int `gorm:"not null;" json:"category"`
	Category Category
	UserId int `gorm:"not null;" json:"user_id"`
	User User `gorm:"foreignKey:UserId"`
}

type Category struct {
	gorm.Model
	Name string `gorm:"size:100;not null;unique" json:"name"`
	Description string `gorm:"size:100;not null;" json:"description"`
	Products []Product
}