package controllers

import (
	"net/http"
	"project/models"

	"github.com/gin-gonic/gin"
)

// GetUsers handles GET requests to retrieve a list of users
func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user models.User

	// Bind JSON to user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Normally you would save the user to a database here

	c.JSON(http.StatusCreated, user)
}
