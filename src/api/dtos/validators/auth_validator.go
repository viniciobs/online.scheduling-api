package api

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	"github.com/online.scheduling-api/src/helpers"
)

const passphraseMinLength = 6

func ValidateAuth(data *dto.AuthCreateOrUpdateRequest) error {
	var errMsg []string

	data.Login = helpers.TrimStartAndEnd(data.Login)
	if data.Login == "" {
		errMsg = append(errMsg, "Informe um login válido")
	}

	data.Passphrase = helpers.TrimStartAndEnd(data.Passphrase)
	if !isValidPassphrase(data.Passphrase) {
		errMsg = append(errMsg, fmt.Sprintf("Informe uma senha válida! A senha deve conter %d caracteres incluindo letras maiúsculas, números e caracteres especiais", passphraseMinLength))
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}

func isValidPassphrase(s string) bool {
	if len(s) < 6 {
		return false
	}

	hasUpperCase := false
	hasSpecialChar := false
	hasNumber := false

	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecialChar = true
		}
	}

	return hasUpperCase && hasSpecialChar && hasNumber
}
