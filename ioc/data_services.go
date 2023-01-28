package ioc

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/online.scheduling-api/config"
	"github.com/online.scheduling-api/src/helpers"
	"github.com/sarulabs/di"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDataServices() []di.Def {
	return []di.Def{
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
	}
}
