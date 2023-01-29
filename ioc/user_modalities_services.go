package ioc

import (
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
)

func GetUserModalitiesServices() []di.Def {
	return []di.Def{
		{
			Name:  "user-modalities-handler",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.UserModalitiesHandler{
					UserModalitiesService: ctn.Get("user-modalities-service").(*services.UserModalitiesService),
				}, nil
			},
		},
		{
			Name:  "user-modalities-service",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.UserModalitiesService{
					ModalityRepository: ctn.Get("modality-repository").(*repository.ModalityRepository),
					UserRepository:     ctn.Get("user-repository").(*repository.UserRepository),
				}, nil
			},
		},
	}
}
