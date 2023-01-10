package models

import "github.com/google/uuid"

type Modality struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
