package ioc

import (
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserServices() []di.Def {
	return []di.Def{
		{
			Name:  "user-handler",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &api.UsersHandler{
					UserService: ctn.Get("user-service").(*services.UserServices),
				}, nil
			},
		},
		{
			Name:  "user-service",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &services.UserServices{
					UserRepository: ctn.Get("user-repository").(*repository.UserRepository),
				}, nil
			},
		},
		{
			Name:  "user-repository",
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				return &repository.UserRepository{
					Client: ctn.Get("mongo").(*mongo.Client),
				}, nil
			},
		},
	}
}
