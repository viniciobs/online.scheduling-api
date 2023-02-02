package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IUserServices interface {
	Get(*models.UserFilter) ([]*models.User, shared.Code)
	GetUserById(uuid *uuid.UUID) (*models.User, shared.Code)
	CreateNewUser(u *models.User) shared.Code
	ActivateUser(uuid *uuid.UUID) shared.Code
	EditUser(uuid *uuid.UUID, u *models.User) shared.Code
	DeleteUserById(uuid *uuid.UUID) shared.Code
}

type UserServices struct {
	UserRepository repository.IUserRepository
}

func (us *UserServices) Get(filter *models.UserFilter) ([]*models.User, shared.Code) {
	result, err := us.UserRepository.Get(filter)

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
	exists, err := us.UserRepository.ExistsByPhone(&u.Id, &u.Phone)
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

func (us *UserServices) ActivateUser(uuid *uuid.UUID) shared.Code {
	err := us.UserRepository.ActivateUser(uuid)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (us *UserServices) EditUser(uuid *uuid.UUID, u *models.User) shared.Code {
	exists, err := us.UserRepository.ExistsByPhone(uuid, &u.Phone)
	if exists {
		return shared.DuplicatedRecord
	}

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	err = us.UserRepository.EditUser(uuid, u)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
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
