package messenger

import "github.com/google/uuid"

type DeleteObjects struct {
	UserId     uuid.UUID
	ModalityId uuid.UUID
}
