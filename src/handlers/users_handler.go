package handlers

import (
	"net/http"

	"github.com/online.scheduling-api/src/business/models"
	"github.com/online.scheduling-api/src/business/services"
	"github.com/online.scheduling-api/src/helpers"
)

type UsersHandler struct {
	UserService services.IUserServices
}

func (uc *UsersHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserService.GetAllUsers()

	if err != nil {
		helpers.JSONResponse(w, http.StatusServiceUnavailable, helpers.NewError(err.Error()))
		return
	}

	helpers.JSONResponse(w, http.StatusOK, &users)
}

func (uc *UsersHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetUUID(r)
	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	user, err := uc.UserService.GetUserById(id)
	if err != nil {
		helpers.JSONResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	if user == nil {
		helpers.JSONResponse(w, http.StatusNotFound, &user)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, user)
}

func (uc *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := helpers.ReadJSONBody(r, &user); err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, helpers.NewError(err.Error()))
		return
	}

	if err := user.Validate(); err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	isDuplicated, err := uc.UserService.CreateNewUser(&user)
	if err != nil {
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}
	if isDuplicated {
		helpers.JSONResponse(w, http.StatusConflict, helpers.NewError("Usuário já cadastrado"))
		return
	}

	helpers.JSONResponse(w, http.StatusCreated, &user)
}

func (uc *UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := helpers.ReadJSONBody(r, &user); err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Validate(); err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	isFound, err := uc.UserService.UpdateUser(&user)

	if err != nil {
		helpers.JSONResponse(w, http.StatusUnprocessableEntity, err)
		return
	}

	if !isFound {
		helpers.JSONResponse(w, http.StatusNotFound, helpers.NewError("Usuário não encontrado"))
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, nil)
}

func (uc *UsersHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.GetUUID(r)
	if err != nil {
		helpers.JSONResponse(w, http.StatusBadRequest, err)
		return
	}

	found, err := uc.UserService.DeleteUserById(id)
	if err != nil {
		helpers.JSONResponse(w, http.StatusServiceUnavailable, err)
		return
	}
	if !found {
		helpers.JSONResponse(w, http.StatusNotFound, nil)
		return
	}

	helpers.JSONResponse(w, http.StatusNoContent, nil)
}
