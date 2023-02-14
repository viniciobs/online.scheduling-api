package di

import (
	"github.com/online.scheduling-api/constants"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserServices() []di.Def {
	return []di.Def{
		{
			Name:  constants.USER_HANDLER,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.UsersHandler{
					UserService: ctn.Get(constants.USER_SERVICE).(*services.UserServices),
				}, nil
			},
		},
		{
			Name:  constants.USER_SERVICE,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.UserServices{
					UserRepository:     ctn.Get(constants.USER_REPOSITORY).(*repository.UserRepository),
					ModalityRepository: ctn.Get(constants.MODALITY_REPOSITORY).(*repository.ModalityRepository),
				}, nil
			},
		},
		{
			Name:  constants.USER_REPOSITORY,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &repository.UserRepository{
					Client: ctn.Get(constants.DB_SERVICE).(*mongo.Client),
				}, nil
			},
		},
	}
}
