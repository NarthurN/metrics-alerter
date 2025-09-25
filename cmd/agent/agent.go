package main

import (
	"log"
	"math/rand/v2"
	"runtime"
	"time"

	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/model"
	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/storage"
	"github.com/samber/lo"
)

type Agent struct {
	//storage  *metricsStorage
	storage  *storage.MemStorage
	reporter *Reporter
}

// storage *metricsStorage
func NewAgent(storage *storage.MemStorage, reporter *Reporter) *Agent {
	return &Agent{storage: storage, reporter: reporter}
}

func (a *Agent) Run(pollCh, reportCh <-chan time.Time) {
	var ms runtime.MemStats

	for {
		select {
		case <-pollCh:
			log.Println("Сбор метрик...")
			runtime.ReadMemStats(&ms)
			metricsInfo := getMetricsFromMemStats(&ms)
			a.storage.UpdateMetricsInStorage(metricsInfo)
		case <-reportCh:
			log.Println("Отправка метрик...")
			metricsToSend := a.storage.GetMetricsFromStorage()
			if err := a.reporter.SendMetrics(metricsToSend); err != nil {
				log.Println("Ошибка отправки метрик:", err)
			}
		}
	}
}

func getMetricsFromMemStats(ms *runtime.MemStats) []*model.Metrics {
	return []*model.Metrics{
		{ID: model.Alloc, MType: model.Gauge, Value: lo.ToPtr(float64(ms.Alloc))},
		{ID: model.BuckHashSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.BuckHashSys))},
		{ID: model.Frees, MType: model.Gauge, Value: lo.ToPtr(float64(ms.Frees))},
		{ID: model.GCCPUFraction, MType: model.Gauge, Value: lo.ToPtr(float64(ms.GCCPUFraction))},
		{ID: model.GCSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.GCSys))},
		{ID: model.HeapAlloc, MType: model.Gauge, Value: lo.ToPtr(float64(ms.HeapAlloc))},
		{ID: model.HeapIdle, MType: model.Gauge, Value: lo.ToPtr(float64(ms.HeapIdle))},
		{ID: model.HeapInuse, MType: model.Gauge, Value: lo.ToPtr(float64(ms.HeapInuse))},
		{ID: model.HeapObjects, MType: model.Gauge, Value: lo.ToPtr(float64(ms.HeapObjects))},
		{ID: model.HeapReleased, MType: model.Gauge, Value: lo.ToPtr(float64(ms.HeapReleased))},
		{ID: model.HeapSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.HeapSys))},
		{ID: model.LastGC, MType: model.Gauge, Value: lo.ToPtr(float64(ms.LastGC))},
		{ID: model.Lookups, MType: model.Gauge, Value: lo.ToPtr(float64(ms.Lookups))},
		{ID: model.MCacheInuse, MType: model.Gauge, Value: lo.ToPtr(float64(ms.MCacheInuse))},
		{ID: model.MCacheSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.MCacheSys))},
		{ID: model.MSpanInuse, MType: model.Gauge, Value: lo.ToPtr(float64(ms.MSpanInuse))},
		{ID: model.MSpanSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.MSpanSys))},
		{ID: model.Mallocs, MType: model.Gauge, Value: lo.ToPtr(float64(ms.Mallocs))},
		{ID: model.NextGC, MType: model.Gauge, Value: lo.ToPtr(float64(ms.NextGC))},
		{ID: model.NumForcedGC, MType: model.Gauge, Value: lo.ToPtr(float64(ms.NumForcedGC))},
		{ID: model.NumGC, MType: model.Gauge, Value: lo.ToPtr(float64(ms.NumGC))},
		{ID: model.OtherSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.OtherSys))},
		{ID: model.PauseTotalNs, MType: model.Gauge, Value: lo.ToPtr(float64(ms.PauseTotalNs))},
		{ID: model.StackInuse, MType: model.Gauge, Value: lo.ToPtr(float64(ms.StackInuse))},
		{ID: model.StackSys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.StackSys))},
		{ID: model.Sys, MType: model.Gauge, Value: lo.ToPtr(float64(ms.Sys))},
		{ID: model.TotalAlloc, MType: model.Gauge, Value: lo.ToPtr(float64(ms.TotalAlloc))},
		{ID: model.RandomValue, MType: model.Gauge, Value: lo.ToPtr(100 * rand.Float64())},
	}
}
