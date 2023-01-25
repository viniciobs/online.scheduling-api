package api

import (
	"errors"
	"regexp"
	"strings"

	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
)

type UserCreateRequest struct {
	Name  string      `json:"name"`
	Phone string      `json:"phone"`
	Role  models.Role `json:"role"`
}

func (u *UserCreateRequest) Validate() error {
	var errMsg []string

	u.Name = helpers.TrimStartAndEnd(u.Name)
	if u.Name == "" {
		errMsg = append(errMsg, "Informe um nome válido")
	}

	u.Phone = helpers.TrimStartAndEnd(u.Phone)
	if !regexp.MustCompile(`^[\d]{11}$`).MatchString(u.Phone) {
		errMsg = append(errMsg, "Informe um celular válido")
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}
