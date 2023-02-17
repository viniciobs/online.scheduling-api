package tests

import (
	"testing"

	"github.com/google/uuid"
	api "github.com/online.scheduling-api/src/api/dtos/responses"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldMapModalitySuccessfully(t *testing.T) {
	// Arrange
	m := models.Modality{
		Id:          uuid.New(),
		Name:        "Manicure",
		Description: "Corte, pintura e unhas de gel",
	}

	// Act
	result := api.MapModalityResponseFrom(&m)

	// Assert
	if result.Id != m.Id ||
		result.Name != m.Name ||
		result.Description != m.Description {
		t.Error("Expected equivalent object but got different properties values")
	}
}
