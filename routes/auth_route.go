package routes

import (
	"tukerin/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
		auth.POST("/validate-token", controller.VerifyOTP)
	}
}
