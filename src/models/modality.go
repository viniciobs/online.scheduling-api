package models

import "github.com/google/uuid"

type Modality struct {
	Id          uuid.UUID `bson:"id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
}

func MapModalityFrom(name, description string) Modality {
	return Modality{
		Id:          uuid.New(),
		Name:        name,
		Description: description,
	}
}
