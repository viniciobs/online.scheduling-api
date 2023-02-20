package services

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/infra/repository"
	infraService "github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/shared"
)

type IScheduleService interface {
	Get(ctx context.Context, filter *models.ScheduleFilter) ([]*models.Schedule, shared.Code)
	Create(ctx context.Context, user, modality *uuid.UUID, availability []*models.Availability) shared.Code
	Edit(ctx context.Context, user, modality *uuid.UUID, availability []*models.Availability) shared.Code
	DeleteBy(ctx context.Context, userId, modalityId *uuid.UUID) shared.Code
}

type ScheduleService struct {
	ScheduleRepository repository.IScheduleRepository
	ModalityRepository repository.IModalityRepository
	UserRespository    repository.IUserRepository
}

func (s *ScheduleService) Get(ctx context.Context, filter *models.ScheduleFilter) ([]*models.Schedule, shared.Code) {
	result, err := s.ScheduleRepository.Get(ctx, filter)

	if err != nil {
		return result, infraService.MapErrorFrom(err)
	}

	return result, shared.Success
}

func (s *ScheduleService) Create(ctx context.Context, userId, modalityId *uuid.UUID, availability []*models.Availability) shared.Code {
	exists, err := s.exists(ctx, userId, modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if exists {
		return shared.DuplicatedRecord
	}

	var (
		wg       sync.WaitGroup
		modality *models.Modality
		user     *models.User
	)

	wg.Add(2)

	go func() {
		modality, err = s.ModalityRepository.GetModalityById(ctx, modalityId)
		wg.Done()
	}()

	go func() {
		user, err = s.UserRespository.GetUserById(ctx, userId)
		wg.Done()
	}()

	wg.Wait()

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if modality == nil || user == nil {
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

	if err := s.ScheduleRepository.Create(ctx, &schedule); err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (s *ScheduleService) Edit(ctx context.Context, userId, modalityId *uuid.UUID, availability []*models.Availability) shared.Code {
	exists, err := s.exists(ctx, userId, modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if !exists {
		return shared.NonExistentRecord
	}

	schedules, err := s.ScheduleRepository.Get(
		ctx,
		&models.ScheduleFilter{
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

	if err := s.ScheduleRepository.Edit(ctx, schedule); err != nil {
		return infraService.MapErrorFrom(err)
	}

	return shared.Success
}

func (s *ScheduleService) DeleteBy(ctx context.Context, userId, modalityId *uuid.UUID) shared.Code {
	isFound, err := s.ScheduleRepository.DeleteBy(ctx, userId, modalityId)

	if err != nil {
		return infraService.MapErrorFrom(err)
	}

	if !isFound {
		return shared.NonExistentRecord
	}

	return shared.Success
}

func (s *ScheduleService) exists(ctx context.Context, user, schedule *uuid.UUID) (bool, error) {
	filter := models.ScheduleFilter{
		UserId:     *user,
		ModalityId: *schedule,
	}

	schedules, err := s.ScheduleRepository.Get(ctx, &filter)

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
