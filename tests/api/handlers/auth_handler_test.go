package tests

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
	mock_repository "github.com/online.scheduling-api/tests/infra/repository"
)

func TestShouldReturnBadRequestWhenMappingBodyIsNil(t *testing.T) {
	// Arrange
	req := createRequest(nil, t)
	rr := httptest.NewRecorder()
	ah := api.AuthHandler{}
	h := http.HandlerFunc(ah.SignIn)

	// Act
	h.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Auth handler returned wrong status code. Wanted %d and got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestShouldReturnUnauthorizedWhenInvalidCredentialsPassed(t *testing.T) {
	// Arrange
	jsonBody := []byte(`{"login": "login_test", "passphrase": "passphrase_test"}`)
	body := bytes.NewReader(jsonBody)

	req := createRequest(body, t)
	rr := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		Authenticate(ctx, "login_test", helpers.Crypt("passphrase_test")).
		Return(false, nil)

	ah := api.AuthHandler{UserRepository: repo}
	h := http.HandlerFunc(ah.SignIn)

	// Act
	h.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Auth handler returned wrong status code. Wanted %d and got %d", rr.Code, http.StatusUnauthorized)
	}
}

func TestShouldReturnSuccessWhenValidCredentialsPassed(t *testing.T) {
	// Arrange
	jsonBody := []byte(`{"login": "login_test", "passphrase": "passphrase_test"}`)
	body := bytes.NewReader(jsonBody)

	req := createRequest(body, t)
	rr := httptest.NewRecorder()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	u := &models.User{
		Id:         uuid.New(),
		Name:       "test",
		Login:      "login_test",
		Passphrase: "passphrase_test",
	}

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		Authenticate(ctx, "login_test", helpers.Crypt("passphrase_test")).
		Return(true, u)

	ah := api.AuthHandler{UserRepository: repo}
	h := http.HandlerFunc(ah.SignIn)

	// Act
	h.ServeHTTP(rr, req)

	// Assert
	if rr.Code != http.StatusOK {
		t.Errorf("Auth handler returned wrong status code. Wanted %d and got %d", rr.Code, http.StatusOK)
	}
}

func createRequest(body io.Reader, t *testing.T) *http.Request {
	req, err := http.NewRequest("POST", "/api/sign-in", body)
	if err != nil {
		t.Fatal(err)
	}

	return req
}
