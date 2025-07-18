package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		fmt.Println("hello world")
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	route.Run(":8080")
}