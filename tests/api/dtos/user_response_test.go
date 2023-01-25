package test_api_dtos

import (
	"testing"

	"github.com/google/uuid"
	api "github.com/online.scheduling-api/src/api/dtos"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldMapSuccessfully(t *testing.T) {
	// Arrange
	id := uuid.New()

	const Name = "João da Silva"
	const Phone = "24999999999"

	u := models.User{
		Id:       id,
		Name:     Name,
		Phone:    Phone,
		IsActive: true,
		Role:     models.Worker,
	}

	// Act
	result := api.MapUserResponseFrom(&u)

	// Assert
	if result.Id != id ||
		result.Name != Name ||
		result.Phone != Phone ||
		!result.IsActive ||
		result.RoleCode != models.Worker {
		t.Error("Expected equivalent object but got different properties values")
	}
}
