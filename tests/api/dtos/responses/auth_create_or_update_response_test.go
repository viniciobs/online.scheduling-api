package tests

import (
	"testing"

	"github.com/google/uuid"
	api "github.com/online.scheduling-api/src/api/dtos/responses"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldMapResonseCorrectly(t *testing.T) {
	// Arrange
	const Token = "ABCQWE@!@#--qwe==9999"

	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    "24999999999",
		Role:     models.Worker,
		IsActive: false,
	}

	// Act
	result := api.MapAuthResponseFrom(&u, Token)

	// Assert
	if u.Id != result.UserId ||
		u.Name != result.Name ||
		u.IsActive != result.IsActive ||
		u.Role != result.RoleCode ||
		u.Role.GetDescription() != result.RoleDescription ||
		result.Token != Token {
		t.Error("Expected equivalent object but got different properties values")
	}
}
