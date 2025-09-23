package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	parseFlags()

	client := &http.Client{
		Timeout: time.Second * 1,
	}

	log.Printf("Отправляем метрики на сервер по адресу: %s", flagRunAddr)
	log.Printf("Частота отправки метрик на сервер: %d", reportInterval)
	log.Printf("Частота опроса метрик: %d", pollInterval)
	reporter := NewReporter(client, flagRunAddr)

	metrics := newMetricsStorage()

	agent := NewAgent(metrics, reporter)

	tickerPollInterval := time.NewTicker(time.Duration(pollInterval) * time.Second)
	defer tickerPollInterval.Stop()

	tickerReportInterval := time.NewTicker(time.Duration(reportInterval) * time.Second)
	defer tickerReportInterval.Stop()

	agent.Run(tickerPollInterval.C, tickerReportInterval.C)
}
