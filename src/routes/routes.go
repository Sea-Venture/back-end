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

	activityRoutes := r.Group("/activities")
	activityRoutes.Use(middleware.AuthMiddleware())
	{
		activityRoutes.POST("/", controller.CreateActivity)
		activityRoutes.GET("/", controller.GetAllActivities)
		activityRoutes.GET("/:id", controller.GetActivityByID)
		activityRoutes.PUT("/:id", controller.UpdateActivity)
		activityRoutes.DELETE("/:id", controller.DeleteActivity)
		activityRoutes.GET("/desc/:id", controller.GetActivityDescriptionByActivityID)
	}

	beachRoutes := r.Group("/beaches")
	beachRoutes.Use(middleware.AuthMiddleware())
	{
		beachRoutes.POST("/", controller.CreateBeach)
		beachRoutes.GET("/", controller.GetAllBeaches)
		beachRoutes.GET("/:id", controller.GetBeachByID)
		beachRoutes.PUT("/:id", controller.UpdateBeach)
		beachRoutes.DELETE("/:id", controller.DeleteBeach)
		beachRoutes.GET("/desc/:id", controller.GetBeachDescriptionByBeachID)
	}

	weatherRoutes := r.Group("/weather")
	weatherRoutes.Use(middleware.AuthMiddleware())
	{
		weatherRoutes.GET("/:id", controller.GetWeatherById)
	}

	guideRoutes := r.Group("/guides")
	guideRoutes.Use(middleware.AuthMiddleware())
	{
		guideRoutes.POST("/", controller.CreateGuide)
		guideRoutes.GET("/", controller.GetAllGuides)
		guideRoutes.GET("/:id", controller.GetGuideByID)
		guideRoutes.PUT("/:id", controller.UpdateGuide)
		guideRoutes.DELETE("/:id", controller.DeleteGuide)
	}

	guideListRoutes := r.Group("/lists/guides")
	guideListRoutes.Use(middleware.AuthMiddleware())
	{
		guideListRoutes.GET("/", controller.GetAllGuides)
		guideListRoutes.GET("/:id", controller.GetGuideByBeachID)
		guideListRoutes.GET("/activity/:id", controller.GetGuideByActivityID)
		guideListRoutes.GET("/:id/activities/:id", controller.GetGuideActivitiesByBeachIDAndActivityID)

	}

}