package metrics

import (
	"context"
	"fmt"
	"strconv"

	"github.com/NarthurN/metrics-alerter/internal/model"
	"github.com/NarthurN/metrics-alerter/internal/repository"
	def "github.com/NarthurN/metrics-alerter/internal/service"
)

var _ def.ServerService = (*MetricsService)(nil)

type MetricsService struct {
	storage repository.ServerRepository
}

func NewMetricsService(s repository.ServerRepository) *MetricsService {
	return &MetricsService{storage: s}
}

// Получение
func (m *MetricsService) GetMetricByName(ctx context.Context, nameMetric, typeMetric string) (model.Metrics, error) {
	if typeMetric == model.Gauge {
		metric, err := m.storage.GetGaugeByName(ctx, nameMetric)
		if err != nil {
			return model.Metrics{}, err
		}

		return metric, nil
	}

	if typeMetric == model.Counter {
		metric, err := m.storage.GetCounterByName(ctx, nameMetric)
		if err != nil {
			return model.Metrics{}, err
		}

		return metric, nil
	}

	return model.Metrics{}, fmt.Errorf("невалидный тип метрик typeMetric %s", typeMetric)
}

// Сохранение
func (m *MetricsService) SetMetricByName(ctx context.Context, nameMetric, typeMetric, valueMetric string) error {
	var value float64
	var err error
	if value, err = strconv.ParseFloat(valueMetric, 64); err != nil {
		return fmt.Errorf("невозможно привести %s к типу float64", valueMetric)
	}

	if typeMetric == model.Gauge {
		err := m.storage.SetGaugeByName(ctx, nameMetric, value)
		if err != nil {
			return err
		}

		return nil
	}

	if typeMetric == model.Counter {
		err := m.storage.SetCounterByName(ctx, nameMetric, int64(value))
		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("невалидный тип метрик typeMetric %s", typeMetric)
}

func (m *MetricsService) GetGaugeMetricByName(ctx context.Context, nameMetric string) (model.Metrics, error) {
	metrics, err := m.storage.GetGaugeByName(ctx, nameMetric)
	if err != nil {
		return model.Metrics{}, err
	}

	return metrics, nil
}

func (m *MetricsService) SetGaugeMetricByName(ctx context.Context, nameMteric string, valueMetric float64) error {
	if err := m.storage.SetGaugeByName(ctx, nameMteric, valueMetric); err != nil {
		return err
	}

	return nil
}

func (m *MetricsService) GetCounterMetricByName(ctx context.Context, nameMetric string) (model.Metrics, error) {
	metrics, err := m.storage.GetCounterByName(ctx, nameMetric)
	if err != nil {
		return model.Metrics{}, err
	}

	return metrics, nil
}

func (m *MetricsService) SetCounterMetricByName(ctx context.Context, nameMetric string, valueMetric int64) error {
	if err := m.storage.SetCounterByName(ctx, nameMetric, valueMetric); err != nil {
		return err
	}

	return nil
}

func (m *MetricsService) AllMetrics(ctx context.Context) ([]model.Metrics, error) {
	metrics, err := m.storage.GetAllMetrics(ctx)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}
