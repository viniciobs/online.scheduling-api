package data

import "go.mongodb.org/mongo-driver/mongo"

type DB struct {
	Client *mongo.Client
}
