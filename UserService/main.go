package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/users", GetUsers)
	route.POST("/users", PostUsers)
	route.GET("/users/:id", GetUserById)
	route.PUT("/users/:id", UpdateUserById)
	route.DELETE("/users/:id", DeleteUserById)
	route.Run("127.0.0.1:5000")
}
