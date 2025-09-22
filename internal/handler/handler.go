package handler

import (
	"log"
	"net/http"

	"github.com/NarthurN/metrics-alerter/internal/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service service.ServerService
}

func NewHandler(s service.ServerService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")
	metricValue := chi.URLParam(r, "metricValue")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	err := h.service.SetMetricByName(r.Context(), metricName, metricType, metricValue)
	if err != nil {
		log.Println("SetMetricByName:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
