package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/tanx-libs/tanx-connector-go/types"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func New(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

/*
Specify max timeout for the http client in seconds

Note: this max timeout is applied to all requests
*/
func (c *Client) Timeout(sec int) *Client {
	c.client.Timeout = time.Duration(sec) * time.Second
	return c
}

/*
Test your connectivity to our REST APIs using a health check endpoint
*/
func (c *Client) Health() (types.HealthResponse, error) {
	resp, err := c.client.Get(c.baseURL + types.HEALTH_PATH)
	if err != nil {
		return types.HealthResponse{}, err
	}
	defer resp.Body.Close()

	var healthResponse types.HealthResponse
	err = json.NewDecoder(resp.Body).Decode(&healthResponse)
	if err != nil {
		return types.HealthResponse{}, err
	}

	return healthResponse, nil
}

// // 24 hour ticker endpoint
// func (c *Client) Ticker() (types.TickerResponse, error) {
// 	resp, err := c.client.Get(c.baseURL + types.TICKER_PATH)
// 	if err != nil {
// 		return types.TickerResponse{}, err
// 	}
// 	defer resp.Body.Close()

// 	var tickerResponse types.TickerResponse
// 	err = json.NewDecoder(resp.Body).Decode(&tickerResponse)
// 	if err != nil {
// 		return types.TickerResponse{}, err
// 	}

// 	return tickerResponse, nil
// }
