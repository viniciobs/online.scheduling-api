package api

import (
	"net/http"

	dtoRequest "github.com/online.scheduling-api/src/api/dtos/requests"
	dtoResponse "github.com/online.scheduling-api/src/api/dtos/responses"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/models"
)

type AuthHandler struct {
	UserRepository repository.IUserRepository
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var requestData dtoRequest.AuthCreateOrUpdateRequest

	if err := helpers.ReadJSONBody(r, &requestData); err != nil {
		helpers.JSONResponseError(w, http.StatusBadRequest, err)
		return
	}

	ok, user := h.UserRepository.Authenticate(requestData.Login, requestData.Passphrase)

	if !ok {
		helpers.JSONResponseError(w, http.StatusNotFound, nil)
		return
	}

	claims := models.MapUserClaimsFrom(user)
	token, err := helpers.CreateTokenFor(claims)
	if err != nil {
		helpers.JSONResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	helpers.JSONResponse(w, http.StatusOK, dtoResponse.MapAuthResponseFrom(user, token))
}
