package tests

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
	mock_repository "github.com/online.scheduling-api/tests/infra/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestShouldReturnIsDuplicatedEqualTrueWhenTryingToCreateUserWithPhoneAlreadyRegisteredToOtherUser(t *testing.T) {
	// Arrange
	phone := "99999999999"
	id := uuid.New()
	login := "test123"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsBy(ctx, &id, &phone, &login).
		Return(true, nil).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	u := models.User{
		Id:       id,
		Name:     "Test",
		Phone:    phone,
		IsActive: true,
		Role:     models.Admin,
		Login:    login,
	}

	// Act
	code, _ := service.CreateNewUser(ctx, &u)

	// Assert
	if code != shared.DuplicatedRecord {
		t.Errorf("Expecting response code to be %s when creating user which phone is already registered to other, But got %s", shared.DuplicatedRecord, code)
	}
}

func TestShouldReturnSuccessWhenTryingToCreateUserWithPhoneNotRegistered(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    "99999999999",
		IsActive: true,
		Role:     models.Admin,
		Login:    "test123",
	}

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsBy(ctx, &u.Id, &u.Phone, &u.Login).
		Return(false, nil).
		Times(1)
	repo.EXPECT().
		CreateNewUser(ctx, &u).
		Return(nil).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	// Act
	code, _ := service.CreateNewUser(ctx, &u)

	// Assert
	if code != shared.Success {
		t.Error("Expecting to succssfully execute when creating user which phone is not registered and has valid fields")
	}
}

func TestShouldReturnDuplicatedRecordWhenTryingToEditUserWithPhoneAlreadyRegisteredToOtherUser(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ctx := context.Background()
	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    "99999999999",
		IsActive: true,
		Role:     models.Admin,
		Login:    "test123",
	}

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsBy(ctx, &u.Id, &u.Phone, &u.Login).
		Return(true, nil).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	// Act
	code, _ := service.CreateNewUser(ctx, &u)

	// Assert
	if code != shared.DuplicatedRecord {
		t.Errorf("Expecting response code to be %s when trying to update users with phone already registered to others. Got %s", shared.DuplicatedRecord, code)
	}
}

func TestShouldReturnThirdPartyFailWhenDatabaseIsNotConnected(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ctx := context.Background()

	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    "99999999999",
		IsActive: true,
		Role:     models.Admin,
		Login:    "test123",
	}

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsBy(ctx, &u.Id, &u.Phone, &u.Login).
		Return(false, mongo.ErrClientDisconnected).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	// Act
	code, _ := service.CreateNewUser(ctx, &u)

	// Assert
	if code != shared.ThirdPartyFail {
		t.Errorf("Expecting response code to be %s when database is unavailable but got %s instead", shared.ThirdPartyFail, code)
	}
}

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

	ctx := context.Background()

	userRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	userRepo.EXPECT().
		GetUserById(ctx, &id).
		Return(&u, nil).
		Times(1)

	service := services.UserServices{
		UserRepository: userRepo,
	}

	// Act
	code := service.EditModalities(ctx, &id, []uuid.UUID{uuid.New()})

	// Assert
	if code != shared.InvalidOperation {
		t.Errorf("Expected response to be %s when trying to edit modalities of a inactive user, but got %s instead", shared.InvalidOperation, code)
	}
}
