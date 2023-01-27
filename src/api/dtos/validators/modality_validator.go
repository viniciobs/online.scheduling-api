package api

import (
	"errors"
	"fmt"
	"strings"

	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	"github.com/online.scheduling-api/src/helpers"
)

const DescriptionMaxLength = 140

func ValidateModality(modality *dto.ModalityCreateOrUpdateRequest) error {
	var errMsg []string

	modality.Name = helpers.TrimStartAndEnd(modality.Name)
	if modality.Name == "" {
		errMsg = append(errMsg, "Informe um nome válido para a modalidade")
	}

	modality.Description = helpers.TrimStartAndEnd(modality.Description)
	if len(modality.Description) > 140 {
		errMsg = append(errMsg, fmt.Sprintf("A descrição da modalidade deve ter até %d caracteres", DescriptionMaxLength))
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}
