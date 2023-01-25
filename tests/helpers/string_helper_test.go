package test_helpers

import (
	"testing"

	"github.com/online.scheduling-api/src/helpers"
)

func TestShouldRemoveWhitespacesAntStartAndEntOfText(t *testing.T) {
	// Arrange
	const Text = "   Hello, world! "

	// Act
	result := helpers.TrimStartAndEnd(Text)

	// Assert
	if result == Text {
		t.Errorf("Should've removed white spaces from start and end. Got %s expected %s", result, Text)
	}
}
