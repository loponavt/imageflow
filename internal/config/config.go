package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	HTTPPort        string        `env:"HTTP_PORT" envDefault:":8080"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"5s"`
	Env             string        `env:"ENV" envDefault:"local"`
}

func Load() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("failed to parse env: %v", err)
	}
	return &cfg
}
