package seeder

import (
	"tukerin/config"
	"tukerin/models"

	"gorm.io/gorm"
)

func SeederCategories() {
	categories := []models.Category{
		{Name: "Electronics", Description: "Devices and gadgets like phones, laptops, etc."},
		{Name: "Fashion", Description: "Clothing, shoes, and accessories"},
		{Name: "Home & Living", Description: "Furniture, home decor, and appliances"},
		{Name: "Beauty & Personal Care", Description: "Skincare, makeup, and personal hygiene"},
		{Name: "Sports & Outdoors", Description: "Fitness, camping, and sports gear"},
		{Name: "Toys & Hobbies", Description: "Toys, games, and hobby supplies"},
		{Name: "Automotive", Description: "Car parts, accessories, and tools"},
		{Name: "Books & Stationery", Description: "Books, journals, and office supplies"},
		{Name: "Health & Wellness", Description: "Supplements, medical supplies"},
		{Name: "Groceries", Description: "Food, beverages, and daily essentials"},
	}
	

	for _, category := range categories {
		var existingCategory models.Category
		if err := config.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				config.DB.Create(&category)
			}
		}
	}
}