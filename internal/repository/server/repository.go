package server

import (
	"context"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

// Интерфейс для работы с базой данных метрик
type ServerRepository interface {
	Get(ctx context.Context, nameMetric string) (metric model.Metrics, err error)
	Set(ctx context.Context, nameMetric string, metric model.Metrics) (err error)
	Delete(ctx context.Context, nameMetric string) (err error)
	Update(ctx context.Context, nameMetric string, metric model.Metrics) (err error)
}
