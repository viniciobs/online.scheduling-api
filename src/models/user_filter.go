package models

import "github.com/google/uuid"

type UserFilter struct {
	UserId       uuid.UUID
	UserName     string
	ModalityId   uuid.UUID
	ModalityName string
}
