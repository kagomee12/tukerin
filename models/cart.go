package models


type Cart struct {
	BaseModel
	UserId    int     `gorm:"not null;" json:"user_id" form:"user_id"`
	TotalPrice float64 `gorm:"not null;" json:"total_price" form:"total_price"`
	CartItems []CartItem `gorm:"foreignKey:CartId" json:"cart_items"`
}