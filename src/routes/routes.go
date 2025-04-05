package routes

import (
	"seaventures/src/controller"
	"seaventures/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	apiRoutes := r.Group("/api")
	{
		// User routes under /api/user
		userRoutes := apiRoutes.Group("/user")
		{
			authRoutes := userRoutes.Group("/auth")
			{
				authRoutes.POST("/register", controller.RegisterUser)
				authRoutes.POST("/login", controller.LoginUser)
				authRoutes.POST("/profile-pic", middleware.AuthMiddleware(), controller.AddProfilePic)
				authRoutes.POST("/protected", middleware.AuthMiddleware(), controller.ProtectedEndpoint)
			}

			// Location routes under /api/user/locations
			locationRoutes := userRoutes.Group("/locations")
			locationRoutes.Use(middleware.AuthMiddleware())
			{
				locationRoutes.POST("/", controller.CreateLocation)
				locationRoutes.GET("/", controller.GetLocations)
				locationRoutes.GET("/:id", controller.GetLocationByID)
				locationRoutes.PUT("/:id", controller.UpdateLocation)
				locationRoutes.DELETE("/:id", controller.DeleteLocation)
			}

			// Activity routes under /api/user/activities
			activityRoutes := userRoutes.Group("/activities")
			activityRoutes.Use(middleware.AuthMiddleware())
			{
				activityRoutes.POST("/", controller.CreateActivity)
				activityRoutes.GET("/", controller.GetAllActivities)
				activityRoutes.GET("/:id", controller.GetActivityByID)
				activityRoutes.PUT("/:id", controller.UpdateActivity)
				activityRoutes.DELETE("/:id", controller.DeleteActivity)
				activityRoutes.GET("/desc/:id", controller.GetActivityDescriptionByActivityID)
			}

			// Event routes under /api/user/events
			eventRoutes := userRoutes.Group("/events")
			eventRoutes.Use(middleware.AuthMiddleware())
			{
				eventRoutes.POST("/", controller.CreateEvent)
				eventRoutes.GET("/", controller.GetEvents)
			}

			// Beach routes under /api/user/beaches
			beachRoutes := userRoutes.Group("/beaches")
			beachRoutes.Use(middleware.AuthMiddleware())
			{
				beachRoutes.POST("/", controller.CreateBeach)
				beachRoutes.GET("/", controller.GetAllBeaches)
				beachRoutes.GET("/:id", controller.GetBeachByID)
				beachRoutes.PUT("/:id", controller.UpdateBeach)
				beachRoutes.DELETE("/:id", controller.DeleteBeach)
				beachRoutes.GET("/desc/:id", controller.GetBeachDescriptionByBeachID)
			}

			// Weather routes under /api/user/weather
			weatherRoutes := userRoutes.Group("/weather")
			weatherRoutes.Use(middleware.AuthMiddleware())
			{
				weatherRoutes.GET("/:id", controller.GetWeatherById)
			}

			// Forecast routes under /api/user/forecast
			forecastRoutes := userRoutes.Group("/forecast")
			{
				forecastRoutes.GET("/", func(c *gin.Context) {
					controller.GetForecastHandler(c.Writer, c.Request)
				})
			}
		}

		// Guide routes under /api/guide
		guideRoutes := apiRoutes.Group("/guide")
		guideRoutes.Use(middleware.AuthMiddleware())
		{
			guideRoutes.POST("/", controller.CreateGuide)
			guideRoutes.GET("/", controller.GetAllGuides)
			guideRoutes.GET("/:id", controller.GetGuideByID)
			guideRoutes.PUT("/:id", controller.UpdateGuide)
			guideRoutes.DELETE("/:id", controller.DeleteGuide)

			// Guide list routes under /api/guide/lists
			guideListRoutes := guideRoutes.Group("/lists")
			{
				guideListRoutes.GET("/", controller.GetAllGuides)
				guideListRoutes.GET("/:id", controller.GetGuideByBeachID)
				guideListRoutes.GET("/activity/:id", controller.GetGuideByActivityID)
				guideListRoutes.GET("/:id/activities/:id", controller.GetGuideActivitiesByBeachIDAndActivityID)
			}
		}

		// Blog routes under /api/blogs
		blogRoutes := apiRoutes.Group("/blogs")
		blogRoutes.Use(middleware.AuthMiddleware())
		{
			blogRoutes.POST("/", controller.CreateBlog)
			blogRoutes.GET("/", controller.GetBlogs)
			blogRoutes.GET("/:id", controller.GetBlogByID)
			blogRoutes.PUT("/:id", controller.UpdateBlog)
			blogRoutes.DELETE("/:id", controller.DeleteBlog)
		}
	}
}
