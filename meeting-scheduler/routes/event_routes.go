package routes

import (
	"meeting-scheduler/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	eventRoutes := router.Group("/events")
	{
		eventRoutes.POST("/", controllers.CreateEvent)
		eventRoutes.GET("/:id", controllers.GetEvent)
		eventRoutes.DELETE("/:id", controllers.DeleteEvent)

		eventRoutes.POST("/:id/availability", controllers.AddUserAvailability)
		eventRoutes.GET("/:id/recommendations", controllers.GetRecommendedSlots)
	}
}
