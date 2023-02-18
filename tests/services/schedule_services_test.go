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

func TestShouldReturnDuplicatedRecordWhenScheduleExists(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userId := uuid.New()
	modalityId := uuid.New()

	repo := mock_repository.NewMockIScheduleRepository(mockCtrl)
	repo.EXPECT().
		Get(gomock.Any()).
		Return([]*models.Schedule{{UserId: userId, ModalityId: modalityId}}, nil)

	service := services.ScheduleService{
		ScheduleRepository: repo,
	}

	// Act
	responseCode := service.Create(&userId, &modalityId, nil)

	// Assert
	if responseCode != shared.DuplicatedRecord {
		t.Error("Expected code to be duplicated record")
	}
}

func TestShouldReturnNonExistentRecordWhenGivenModalityDoesNotExists(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userId := uuid.New()
	modalityId := uuid.New()

	repo := mock_repository.NewMockIScheduleRepository(mockCtrl)
	repo.EXPECT().
		Get(gomock.Any()).
		Return([]*models.Schedule{}, nil)

	modalityRepo := mock_repository.NewMockIModalityRepository(mockCtrl)
	modalityRepo.EXPECT().
		GetModalityById(&modalityId).
		Return(nil, nil)

	service := services.ScheduleService{
		ScheduleRepository: repo,
		ModalityRepository: modalityRepo,
	}

	// Act
	responseCode := service.Create(&userId, &modalityId, nil)

	// Assert
	if responseCode != shared.NonExistentRecord {
		t.Error("Expected code to be non existent record")
	}
}

func TestShouldReturnNonExistentRecordWhenGivenUserDoesNotExists(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userId := uuid.New()
	modalityId := uuid.New()

	repo := mock_repository.NewMockIScheduleRepository(mockCtrl)
	repo.EXPECT().
		Get(gomock.Any()).
		Return([]*models.Schedule{}, nil)

	modalityRepo := mock_repository.NewMockIModalityRepository(mockCtrl)
	modalityRepo.EXPECT().
		GetModalityById(&modalityId).
		Return(&models.Modality{Id: modalityId}, nil)

	userRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	userRepo.EXPECT().
		GetUserById(&userId).
		Return(nil, nil)

	service := services.ScheduleService{
		ScheduleRepository: repo,
		ModalityRepository: modalityRepo,
		UserRespository:    userRepo,
	}

	// Act
	responseCode := service.Create(&userId, &modalityId, nil)

	// Assert
	if responseCode != shared.NonExistentRecord {
		t.Error("Expected code to be non existent record")
	}
}

func TestShouldReturnInvalidOperationWhenGivenUserDoesNotWorkWithGivenModality(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userId := uuid.New()
	modalityId := uuid.New()

	repo := mock_repository.NewMockIScheduleRepository(mockCtrl)
	repo.EXPECT().
		Get(gomock.Any()).
		Return([]*models.Schedule{}, nil)

	modalityRepo := mock_repository.NewMockIModalityRepository(mockCtrl)
	modalityRepo.EXPECT().
		GetModalityById(&modalityId).
		Return(&models.Modality{Id: modalityId}, nil)

	userRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	userRepo.EXPECT().
		GetUserById(&userId).
		Return(
			&models.User{
				Id:         userId,
				Modalities: []models.Modality{},
			},
			nil)

	service := services.ScheduleService{
		ScheduleRepository: repo,
		ModalityRepository: modalityRepo,
		UserRespository:    userRepo,
	}

	// Act
	responseCode := service.Create(&userId, &modalityId, nil)

	// Assert
	if responseCode != shared.InvalidOperation {
		t.Error("Expected code to be invalid operation")
	}
}
