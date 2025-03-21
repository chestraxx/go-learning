package main

import (
	"example.com/pj-rest-api/db"
	"example.com/pj-rest-api/route"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	route.InitRoute(server)

	server.Run(":8080")
}
