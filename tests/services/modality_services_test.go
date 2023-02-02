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

func TestShouldReturnIsDuplicatedEqualTrueWhenTryingToCreateMOdalityWithNameAlreadyRegisteredToOtherModality(t *testing.T) {
	// Arrange
	name := "Test"
	id := uuid.New()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_repository.NewMockIModalityRepository(mockCtrl)
	repo.EXPECT().
		ExistsByName(&id, &name).
		Return(true, nil).
		Times(1)

	service := services.ModalityService{
		ModalityRepository: repo,
	}

	m := models.Modality{
		Id:          id,
		Name:        "Test",
		Description: "Lorem Ipsum",
	}

	// Act
	code := service.CreateNewModality(&m)

	// Assert
	if code != shared.DuplicatedRecord {
		t.Errorf("Expecting response code to be %s when creating modality which name is already registered to other, But got %s", shared.DuplicatedRecord, code)
	}
}

func TestShouldReturnSuccessWhenTryingToCreateModalityWithNameNotRegistered(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := models.Modality{
		Id:          uuid.New(),
		Name:        "Test",
		Description: "Lorem Ipsum",
	}

	repo := mock_repository.NewMockIModalityRepository(mockCtrl)
	repo.EXPECT().
		ExistsByName(&m.Id, &m.Name).
		Return(false, nil).
		Times(1)
	repo.EXPECT().
		CreateNewModality(&m).
		Return(nil).
		Times(1)

	service := services.ModalityService{
		ModalityRepository: repo,
	}

	// Act
	code := service.CreateNewModality(&m)

	// Assert
	if code != shared.Success {
		t.Error("Expecting to succssfully execute when creating modality which name is not registered and has valid fields")
	}
}

func TestShouldReturnDuplicatedRecordWhenTryingToEditModalityWithNameAlreadyRegisteredToOtherModality(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var emptyUsers []*models.User

	m := models.Modality{
		Id:          uuid.New(),
		Name:        "Test",
		Description: "Lorem Ipsum",
	}

	repo := mock_repository.NewMockIModalityRepository(mockCtrl)
	repo.EXPECT().
		ExistsByName(&m.Id, &m.Name).
		Return(true, nil).
		Times(1)

	uRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	uRepo.EXPECT().
		Get(&models.UserFilter{ModalityId: m.Id}).
		Return(emptyUsers, nil).
		Times(1)

	service := services.ModalityService{
		ModalityRepository: repo,
		UserRepository:     uRepo,
	}

	// Act
	code := service.EditModality(&m.Id, &m)

	// Assert
	if code != shared.DuplicatedRecord {
		t.Errorf("Expecting response code to be %s when trying to update modality with name already registered to others. Got %s", shared.DuplicatedRecord, code)
	}
}

func TestShouldReturnInvalidOperationWhenTryingToDeleteModalityInUseByUsers(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	users := []*models.User{{Id: uuid.New()}}

	m := models.Modality{
		Id:          uuid.New(),
		Name:        "Test",
		Description: "Lorem Ipsum",
	}

	uRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	uRepo.EXPECT().
		Get(&models.UserFilter{ModalityId: m.Id}).
		Return(users, nil).
		Times(1)

	service := services.ModalityService{
		UserRepository: uRepo,
	}

	// Act
	code := service.DeleteModalityById(&m.Id)

	// Assert
	if code != shared.InvalidOperation {
		t.Errorf("Expecting response code to be %s when trying to delete modality which is in use by users. Got %s", shared.InvalidOperation, code)
	}
}

func TestShouldReturnInvalidOperationWhenTryingToEditModalityInUseByUsers(t *testing.T) {
	// Arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	users := []*models.User{{Id: uuid.New()}}

	m := models.Modality{
		Id:          uuid.New(),
		Name:        "Test",
		Description: "Lorem Ipsum",
	}

	uRepo := mock_repository.NewMockIUserRepository(mockCtrl)
	uRepo.EXPECT().
		Get(&models.UserFilter{ModalityId: m.Id}).
		Return(users, nil).
		Times(1)

	service := services.ModalityService{
		UserRepository: uRepo,
	}

	// Act
	code := service.EditModality(&m.Id, &m)

	// Assert
	if code != shared.InvalidOperation {
		t.Errorf("Expecting response code to be %s when trying to edit modality which is in use by users. Got %s", shared.InvalidOperation, code)
	}
}
