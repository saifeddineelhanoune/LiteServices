package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Product service running"})
	})
	route.GET("/products", GetProducts)
	route.POST("/products", PostProducts)
	route.GET("/products/:id", GetProductById)
	route.PUT("/products/:id", UpdateProductById)
	route.DELETE("/products/:id", DeleteProductById)
	route.Run("0.0.0.0:5001")
}
