package models

import "github.com/google/uuid"

type UserFilter struct {
	Name         string
	ModalityId   uuid.UUID
	ModalityName string
}
