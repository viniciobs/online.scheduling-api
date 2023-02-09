package models

import "github.com/google/uuid"

type ScheduleFilter struct {
	ModalityId   uuid.UUID
	ModalityName string
	UserId       uuid.UUID
	UserName     string
	Available    bool
	ReservedTo   uuid.UUID
}
