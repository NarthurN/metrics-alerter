package middleware

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/NarthurN/metrics-alerter/internal/model"
)

func ValidateMetrics(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверка Content-Type на text/plain
		contentType := r.Header.Get("Content-Type")
		if contentType != "text/plain" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		metricType := r.PathValue("metricType")
		metricName := r.PathValue("metricName")
		metricValue := r.PathValue("metricValue")

		// Проверка metricName на пустоту
		if metricName == "" {
			http.NotFound(w, r)
			return
		}

		// Проверка metricType на валидность
		if metricType != model.Counter && metricType != model.Gauge {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Проверка metricValue. Если gauge, то float64, если counter, то int
		parsedValue, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			log.Println("metricValue не является числом")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Проверка что при counter metricValue является целым
		if metricType == model.Counter {
			if parsedValue != math.Trunc(parsedValue) {
				log.Println("metricValue не является целым числом")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
