package main

import (
	"net/http"

	"example.com/pj-rest-api/db"
	"example.com/pj-rest-api/model"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/event", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})

	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
