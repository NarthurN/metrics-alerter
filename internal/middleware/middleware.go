package middleware

import (
	"log"
	"net/http"
	"strconv"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

func ValidateMetrics(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "text/plain" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}


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

		if parsedValue, err := strconv.ParseFloat(metricValue, 64); err != nil {
			log.Println("gauge не парсится")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ошибка strconv.ParseFloat(metricValue, 64)"))
			return
		}

		if metricType == model.Counter {
			parsedValueInt, ok := parsedValue
		}

		switch metricType {
		case model.Gauge:
		case model.Counter:

			if _, err := strconv.ParseInt(metricValue, 10, 64); err != nil {
				log.Println("counter не парсится")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ошибка strconv.ParseInt(metricValue, 10, 64)"))
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
