package routes

import (
	"seanThakur/go-restapi/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticatedEvent := server.Group("/")
	authenticatedEvent.Use(middlewares.ProtectedAuth)
	authenticatedEvent.POST("/events", createEvent)
	authenticatedEvent.PUT("/event/:id", updateEvent)
	authenticatedEvent.DELETE("/event/:id", deleteEvent)
	authenticatedEvent.POST("/events/:id/register", registerEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
