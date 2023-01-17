package models

import (
	"errors"
	"regexp"
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

func (u *User) removeWhiteSpaces() {
	u.Name = strings.Trim(u.Name, " ")
	u.Phone = strings.Trim(u.Phone, " ")
}

func (u *User) Validate() error {
	var errMsg []string

	u.removeWhiteSpaces()

	if u.Name == "" {
		errMsg = append(errMsg, "Informe um nome válido")
	}

	if !regexp.MustCompile(`^[\d]{11}$`).MatchString(u.Phone) {
		errMsg = append(errMsg, "Informe um celular válido")
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}
