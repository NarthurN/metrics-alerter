package repository

import (
	"context"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

// Интерфейс для работы с базой данных метрик
type ServerRepository interface {
	GetMetricByName(_ context.Context, nameMetrics string) (model.Metrics, error)
	SetGaugeByName(ctx context.Context, nameMetrics string, valueMetrics float64) error
	SetCounterByName(ctx context.Context, nameMetrics string, valueMetrics int64) error
	GetAllMetrics(ctx context.Context) ([]model.Metrics, error)
}
