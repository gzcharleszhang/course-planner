package env

import (
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func LoadTestEnv() error {
	return godotenv.Load("../../../../test.env")
}
