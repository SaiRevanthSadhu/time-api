package main

import (
	"log"

	"github.com/SaiRevanthSadhu/time-api/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config, err := database.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to the database
	db, err := database.ConnectDB(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize the Gin router
	router := gin.Default()

	// Define endpoints
	router.GET("/current-time", func(c *gin.Context) {
		handlers.HandleCurrentTime(c, db)
	})

	router.GET("/all-times", func(c *gin.Context) {
		handlers.HandleAllTimes(c, db)
	})

	// Start server
	router.Run(":8080")
}
