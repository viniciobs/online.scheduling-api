package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoUri() string {
	verifyEnvFile()
	return os.Getenv("MONGO_URI")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetUsersCollection() string {
	return os.Getenv("USERS_COLLECTION")
}

func verifyEnvFile() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
