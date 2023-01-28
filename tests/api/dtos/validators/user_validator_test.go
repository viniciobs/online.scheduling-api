package tests

import (
	"testing"

	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
	"github.com/online.scheduling-api/src/models"
)

func TestShouldReturnErrorWhenNameIsEmptyString(t *testing.T) {
	// Arrange
	u := dto.UserCreateOrUpdateRequest{
		Name:  "",
		Phone: "24999999999",
		Role:  models.Customer,
	}

	// Act
	err := validator.ValidateUser(&u)

	// Assert
	if err == nil {
		t.Error("Expected error when validating string empty for name")
	}
}

func TestShouldReturnErrorWhenPhoneIsEmptyString(t *testing.T) {
	// Arrange
	u := dto.UserCreateOrUpdateRequest{
		Name:  "Test",
		Phone: "",
		Role:  models.Customer,
	}

	// Act
	err := validator.ValidateUser(&u)

	// Assert
	if err == nil {
		t.Error("Expected error when validating string empty for phone")
	}
}

func TestShouldReturnErrorWhenPhoneNumberIsNotAValidPhoneNumber(t *testing.T) {
	// Arrange
	u := dto.UserCreateOrUpdateRequest{
		Name:  "Test",
		Phone: "123ABC-_ ",
		Role:  models.Customer,
	}

	// Act
	err := validator.ValidateUser(&u)

	// Assert
	if err == nil {
		t.Errorf("Expected error when validating \"%s\" for phone", u.Phone)
	}
}

func TestShouldReturnErrorWheRoleIsInvalid(t *testing.T) {
	// Arrange
	u := dto.UserCreateOrUpdateRequest{
		Name:  "Lorem Ipsum",
		Phone: "24999999999",
		Role:  -1,
	}

	// Act
	err := validator.ValidateUser(&u)

	// Assert
	if err == nil {
		t.Error("Expected error when role is invalid")
	}
}

func TestShouldNotReturnErrorWhenGivenUserIsValid(t *testing.T) {
	// Arrange
	u := dto.UserCreateOrUpdateRequest{
		Name:  "Lorem Ipsum",
		Phone: "24999999999",
		Role:  models.Customer,
	}

	// Act
	err := validator.ValidateUser(&u)

	// Assert
	if err != nil {
		t.Error("Expected validation to pass")
	}
}
