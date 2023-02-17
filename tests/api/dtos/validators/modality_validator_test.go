package tests

import (
	"testing"

	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
)

func TestShouldReturnErrorWhenModalityNameIsEmptyString(t *testing.T) {
	// Arrange
	m := dto.ModalityCreateOrUpdateRequest{
		Name:        "",
		Description: "Lorem Ipsum",
	}

	// Act
	err := validator.ValidateModality(&m)

	// Assert
	if err == nil {
		t.Error("Expected error when validating string empty for name")
	}
}

func TestShouldReturnErrorWhenDescriptionLengthIsHigherThan140(t *testing.T) {
	// Arrange
	m := dto.ModalityCreateOrUpdateRequest{
		Name:        "Test",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
	}

	// Act
	err := validator.ValidateModality(&m)

	// Assert
	if err == nil {
		t.Error("Expected error when modality description is higher than 140")
	}
}

func TestShouldNotReturnErrorWhenGivenModalityIsValid(t *testing.T) {
	// Arrange
	m := dto.ModalityCreateOrUpdateRequest{
		Name:        "Lorem Ipsum",
		Description: "Lorem ipsum dolor sit amet",
	}

	// Act
	err := validator.ValidateModality(&m)

	// Assert
	if err != nil {
		t.Error("Expected validation to pass")
	}
}
