package main

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Addr string `env:"ADDRESS"`
}

var flagRunAddr string

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")

	flag.Parse()
}

func NewConfig() *Config {
	parseFlags()

	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Ошибка чтения env файла: %v", err)
	}

	if cfg.Addr == "" {
		cfg.Addr = flagRunAddr
	}

	return &cfg
}
