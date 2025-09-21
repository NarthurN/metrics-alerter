package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NarthurN/metrics-alerter/internal/handler"
	"github.com/NarthurN/metrics-alerter/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockServerService struct {
	mock.Mock
}

func (m *MockServerService) AllMetrics(ctx context.Context) ([]model.Metrics, error) {
	args := m.Called(ctx)
	return nil, args.Error(1)
}

func (m *MockServerService) GetMetricByName(ctx context.Context, nameMetric string) (model.Metrics, error) {
	args := m.Called(ctx, nameMetric)
	return model.Metrics{}, args.Error(1)
}

func (m *MockServerService) SetMetricByName(ctx context.Context, nameMetric, typeMetric, valueMetric string) error {
	args := m.Called(ctx, nameMetric, typeMetric, valueMetric)
	return args.Error(0)
}

func TestHandler_UpdateMetric(t *testing.T) {
	mockService := new(MockServerService)
	tests := []struct {
		name                string
		metricType          string
		metricName          string
		metricValue         string
		mockSetup           func()
		expectedStatus      int
		expectedContentType string
	}{
		{
			name:        "Успешный запрос (StatusOK)",
			metricType:  "gauge",
			metricName:  "TestMetric",
			metricValue: "123.45",
			mockSetup: func() {
				mockService.On("SetMetricByName", mock.Anything, "TestMetric", "gauge", "123.45").
					Return(nil).
					Once()
			},
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/plain; charset=utf-8",
		},
		{
			name:        "Ошибка от сервиса (InternalServerError)",
			metricType:  "counter",
			metricName:  "FailedMetric",
			metricValue: "10",
			mockSetup: func() {
				mockService.On("SetMetricByName", mock.Anything, "FailedMetric", "counter", "10").
					Return(errors.New("something went wrong")).
					Once()
			},
			expectedStatus:      http.StatusInternalServerError,
			expectedContentType: "text/plain; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			h := handler.NewHandler(mockService)
			req := httptest.NewRequest(http.MethodPost, "/update/", nil)
			req.SetPathValue("metricType", tt.metricType)
			req.SetPathValue("metricName", tt.metricName)
			req.SetPathValue("metricValue", tt.metricValue)

			rr := httptest.NewRecorder()

			h.UpdateMetric(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			if tt.expectedContentType != "" {
				assert.Equal(t, tt.expectedContentType, rr.Header().Get("Content-Type"))
			}

			mockService.AssertExpectations(t)
		})
	}
}
