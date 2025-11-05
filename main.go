package main

import (
	"microservices/UserService"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/users", UserService.GetUsers)
	route.Run(":8000")
}
