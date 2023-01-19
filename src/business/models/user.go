package models

import (
	"errors"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" bson:"id"`
	Name     string    `json:"name" bson:"name"`
	Phone    string    `json:"phone" bson:"phone"`
	Role     Role      `json:"role" bson:"role"`
	IsActive bool      `json:"isActive" bson:"isActive"`
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
