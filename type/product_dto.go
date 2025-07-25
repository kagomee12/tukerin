package types

type ProductsDTO struct {
	Name        string      `json:"name" form:"name"`
	Description string      `json:"description" form:"description"`
	Price       float64     `json:"price" form:"price"`
	CategoryId  int         `json:"category_id" form:"category_id"`
	Category    CategoryDTO `json:"category" form:"category"`
	UserId      int         `json:"user_id" form:"user_id"`
	User        UserDTO     `json:"user" form:"user"`
	Images      []ImageDTO    `json:"image" form:"image"`
}
