package controller

import (
	"fmt"
	"math"
	"strconv"
	"tukerin/config"
	"tukerin/models"
	"tukerin/type"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	var data []types.ProductsDTO
	
	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64

	name_product := c.Query("name_product")
	category := c.Query("category")
	category_id, _ := strconv.Atoi(category)

	query := config.DB.Model(&models.Product{})

	if name_product != "" {
		query = query.Where("name ILIKE ?", "%"+name_product+"%")
	}
	if category != "" {
		query = query.Where("category_id =?", category_id)
	}

	if err := query.Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}

	if err := query.Limit(limit).Offset(offset).Preload("Category").Preload("User").Preload("User.Role").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	for _, product := range products {
		data = append(data, types.ProductsDTO{
			Name: 	  product.Name,
			Description: product.Description,
			Price:	  product.Price,
			CategoryId: product.CategoryId,
			UserId:	  product.UserId,
			Category: types.CategoryDTO{
				ID:   product.Category.ID,
				Name: product.Category.Name,
			},
			User: types.UserDTO{
				ID:    product.User.ID,
				Name:  product.User.Name,
				Email: product.User.Email,
				RoleId: strconv.Itoa(product.User.RoleId),
				Role: types.RoleDTO{
					ID:   product.User.Role.ID,
					Name: product.User.Role.Name,
				},
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"limit": limit,
		"page": page,
		"total": total,
		"totalPages": int(math.Ceil(float64(total) / float64(limit))),
	})
}

func GetProductByID(c *gin.Context) {
	var product models.Product
	var data types.ProductsDTO
	productID := c.Param("id")

	_, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	if err := config.DB.Where("id = ?", productID).Preload("Category").Preload("User").Preload("User.Role").First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	data = types.ProductsDTO{
		Name: 	  product.Name,
		Description: product.Description,
		Price:	  product.Price,
		CategoryId: product.CategoryId,
		UserId:	  product.UserId,
		Category: types.CategoryDTO{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
		User: types.UserDTO{
			ID:    product.User.ID,
			Name:  product.User.Name,
			Email: product.User.Email,
			RoleId: strconv.Itoa(product.User.RoleId),
			Role: types.RoleDTO{
				ID:   product.User.Role.ID,
				Name: product.User.Role.Name,
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": data} )
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	product.UserId = userID.(int)

	if err := config.DB.Where("id = ?", product.CategoryId).First(&models.Category{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := config.DB.Select("name", "description", "price", "category_id", "user_id").Create(&product).Error; err != nil {
		fmt.Println("Error creating product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"success": "Product created successfully"})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	productID := c.Param("id")

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	if err := config.DB.Model(&product).Where("id = ? AND user_id = ?", productID, userID).Updates(product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	if err := config.DB.Where("id = ? AND user_id = ?", productID, userID).Delete(&models.Product{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}