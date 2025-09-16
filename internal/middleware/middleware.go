package middleware

import (
	"net/http"
	"strconv"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

func ValidateMetrics(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metricType := r.PathValue("metricType")
		metricName := r.PathValue("metricName")
		metricValue := r.PathValue("metricValue")

		if metricName == "" {
			http.NotFound(w, r)
			return
		}

		if metricType != model.Counter && metricType != model.Gauge {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch metricType {
		case model.Gauge:
			if _, err := strconv.ParseFloat(metricValue, 64); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case model.Counter:
			if _, err := strconv.ParseInt(metricValue, 10, 64); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
