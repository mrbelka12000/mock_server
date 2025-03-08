package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	//PathToDB       string `env:"PATH_TO_DB,required"`
	PGURL      string `env:"PG_URL,required"`
	ServerPort int    `env:"SERVER_PORT, default=5552"`
	ClientPort int    `env:"CLIENT_PORT, default=5553"`
}

func Get() (Config, error) {
	return parseConfig()
}

func parseConfig() (cfg Config, err error) {
	godotenv.Load()

	err = envconfig.Process(context.Background(), &cfg)
	if err != nil {
		return cfg, fmt.Errorf("fill config: %w", err)
	}

	return cfg, nil
}
