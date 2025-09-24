package service

import (
	"context"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

type ServerService interface {
	GetMetricByName(ctx context.Context, nameMetric string, typeMetric string) (model.Metrics, error)
	SetMetricByName(ctx context.Context, nameMetric, typeMetric string, valueMetric float64) error
	GetAllMetrics(ctx context.Context) ([]model.Metrics, error)
}
