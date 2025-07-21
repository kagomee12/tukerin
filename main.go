package main

import (
	"fmt"
	"tukerin/config"
	"tukerin/models"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main(){

	gotenv.Load()

	config.ConnectDB()

	// config.DB.Migrator().DropTable(
	// 	&models.User{},
	// 	&models.Role{},
	// 	&models.Order{},
	// 	&models.OrderItem{},
	// 	&models.Product{},
	// 	&models.Payment{},
	// 	&models.Cart{},
	// 	&models.CartItem{},
	// 	&models.Category{},
	// 	&models.Otp{},
	// )

	config.DB.AutoMigrate(
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

	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		fmt.Println("hello world")
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	route.Run(":8080")
}