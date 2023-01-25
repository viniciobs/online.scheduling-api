package api

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
)

type UserCreateResponse struct {
	Id              uuid.UUID   `json:"id"`
	Name            string      `json:"name"`
	Phone           string      `json:"phone"`
	RoleCode        models.Role `json:"role-code"`
	RoleDescription string      `json:"role-description"`
	IsActive        bool        `json:"isActive"`
}

func MapUserResponseFrom(u *models.User) UserCreateResponse {
	return UserCreateResponse{
		Id:              u.Id,
		Name:            u.Name,
		Phone:           u.Phone,
		IsActive:        u.IsActive,
		RoleCode:        u.Role,
		RoleDescription: u.Role.GetDescription(),
	}
}
