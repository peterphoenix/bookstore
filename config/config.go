package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port    int    `envconfig:"PORT" default:"9001"`
	Host    string `envconfig:"HOST" default:"localhost:9001"`
	BaseURL string `envconfig:"BASE_URL" default:"http://localhost:9001"`

	PostgresConfig *PostgresConfig `envconfig:"POSTGRES"`
}

type PostgresConfig struct {
	Host     string `envconfig:"HOST" default:"localhost"`
	Port     int    `envconfig:"PORT" default:"5432"`
	Database string `envconfig:"DATABASE" default:"bookstore"`
	User     string `envconfig:"USER" default:"root"`
	Password string `envconfig:"PASSWORD" default:"password"`
}

func NewConfig(ctx context.Context) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error parse config: %v", err)
	}

	var cfg Config
	envconfig.MustProcess("", &cfg)

	return &cfg
}
