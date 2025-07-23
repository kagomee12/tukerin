package routes

import (
	"tukerin/controller"
	"tukerin/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoute(r *gin.Engine) {
	product := r.Group("/product")
	product.Use(middleware.AuthMiddleware())
	{
		product.GET("/", controller.GetProducts)
		product.GET("/:id", controller.GetProductByID)
		product.POST("/", controller.CreateProduct)
		product.PATCH("/:id", controller.UpdateProduct)
		product.DELETE("/:id", controller.DeleteProduct)
	}
}