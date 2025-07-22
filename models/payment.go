package models

type Payment struct {
	BaseModel
	UserId int `gorm:"not null;" json:"user" form:"user_id"`
	User User
	ProductId int `gorm:"not null;" json:"product" form:"product_id"`
	Product Product
	Amount float64 `gorm:"not null;" json:"amount" form:"amount"`
	Status string `gorm:"not null;" json:"status" form:"status"`
	OrderId int `gorm:"not null;" json:"order" form:"order_id"`
}