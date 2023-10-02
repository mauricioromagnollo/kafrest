package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	AppEnv    string
	AppName   string
	AppPort   string
	KafkaHost string
}

func NewEnvironment() *Environment {
	loadDotEnvFile()

	return &Environment{
		AppEnv:    getEnv("KREST_APP_ENV"),
		AppName:   getEnv("KREST_APP_NAME"),
		AppPort:   getEnv("KREST_APP_PORT"),
		KafkaHost: getEnv("KREST_KAFKA_HOST"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}

func loadDotEnvFile() {
	appEnv := getEnv("APP_ENV")

	if appEnv == "development" || appEnv == "test" {
		envFile := ".env." + appEnv
		if err := godotenv.Load(envFile); err != nil {
			log.Fatalf("Error loading .env.%v file: %v", appEnv, err)
		}
	}
}
