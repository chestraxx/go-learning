package route

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.Engine) {
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/event", createEvent)
	r.PUT("/event/:id", updateEvent)
	r.DELETE("/event/:id", deleteEvent)
	r.POST("/signup", signup)
}
