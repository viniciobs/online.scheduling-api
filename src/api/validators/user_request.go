package api

import (
	"errors"
	"strings"

	"github.com/online.scheduling-api/src/business/models"
)

func Validate(u *models.User) error {
	var errMsg []string

	if strings.Trim(u.Name, " ") == "" {
		errMsg = append(errMsg, "Informe um nome válido")
	}

	if strings.Trim(u.Phone, " ") == "" {
		errMsg = append(errMsg, "Informe um celular válido")
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}
