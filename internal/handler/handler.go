package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/NarthurN/metrics-alerter/internal/model"
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

func (h *Handler) GetMetricByName(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	metricName := chi.URLParam(r, "metricName")

	metric, err := h.service.GetMetricByName(r.Context(), metricName, metricType)
	if err != nil {
		log.Println("GetMetricByName:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	var value string
	switch metric.MType {
	case model.Counter:
		value = fmt.Sprintf("%d", *metric.Delta)
	case model.Gauge:
		value = fmt.Sprintf("%g", *metric.Value)
	}

	w.Write([]byte(value))
}

const metricsTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Metrics</title>
</head>
<body>
    <h1>All Metrics</h1>
    <ul>
        {{range .}}
            <li>{{.ID}}: {{if eq .MType "gauge"}}{{.Value}}{{else}}{{.Delta}}{{end}}</li>
        {{end}}
    </ul>
</body>
</html>
`

func (h *Handler) GetAllMetrics(w http.ResponseWriter, r *http.Request) {
	metrics, err := h.service.GetAllMetrics(r.Context())
	if err != nil {
		http.Error(w, "Failed to get metrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tmpl, err := template.New("metrics").Parse(metricsTemplate)
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, metrics)
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}
