package api

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
)

type ModalityCreateOrUpdateResponse struct {
	Id          uuid.UUID `string:"id"`
	Name        string    `string:"name"`
	Description string    `string:"description"`
}

func MapModalityResponseFrom(m *models.Modality) ModalityCreateOrUpdateResponse {
	return ModalityCreateOrUpdateResponse{
		Id:          m.Id,
		Name:        m.Name,
		Description: m.Description,
	}
}
