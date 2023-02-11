package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/constants"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/messenger"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IUserServices interface {
	Get(*models.UserFilter) ([]*models.User, shared.Code)
	GetUserById(uuid *uuid.UUID) (*models.User, shared.Code)
	CreateNewUser(u *models.User) (code shared.Code, token string)
	ActivateUser(uuid *uuid.UUID) shared.Code
	EditUser(uuid *uuid.UUID, u *models.User) shared.Code
	EditModalities(userId *uuid.UUID, modalitiesIds []uuid.UUID) shared.Code
	DeleteUserById(uuid *uuid.UUID) shared.Code
}

type UserServices struct {
	UserRepository     repository.IUserRepository
	ModalityRepository repository.IModalityRepository
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

func (us *UserServices) CreateNewUser(u *models.User) (code shared.Code, token string) {
	exists, err := us.UserRepository.ExistsBy(&u.Id, &u.Phone, &u.Login)
	if err != nil {
		return infraService.MapErrorFrom(err), token
	}
	if exists {
		return shared.DuplicatedRecord, token
	}

	u.Passphrase = helpers.Crypt(u.Passphrase)

	if err = us.UserRepository.CreateNewUser(u); err != nil {
		return infraService.MapErrorFrom(err), token
	}

	claims := models.MapUserClaimsFrom(u)
	if token, err = helpers.CreateTokenFor(claims); err != nil {
		return infraService.MapErrorFrom(err), token
	}

	return shared.Success, token
}

func (us *UserServices) ActivateUser(uuid *uuid.UUID) shared.Code {
	err := us.UserRepository.ActivateUser(uuid)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (us *UserServices) EditUser(uuid *uuid.UUID, u *models.User) shared.Code {
	exists, err := us.UserRepository.ExistsBy(uuid, &u.Phone, &u.Login)
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

func (us *UserServices) EditModalities(userId *uuid.UUID, modalitiesIds []uuid.UUID) shared.Code {
	user, err := us.UserRepository.GetUserById(userId)
	if err != nil {
		return shared.NonExistentRecord
	}

	if !user.IsActive {
		return shared.InvalidOperation
	}

	modalities, err := us.ModalityRepository.GetModalities(&models.ModalityFilter{Ids: modalitiesIds})
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	// TODO: the above code could be executed at same time

	user.Modalities = modalities

	err = us.UserRepository.EditUserModalities(userId, user)
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

	go messenger.Produce(
		context.TODO(),
		constants.DELETED_OBJECTS_TOPIC,
		messenger.DeleteObjects{
			UserId: *uuid,
		},
	)

	return shared.Success
}
