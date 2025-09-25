package main

import (
	"net/http"
	"time"
)

type Client struct {
	Client *http.Client
}

func NewClient() *Client {
	client := &http.Client{
		Timeout: time.Second * 1,
	}

	return &Client{
		Client: client,
	}
}
