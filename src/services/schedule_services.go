package services

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IScheduleService interface {
	Get(*models.ScheduleFilter) ([]*models.Schedule, shared.Code)
	Create(user, modality *uuid.UUID, availability []*models.Availability) shared.Code
	Edit(user, modality *uuid.UUID, availability []*models.Availability) shared.Code
	DeleteBy(userId, modalityId *uuid.UUID) shared.Code
}

type ScheduleService struct {
	ScheduleRepository repository.IScheduleRepository
	ModalityRepository repository.IModalityRepository
	UserRespository    repository.IUserRepository
}

func (s *ScheduleService) Get(filter *models.ScheduleFilter) ([]*models.Schedule, shared.Code) {
	result, err := s.ScheduleRepository.Get(filter)

	if err != nil {
		return result, infraService.MapErrorFrom(err)
	}

	return result, shared.Success
}

func (s *ScheduleService) Create(userId, modalityId *uuid.UUID, availability []*models.Availability) shared.Code {
	exists, err := s.exists(userId, modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if exists {
		return shared.DuplicatedRecord
	}

	modality, err := s.ModalityRepository.GetModalityById(modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if modality == nil {
		return shared.NonExistentRecord
	}

	user, err := s.UserRespository.GetUserById(userId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if user == nil {
		return shared.NonExistentRecord
	}

	if !userHasModality(user, &modality.Id) {
		return shared.InvalidOperation
	}

	schedule := models.Schedule{
		UserId:       user.Id,
		UserName:     user.Name,
		ModalityId:   modality.Id,
		ModalityName: modality.Name,
		Availability: availability,
	}

	if err := s.ScheduleRepository.Create(&schedule); err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (s *ScheduleService) Edit(userId, modalityId *uuid.UUID, availability []*models.Availability) shared.Code {
	exists, err := s.exists(userId, modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if !exists {
		return shared.NonExistentRecord
	}

	schedules, err := s.ScheduleRepository.Get(&models.ScheduleFilter{
		ModalityId: *modalityId,
		UserId:     *userId,
	})

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if len(schedules) == 0 {
		return shared.NonExistentRecord
	}

	if len(schedules) != 1 {
		return shared.Unkown
	}

	schedule := schedules[0]
	schedule.Availability = availability

	if err := s.ScheduleRepository.Edit(schedule); err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (s *ScheduleService) DeleteBy(userId, modalityId *uuid.UUID) shared.Code {
	isFound, err := s.ScheduleRepository.DeleteBy(userId, modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if !isFound {
		return shared.NonExistentRecord
	}

	return shared.Success
}

func (s *ScheduleService) exists(user, schedule *uuid.UUID) (bool, error) {
	filter := models.ScheduleFilter{
		UserId:     *user,
		ModalityId: *schedule,
	}

	schedules, err := s.ScheduleRepository.Get(&filter)

	if err != nil {
		return false, err
	}

	if len(schedules) > 0 {
		return true, nil
	}

	return false, nil
}

func userHasModality(u *models.User, modalityId *uuid.UUID) bool {
	for _, m := range u.Modalities {
		if m.Id == *modalityId {
			return true
		}
	}

	return false
}
