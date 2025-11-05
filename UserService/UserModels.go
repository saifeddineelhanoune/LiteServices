package UserService

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var Users = []User{
	{ID: "1", Name: "Saif", Email: "saif@example.com"},
	{ID: "2", Name: "khalil", Email: "khalil@example.com"},
	{ID: "3", Name: "ider", Email: "ider@example.com"},
}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Users)
}

func PostUsers(ctx	*gin.Context) {
	var newUser User

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	Users = append(Users, newUser)
	ctx.JSON(http.StatusCreated, Users)
}

func GetUserById(ctx *gin.Context) {
	for _, a := range Users {
		if a.ID == ctx.Param("id") {
			ctx.JSON(http.StatusOK, a)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "the user you are looking for is not found"})
}

