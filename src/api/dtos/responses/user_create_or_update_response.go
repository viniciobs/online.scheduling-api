package api

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
)

type UserCreateOrUpdateResponse struct {
	Id              uuid.UUID         `json:"id"`
	Name            string            `json:"name"`
	Phone           string            `json:"phone"`
	RoleCode        models.Role       `json:"role-code"`
	RoleDescription string            `json:"role-description"`
	IsActive        bool              `json:"isActive"`
	Modalities      []models.Modality `json:"modalities,omitempty"`
	Token           string            `json:"token,omitempty"`
}

func MapUserResponseFrom(u *models.User) UserCreateOrUpdateResponse {
	modalities := []models.Modality{}

	if len(u.Modalities) > 0 {
		modalities = u.Modalities
	}

	return UserCreateOrUpdateResponse{
		Id:              u.Id,
		Name:            u.Name,
		Phone:           u.Phone,
		IsActive:        u.IsActive,
		RoleCode:        u.Role,
		RoleDescription: u.Role.GetDescription(),
		Modalities:      modalities,
	}
}

func MapNewUserResponseFrom(u *models.User, token string) UserCreateOrUpdateResponse {
	return UserCreateOrUpdateResponse{
		Id:              u.Id,
		Name:            u.Name,
		Phone:           u.Phone,
		IsActive:        u.IsActive,
		RoleCode:        u.Role,
		RoleDescription: u.Role.GetDescription(),
		Token:           token,
	}
}
