package routes

import (
	controllers "seaventures/src/controller"
	"seaventures/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.PUT("/", controllers.UpdateUser)
	}

	// Apply AuthMiddleware to routes that require authentication
	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/protected", controllers.ProtectedEndpoint)
	}
}
