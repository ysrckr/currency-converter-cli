package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/ysrckr/currency-converter-cli/lib/infisical"
)

type Config struct {
	Envs            map[string]string
	InfisicalClient infisical.SecretStore
}

func NewConfig(envPath string) *Config {
	c := &Config{
		Envs: make(map[string]string),
	}
	err := godotenv.Load(envPath)
	if err != nil {
		log.Panic("Error loading .env file")
	}

	c.Envs, err = godotenv.Read(envPath)
	if err != nil {
		log.Panic("Error mapping envs")
	}

	c.InfisicalClient = infisical.NewInfisicalClient(c.Envs["INFISICAL_PROJECT_ID"])

	c.InfisicalClient.Login(c.Envs["INFISICAL_CLIENT_ID"], c.Envs["INFISICAL_CLIENT_SECRET"])

	return c
}
