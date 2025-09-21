package handler

import (
	"log"
	"net/http"

	"github.com/NarthurN/metrics-alerter/internal/service"
)

type Handler struct {
	service service.ServerService
}

func NewHandler(s service.ServerService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) UpdateMetric(w http.ResponseWriter, r *http.Request) {
	metricType := r.PathValue("metricType")
	metricName := r.PathValue("metricName")
	metricValue := r.PathValue("metricValue")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	err := h.service.SetMetricByName(r.Context(), metricName, metricType, metricValue)
	if err != nil {
		log.Println("SetMetricByName:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
