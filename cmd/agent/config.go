package main

import (
	"flag"
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Addr           string `env:"ADDRESS"`
	ReportInterval int    `env:"REPORT_INTERVAL"`
	PollInterval   int    `env:"POLL_INTERVAL"`
}

var flagRunAddr string
var pollInterval int
var reportInterval int

func parseFlags() {
	flag.StringVar(&flagRunAddr, "a", "localhost:8080", "address and port to run server")
	flag.IntVar(&reportInterval, "r", 10, "частота отправки метрик на сервер в секундах")
	flag.IntVar(&pollInterval, "p", 2, "частота опроса метрик из пакета runtime в секундах")

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

	if cfg.ReportInterval == 0 {
		cfg.ReportInterval = reportInterval
	}

	if cfg.PollInterval == 0 {
		cfg.PollInterval = pollInterval
	}

	return &cfg
}
