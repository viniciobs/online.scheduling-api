package config

import (
	"os"
)

var DeletedObjects = "DELETED_OBJECTS"

func GetMongoUri() string {
	return os.Getenv("MONGO_URI")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetMessengerBroker() string {
	return os.Getenv("MESSENGER_BROKER")
}
