package tests

import (
	"testing"

	"github.com/online.scheduling-api/src/helpers"
)

func TestShouldCryptAndDecryptReturnExpectedValue(t *testing.T) {
	// Arrange
	value := "123qweASD!@#"

	// Act
	encrypted := helpers.Crypt(value)
	decrypted := helpers.Decrypt(encrypted)

	// Assert
	if encrypted == value {
		t.Error("Expected value to be encrypted")
	}

	if decrypted != value {
		t.Error("Expected value do be decrypted correctly")
	}
}
