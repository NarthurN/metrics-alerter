package metrics

import (
	"context"
	"fmt"
	"sync"

	"github.com/NarthurN/metrics-alerter/internal/model"
	def "github.com/NarthurN/metrics-alerter/internal/repository"
	"github.com/samber/lo"
)

// Проверяем, что MemStorage реализует интерфейс ServerRepository
var _ def.ServerRepository = (*MemStorage)(nil)

type MemStorage struct {
	mu       sync.RWMutex
	gauges   map[string]model.Metrics
	counters map[string]model.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		gauges:   make(map[string]model.Metrics),
		counters: make(map[string]model.Metrics),
	}
}

// Получить метрику gauge по имени
func (m *MemStorage) GetGaugeByName(_ context.Context, nameMetrics string) (model.Metrics, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	metrics, ok := m.gauges[nameMetrics]
	if !ok {
		return model.Metrics{}, fmt.Errorf("в базе нет метрики с именем %s", nameMetrics)
	}

	return metrics, nil
}

// Установить новое значение для метрики (заменяет старое значение)
func (m *MemStorage) SetGaugeByName(_ context.Context, nameMetrics string, valueMetrics float64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	newGaugeMetric := model.Metrics{
		Value: &valueMetrics,
	}

	m.gauges[nameMetrics] = newGaugeMetric

	return nil
}

// Получить метрику counter по имени
func (m *MemStorage) GetCounterByName(_ context.Context, nameMetrics string) (model.Metrics, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	metrics, ok := m.counters[nameMetrics]
	if !ok {
		return model.Metrics{}, fmt.Errorf("в базе нет метрики с именем %s", nameMetrics)
	}

	return metrics, nil
}

// Установить новое значение для метрики (прибавляет к старому значению)
func (m *MemStorage) SetCounterByName(_ context.Context, nameMetrics string, valueMetrics int64) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	metrics, ok := m.gauges[nameMetrics]
	if !ok {
		return fmt.Errorf("в базе нет метрики с именем %s", nameMetrics)
	}

	if metrics.Value == nil {
		metrics.Value = lo.ToPtr(float64(valueMetrics))
	} else {
		*metrics.Value += float64(valueMetrics)
	}

	m.counters[nameMetrics] = metrics

	return nil
}

// Получение всех метрик
func (m *MemStorage) GetAllMetrics(_ context.Context) ([]model.Metrics, error) {
	countMetrics := len(m.counters) + len(m.gauges)
	allMetrics := make([]model.Metrics, 0, countMetrics)

	for _, v := range m.counters {
		allMetrics = append(allMetrics, v)
	}

	for _, v := range m.gauges {
		allMetrics = append(allMetrics, v)
	}

	return allMetrics, nil
}
