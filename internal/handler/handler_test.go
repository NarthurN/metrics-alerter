package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NarthurN/metrics-alerter/internal/handler"
	"github.com/NarthurN/metrics-alerter/internal/model"
	"github.com/NarthurN/metrics-alerter/internal/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-resty/resty/v2"
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

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("metricType", tt.metricType)
			rctx.URLParams.Add("metricName", tt.metricName)
			rctx.URLParams.Add("metricValue", tt.metricValue)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

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

func TestHandler_UpdateMetric_success(t *testing.T) {
	mockService := new(MockServerService)
	mockService.On("SetMetricByName", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	r := router.NewRouter(mockService)
	srv := httptest.NewServer(r)
	defer srv.Close()

	testCases := []struct {
		method       string
		expectedCode int
		metricType   string
		metricName   string
		metricValue  string
	}{
		{method: http.MethodPost, expectedCode: http.StatusOK, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
		{method: http.MethodPatch, expectedCode: http.StatusMethodNotAllowed, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
		{method: http.MethodOptions, expectedCode: http.StatusMethodNotAllowed, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
		{method: http.MethodHead, expectedCode: http.StatusMethodNotAllowed, metricType: "gauge", metricName: "TestMetric", metricValue: "123.45"},
	}

	for _, tt := range testCases {
		t.Run(tt.method, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tt.method
			req.URL = srv.URL + "/update/" + tt.metricType + "/" + tt.metricName + "/" + tt.metricValue
			req.SetHeader("Content-Type", "text/plain")

			resp, err := req.Send()
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, resp.StatusCode())
		})
	}
}
