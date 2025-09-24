package storage

import (
	"context"
	"fmt"
	"sync"

	def "github.com/NarthurN/metrics-alerter/pkg/storage"
	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/model"
)

// Проверяем, что MemStorage реализует интерфейс ServerRepository
var _ def.Storage = (*MemStorage)(nil)

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

func (m *MemStorage) UpdateMetricsInStorage(metrics []*model.Metrics) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, mm := range metrics {
		mtrc, ok := m.metrics[mm.ID]
		if !ok {
			m.metrics[mm.ID] = model.Metrics{
				ID:    mm.ID,
				MType: mm.MType,
				Value: mm.Value,
			}

			continue
		}

		switch mm.MType {
		case model.Gauge:
			mtrc.Value = mm.Value
		case model.Counter:
			*mtrc.Value += *mm.Value
		}
	}

	mtrc, ok := m.metrics[model.PollCount]
	if !ok {
		m.metrics[model.PollCount] = model.Metrics{
			ID:    model.PollCount,
			MType: model.Counter,
			Value: new(float64),
		}

		return
	}
	*mtrc.Value++
}

func (m *MemStorage) GetMetricsFromStorage() []*model.Metrics {
	m.mu.RLock()
	defer m.mu.RUnlock()

	res := make([]*model.Metrics, 0, len(m.metrics))
	for _, mtrc := range m.metrics {
		res = append(res, &mtrc)
	}

	return res
}
