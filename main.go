package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/go-microservice/db"
	"github.com/yourusername/go-microservice/routes"
)

func main() {
	// Setup MongoDB connection
	db.SetupDatabase()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}
