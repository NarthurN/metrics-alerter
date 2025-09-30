package main

import (
	"context"
	"log"
	"time"

	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/storage"
	"go.uber.org/fx"
)

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
