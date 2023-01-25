package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IUserServices interface {
	GetAllUsers() ([]*models.User, shared.Code)
	GetUserById(uuid *uuid.UUID) (*models.User, shared.Code)
	CreateNewUser(u *models.User) shared.Code
	UpdateUser(uuid *uuid.UUID, u *models.User) (isFound bool, err error)
	DeleteUserById(uuid *uuid.UUID) shared.Code
}

type UserServices struct {
	UserRepository repository.IUserRepository
}

func (us *UserServices) GetAllUsers() ([]*models.User, shared.Code) {
	result, err := us.UserRepository.GetAllUsers()

	if err != nil {
		return result, infraService.MapErrorFrom(err)
	}

	return result, shared.Success
}

func (us *UserServices) GetUserById(uuid *uuid.UUID) (*models.User, shared.Code) {
	result, err := us.UserRepository.GetUserById(uuid)
	if err != nil {
		return nil, infraService.MapErrorFrom(err)
	}

	if result == nil {
		return nil, shared.NonExistentRecord
	}

	return result, shared.Success
}

func (us *UserServices) CreateNewUser(u *models.User) shared.Code {
	exists, err := us.UserRepository.ExistsByPhone(u.Phone)
	if exists {
		return shared.DuplicatedRecord
	}

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	err = us.UserRepository.CreateNewUser(u)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (us *UserServices) UpdateUser(uuid *uuid.UUID, u *models.User) (isFound bool, err error) {
	return us.UserRepository.UpdateUser(uuid, u)
}

func (us *UserServices) DeleteUserById(uuid *uuid.UUID) shared.Code {
	isFound, err := us.UserRepository.DeleteUserById(uuid)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if !isFound {
		return shared.NonExistentRecord
	}

	return shared.Success
}
