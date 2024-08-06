package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ysrckr/currency-converter-cli/lib/infisical"
)

type Config struct {
	Envs            map[string]string
	InfisicalClient infisical.Client
}

func NewConfig() *Config {
	c := &Config{
		Envs: make(map[string]string),
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.Envs, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error mapping envs")
	}

	InfisicalClient := infisical.NewClient()

	InfisicalClient.Login(c.Envs["INFISICAL_CLIENT_ID"], c.Envs["INFISICAL_CLIENT_SECRET"])

	return c
}
