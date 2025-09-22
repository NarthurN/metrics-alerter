package router

import (
	"net/http"

	"github.com/NarthurN/metrics-alerter/internal/handler"
	m "github.com/NarthurN/metrics-alerter/internal/middleware"
	"github.com/NarthurN/metrics-alerter/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(s service.ServerService) http.Handler {
	r := chi.NewRouter()

	h := handler.NewHandler(s)

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.With(m.ValidateMetrics).Post("/update/{metricType}/{metricName}/{metricValue}", h.UpdateMetric)

	return r
}
