package service

import (
	"context"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

type ServerService interface {
	GetMetric(ctx context.Context, nameMetric string) (metric model.Metrics, err error)
}
