package routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserRoutes(t *testing.T) {
	// Setup
	router := gin.Default()
	UserRoutes(router)

	// Test GET route
	req, _ := http.NewRequest("GET", "/users/", nil) // Added trailing slash
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")

	// Test POST route
	body := `{"id": 1, "name": "New User", "email": "newuser@example.com"}`
	req, _ = http.NewRequest("POST", "/users/", bytes.NewBufferString(body)) // Added trailing slash
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New User")
}
