package services

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/constants"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/messenger"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IModalityService interface {
	GetModalities(ctx context.Context, filter *models.ModalityFilter) ([]models.Modality, shared.Code)
	GetModalityById(ctx context.Context, uuid *uuid.UUID) (*models.Modality, shared.Code)
	CreateNewModality(ctx context.Context, m *models.Modality) shared.Code
	EditModality(ctx context.Context, uuid *uuid.UUID, m *models.Modality) shared.Code
	DeleteModalityById(ctx context.Context, uuid *uuid.UUID) shared.Code
}

type ModalityService struct {
	ModalityRepository repository.IModalityRepository
	UserRepository     repository.IUserRepository
}

func (ms *ModalityService) GetModalities(ctx context.Context, filter *models.ModalityFilter) ([]models.Modality, shared.Code) {
	result, err := ms.ModalityRepository.GetModalities(ctx, filter)

	if err != nil {
		return result, infraService.MapErrorFrom(err)
	}

	return result, shared.Success
}

func (ms *ModalityService) GetModalityById(ctx context.Context, uuid *uuid.UUID) (*models.Modality, shared.Code) {
	result, err := ms.ModalityRepository.GetModalityById(ctx, uuid)
	if err != nil {
		return nil, infraService.MapErrorFrom(err)
	}

	if result == nil {
		return nil, shared.NonExistentRecord
	}

	return result, shared.Success
}

func (ms *ModalityService) CreateNewModality(ctx context.Context, m *models.Modality) shared.Code {
	exists, err := ms.ModalityRepository.ExistsByName(ctx, &m.Id, &m.Name)
	if exists {
		return shared.DuplicatedRecord
	}

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	err = ms.ModalityRepository.CreateNewModality(ctx, m)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (ms *ModalityService) EditModality(ctx context.Context, uuid *uuid.UUID, m *models.Modality) shared.Code {
	if ms.isInUse(ctx, uuid) {
		return shared.InvalidOperation
	}

	exists, err := ms.ModalityRepository.ExistsByName(ctx, uuid, &m.Name)
	if exists {
		return shared.DuplicatedRecord
	}

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	err = ms.ModalityRepository.EditModality(ctx, uuid, m)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (ms *ModalityService) DeleteModalityById(ctx context.Context, uuid *uuid.UUID) shared.Code {
	if ms.isInUse(ctx, uuid) {
		return shared.InvalidOperation
	}

	isFound, err := ms.ModalityRepository.DeleteModalityById(ctx, uuid)

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
			ModalityId: *uuid,
		},
	)

	return shared.Success
}

func (ms *ModalityService) isInUse(ctx context.Context, uuid *uuid.UUID) bool {
	users, err := ms.UserRepository.Get(
		ctx,
		&models.UserFilter{
			ModalityId: *uuid,
		})

	if err != nil {
		log.Panic(err.Error())
	}

	return len(users) > 0
}
