package models


type Order struct {
	BaseModel
	User_id    int     `gorm:"not null;" json:"user" form:"user_id"`
	User       User    `gorm:"foreignKey:User_id"`
	Status     string  `gorm:"not null;" json:"status" form:"status"`
	TotalPrice float64 `gorm:"not null;" json:"total_price" form:"total_price"`
	Payment Payment `gorm:"foreignKey:OrderId"`
}