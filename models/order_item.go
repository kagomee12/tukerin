package models

type OrderItem struct {
	BaseModel
	OrderId    int     `gorm:"not null;" json:"order_id" form:"order_id"`
	Order      Order   `gorm:"foreignKey:OrderId"`
	ProductId  int     `gorm:"not null;" json:"product_id" form:"product_id"`
	Product    Product `gorm:"foreignKey:ProductId"`
	Quantity   int     `gorm:"not null;" json:"quantity" form:"quantity"`
	UnitPrice  float64 `gorm:"not null;" json:"unit_price" form:"unit_price"`
	UserId     int     `gorm:"not null;" json:"user_id" form:"user_id"`
	User       User    `gorm:"foreignKey:UserId"`
}