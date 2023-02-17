package tests

import (
	"testing"

	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
)

func TestShouldReturnErrorWhenLoginIsEmptyString(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "",
		Passphrase: "123qweASD!@#",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when login is an empty string")
	}
}

func TestShouldReturnErrorWhenPassPhraseIsEmptyString(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "TestLogin",
		Passphrase: "",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when passphrase is an empty string")
	}
}

func TestShouldReturnErrorWhenPassPhraseIsSmallerThan6Chars(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "TestLogin",
		Passphrase: "12345",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when passphrase is smaller than 6 char")
	}
}

func TestShouldReturnErrorWhenPassPhraseHasNoNumber(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "TestLogin",
		Passphrase: "ABCdef!@#",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when passphrase has no number")
	}
}

func TestShouldReturnErrorWhenPassPhraseHasNoUpperCaseLetter(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "TestLogin",
		Passphrase: "def!@#123",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when passphrase has no uppercase letter")
	}
}

func TestShouldReturnErrorWhenPassPhraseHasNoSpecialChar(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "TestLogin",
		Passphrase: "ABCdef123",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when passphrase has no special char")
	}
}

func TestShouldNotReturnErrorWhenAuthIsValid(t *testing.T) {
	// Arrange
	data := dto.AuthCreateOrUpdateRequest{
		Login:      "test-login",
		Passphrase: "ABCdef123@!@#",
	}

	// Act
	err := validator.ValidateAuth(&data)

	// Assert
	if err != nil {
		t.Error("Expected validation to pass")
	}
}
