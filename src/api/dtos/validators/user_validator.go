package api

import (
	"errors"
	"regexp"
	"strings"

	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	"github.com/online.scheduling-api/src/helpers"
)

func ValidateUser(user *dto.UserCreateOrUpdateRequest) error {
	var errMsg []string

	user.Name = helpers.TrimStartAndEnd(user.Name)
	if user.Name == "" {
		errMsg = append(errMsg, "Informe um nome válido")
	}

	user.Phone = helpers.TrimStartAndEnd(user.Phone)
	if !regexp.MustCompile(`^[\d]{11}$`).MatchString(user.Phone) {
		errMsg = append(errMsg, "Informe um celular válido")
	}

	if !user.Role.IsDefined() {
		errMsg = append(errMsg, "Informe um perfil válido")
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}
