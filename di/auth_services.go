package di

import (
	"github.com/online.scheduling-api/constants"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/sarulabs/di"
)

func GetAuthServices() []di.Def {
	return []di.Def{
		{
			Name:  constants.AUTH_HANDLER,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.AuthHandler{
					UserRepository: ctn.Get(constants.USER_REPOSITORY).(*repository.UserRepository),
				}, nil
			},
		},
	}
}
