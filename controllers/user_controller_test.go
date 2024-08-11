package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", GetUsers)
	router.POST("/users", CreateUser)
	return router
}

func TestGetUsers(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "Jane Doe")
}

func TestCreateUser_Success(t *testing.T) {
	router := setupRouter()

	user := models.User{
		Name:  "New User",
		Email: "newuser@example.com",
	}
	userJSON, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New User")
	assert.Contains(t, w.Body.String(), "newuser@example.com")
}

func TestCreateUser_InvalidInput(t *testing.T) {
	router := setupRouter()

	// Send invalid JSON
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
