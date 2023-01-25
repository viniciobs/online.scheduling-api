package test_services

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	mock_repository "github.com/online.scheduling-api/tests/repository"
)

func TestShouldReturnIsDuplicatedEqualTrueWhenTryingToCreateUserWithPhoneAlreadyRegistered(t *testing.T) {
	// Arrange
	const Phone = "99999999999"

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_repository.NewMockIUserRepository(mockCtrl)
	repo.EXPECT().
		ExistsByPhone(Phone).
		Return(true, nil).
		Times(1)

	service := services.UserServices{
		UserRepository: repo,
	}

	u := models.User{
		Id:       uuid.New(),
		Name:     "Test",
		Phone:    Phone,
		IsActive: true,
		Role:     models.Admin,
	}

	// Act
	isDuplicated, _ := service.CreateNewUser(&u)

	// Assert
	if !isDuplicated {
		t.Errorf("Expecting isDuplicated to be true when creating user which phone is already registered, But got %s", strconv.FormatBool(isDuplicated))
	}
}

func TestShouldReturnIsDuplicatedEqualFalseWhenTryingToCreateUserWithPhoneNotRegistered(t *testing.T) {
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
		ExistsByPhone(u.Phone).
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
	isDuplicated, _ := service.CreateNewUser(&u)

	// Assert
	if isDuplicated {
		t.Errorf("Expecting isDuplicated to be false when creating user which phone is not registered, But got %s", strconv.FormatBool(isDuplicated))
	}
}
