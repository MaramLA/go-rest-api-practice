package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// users routes
	server.POST("/signup", signup)
	server.POST("/login", login)

	// events routes
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)
	server.POST("/event", createEvent)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)

	// register user for event
	// cancel registration}
}
