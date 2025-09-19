package main

import "sync"

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

func (m *metricsStorage) UpdateMetricsInStorage(metrics []*metric) {
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
