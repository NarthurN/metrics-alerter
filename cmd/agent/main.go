package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

type metric struct {
	nameM  string
	typeM  string
	valueM float64
}

type metricsStorage struct {
	mu      sync.RWMutex
	storage map[string]*metric
}

func newMetricsStorage() *metricsStorage {
	return &metricsStorage{
		storage: make(map[string]*metric),
	}
}

func (m *metricsStorage) UpdateMetricsInStorage(metrics []metric) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, mm := range metrics {
		mtrc, ok := m.storage[mm.nameM]
		if !ok {
			m.storage[mm.nameM] = &metric{
				nameM:  mm.nameM,
				typeM:  mm.typeM,
				valueM: mm.valueM,
			}

			continue
		}

		switch mm.typeM {
		case gauge:
			mtrc.valueM = mm.valueM
		case counter:
			mtrc.valueM += 1
		}
	}

	mtrc, ok := m.storage[PollCount]
	if !ok {
		m.storage[PollCount] = &metric{
			nameM:  PollCount,
			typeM:  counter,
			valueM: 1,
		}

		return
	}
	mtrc.valueM++
}

func (m *metricsStorage) GetMetricsFromStorage() []*metric {
	m.mu.RLock()
	defer m.mu.RUnlock()

	res := make([]*metric, 0, len(m.storage))
	for _, mtrc := range m.storage {
		res = append(res, mtrc)
	}

	return res
}

func main() {
	var ms runtime.MemStats
	metrics := newMetricsStorage()

	tickerPollInterval := time.NewTicker(pollInterval)
	defer tickerPollInterval.Stop()

	tickerReportInterval := time.NewTicker(reportInterval)
	defer tickerPollInterval.Stop()

	for {
		select {
		case <-tickerPollInterval.C:
			runtime.ReadMemStats(&ms)

			metricsInfo := getMetricsFromMemStats(&ms)
			metrics.UpdateMetricsInStorage(metricsInfo)
		case <-tickerReportInterval.C:
			for _, mtrc := range metrics.GetMetricsFromStorage() {
				url := baseURL + fmt.Sprintf("%s/%s/%f", mtrc.nameM, mtrc.typeM, mtrc.valueM)
				resp, err := http.Post(url, contentType, http.NoBody)
				if err != nil {
					log.Println("ошибка http.Post:", err)
					return
				}
				defer resp.Body.Close()

				_, err = io.Copy(io.Discard, resp.Body)
				if err != nil {
					log.Println("ошибка io.Copy:", err)
					return
				}
			}
		}
	}

}

func getMetricsFromMemStats(ms *runtime.MemStats) []metric {
	return []metric{
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
