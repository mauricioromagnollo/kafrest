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

	var appPort string

	if getEnv("PORT") == "" {
		appPort = "8080"
	} else {
		appPort = getEnv("PORT")
	}

	return &Environment{
		AppEnv:    getEnv("KREST_APP_ENV"),
		AppName:   "kafka-rest",
		AppPort:   appPort,
		KafkaHost: getEnv("KAFKA_BROKERCONNECT"),
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
