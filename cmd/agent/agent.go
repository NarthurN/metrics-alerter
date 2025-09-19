package main

import (
	"log"
	"math/rand/v2"
	"runtime"
	"time"
)

type Agent struct {
	storage  *metricsStorage
	reporter *Reporter
}

func NewAgent(storage *metricsStorage, reporter *Reporter) *Agent {
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

func getMetricsFromMemStats(ms *runtime.MemStats) []*metric {
	return []*metric{
		{nameM: Alloc, typeM: gauge, valueM: float64(ms.Alloc)},
		{nameM: BuckHashSys, typeM: gauge, valueM: float64(ms.BuckHashSys)},
		{nameM: Frees, typeM: gauge, valueM: float64(ms.Frees)},
		{nameM: GCCPUFraction, typeM: gauge, valueM: float64(ms.GCCPUFraction)},
		{nameM: GCSys, typeM: gauge, valueM: float64(ms.GCSys)},
		{nameM: HeapAlloc, typeM: gauge, valueM: float64(ms.HeapAlloc)},
		{nameM: HeapIdle, typeM: gauge, valueM: float64(ms.HeapIdle)},
		{nameM: HeapInuse, typeM: gauge, valueM: float64(ms.HeapInuse)},
		{nameM: HeapObjects, typeM: gauge, valueM: float64(ms.HeapObjects)},
		{nameM: HeapReleased, typeM: gauge, valueM: float64(ms.HeapReleased)},
		{nameM: HeapSys, typeM: gauge, valueM: float64(ms.HeapSys)},
		{nameM: LastGC, typeM: gauge, valueM: float64(ms.LastGC)},
		{nameM: Lookups, typeM: gauge, valueM: float64(ms.Lookups)},
		{nameM: MCacheInuse, typeM: gauge, valueM: float64(ms.MCacheInuse)},
		{nameM: MCacheSys, typeM: gauge, valueM: float64(ms.MCacheSys)},
		{nameM: MSpanInuse, typeM: gauge, valueM: float64(ms.MSpanInuse)},
		{nameM: MSpanSys, typeM: gauge, valueM: float64(ms.MSpanSys)},
		{nameM: Mallocs, typeM: gauge, valueM: float64(ms.Mallocs)},
		{nameM: NextGC, typeM: gauge, valueM: float64(ms.NextGC)},
		{nameM: NumForcedGC, typeM: gauge, valueM: float64(ms.NumForcedGC)},
		{nameM: NumGC, typeM: gauge, valueM: float64(ms.NumGC)},
		{nameM: OtherSys, typeM: gauge, valueM: float64(ms.OtherSys)},
		{nameM: PauseTotalNs, typeM: gauge, valueM: float64(ms.PauseTotalNs)},
		{nameM: StackInuse, typeM: gauge, valueM: float64(ms.StackInuse)},
		{nameM: StackSys, typeM: gauge, valueM: float64(ms.StackSys)},
		{nameM: Sys, typeM: gauge, valueM: float64(ms.Sys)},
		{nameM: TotalAlloc, typeM: gauge, valueM: float64(ms.TotalAlloc)},
		{nameM: RandomValue, typeM: gauge, valueM: 100 * rand.Float64()},
	}
}
