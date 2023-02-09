package api

import "github.com/google/uuid"

type ScheduleDeleteRequest struct {
	UserId     *uuid.UUID `json:"user-id"`
	ModalityId *uuid.UUID `json:"modality-id"`
}
