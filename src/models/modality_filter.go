package models

import "github.com/google/uuid"

type ModalityFilter struct {
	Ids  []uuid.UUID
	Name string
}
