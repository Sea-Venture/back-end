package routes

import (
    "seaventures/src/controller"
    "seaventures/src/middleware"

    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

    userRoutes := r.Group("/users")
    {
        userRoutes.POST("/", controllers.RegisterUser)
        userRoutes.POST("/login", controllers.LoginUser)
        userRoutes.PUT("/profile-pic", middleware.AuthMiddleware(), controllers.AddProfilePic)
    }

    authRoutes := r.Group("/auth")
    authRoutes.Use(middleware.AuthMiddleware())
    {
        authRoutes.POST("/protected", controllers.ProtectedEndpoint)
    }
}

