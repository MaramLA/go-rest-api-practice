package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", getEvents)
	server.GET("/:id", getEventById)
	server.POST("/", createEvent)

	// update by id
	// delete by id
	// sign up
	// login
	// register user for event
	// cancel registration}
}
