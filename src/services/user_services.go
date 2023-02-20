package services

import (
	"context"
	"sync"

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
	Get(ctx context.Context, filter *models.UserFilter) ([]*models.User, shared.Code)
	GetUserById(ctx context.Context, uuid *uuid.UUID) (*models.User, shared.Code)
	CreateNewUser(ctx context.Context, u *models.User) (code shared.Code, token string)
	ActivateUser(ctx context.Context, uuid *uuid.UUID) shared.Code
	EditUser(ctx context.Context, uuid *uuid.UUID, u *models.User) shared.Code
	EditModalities(ctx context.Context, userId *uuid.UUID, modalitiesIds []uuid.UUID) shared.Code
	DeleteUserById(ctx context.Context, uuid *uuid.UUID) shared.Code
}

type UserServices struct {
	UserRepository     repository.IUserRepository
	ModalityRepository repository.IModalityRepository
}

func (us *UserServices) Get(ctx context.Context, filter *models.UserFilter) ([]*models.User, shared.Code) {
	result, err := us.UserRepository.Get(ctx, filter)

	if err != nil {
		return result, infraService.MapErrorFrom(err)
	}

	return result, shared.Success
}

func (us *UserServices) GetUserById(ctx context.Context, uuid *uuid.UUID) (*models.User, shared.Code) {
	result, err := us.UserRepository.GetUserById(ctx, uuid)
	if err != nil {
		return nil, infraService.MapErrorFrom(err)
	}

	if result == nil {
		return nil, shared.NonExistentRecord
	}

	return result, shared.Success
}

func (us *UserServices) CreateNewUser(ctx context.Context, u *models.User) (code shared.Code, token string) {
	exists, err := us.UserRepository.ExistsBy(ctx, &u.Id, &u.Phone, &u.Login)
	if err != nil {
		return infraService.MapErrorFrom(err), token
	}
	if exists {
		return shared.DuplicatedRecord, token
	}

	u.Passphrase = helpers.Crypt(u.Passphrase)

	if err = us.UserRepository.CreateNewUser(ctx, u); err != nil {
		return infraService.MapErrorFrom(err), token
	}

	claims := models.MapUserClaimsFrom(u)
	if token, err = helpers.CreateTokenFor(claims); err != nil {
		return infraService.MapErrorFrom(err), token
	}

	return shared.Success, token
}

func (us *UserServices) ActivateUser(ctx context.Context, uuid *uuid.UUID) shared.Code {
	err := us.UserRepository.ActivateUser(ctx, uuid)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (us *UserServices) EditUser(ctx context.Context, uuid *uuid.UUID, u *models.User) shared.Code {
	exists, err := us.UserRepository.ExistsBy(ctx, uuid, &u.Phone, &u.Login)
	if exists {
		return shared.DuplicatedRecord
	}

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	err = us.UserRepository.EditUser(ctx, uuid, u)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (us *UserServices) EditModalities(ctx context.Context, userId *uuid.UUID, modalitiesIds []uuid.UUID) shared.Code {
	var (
		wg         sync.WaitGroup
		user       *models.User
		modalities []models.Modality
		err        error
		wgResult   shared.Code
	)

	wg.Add(2)

	wgResult = shared.Success

	go func() {
		user, err = us.UserRepository.GetUserById(ctx, userId)
		if err != nil {
			wgResult = shared.NonExistentRecord
			wg.Done()
		}

		if !user.IsActive {
			wgResult = shared.InvalidOperation
		}

		wg.Done()
	}()

	go func() {
		modalities, err = us.ModalityRepository.GetModalities(
			ctx,
			&models.ModalityFilter{
				Ids: modalitiesIds,
			},
		)

		if err != nil {
			wgResult = infraService.MapErrorFrom(err)
		}

		wg.Done()
	}()

	wg.Wait()

	if wgResult != shared.Success {
		return wgResult
	}

	// TODO: the above code could be executed at same time

	user.Modalities = modalities

	err = us.UserRepository.EditUserModalities(ctx, userId, user)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (us *UserServices) DeleteUserById(ctx context.Context, uuid *uuid.UUID) shared.Code {
	isFound, err := us.UserRepository.DeleteUserById(ctx, uuid)

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
