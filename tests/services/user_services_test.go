package tests

import (
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

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsByPhone(&id, &phone).
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
	}

	// Act
	code := service.CreateNewUser(&u)

	// Assert
	if code != shared.DuplicatedRecord {
		t.Errorf("Expecting response code to be %s when creating user which phone is already registered to other, But got %s", shared.DuplicatedRecord, code)
	}
}

func TestShouldReturnSuccessWhenTryingToCreateUserWithPhoneNotRegistered(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    "99999999999",
		IsActive: true,
		Role:     models.Admin,
	}

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsByPhone(&u.Id, &u.Phone).
		Return(false, nil).
		Times(1)
	repo.EXPECT().
		CreateNewUser(&u).
		Return(nil).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	// Act
	code := service.CreateNewUser(&u)

	// Assert
	if code != shared.Success {
		t.Error("Expecting to succssfully execute when creating user which phone is not registered and has valid fields")
	}
}

func TestShouldReturnThirdPartyFailWhenDatabaseIsNotConnected(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    "99999999999",
		IsActive: true,
		Role:     models.Admin,
	}

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsByPhone(&u.Id, &u.Phone).
		Return(false, mongo.ErrClientDisconnected).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	// Act
	code := service.CreateNewUser(&u)

	// Assert
	if code != shared.ThirdPartyFail {
		t.Errorf("Expecting response code to be %s when database is unavailable but got %s instead", shared.ThirdPartyFail, code)
	}
}
