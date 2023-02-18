package tests

import (
	"errors"
	"testing"

	"github.com/online.scheduling-api/src/infra/services"
	"github.com/online.scheduling-api/src/shared"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestShouldMapResponseCodeFromError(t *testing.T) {
	// Arrange
	err := map[error]shared.Code{
		mongo.ErrNoDocuments:        shared.NonExistentRecord,
		mongo.ErrClientDisconnected: shared.ThirdPartyFail,
		errors.New("Error"):         shared.Unkown,
	}

	// Act

	// Assert
	for k, v := range err {
		e := services.MapErrorFrom(k)
		if e != v {
			t.Error("Maping error")
		}
	}
}
