package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"project/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("DATABASE_URL", "postgres://user:pass@localhost/db")
	// Setup the router
	router := setupRouter()

	// Test a simple GET request
	req, _ := http.NewRequest("GET", "/users/", nil) // Added trailing slash
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

// setupRouter is a helper function to set up the router for testing
func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.UserRoutes(router)
	return router
}

func TestMainIntegration(t *testing.T) {
	go main() // Run the main function in a goroutine

	// Allow some time for the server to start up
	time.Sleep(1 * time.Second)

	// Test the running server
	resp, err := http.Get("http://localhost:8080/users/")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
