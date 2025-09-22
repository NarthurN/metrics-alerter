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
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		metrics: make(map[string]model.Metrics),
	}
}

// Получить метрику gauge по имени
func (m *MemStorage) GetMetricByName(_ context.Context, nameMetrics string) (model.Metrics, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	metrics, ok := m.metrics[nameMetrics]
	if !ok {
		return model.Metrics{}, fmt.Errorf("в базе нет метрики с именем %s", nameMetrics)
	}

	return metrics, nil
}

// Установить новое значение для метрики (заменяет старое значение)
func (m *MemStorage) SetGaugeByName(_ context.Context, nameMetrics string, valueMetrics float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.metrics[nameMetrics] = model.Metrics{
		ID:    nameMetrics,
		MType: model.Gauge,
		Value: &valueMetrics,
	}

	return nil
}

// Установить новое значение для метрики (прибавляет к старому значению)
func (m *MemStorage) SetCounterByName(_ context.Context, nameMetrics string, valueMetrics int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	metric, ok := m.metrics[nameMetrics]
	if !ok {
		metric = model.Metrics{
			ID:    nameMetrics,
			MType: model.Counter,
			Delta: new(int64),
		}
	}

	*metric.Delta += valueMetrics
	m.metrics[nameMetrics] = metric

	return nil
}

// Получение всех метрик
func (m *MemStorage) GetAllMetrics(_ context.Context) ([]model.Metrics, error) {
	countMetrics := len(m.metrics)
	allMetrics := make([]model.Metrics, 0, countMetrics)

	for _, v := range m.metrics {
		allMetrics = append(allMetrics, v)
	}

	return allMetrics, nil
}
