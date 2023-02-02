package ioc

import (
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetModalityServices() []di.Def {
	return []di.Def{
		{
			Name:  "modality-handler",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.ModalityHandler{
					ModalityService: ctn.Get("modality-service").(*services.ModalityService),
				}, nil
			},
		},
		{
			Name:  "modality-service",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.ModalityService{
					ModalityRepository: ctn.Get("modality-repository").(*repository.ModalityRepository),
					UserRepository:     ctn.Get("user-repository").(*repository.UserRepository),
				}, nil
			},
		},
		{
			Name:  "modality-repository",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &repository.ModalityRepository{
					Client: ctn.Get("mongo").(*mongo.Client),
				}, nil
			},
		},
	}
}
