package tests

import (
	"testing"

	api "github.com/online.scheduling-api/src/api/dtos/requests"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldMapAuthRequestCorrectly(t *testing.T) {
	// Arrange
	u := api.UserCreateOrUpdateRequest{
		Name:       "Test",
		Phone:      "24999999999",
		Role:       models.Customer,
		Login:      "test-login",
		Passphrase: "123qweASD!@#",
	}

	// Act
	response := api.MapAuthRequestFrom(&u)

	// Assert
	if response.Login != u.Login ||
		response.Passphrase != u.Passphrase {
		t.Error("Expected equivalent object but got different properties values")
	}
}
