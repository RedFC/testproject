package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/RedFC/testproject/database"
	"github.com/RedFC/testproject/routes"
)

func main() {
	// Initialize the database connection
	database.InitDB()

	// Set up the Gin router
	router := gin.Default()

	var apiVersion string = "/api/v1"

	// Define routes
	routes.SetupProductRoutes(router, apiVersion)
	routes.SetupUserRoutes(router, apiVersion)

	// Run the server
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
}
