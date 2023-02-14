package api

import (
	"github.com/online.scheduling-api/src/models"
)

type UserCreateOrUpdateRequest struct {
	Name       string      `json:"name"`
	Phone      string      `json:"phone"`
	Role       models.Role `json:"role"`
	Login      string      `json:"login"`
	Passphrase string      `json:"passphrase"`
}
