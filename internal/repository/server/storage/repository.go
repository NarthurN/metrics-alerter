package storage

import (
	"github.com/NarthurN/metrics-alerter/internal/model"
)

type MemStorage struct {
	//mu      sync.RWMutex
	metrics map[string]model.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]model.Metrics),
	}
}

// func (m *MemStorage) Get(_ context.Context, key string) (metric model.Metrics, err error) {
// 	m.mu.RLock()
// 	defer m.mu.RUnlock()

// 	metric, ok := m.metrics[key]
// 	if !ok {
// 		return model.Metrics{}, fmt.Errorf("метрики по ключу %s нет", key)
// 	}

// 	return metric, nil
// }

// func (m *MemStorage) Set(_ context.Context, key string, metric model.Metrics) (err error) {
// 	m.mu.Lock()
// 	defer m.mu.Unlock()

// 	_, ok := m.metrics[key]
// 	if ok {
// 		return fmt.Errorf("Метркиа с ключем %s уже существует", key)
// 	}

// 	m.metrics[key] = metric

// 	return nil
// }
