package main

import (
	"context"
	"log"
	"time"

	"go.uber.org/fx"
)

// Config хранит всю конфигурацию агента
type Config struct {
	Addr           string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

// NewConfig является провайдером для нашей конфигурации.
// Он парсит флаги и возвращает готовую структуру.
func NewConfig() *Config {
	parseFlags()
	return &Config{
		Addr:           flagRunAddr,
		ReportInterval: time.Duration(reportInterval) * time.Second,
		PollInterval:   time.Duration(pollInterval) * time.Second,
	}
}

func main() {
	app := fx.New(
		fx.Provide(
			NewConfig,
			NewClient,
			NewAgent,
			NewReporter,
			newMetricsStorage,
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
