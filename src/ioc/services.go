package ioc

import (
	"context"
	"time"

	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/business/services"
	"github.com/online.scheduling-api/src/handlers"
	"github.com/online.scheduling-api/src/infra/repository"
)

var Services = []di.Def{
	{
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			client, _ := mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongoUri()))
			err := client.Ping(ctx, readpref.Primary())

			return client, err
		},
		Close: func(obj interface{}) error {
			return obj.(*mongo.Client).Disconnect(context.TODO())
		},
	},
	{
		Name:  "user-handler",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &handlers.UsersHandler{
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