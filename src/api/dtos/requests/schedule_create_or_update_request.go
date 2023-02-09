package api

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
)

type ScheduleCreateOrUpdateRequest struct {
	ModalityId   uuid.UUID              `json:"modality"`
	UserId       uuid.UUID              `json:"user"`
	Availability []*models.Availability `json:"availability"`
}
