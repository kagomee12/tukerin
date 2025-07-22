package models


type CartItem struct {
	BaseModel
	CartId    int     `gorm:"not null;" json:"cart_id" form:"cart_id"`
	Cart      Cart    `gorm:"foreignKey:CartId"`
	ProductId int     `gorm:"not null;" json:"product_id" form:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId"`
	Quantity  int     `gorm:"not null;" json:"quantity" form:"quantity"`
	UnitPrice float64 `gorm:"not null;" json:"unit_price" form:"unit_price"`
}