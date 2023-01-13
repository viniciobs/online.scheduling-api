package models

import (
	"testing"

	"github.com/online.scheduling-api/src/business/models"
)

func TestShouldRemoveWhiteSpacesAtStartAndAtEnd(t *testing.T) {
	// Arrange
	u := models.User{
		Name:  "  first-name last-name  ",
		Phone: "   24999999999   ",
	}

	// Act
	u.RemoveWhiteSpaces()

	// Assert
	if u.Name != "first-name last-name" || u.Phone != "24999999999" {
		t.Error("Expected to remove white spaces at start and at the end")
	}
}
