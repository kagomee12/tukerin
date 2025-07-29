package types

type UserDTO struct {
	ID     int     `json:"id" form:"id"`
	Name   string  `json:"name" form:"name"`
	Email  string  `json:"email" form:"email"`
	RoleId string  `json:"role_id" form:"role_id"`
	Role   RoleDTO `json:"role" form:"role"`
}
