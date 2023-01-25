package services

import (
	"github.com/online.scheduling-api/src/shared"
	"go.mongodb.org/mongo-driver/mongo"
)

func MapErrorFrom(err error) shared.Code {
	switch err {
	case mongo.ErrNoDocuments:
		return shared.NonExistentRecord
	case mongo.ErrClientDisconnected:
		return shared.ThirdPartyFail
	default:
		return shared.Unkown
	}
}
