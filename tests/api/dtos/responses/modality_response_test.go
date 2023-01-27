package tests

import (
	"testing"

	"github.com/google/uuid"
	api "github.com/online.scheduling-api/src/api/dtos/responses"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldMapModalitySuccessfully(t *testing.T) {
	// Arrange
	id := uuid.New()

	const Name = "Manicure"
	const Description = "Corte, pintura e unhas de gel"

	m := models.Modality{
		Id:          id,
		Name:        Name,
		Description: Description,
	}

	// Act
	result := api.MapModalityResponseFrom(&m)

	// Assert
	if result.Id != id ||
		result.Name != Name ||
		result.Description != Description {
		t.Error("Expected equivalent object but got different properties values")
	}
}
