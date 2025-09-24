package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Reporter struct {
	client *Client
	addr   string
}

func NewReporter(client *Client, config *Config) *Reporter {
	return &Reporter{
		client: client,
		addr:   config.Addr,
	}
}

func (r *Reporter) SendMetrics(metrics []*metric) error {
	for _, mtrc := range metrics {
		url := fmt.Sprintf("http://%s/update/%s/%s/%f", r.addr, mtrc.typeM, mtrc.nameM, mtrc.valueM)
		log.Println(url)
		resp, err := r.client.Client.Post(url, contentType, http.NoBody)
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
