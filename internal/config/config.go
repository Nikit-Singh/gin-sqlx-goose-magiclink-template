package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	SERVER_PORT string `env:"SERVER_PORT,required"`
	DB_HOST     string `env:"DB_HOST,required"`
	DB_PORT     string `env:"DB_PORT,required"`
	DB_NAME     string `env:"DB_NAME,required"`
	DB_USER     string `env:"DB_USER,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_SSLMODE  string `env:"DB_SSLMODE,required"`
	DB_URL      string `env:"DB_URL,required"`
	ENV         string `env:"ENV,required"`
}

var Envs = newEnvConfig()

func newEnvConfig() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Failed to load environment variables from .env: %v", err)
	}

	return config
}
