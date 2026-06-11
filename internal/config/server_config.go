package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Port   string `validate:"required"`
	AppEnv string `validate:"required"`
}

func LoadServerConfig() (*ServerConfig, error) {
	// load the env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Something went wrong while loading the env vars:", err.Error())
		return nil, err
	}

	// load the envs & create the config instance
	cfg := &ServerConfig{
		Port:   LoadSingleEnvVar("PORT", "3000"),
		AppEnv: LoadSingleEnvVar("APP_ENV", "development"),
	}

	return cfg, nil
}
