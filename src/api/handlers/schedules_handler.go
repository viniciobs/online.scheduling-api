package api

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	dtoRequest "github.com/online.scheduling-api/src/api/dtos/requests"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
)

type SchedulesHandler struct {
	ScheduleService services.IScheduleService
}

func (h *SchedulesHandler) Get(w http.ResponseWriter, r *http.Request) {
	filter := models.ScheduleFilter{}

	if name := r.URL.Query().Get("user"); name != "" {
		filter.UserName = name
	}

	if modality := r.URL.Query().Get("modality"); modality != "" {
		filter.ModalityName = modality
	}

	if available, err := strconv.ParseBool(r.URL.Query().Get("available")); err == nil {
		filter.Available = available
	}

	if reservedTo, err := uuid.Parse(r.URL.Query().Get("reserved-to")); err == nil {
		filter.ReservedTo = reservedTo
	}

	result, responseCode := h.ScheduleService.Get(&filter)

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, &result)
}

func (h *SchedulesHandler) Create(w http.ResponseWriter, r *http.Request) {
	requestData, err := getRequestData(r)

	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	responseCode := h.ScheduleService.Create(
		&requestData.UserId,
		&requestData.ModalityId,
		requestData.Availability,
	)

	if responseCode == shared.NonExistentRecord {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, nil)
}

func (h *SchedulesHandler) Edit(w http.ResponseWriter, r *http.Request) {
	requestData, err := getRequestData(r)

	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	responseCode := h.ScheduleService.Edit(
		&requestData.UserId,
		&requestData.ModalityId,
		requestData.Availability,
	)

	if responseCode == shared.NonExistentRecord {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, nil)
}

func getRequestData(r *http.Request) (dtoRequest.ScheduleCreateOrUpdateRequest, error) {
	var requestData dtoRequest.ScheduleCreateOrUpdateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		return requestData, err
	}

	if err := validator.ValidateSchedule(&requestData); err != nil {
		return requestData, err
	}

	return requestData, nil
}
