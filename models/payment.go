package models

type Payment struct {
	BaseModel
	UserId int `gorm:"not null;" json:"user"`
	User User
	ProductId int `gorm:"not null;" json:"product"`
	Product Product
	Amount float64 `gorm:"not null;" json:"amount"`
	Status string `gorm:"not null;" json:"status"`
	OrderId int `gorm:"not null;" json:"order"`
}