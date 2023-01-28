package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IModalityService interface {
	GetAllModalities() ([]*models.Modality, shared.Code)
	GetModalityById(uuid *uuid.UUID) (*models.Modality, shared.Code)
	CreateNewModality(m *models.Modality) shared.Code
	EditModality(uuid *uuid.UUID, m *models.Modality) shared.Code
	DeleteModalityById(uuid *uuid.UUID) shared.Code
}

type ModalityService struct {
	ModalityRepository *repository.ModalityRepository
}

func (ms *ModalityService) GetAllModalities() ([]*models.Modality, shared.Code) {
	result, err := ms.ModalityRepository.GetAllModalities()

	if err != nil {
		return result, infraService.MapErrorFrom(err)
	}

	return result, shared.Success
}

func (ms *ModalityService) GetModalityById(uuid *uuid.UUID) (*models.Modality, shared.Code) {
	result, err := ms.ModalityRepository.GetModalityById(uuid)
	if err != nil {
		return nil, infraService.MapErrorFrom(err)
	}

	if result == nil {
		return nil, shared.NonExistentRecord
	}

	return result, shared.Success
}

func (ms *ModalityService) CreateNewModality(m *models.Modality) shared.Code {
	err := ms.ModalityRepository.CreateNewModality(m)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (ms *ModalityService) EditModality(uuid *uuid.UUID, m *models.Modality) shared.Code {
	err := ms.ModalityRepository.EditModality(uuid, m)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (ms *ModalityService) DeleteModalityById(uuid *uuid.UUID) shared.Code {
	isFound, err := ms.ModalityRepository.DeleteSModalityById(uuid)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if !isFound {
		return shared.NonExistentRecord
	}

	return shared.Success
}
