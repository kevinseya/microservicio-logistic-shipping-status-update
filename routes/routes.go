package routes

import (
	"shippment-asignment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api/shippment")
	{
		api.PATCH("/update-state/:id", controllers.UpdateShipmentStatus)
	}
}
