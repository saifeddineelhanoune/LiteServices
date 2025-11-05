package usermodels

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

var Users = []User{
	{ID:"1", Name:"Saif", Email:"saif@example.com"},
	{ID:"2", Name:"khalil", Email:"khalil@example.com"},
	{ID:"3", Name:"ider", Email:"ider@example.com"},
}

func getUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Users)
}