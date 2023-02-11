package helpers

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/constants"
)

func CreateTokenFor(_claims map[string]any) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	for i, c := range _claims {
		claims[i] = c
	}

	return token.SignedString(config.GetSecretKey())
}

func RetrieveToken(r *http.Request) (*jwt.Token, error) {
	if r.Header[constants.HEADER_AUTH] == nil {
		return nil, errors.New("authorization header not present")
	}

	return jwt.Parse(
		string(r.Header[constants.HEADER_AUTH][0][7:]),
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("token not found")
			}

			return config.GetSecretKey(), nil
		})
}
