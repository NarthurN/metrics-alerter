package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_metricsStorage_GetMetricsFromStorage(t *testing.T) {
	tests := []struct {
		name string
		want []*metric
	}{
		{
			name: "GetMetricsFromStorageSucces",
			want: []*metric{
				{nameM: "a", typeM: "a", valueM: 1},
				{nameM: "b", typeM: "b", valueM: 2},
				{nameM: "c", typeM: "c", valueM: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMetricsStorage()
			metricsInStorage := []*metric{
				{nameM: "a", typeM: "a", valueM: 1},
				{nameM: "b", typeM: "b", valueM: 2},
				{nameM: "c", typeM: "c", valueM: 3},
			}

			m.UpdateMetricsInStorage(metricsInStorage)
			got := m.GetMetricsFromStorage()
			
			assert.Equal(t, len(got), 4)
			assert.Subset(t, got, metricsInStorage)
		})
	}
}

func Test_metricsStorage_UpdateMetricsInStorage(t *testing.T) {
	tests := []struct {
		name string
		metrics []*metric
	}{
		{
			name: "UpdateMetricsInStorage succes",
			metrics: []*metric{
				{nameM: "a", typeM: "a", valueM: 1},
				{nameM: "b", typeM: "b", valueM: 2},
				{nameM: "c", typeM: "c", valueM: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := newMetricsStorage()
			m.UpdateMetricsInStorage(tt.metrics)

			assert.Equal(t, len(m.storage), 4)
			assert.Equal(t, float64(1), m.storage[PollCount].valueM)

			m.UpdateMetricsInStorage(tt.metrics)
			m.UpdateMetricsInStorage(tt.metrics)
			m.UpdateMetricsInStorage(tt.metrics)
			assert.Equal(t, float64(4), m.storage[PollCount].valueM)
		})
	}
}
