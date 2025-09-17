package service

import (
	"context"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

type ServerService interface {
	AllMetrics(ctx context.Context) ([]model.Metrics, error)
	GetMetricByName(ctx context.Context, nameMetric, typeMetric string) (model.Metrics, error)
	SetMetricByName(ctx context.Context, nameMetric, typeMetric, valueMetric string) error
}
