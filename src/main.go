package main

import (
	"log"
	"seaventures/src/config"
	"seaventures/src/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()

	// Initialize database
	config.ConnectDB()

	// Register routes
	routes.RegisterRoutes(r)

	r.Run(":8080") // Run server on port 8080
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
