package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	dtoRequest "github.com/online.scheduling-api/src/api/dtos/requests"
	dtoResponse "github.com/online.scheduling-api/src/api/dtos/responses"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
)

type ModalityHandler struct {
	ModalityService services.IModalityService
}

func (h *ModalityHandler) Get(w http.ResponseWriter, r *http.Request) {
	filter := models.ModalityFilter{}

	name := r.URL.Query().Get("name")
	if name != "" {
		filter.Name = name
	}

	modalities, responseCode := h.ModalityService.GetModalities(r.Context(), &filter)

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	var result []dtoResponse.ModalityCreateOrUpdateResponse

	for i := range modalities {
		result = append(result, dtoResponse.MapModalityResponseFrom(&modalities[i]))
	}

	helpers.JSONResponse(w, http.StatusOK, &result)
}

func (h *ModalityHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	modality, responseCode := h.ModalityService.GetModalityById(r.Context(), &id)

	if responseCode == shared.NonExistentRecord {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, dtoResponse.MapModalityResponseFrom(modality))
}

func (h *ModalityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestData dtoRequest.ModalityCreateOrUpdateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateModality(&requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	m := models.MapModalityFrom(
		requestData.Name,
		requestData.Description)

	responseCode := h.ModalityService.CreateNewModality(r.Context(), &m)
	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, dtoResponse.MapModalityResponseFrom(&m))
}

func (h *ModalityHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	responseCode := h.ModalityService.DeleteModalityById(r.Context(), &id)

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

func (h *ModalityHandler) Edit(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}
	var requestData dtoRequest.ModalityCreateOrUpdateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateModality(&requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	m := models.MapModalityFrom(
		requestData.Name,
		requestData.Description)

	responseCode := h.ModalityService.EditModality(r.Context(), &id, &m)

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
