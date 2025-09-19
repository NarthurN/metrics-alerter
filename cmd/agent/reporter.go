package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Reporter struct {
	client *http.Client
}

func NewReporter(client *http.Client) *Reporter {
	return &Reporter{
		client: client,
	}
}

func (r *Reporter) SendMetrics(metrics []*metric) error {
	for _, mtrc := range metrics {
		url := baseURL + fmt.Sprintf("%s/%s/%f", mtrc.typeM, mtrc.nameM, mtrc.valueM)
		log.Println(url)
		resp, err := r.client.Post(url, contentType, http.NoBody)
		if err != nil {
			return fmt.Errorf("ошибка http.Post: %w", err)
		}
		
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("неожиданный статус ответа: %s", resp.Status)
		}
	}
	return nil
}
