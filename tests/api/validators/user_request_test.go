package api

import (
	"testing"

	validator "github.com/online.scheduling-api/src/api/validators"
	"github.com/online.scheduling-api/src/business/models"
)

func TestShouldReturnErrorWhenNameIsEmptyString(t *testing.T) {
	// Arrange
	u := models.User{
		Name:  "",
		Phone: "24999999999",
	}

	// Act
	err := validator.Validate(&u)

	// Assert
	if err == nil {
		t.Error("Expected error when validating string empty for name")
	}
}

func TestShouldReturnErrorWhenPhoneIsEmptyString(t *testing.T) {
	// Arrange
	u := models.User{
		Name:  "Test",
		Phone: "",
	}

	// Act
	err := validator.Validate(&u)

	// Assert
	if err == nil {
		t.Error("Expected error when validating string empty for phone")
	}
}

func TestShouldReturnErrorWhenPhoneNumberIsNotAValidPhoneNumber(t *testing.T) {
	// Arrange
	u := models.User{
		Name:  "Test",
		Phone: "123ABC-_ ",
	}

	// Act
	err := validator.Validate(&u)

	// Assert
	if err == nil {
		t.Errorf("Expected error when validating \"%s\" for phone", u.Phone)
	}
}

func TestShouldNotReturnErrorWhenGivenUserIsValid(t *testing.T) {
	// Arrange
	u := models.User{
		Name:  "Lorem Ipsum",
		Phone: "24999999999",
	}

	// Act
	err := validator.Validate(&u)

	// Assert
	if err != nil {
		t.Error("Expected validation to pass")
	}
}
