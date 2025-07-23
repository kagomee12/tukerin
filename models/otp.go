package models

import (
	"time"
)

type Otp struct {
	BaseModel
	Email string `json:"email" form:"email"`
	Code string `json:"otp" form:"otp"`
	ExpiresAt time.Time `json:"expires_at" form:"expires_at"`
	UserId int `gorm:"not null;" json:"user_id" form:"user_id"`
	User User `gorm:"foreignKey:UserId"`
	Used bool `gorm:"default:false" json:"used" form:"used"`
	Type string `gorm:"size:20;not null;" json:"type" form:"type"` // e.g., "email", "sms"
}