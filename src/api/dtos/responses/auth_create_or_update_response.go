package api

import (
	"github.com/google/uuid"
	"github.com/online.scheduling-api/src/models"
)

type AuthCreateOrUpdateResponse struct {
	UserId          uuid.UUID   `json:"user-id"`
	Name            string      `json:"user-name"`
	RoleCode        models.Role `json:"role-code"`
	RoleDescription string      `json:"role-description"`
	IsActive        bool        `json:"isActive"`
	Token           string      `json:"auth-token"`
}

func MapAuthResponseFrom(u *models.User, token string) AuthCreateOrUpdateResponse {
	return AuthCreateOrUpdateResponse{
		UserId:          u.Id,
		Name:            u.Name,
		IsActive:        u.IsActive,
		RoleCode:        u.Role,
		RoleDescription: u.Role.GetDescription(),
		Token:           token,
	}
}
