package routes

import (
	"net/http"
	"strconv"
	"time"

	"example.com/go-api-practice/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "events could not be fetched", "error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "events retrieved successfully", "data": events})
}

func getEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event fetched successfully", "data": event})
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")
	event.UserId = userId
	event.DataTime = time.Now()

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event to database", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "data": event})
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id", "error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch the event", "error": err.Error()})
		return
	}

	if(event.UserId != userId){
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized to update this event"})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.UpdateEventById()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid event id", "error": err.Error()})
		return
	}
	foundEvent, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event", "error": err.Error()})
		return
	}

	if(foundEvent.UserId != userId){
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized to delete this event"})
		return
	}

	err = foundEvent.DeleteEventById()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}



