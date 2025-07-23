package main

import (
	"tukerin/config"
	"tukerin/config/seeder"
	"tukerin/routes"

	// "tukerin/models"
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

	config.Migrate()
	seeder.SeederRoles()

	routes.IndexRoute()

}