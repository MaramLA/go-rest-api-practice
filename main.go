package main

import (
	"net/http"
	"strconv"
	"time"

	"example.com/go-api-practice/db"
	"example.com/go-api-practice/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", getEvents)
	server.GET("/:id", getEventById)

	server.POST("/", createEvent)
	// update by id
	// delete by id
	// sign up
	// login
	// register user for event
	// cancel registration

	server.Run(":8080")
}

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

	event.ID = 1
	event.UserId = 1
	event.DataTime = time.Now()

	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event to database", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "data": event})
}
