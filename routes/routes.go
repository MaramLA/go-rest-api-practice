package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// events routes
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)
	server.POST("/event", createEvent)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)

	// users routes
	server.POST("/signup", signup)

	// login
	// register user for event
	// cancel registration}
}
