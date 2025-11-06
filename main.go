package main

import (
	"microservices/UserService"
	"microservices/ProductService"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/users", UserService.GetUsers)
	route.POST("/users", UserService.PostUsers)
	route.GET("/users/:id", UserService.GetUserById)
	route.PUT("/users/:id", UserService.UpdateUserById)
	route.DELETE("/users/:id", UserService.DeleteUserById)
	route.GET("/product", ProductService.GetProducts)
	route.Run(":5000")
}
