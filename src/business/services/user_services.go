package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/business/models"
	"github.com/online.scheduling-api/src/infra/repository"
)

func GetAllUsers() (*[]models.User, error) {
	return repository.GetAllUsers()
}

func GetUserById(uuid *uuid.UUID) (*models.User, error) {
	return repository.GetUserById(uuid)
}

func CreateNewUser(u *models.User) (isDuplicated bool, err error) {
	isDuplicated, err = repository.ExistsByPhone(u.Phone)
	if err != nil || isDuplicated {
		return isDuplicated, err
	}

	return false, repository.CreateNewUser(u)
}

func UpdateUser(u *models.User) error {
	return repository.UpdateUser(u)
}

func DeleteUserById(uuid *uuid.UUID) (found bool, err error) {
	user, err := repository.GetUserById(uuid)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}

	return true, repository.DeleteUserById(uuid)
}
