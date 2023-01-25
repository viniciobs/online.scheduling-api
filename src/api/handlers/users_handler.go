package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	dto "github.com/online.scheduling-api/src/api/dtos"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
)

type UsersHandler struct {
	UserService services.IUserServices
}

func (uc *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, responseCode := uc.UserService.GetAllUsers()

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	var result []dto.UserCreateResponse

	for i := range users {
		result = append(result, dto.MapUserResponseFrom(users[i]))
	}

	helpers.JSONResponse(w, http.StatusOK, &result)
}

func (uc *UsersHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	user, responseCode := uc.UserService.GetUserById(&id)

	if responseCode == shared.NonExistentRecord {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, dto.MapUserResponseFrom(user))
}

func (uc *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestData dto.UserCreateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := requestData.Validate(); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	u := models.MapUserFrom(
		requestData.Name,
		requestData.Phone,
		requestData.Role,
		false)

	responseCode := uc.UserService.CreateNewUser(&u)
	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, dto.MapUserResponseFrom(&u))
}

func (uc *UsersHandler) Activate(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	responseCode := uc.UserService.ActivateUser(&id)

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

func (uc *UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	responseCode := uc.UserService.DeleteUserById(&id)

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
