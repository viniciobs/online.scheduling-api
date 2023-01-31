package config

import (
	"os"
)

func GetMongoUri() string {
	return os.Getenv("MONGO_URI")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetUsersCollection() string {
	return os.Getenv("USERS_COLLECTION")
}

func GetModalitiesCollection() string {
	return os.Getenv("MODALITIES_COLLECTION")
}

func GetMessengerBroker() string {
	return os.Getenv("MESSENGER_BROKER")
}

func GetMessengerModalitiesEditTopic() string {
	return os.Getenv("MESSENGER_MODALITIES_EDIT_TOPIC")
}
