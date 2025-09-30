package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/NarthurN/metrics-alerter/pkg/storage/metrics/model"
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

func (r *Reporter) SendMetrics(metrics []*model.Metrics) error {
	for _, mtrc := range metrics {
		rawURL := fmt.Sprintf("http://%s/update/%s/%s/%f", r.addr, mtrc.MType, mtrc.ID, *mtrc.Value)
		URL, err := url.Parse(rawURL)
		if err != nil {
			log.Println("rawUrl: ", err.Error())
			return fmt.Errorf("невозможно распарсить rawUrl: %s", rawURL)
		}

		log.Println("URL", rawURL)
		resp, err := r.client.Client.Post(URL.String(), model.ContentType, http.NoBody)
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
