package tests

import (
	"testing"

	"github.com/google/uuid"
	api "github.com/online.scheduling-api/src/api/dtos/responses"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldMapUserSuccessfully(t *testing.T) {
	// Arrange
	u := models.User{
		Id:       uuid.New(),
		Name:     "João da Silva",
		Phone:    "24999999999",
		IsActive: true,
		Role:     models.Worker,
		Modalities: []models.Modality{
			{
				Id:          uuid.New(),
				Name:        "Manicure",
				Description: "Corte, pintura e unhas de gel",
			},
		},
	}

	// Act
	result := api.MapUserResponseFrom(&u)

	// Assert
	if result.Id != u.Id ||
		result.Name != u.Name ||
		result.Phone != u.Phone ||
		!result.IsActive ||
		result.RoleCode != models.Worker ||
		len(result.Modalities) != len(u.Modalities) ||
		result.Modalities[0] != u.Modalities[0] {
		t.Error("Expected equivalent object but got different properties values")
	}
}
