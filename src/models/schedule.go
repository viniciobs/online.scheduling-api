package models

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	UserId       uuid.UUID       `bson:"user-id"`
	UserName     string          `bson:"user-name"`
	ModalityId   uuid.UUID       `bson:"modality-id"`
	ModalityName string          `bson:"modality-name"`
	Availability []*Availability `bson:"availability"`
}

type Availability struct {
	Time       time.Time `bson:"datetime" json:"datetime"`
	ReservedTo uuid.UUID `bson:"reserved-to" json:"reserved-to"`
	Payment    Payment   `bson:"payment"`
}
