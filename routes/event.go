package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events_api/models"
	"github.com/gin-gonic/gin"
)

func addEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse request data."})
		return
	}
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot save event."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event added."})
}

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot fetch events."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Events fetched.", "data": events})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse request id."})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot fetch event with id."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event fetched.", "data": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse request id."})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		fmt.Println(event)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot fetch event with id."})
		return
	}
	err = context.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot parse request data."})
		return
	}
	userId := context.GetInt64("userId")
	if userId != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action !!"})
		return
	}
	err = event.Update()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot update event."})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event updated."})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "Couldnot parse request id."})
		return
	}
	event, err := models.GetEvent(eventId)
	if err != nil {
		fmt.Println(event)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot fetch event with id."})
		return
	}
	userId := context.GetInt64("userId")
	if userId != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action !!"})
		return
	}
	err = event.Delete()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnot delete event."})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"message": "Event deleted."})
}
