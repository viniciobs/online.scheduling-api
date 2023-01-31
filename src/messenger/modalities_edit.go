package messenger

import "github.com/google/uuid"

type Action string

const (
	Delete Action = "Delete"
	Update Action = "Update"
)

type ModalitiesEdit struct {
	ModalityId uuid.UUID `json:"modality-id"`
	Action     Action    `json:"action"`
}
