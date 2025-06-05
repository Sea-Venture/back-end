package main

import (
	"log"
	"seaventures/src/config"
	"seaventures/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Allow requests from your frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Headers exposed to the client
		AllowCredentials: true,                                                // Allow cookies and credentials
	}))

	db := config.ConnectDB()

	routes.RegisterRoutes(r, db)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load("./.env.local")
	if err != nil {
		log.Println("env not found")
		err = godotenv.Load()
		if err != nil {
			log.Fatal("error loading env")
		}
	}
}


