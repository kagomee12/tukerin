package models

type Product struct {
	BaseModel
	Name string `gorm:"size:100;not null" json:"name" form:"name"`
	Description string `gorm:"size:100;not null;" json:"description" form:"description"`
	Price float64 `gorm:"not null;" json:"price" form:"price"`
	CategoryId int `gorm:"not null;" json:"category" form:"category_id"`
	Category Category
	UserId int `gorm:"not null;" json:"user_id" form:"user_id"`
	User User `gorm:"foreignKey:UserId"`
}

type Category struct {
	BaseModel
	Name string `gorm:"size:100;not null;unique" json:"name" form:"name"`
	Description string `gorm:"size:100;not null;" json:"description" form:"description"`
	Products []Product
}