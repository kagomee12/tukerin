package routes

import (
	"github.com/gin-gonic/gin"
)

func IndexRoute() {
	route := gin.Default()

	AuthRoute(route)

	route.Run(":8081") // Start the server on port 8080
}