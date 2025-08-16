package configuration

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Configuration struct {
	AppPort string
}

func Configurations() *Configuration {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("/************************* No .env file found")
	}

	return &Configuration{AppPort: getEnv("APP_PORT", "8080")}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
