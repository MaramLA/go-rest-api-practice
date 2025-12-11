package routes

import (
	"net/http"
	"strconv"

	"example.com/go-api-practice/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event", "error": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered for event successfully"})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid event id", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId
	event.UserId = userId
	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized to cancel registration for this event"})
	}
	err = event.CancelRegister(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration for event", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration cancelled successfully"})
}
