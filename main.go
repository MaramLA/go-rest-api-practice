package main

import (
	"net/http"
	"time"

	"example.com/go-api-practice/models"
	"github.com/gin-gonic/gin"
)



func main(){

	server := gin.Default()

	server.GET("/", getEvents)
	server.POST("/", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context){
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, gin.H{"message": "events retrieved successfully", "data": events})
}

func createEvent(c *gin.Context){

	var event models.Event

	err := c.ShouldBindJSON(&event)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body", "error": err.Error()})
		return
	}

	event.ID = len(models.GetAllEvents())+1
	event.UserId = len(models.GetAllEvents())+1
	event.DataTime = time.Now()

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "event created successfully", "data": event})

	
}