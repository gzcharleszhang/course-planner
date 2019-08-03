package env

import (
	"github.com/joho/godotenv"
)

const (
	MongoURIEnvKey    string = "MONGO_URI"
	MongoDBNameEnvKey string = "MONGO_DB_NAME"
	JWTSecretEnvKey   string = "JWT_SECRET"
)

func LoadEnv() error {
	return godotenv.Load()
}

func LoadTestEnv() error {
	return godotenv.Load("../../../../test.env")
}
