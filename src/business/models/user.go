package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Role     Role      `json:"role"`
	IsActive bool      `json:"isActive"`
}
