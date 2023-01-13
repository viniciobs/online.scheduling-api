package models

import (
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Role     Role      `json:"role"`
	IsActive bool      `json:"isActive"`
}

func (u *User) RemoveWhiteSpaces() {
	u.Name = strings.Trim(u.Name, " ")
	u.Phone = strings.Trim(u.Phone, " ")
}
