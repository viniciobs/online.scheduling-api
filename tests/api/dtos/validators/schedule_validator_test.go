package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	dto "github.com/online.scheduling-api/src/api/dtos/requests"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
	"github.com/online.scheduling-api/src/models"
)

var schedule = dto.ScheduleCreateOrUpdateRequest{
	ModalityId: uuid.New(),
	UserId:     uuid.New(),
	Availability: []*models.Availability{
		{
			Time: time.Now().AddDate(0, 0, 1),
		},
	},
}

func TestShouldReturnrErrorWhenModalityIdIsNil(t *testing.T) {
	// Arrange
	data := schedule
	data.ModalityId = uuid.Nil

	// Act
	err := validator.ValidateSchedule(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when modality is not set")
	}
}

func TestShouldReturnrErrorWhenUserIdIsNil(t *testing.T) {
	// Arrange
	data := schedule
	data.UserId = uuid.Nil

	// Act
	err := validator.ValidateSchedule(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when user is not set")
	}
}

func TestShouldReturnrErrorWhenSomeAvailabilityHasInvalidDateSetted(t *testing.T) {
	// Arrange
	data := schedule
	data.Availability = []*models.Availability{
		{
			Time: time.Now().AddDate(0, 0, -1),
		},
	}

	// Act
	err := validator.ValidateSchedule(&data)

	// Assert
	if err == nil {
		t.Error("Expected validation to fail when availability setted earlier date")
	}
}
