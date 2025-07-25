package models

import "time"

type BaseModel struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}