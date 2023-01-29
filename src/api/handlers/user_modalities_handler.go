package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	dtoRequest "github.com/online.scheduling-api/src/api/dtos/requests"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
)

type UserModalitiesHandler struct {
	UserModalitiesService services.IUserModalitiesService
}

func (h *UserModalitiesHandler) Edit(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	var requestData dtoRequest.UserModalitiesRequest
	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	responseCode := h.UserModalitiesService.Edit(&userId, requestData.Modalities)

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, nil)
}
