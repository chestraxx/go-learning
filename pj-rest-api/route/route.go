package route

import (
	"example.com/pj-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/event/:id", getEvent)
	r.POST("/signup", signup)
	r.POST("/login", login)

	auth := r.Group("/")
	auth.Use(middleware.Auth)
	auth.POST("/event", createEvent)
	auth.PUT("/event/:id", updateEvent)
	auth.DELETE("/event/:id", deleteEvent)
	auth.POST("/event/:id/register", registerForEvent)
	auth.DELETE("/event/:id/register", unregisterForEvent)
}
