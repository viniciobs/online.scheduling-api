package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/online.scheduling-api/constants"
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

	token, err := helpers.RetrieveToken(r)
	if err != nil {
		helpers.JSONResponseError(w, http.StatusUnauthorized, errors.New("missing auth"))
	}

	claims := token.Claims.(jwt.MapClaims)
	isCustomer := claims[constants.CLAIM_USER_ROLE].(float64) == float64(models.Customer)

	if isCustomer {
		filter.UserId, _ = uuid.Parse(claims[constants.CLAIM_USER_ID].(string))
	}

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

	result, responseCode := h.ScheduleService.Get(r.Context(), &filter)

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
		r.Context(),
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
		r.Context(),
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

func (h *SchedulesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	var requestData dtoRequest.ScheduleDeleteRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if *requestData.ModalityId == uuid.Nil || *requestData.UserId == uuid.Nil {
		helpers.JSONResponse(w, http.StatusBadRequest, nil)
		return
	}

	responseCode := h.ScheduleService.DeleteBy(r.Context(), requestData.UserId, requestData.ModalityId)

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
