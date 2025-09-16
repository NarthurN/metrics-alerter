package handler

import (
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
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
