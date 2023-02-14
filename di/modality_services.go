package di

import (
	"github.com/online.scheduling-api/constants"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/data"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
)

func GetModalityServices() []di.Def {
	return []di.Def{
		{
			Name:  constants.MODALITY_HANDLER,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.ModalityHandler{
					ModalityService: ctn.Get(constants.MODALITY_SERVICE).(*services.ModalityService),
				}, nil
			},
		},
		{
			Name:  constants.MODALITY_SERVICE,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.ModalityService{
					ModalityRepository: ctn.Get(constants.MODALITY_REPOSITORY).(*repository.ModalityRepository),
					UserRepository:     ctn.Get(constants.USER_REPOSITORY).(*repository.UserRepository),
				}, nil
			},
		},
		{
			Name:  constants.MODALITY_REPOSITORY,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &repository.ModalityRepository{
					DB: ctn.Get(constants.DB_SERVICE).(*data.DB),
				}, nil
			},
		},
	}
}
