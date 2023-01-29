package api

import (
	"github.com/google/uuid"
)

type UserModalitiesRequest struct {
	Modalities []uuid.UUID `json:"modalities-ids"`
}
