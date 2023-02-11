package config

import (
	"os"
)

func GetConnection() string {
	return os.Getenv("CONN_STR")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetMessengerBroker() string {
	return os.Getenv("MESSENGER_BROKER")
}

func GetSecretKey() []byte {
	return []byte(os.Getenv("SECRET_KEY"))
}
