package handlers

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
)

type UsersHandler struct {
	UserService services.IUserServices
}

func (uc *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserService.GetAllUsers()

	if err != nil {
		helpers.JSONResponseError(w, http.StatusServiceUnavailable, err)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, &users)
}

func (uc *UsersHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	user, err := uc.UserService.GetUserById(&id)
	if err != nil {
		helpers.JSONResponseError(w, http.StatusServiceUnavailable, err)
		return
	}
	if user == nil {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, user)
}

func (uc *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := helpers.ReadJSONBody(r, &user); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate(); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	isDuplicated, err := uc.UserService.CreateNewUser(&user)
	if err != nil {
		helpers.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}
	if isDuplicated {
		helpers.JSONResponse(w, http.StatusConflict, errors.New("usuário já cadastrado"))
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, &user)
}

func (uc *UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	var user models.User

	if err := helpers.ReadJSONBody(r, &user); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate(); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	isFound, err := uc.UserService.UpdateUser(&id, &user)

	if err != nil {
		helpers.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if !isFound {
		helpers.JSONResponseError(w, http.StatusNotFound, errors.New("usuário não encontrado"))
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

	found, err := uc.UserService.DeleteUserById(&id)
	if err != nil {
		helpers.JSONResponseError(w, http.StatusServiceUnavailable, err)
		return
	}
	if !found {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, nil)
}
