package main

import (
	"example.com/go-api-practice/db"
	"example.com/go-api-practice/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
