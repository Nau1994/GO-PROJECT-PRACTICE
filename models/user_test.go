package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInitialization(t *testing.T) {
	// Create a User instance
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Verify the field values
	assert.Equal(t, 1, user.ID, "ID should be 1")
	assert.Equal(t, "John Doe", user.Name, "Name should be 'John Doe'")
	assert.Equal(t, "john.doe@example.com", user.Email, "Email should be 'john.doe@example.com'")
}

func TestUserEmptyFields(t *testing.T) {
	// Create a User instance with empty fields
	user := User{
		ID:    0,
		Name:  "",
		Email: "",
	}

	// Verify the field values
	assert.Equal(t, 0, user.ID, "ID should be 0")
	assert.Equal(t, "", user.Name, "Name should be empty")
	assert.Equal(t, "", user.Email, "Email should be empty")
}
