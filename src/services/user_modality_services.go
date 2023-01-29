package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IUserModalitiesService interface {
	Edit(userId *uuid.UUID, modalitiesIds []uuid.UUID) shared.Code
}

type UserModalitiesService struct {
	UserRepository     repository.IUserRepository
	ModalityRepository repository.IModalityRepository
}

func (s *UserModalitiesService) Edit(userId *uuid.UUID, modalitiesIds []uuid.UUID) shared.Code {
	user, responseCode := s.getUser(userId)
	if responseCode != shared.Success {
		return responseCode
	}

	if !user.IsActive {
		return shared.InvalidOperation
	}

	modalities, responseCode := s.getModalities(modalitiesIds)
	if responseCode != shared.Success {
		return responseCode
	}

	// TODO: the above code could be executed at same time

	user.Modalities = modalities

	err := s.UserRepository.EditUserModalities(userId, user)
	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (s *UserModalitiesService) getUser(userId *uuid.UUID) (*models.User, shared.Code) {
	user, err := s.UserRepository.GetUserById(userId)
	if err != nil {
		return nil, infraService.MapErrorFrom(err)
	}

	if user == nil {
		return nil, shared.NonExistentRecord
	}

	return user, shared.Success
}

func (s *UserModalitiesService) getModalities(ids []uuid.UUID) ([]models.Modality, shared.Code) {
	filter := models.ModalityFilter{
		Ids: ids,
	}

	modalities, err := s.ModalityRepository.GetModalities(&filter)
	if err != nil {
		return nil, infraService.MapErrorFrom(err)
	}

	if len(modalities) == 0 {
		return nil, shared.NonExistentRecord
	}

	return modalities, shared.Success
}
