package main

import (
	"microservices/UserService"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/users", UserService.GetUsers)
	route.POST("/users", UserService.PostUsers)
	route.Run(":5000")
}
