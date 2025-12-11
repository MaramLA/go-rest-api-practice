package routes

import (
	"example.com/go-api-practice/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// users routes
	server.POST("/signup", signup)
	server.POST("/login", login)

	// protected routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	// events routes
	authenticated.GET("/events", getEvents)
	authenticated.GET("/event/:id", getEventById)
	authenticated.POST("/event", createEvent)
	authenticated.PUT("/event/:id", updateEvent)
	authenticated.DELETE("/event/:id", deleteEvent)

	// registrations routes
	authenticated.POST("/event/:id/register", registerForEvent)
	authenticated.DELETE("/event/:id/register", cancelRegistration)
}
