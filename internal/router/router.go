package router

import (
	"net/http"

	"github.com/NarthurN/metrics-alerter/internal/handler"
	m "github.com/NarthurN/metrics-alerter/internal/middleware"
	"github.com/NarthurN/metrics-alerter/internal/service"
)

func NewRouter(s service.ServerService) http.Handler {
	mux := http.NewServeMux()
	h := handler.NewHandler(s)

	mux.HandleFunc("/", h.NotFound)
	mux.HandleFunc("POST /update/{metricType}/{metricName}/{metricValue}", m.ValidateMetrics(h.UpdateMetric))

	return mux
}
