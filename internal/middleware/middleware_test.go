package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NarthurN/metrics-alerter/internal/middleware"
	"github.com/NarthurN/metrics-alerter/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestValidateMetrics(t *testing.T) {
	// "Шпион" для проверки, был ли вызван следующий обработчик
	var nextHandlerCalled bool
	mockNextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextHandlerCalled = true
		w.WriteHeader(http.StatusOK) // Симулируем успешный ответ от основного хендлера
	})

	// Определяем тестовые случаи
	tests := []struct {
		name             string
		requestSetup     func(r *http.Request) // Функция для настройки запроса
		expectedStatus   int
		expectNextCalled bool
	}{
		{
			name: "Успешный запрос (gauge)",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				r.SetPathValue("metricType", model.Gauge)
				r.SetPathValue("metricName", "TestGauge")
				r.SetPathValue("metricValue", "123.45")
			},
			expectedStatus:   http.StatusOK, // Ожидаем статус от mockNextHandler
			expectNextCalled: true,
		},
		{
			name: "Успешный запрос (counter)",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				r.SetPathValue("metricType", model.Counter)
				r.SetPathValue("metricName", "TestCounter")
				r.SetPathValue("metricValue", "123")
			},
			expectedStatus:   http.StatusOK,
			expectNextCalled: true,
		},
		{
			name: "Неверный Content-Type",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "application/json")
			},
			expectedStatus:   http.StatusBadRequest,
			expectNextCalled: false,
		},
		{
			name: "Пустое имя метрики",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				r.SetPathValue("metricType", model.Gauge)
				r.SetPathValue("metricName", "") // Пустое имя
			},
			expectedStatus:   http.StatusNotFound,
			expectNextCalled: false,
		},
		{
			name: "Невалидный тип метрики",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				r.SetPathValue("metricType", "invalid_type")
				r.SetPathValue("metricName", "SomeMetric")
			},
			expectedStatus:   http.StatusBadRequest,
			expectNextCalled: false,
		},
		{
			name: "Нечисловое значение метрики",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				r.SetPathValue("metricType", model.Gauge)
				r.SetPathValue("metricName", "SomeMetric")
				r.SetPathValue("metricValue", "abc") // Не число
			},
			expectedStatus:   http.StatusBadRequest,
			expectNextCalled: false,
		},
		{
			name: "Дробное значение для counter",
			requestSetup: func(r *http.Request) {
				r.Header.Set("Content-Type", "text/plain")
				r.SetPathValue("metricType", model.Counter)
				r.SetPathValue("metricName", "SomeCounter")
				r.SetPathValue("metricValue", "123.45") // Дробное
			},
			expectedStatus:   http.StatusBadRequest,
			expectNextCalled: false,
		},
	}

	// Запускаем цикл по тестам
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Сбрасываем флаг перед каждым тестом
			nextHandlerCalled = false

			// Создаем запрос и рекордер
			req := httptest.NewRequest(http.MethodPost, "/update/", nil)
			rr := httptest.NewRecorder()

			// Применяем настройки для текущего тестового случая
			tt.requestSetup(req)

			// Создаем тестируемый обработчик (нашу middleware)
			handlerToTest := middleware.ValidateMetrics(mockNextHandler)

			// Выполняем запрос через middleware
			handlerToTest.ServeHTTP(rr, req)

			// Проверяем результаты
			assert.Equal(t, tt.expectedStatus, rr.Code, "статус код не совпадает с ожидаемым")
			assert.Equal(t, tt.expectNextCalled, nextHandlerCalled, "вызов next-хендлера не соответствует ожиданиям")
		})
	}
}
