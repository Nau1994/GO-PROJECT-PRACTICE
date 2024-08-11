package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_DefaultPort(t *testing.T) {
	// Set up the environment variables
	os.Setenv("DATABASE_URL", "postgres://user:pass@localhost/db")
	os.Unsetenv("PORT")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.Equal(t, "8080", config.Port)
	assert.Equal(t, "postgres://user:pass@localhost/db", config.DatabaseURL)
}

func TestLoadConfig_CustomPort(t *testing.T) {
	// Set custom environment variables for testing
	os.Setenv("PORT", "9090")
	os.Setenv("DATABASE_URL", "postgres://user:pass@localhost/db")

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	assert.Equal(t, "9090", config.Port)
	assert.Equal(t, "postgres://user:pass@localhost/db", config.DatabaseURL)
}

func TestLoadConfig_MissingDatabaseURL(t *testing.T) {
	// Clear any existing environment variables to simulate a missing DATABASE_URL
	os.Unsetenv("DATABASE_URL")

	_, err := LoadConfig()
	if err == nil {
		t.Fatalf("Expected an error but got nil")
	}

	assert.Contains(t, err.Error(), "DATABASE_URL is not set")
}
