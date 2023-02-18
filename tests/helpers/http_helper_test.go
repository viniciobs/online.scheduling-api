package tests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/shared"
)

func TestShouldNotThrowErrorWhenNilError(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()

	// Act
	helpers.JSONResponseError(w, 404, nil)

	res := w.Result()
	defer res.Body.Close()

	_, err := io.ReadAll(res.Body)

	// Assert
	if err != nil {
		t.Errorf("Expected to not throw errors but got %v", err)
	}
}

func TestShouldReturnExpectedStructureForError(t *testing.T) {
	// Arrange
	const ErrorMessage = "Something wrent wrong"

	w := httptest.NewRecorder()

	// Act
	helpers.JSONResponseError(w, 500, errors.New(ErrorMessage))

	res := w.Result()
	defer res.Body.Close()

	var e helpers.Error
	data, _ := io.ReadAll(res.Body)
	json.Unmarshal(data, &e)

	// Assert
	if e.Message != ErrorMessage {
		t.Errorf("Expected error message to be %v, got %v", ErrorMessage, e.Message)
	}
}

func TestShouldNotThrowErrorWhenNilBody(t *testing.T) {
	// Arrange
	w := httptest.NewRecorder()

	// Act
	helpers.JSONResponse(w, 204, nil)

	res := w.Result()
	defer res.Body.Close()

	_, err := io.ReadAll(res.Body)

	// Assert
	if err != nil {
		t.Errorf("Expected to not throw errors but got %v", err)
	}
}

func TestShouldReturnExpectedStatusCodeFromResponseCode(t *testing.T) {
	// Arrange
	errors := map[shared.Code]int{
		shared.DuplicatedRecord:  http.StatusUnprocessableEntity,
		shared.InvalidOperation:  http.StatusUnprocessableEntity,
		shared.NonExistentRecord: http.StatusNotFound,
		shared.ThirdPartyFail:    http.StatusBadGateway,
	}

	// Act

	// Assert
	for responseCode, statusResponse := range errors {
		result := helpers.GetErrorStatusCodeFrom(responseCode)
		if result != statusResponse {
			t.Errorf("Expected %d and got %d", statusResponse, result)
		}
	}
}
