package main

import (
	"fmt"
	"project/config"
	"project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	config, _ := config.LoadConfig()

	// Set up the Gin router
	router := gin.Default()

	// Set up routes
	routes.UserRoutes(router)

	// Start the server
	err := router.Run(fmt.Sprintf(":%s", config.Port))
	if err != nil {
		panic(err)
	}
}
