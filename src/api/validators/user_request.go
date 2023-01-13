package api

import (
	"errors"
	"regexp"
	"strings"

	"github.com/online.scheduling-api/src/business/models"
)

func Validate(u *models.User) error {
	var errMsg []string

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
