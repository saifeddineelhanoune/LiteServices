package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saifeddineelhanoune/LiteServices/UserService"
)

func main() {
	route := gin.Default()
	route.GET("/users", getUsers)
	route.Run(":8000")
}