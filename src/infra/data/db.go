package data

import (
	"context"
	"log"
	"time"

	"github.com/online.scheduling-api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Client *mongo.Client
}

func NewDB() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opt := options.Client().
		ApplyURI(config.GetConnection()).
		SetRegistry(MongoRegistry)

	client, _ := mongo.Connect(ctx, opt)
	err := client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("Server unavailable")
	}

	return DB{Client: client}, err
}

func (db *DB) Close() error {
	return db.Client.Disconnect(context.TODO())
}
