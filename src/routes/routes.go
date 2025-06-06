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

		userRoutes := apiRoutes.Group("/user")
		{
			authRoutes := userRoutes.Group("/auth")
			{
				authRoutes.POST("/register", controller.RegisterUser)
				authRoutes.POST("/login", controller.LoginUser)
				authRoutes.POST("/protected", middleware.AuthMiddleware(), controller.ProtectedEndpoint)
			}

			profileRoutes := userRoutes.Group("/profile")
			{
				profileRoutes.POST("/",  controller.GetUserByEmail)
				profileRoutes.POST("/profile-pic", controller.AddProfilePic)
				profileRoutes.PUT("/role/:id", controller.UpdateUserRoleById)
				profileRoutes.GET("/getid",  controller.GetUserIdByEmail)
			}

			locationRoutes := userRoutes.Group("/locations")
			{
				locationRoutes.POST("/", controller.CreateLocation)
				locationRoutes.GET("/", controller.GetLocations)
				locationRoutes.GET("/:id", controller.GetLocationByID)
				locationRoutes.PUT("/:id", controller.UpdateLocation)
				locationRoutes.DELETE("/:id", controller.DeleteLocation)
			}

			activityRoutes := userRoutes.Group("/activities")
			{
				activityRoutes.POST("/", controller.CreateActivity)
				activityRoutes.GET("/", controller.GetAllActivities)
				activityRoutes.GET("/:id", controller.GetActivityByID)
				activityRoutes.PUT("/:id", controller.UpdateActivity)
				activityRoutes.DELETE("/:id", controller.DeleteActivity)
				activityRoutes.GET("/desc/:id", controller.GetActivityDescriptionByActivityID)
			}

			eventRoutes := userRoutes.Group("/events")
			{
				eventRoutes.POST("/", controller.CreateEvent)
				eventRoutes.GET("/", controller.GetEvents)
				eventRoutes.GET("/activity/:id", controller.GetEventByActivityID)
				eventRoutes.GET("/location/:id", controller.GetEventByLocationID)
				eventRoutes.GET("/activity/location/:location_id/:activity_id", controller.GetEventByLocationIDAndActivityID)
			}

			beachRoutes := userRoutes.Group("/beaches")
			{
				beachRoutes.POST("/", controller.CreateBeach)
				beachRoutes.GET("/", controller.GetAllBeaches)
				beachRoutes.GET("/:id", controller.GetBeachByID)
				beachRoutes.PUT("/:id", controller.UpdateBeach)
				beachRoutes.DELETE("/:id", controller.DeleteBeach)
				beachRoutes.GET("/desc/:id", controller.GetBeachDescriptionByBeachID)
			}

			weatherRoutes := userRoutes.Group("/weather")
			{
				weatherRoutes.GET("/:id", controller.GetWeatherById)
			}

			forecastRoutes := userRoutes.Group("/forecast")
			{
				forecastRoutes.GET("/beach", controller.GetForecastHandler)
				forecastRoutes.GET("/advanced", controller.GetAdvancedForecastHandler)
			}
		}

		guideRoutes := apiRoutes.Group("/guide")
		{
			guideRoutes.POST("/", controller.CreateGuide)
			guideRoutes.GET("/", controller.GetAllGuides)
			guideRoutes.GET("/:id", controller.GetGuideByID)
			guideRoutes.PUT("/:id", controller.UpdateGuide)
			guideRoutes.DELETE("/:id", controller.DeleteGuide)

			guideListRoutes := guideRoutes.Group("/lists")
			{
				guideListRoutes.GET("/", controller.GetAllGuides)
				guideListRoutes.GET("/:id", controller.GetGuideByBeachID)
				guideListRoutes.GET("/activity/:id", controller.GetGuideByActivityID)
				guideListRoutes.GET("/:id/activities/:id", controller.GetGuideActivitiesByBeachIDAndActivityID)
			}
		}

		blogRoutes := apiRoutes.Group("/blogs")
		{
			blogRoutes.POST("/", controller.CreateBlog)
			blogRoutes.GET("/", controller.GetBlogs)
			blogRoutes.GET("/:id", controller.GetBlogByID)
			blogRoutes.PUT("/:id", controller.UpdateBlog)
			blogRoutes.DELETE("/:id", controller.DeleteBlog)
		}
	}
}
