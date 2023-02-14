package api

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/online.scheduling-api/constants"
	dtoRequest "github.com/online.scheduling-api/src/api/dtos/requests"
	dtoResponse "github.com/online.scheduling-api/src/api/dtos/responses"
	validator "github.com/online.scheduling-api/src/api/dtos/validators"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
	"github.com/online.scheduling-api/src/services"
	"github.com/online.scheduling-api/src/shared"
)

type UsersHandler struct {
	UserService services.IUserServices
}

func (uc *UsersHandler) Get(w http.ResponseWriter, r *http.Request) {
	filter := models.UserFilter{}

	token, err := helpers.RetrieveToken(r)
	if err != nil {
		helpers.JSONResponseError(w, http.StatusUnauthorized, errors.New("missing auth"))
	}

	claims := token.Claims.(jwt.MapClaims)
	isCustomer := claims[constants.CLAIM_USER_ROLE].(float64) == float64(models.Customer)

	if isCustomer {
		filter.UserId, _ = uuid.Parse(claims[constants.CLAIM_USER_ID].(string))
	}

	if name := r.URL.Query().Get("name"); name != "" {
		filter.UserName = name
	}

	if modality := r.URL.Query().Get("modality"); modality != "" {
		filter.ModalityName = modality
	}

	users, responseCode := uc.UserService.Get(&filter)

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	var result []dtoResponse.UserCreateOrUpdateResponse

	for i := range users {
		result = append(result, dtoResponse.MapUserResponseFrom(users[i]))
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

	helpers.JSONResponse(w, http.StatusOK, dtoResponse.MapUserResponseFrom(user))
}

func (uc *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requestData dtoRequest.UserCreateOrUpdateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateUser(&requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	authData := dtoRequest.MapAuthRequestFrom(&requestData)
	if err := validator.ValidateAuth(authData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	u := models.MapNewUserFrom(
		requestData.Name,
		requestData.Login,
		requestData.Passphrase,
		requestData.Phone,
		requestData.Role,
		false)

	responseCode, token := uc.UserService.CreateNewUser(&u)

	if responseCode != shared.Success {
		helpers.JSONResponseError(w, helpers.GetErrorStatusCodeFrom(responseCode), nil)
		return
	}

	helpers.JSONResponse(
		w,
		http.StatusCreated,
		dtoResponse.MapNewUserResponseFrom(&u, token))
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

	helpers.JSONResponse(w, http.StatusOK, nil)
}

func (uc *UsersHandler) Edit(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}
	var requestData dtoRequest.UserCreateOrUpdateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateUser(&requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	u := models.MapUserFrom(
		requestData.Name,
		requestData.Phone,
		requestData.Role,
		false)

	responseCode := uc.UserService.EditUser(&id, &u)

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

func (h *UsersHandler) EditModalities(w http.ResponseWriter, r *http.Request) {
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

	responseCode := h.UserService.EditModalities(&userId, requestData.Modalities)

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
