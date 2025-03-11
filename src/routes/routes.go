package routes

import (
	"seaventures/src/controller"
	"seaventures/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controller.RegisterUser)
		userRoutes.POST("/login", controller.LoginUser)
		userRoutes.PUT("/profile-pic", middleware.AuthMiddleware(), controller.AddProfilePic)
	}

	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/protected", controller.ProtectedEndpoint)
	}

	blogRoutes := r.Group("/blogs")
	blogRoutes.Use(middleware.AuthMiddleware())
	{
		blogRoutes.POST("/", controller.CreateBlog)
		blogRoutes.GET("/", controller.GetBlogs)
		blogRoutes.GET("/:id", controller.GetBlogByID)
		blogRoutes.PUT("/:id", controller.UpdateBlog)
		blogRoutes.DELETE("/:id", controller.DeleteBlog)
	}

	eventRoutes := r.Group("/events")
	eventRoutes.Use(middleware.AuthMiddleware())
	{
		eventRoutes.POST("/", controller.CreateEvent)
		eventRoutes.GET("/", controller.GetEvents)
    }

	locationRoutes := r.Group("/locations")
	locationRoutes.Use(middleware.AuthMiddleware())
	{
		locationRoutes.POST("/", controller.CreateLocation)
		locationRoutes.GET("/", controller.GetLocations)
		locationRoutes.GET("/:id", controller.GetLocationByID)
		locationRoutes.PUT("/:id", controller.UpdateLocation)
		locationRoutes.DELETE("/:id", controller.DeleteLocation)
	}

}