package seeder

import (
	"tukerin/config"
	"tukerin/models"
)

func SeederRoles() {
	roles := []models.Role{
		{Name: "customer", Description: "Customer with access to basic features"},
		{Name: "seller", Description: "Seller with access to product management"},
	}

	config.DB.Create(&roles)
}
