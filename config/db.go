package config

import (
	"fmt"
	"os"
	"tukerin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"), 
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Order{},
		&models.OrderItem{},
		&models.Product{},
		&models.Payment{},
		&models.Cart{},
		&models.CartItem{},
		&models.Category{},
		&models.Otp{},
	)
	if err != nil {
		panic("Failed to migrate database")
	}
}
