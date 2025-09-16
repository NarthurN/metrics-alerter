package metrics

import (
	"context"

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

func (s *MetricsService) GetMetric(ctx context.Context, nameMetric string) (metric model.Metrics, err error) {
	return s.storage.Get(ctx, nameMetric)
}
