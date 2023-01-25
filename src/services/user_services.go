package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/repository"
)

type IUserServices interface {
	GetAllUsers() ([]*models.User, error)
	GetUserById(uuid *uuid.UUID) (*models.User, error)
	CreateNewUser(u *models.User) (isDuplicated bool, err error)
	UpdateUser(uuid *uuid.UUID, u *models.User) (isFound bool, err error)
	DeleteUserById(uuid *uuid.UUID) (found bool, err error)
}

type UserServices struct {
	UserRepository repository.IUserRepository
}

func (us *UserServices) GetAllUsers() ([]*models.User, error) {
	return us.UserRepository.GetAllUsers()
}

func (us *UserServices) GetUserById(uuid *uuid.UUID) (*models.User, error) {
	return us.UserRepository.GetUserById(uuid)
}

func (us *UserServices) CreateNewUser(u *models.User) (isDuplicated bool, err error) {
	isDuplicated, err = us.UserRepository.ExistsByPhone(u.Phone)
	if err != nil || isDuplicated {
		return isDuplicated, err
	}

	return false, us.UserRepository.CreateNewUser(u)
}

func (us *UserServices) UpdateUser(uuid *uuid.UUID, u *models.User) (isFound bool, err error) {
	return us.UserRepository.UpdateUser(uuid, u)
}

func (us *UserServices) DeleteUserById(uuid *uuid.UUID) (isFound bool, err error) {
	return us.UserRepository.DeleteUserById(uuid)
}
