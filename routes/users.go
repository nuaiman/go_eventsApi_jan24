package routes

import (
	"net/http"

	"example.com/events_api/models"
	"example.com/events_api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse request data."})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user."})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "User created."})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse request data."})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials."})
		return
	}
	token, err := utils.GenerateJWT(user.Id, user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Invalid credentials."})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Logged in.", "token": token})
}
