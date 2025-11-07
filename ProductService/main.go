package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/products", GetProducts)
	route.POST("/products", PostProducts)
	route.GET("/products/:id", GetProductById)
	route.PUT("/products/:id", UpdateProductById)
	route.DELETE("/products/:id", DeleteProductById)
	route.Run("127.0.0.1:5001")
}
