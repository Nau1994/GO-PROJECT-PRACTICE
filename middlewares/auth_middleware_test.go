package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_ValidToken(t *testing.T) {
	// Setup
	router := gin.Default()
	router.Use(AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Authorized")
	})

	// Create a request with a valid token
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "valid-token")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Authorized", w.Body.String())
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	// Setup
	router := gin.Default()
	router.Use(AuthMiddleware())
	router.GET("/protected", func(c *gin.Context) {
		c.String(http.StatusOK, "Authorized")
	})

	// Create a request with an invalid token
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "invalid-token")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Unauthorized")
}
