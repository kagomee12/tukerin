package models

import "time"

type Otp struct {
	Email string `json:"email"`
	Code string `json:"otp"`
	ExpiresAt time.Time `json:"expires_at"`
	UserId int `gorm:"not null;" json:"user_id"`
	User User `gorm:"foreignKey:UserId"`
	Used bool `gorm:"default:false" json:"used"`
	Type string `gorm:"size:20;not null;" json:"type"` // e.g., "email", "sms"
}