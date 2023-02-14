package di

import (
	"context"
	"log"
	"time"

	"github.com/online.scheduling-api/constants"
	"github.com/online.scheduling-api/src/infra/data"
	"github.com/sarulabs/di"

	"github.com/online.scheduling-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDataServices() []di.Def {
	return []di.Def{
		{
			Name:  constants.DB_SERVICE,
			Scope: di.Request,
			Build: func(ctn di.Container) (interface{}, error) {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				opt := options.Client().
					ApplyURI(config.GetConnection()).
					SetRegistry(data.MongoRegistry)

				client, _ := mongo.Connect(ctx, opt)
				err := client.Ping(ctx, readpref.Primary())

				if err != nil {
					log.Fatal("Server unavailable")
				}

				return &data.DB{Client: client}, err

				// return data.NewDB()
			},
			Close: func(obj interface{}) error {
				return obj.(*data.DB).Close()
			},
		},
	}
}
