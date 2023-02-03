package api

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	dto "github.com/online.scheduling-api/src/api/dtos/requests"
)

func ValidateSchedule(r *dto.ScheduleCreateOrUpdateRequest) error {
	errMsg := []string{}

	if r.ModalityId == uuid.Nil {
		errMsg = append(errMsg, "Informe a modalidade")
	}

	if r.UserId == uuid.Nil {
		errMsg = append(errMsg, "Informe de quem é a agenda")
	}

	for _, a := range r.Availability {
		if a.Time.Before(time.Now()) {
			errMsg = append(errMsg, fmt.Sprintf("A data %v não é válida", a.Time.Format("dd/MM/yyyy HH:mm")))
		}
	}

	if len(errMsg) > 0 {
		return errors.New(strings.Join(errMsg, ", "))
	}

	return nil
}
