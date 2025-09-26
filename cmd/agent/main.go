package main

import (
	"context"
	"log"
	"time"

	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/storage"
	"github.com/caarlos0/env"
	"go.uber.org/fx"
)

type Config struct {
	Addr           string        `env:"ADDRESS"`
	ReportInterval time.Duration `env:"REPORT_INTERVAL"`
	PollInterval   time.Duration `env:"POLL_INTERVAL"`
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
		cfg.ReportInterval = time.Duration(reportInterval) * time.Second
	}

	if cfg.PollInterval == 0 {
		cfg.PollInterval = time.Duration(pollInterval) * time.Second
	}

	return &cfg
}

func main() {
	app := fx.New(
		fx.Provide(
			NewConfig,
			NewClient,
			NewAgent,
			NewReporter,
			storage.NewMemStorage,
		),
		fx.Invoke(runAgent),
	)

	app.Run()
}

func runAgent(lc fx.Lifecycle, agent *Agent, config *Config) {
	var pollTicker *time.Ticker
	var reportTicker *time.Ticker

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Printf("Отправляем метрики на сервер по адресу: %s", config.Addr)
				log.Printf("Частота отправки метрик на сервер: %s", config.ReportInterval)
				log.Printf("Частота опроса метрик: %s", config.PollInterval)

				pollTicker = time.NewTicker(config.PollInterval)
				reportTicker = time.NewTicker(config.ReportInterval)

				go agent.Run(pollTicker.C, reportTicker.C)

				return nil
			},
			OnStop: func(ctx context.Context) error {
				pollTicker.Stop()
				reportTicker.Stop()
				log.Println("Агент остановлен")
				return nil
			},
		},
	)
}
