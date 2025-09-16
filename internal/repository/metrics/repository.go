package metrics

import (
	"context"
	"fmt"
	"sync"

	"github.com/NarthurN/metrics-alerter/internal/model"
	def "github.com/NarthurN/metrics-alerter/internal/repository"
)

// Проверяем, что MemStorage реализует интерфейс ServerRepository
var _ def.ServerRepository = (*MemStorage)(nil)

type MemStorage struct {
	mu      sync.RWMutex
	metrics map[string]model.Metrics
	// gauges map[string]float64
	// counters map[string]int64
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]model.Metrics),
	}
}

func (m *MemStorage) Get(_ context.Context, key string) (metric model.Metrics, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	metric, ok := m.metrics[key]
	if !ok {
		return model.Metrics{}, fmt.Errorf("метрики по ключу %s нет", key)
	}

	return metric, nil
}

func (m *MemStorage) Set(_ context.Context, key string, metric model.Metrics) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.metrics[key]
	if ok {
		return fmt.Errorf("Метркиа с ключем %s уже существует", key)
	}

	m.metrics[key] = metric

	return nil
}

func (m *MemStorage) Delete(_ context.Context, key string) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.metrics, key)

	return nil
}

func (m *MemStorage) Update(_ context.Context, key string, metric model.Metrics) (err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.metrics[key] = metric

	return nil
}
