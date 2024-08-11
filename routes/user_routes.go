package routes

import (
	"project/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// Define routes
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
	}
}
