package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
	mock_repository "github.com/online.scheduling-api/tests/infra/repository"
)

func TestShouldNotEditUserModalitiesWhenUserIsNotActive(t *testing.T) {
	// Arrange
	id := uuid.New()
	u := models.User{
		Id:       id,
		Name:     "Test",
		Phone:    "99999999999",
		Role:     models.Worker,
		IsActive: false,
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	userRepo.EXPECT().
		GetUserById(&id).
		Return(&u, nil).
		Times(1)

	service := services.UserModalitiesService{
		UserRepository: userRepo,
	}

	// Act
	code := service.Edit(&id, []uuid.UUID{uuid.New()})

	// Assert
	if code != shared.InvalidOperation {
		t.Errorf("Expected response to be %s when trying to edit modalities of a inactive user, but got %s instead", shared.InvalidOperation, code)
	}
}
