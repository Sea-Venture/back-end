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
	config.InitFirebase()
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},                  
		AllowCredentials: true,                                                
	}))

	db := config.ConnectDB()

	routes.RegisterRoutes(r, db)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("env not found")
		err = godotenv.Load()
		if err != nil {
			log.Fatal("error loading env")
		}
	}
}
