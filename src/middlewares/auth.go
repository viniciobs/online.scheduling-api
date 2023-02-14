package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/online.scheduling-api/constants"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/models"
)

type Next func(w http.ResponseWriter, r *http.Request)

func EnsureAuth(next Next) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token *jwt.Token
		var err error

		if token, err = helpers.RetrieveToken(r); err != nil {
			helpers.JSONResponseError(w, http.StatusUnauthorized, err)
			return
		}

		if !token.Valid {
			helpers.JSONResponse(w, http.StatusForbidden, nil)
			return
		}

		next(w, r)
	})
}

func EnsureRole(next Next, roles []models.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token *jwt.Token
		var err error

		if token, err = helpers.RetrieveToken(r); err != nil {
			helpers.JSONResponseError(w, http.StatusUnauthorized, err)
			return
		}

		isAllowed := func() bool {
			uRole := token.Claims.(jwt.MapClaims)[constants.CLAIM_USER_ROLE].(float64)
			for _, role := range roles {
				if uRole == float64(role) {
					return true
				}
			}
			return false
		}()

		if !isAllowed || !token.Valid {
			helpers.JSONResponse(w, http.StatusForbidden, nil)
			return
		}

		next(w, r)
	})
}
