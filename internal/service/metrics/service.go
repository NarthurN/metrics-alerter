package metrics

import (
	"context"

	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/model"
	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/storage"
	def "github.com/NarthurN/metrics-alerter/internal/service"
)

var _ def.ServerService = (*MetricsService)(nil)

type MetricsService struct {
	storage *storage.MemStorage
}

func NewMetricsService(s *storage.MemStorage) *MetricsService {
	return &MetricsService{storage: s}
}

// Получение
func (m *MetricsService) GetMetricByName(ctx context.Context, nameMetric string, typeMetric string) (model.Metrics, error) {
	metric, err := m.storage.GetMetricByName(ctx, nameMetric)
	if err != nil {
		return model.Metrics{}, err
	}

	return metric, nil
}

// Сохранение
func (m *MetricsService) SetMetricByName(ctx context.Context, nameMetric, typeMetric string, valueMetric float64) error {
	if typeMetric == model.Gauge {
		err := m.storage.SetGaugeByName(ctx, nameMetric, valueMetric)
		if err != nil {
			return err
		}
	}

	if typeMetric == model.Counter {
		err := m.storage.SetCounterByName(ctx, nameMetric, int64(valueMetric))
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MetricsService) GetAllMetrics(ctx context.Context) ([]model.Metrics, error) {
	metrics, err := m.storage.GetAllMetrics(ctx)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
