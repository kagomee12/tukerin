package models

type User struct {
	BaseModel
	Name      string `gorm:"size:255;not null;unique" json:"name"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	RoleId    int    `gorm:"not null;" json:"role"`
	Role      Role
	Password  string     `gorm:"size:100;not null;" json:"password"`
	Product   []Product  `gorm:"foreignKey:UserId"`
	Cart      Cart       `gorm:"foreignKey:UserId"`
	Otp       []Otp      `gorm:"foreignKey:UserId"`
}

type Role struct {
	BaseModel
	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Description string `gorm:"size:100;not null;" json:"description"`
	Users       []User
}
