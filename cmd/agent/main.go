package main

import (
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 1,
	}

	reporter := NewReporter(client)

	metrics := newMetricsStorage()

	agent := NewAgent(metrics, reporter)

	tickerPollInterval := time.NewTicker(pollInterval)
	defer tickerPollInterval.Stop()

	tickerReportInterval := time.NewTicker(reportInterval)
	defer tickerReportInterval.Stop()

	agent.Run(tickerPollInterval.C, tickerReportInterval.C)
}
