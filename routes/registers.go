package routes

import (
	"net/http"
	"strconv"

	"example.com/events_api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse id."})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot fetch event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot register"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered!!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse id."})
		return
	}
	var event models.Event
	event.Id = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot cancel registration."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Cancelled!!"})
}
