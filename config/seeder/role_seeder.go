package seeder

import (
	"tukerin/config"
	"tukerin/models"

	"gorm.io/gorm"
)

func SeederRoles() {
	roles := []models.Role{
		{Name: "customer", Description: "Customer with access to basic features"},
		{Name: "seller", Description: "Seller with access to product management"},
	}

	for _, role := range roles {
		var existingRole models.Role
		if err := config.DB.Where("name = ?", role.Name).First(&existingRole).Error; err != nil {
			if err == gorm.ErrRecordNotFound{
				config.DB.Create(&role)
			}
		}
			
	}
}
