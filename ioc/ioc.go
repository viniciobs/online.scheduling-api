package ioc

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/online.scheduling-api/config"
	api "github.com/online.scheduling-api/src/api/handlers"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/online.scheduling-api/src/infra/repository"
	"github.com/online.scheduling-api/src/services"
)

var Services = []di.Def{
	{
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			if err := godotenv.Load(); err != nil {
				log.Fatal("Error loading .env file")
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			opt := options.Client().
				ApplyURI(config.GetMongoUri()).
				SetRegistry(helpers.MongoRegistry)

			client, _ := mongo.Connect(ctx, opt)
			err := client.Ping(ctx, readpref.Primary())

			if err != nil {
				log.Fatal("Server unavailable")
			}

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
