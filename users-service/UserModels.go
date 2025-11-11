package main

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

func UpdateUserById(ctx *gin.Context) {
	var UpdatedUser User
	if err := ctx.BindJSON(&UpdatedUser); err != nil {
		return
	}
	for i, a := range Users {
		if a.ID == ctx.Param("id") {
			Users[i].ID = UpdatedUser.ID
			Users[i].Name = UpdatedUser.Name 
			Users[i].Email = UpdatedUser.Email 
			ctx.JSON(http.StatusAccepted, Users[i])
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "can't find the user id you are searching for"})
}

func DeleteUserById(ctx *gin.Context) {
	var user User
	if err := ctx.BindJSON(&user); err != nil {
		return
	}
	for i, u := range Users {
		if u.ID == ctx.Param("id") {
			Users = append(Users[:i], Users[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "user Deleted successfuly"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
