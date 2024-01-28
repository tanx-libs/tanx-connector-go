package client

import (
	"net/http"
	"net/url"
	"time"

	"github.com/tanx-libs/tanx-connector-go/types"
)

// Cient stores all the urls and the go standard http client
type Client struct {
	baseURL   *url.URL
	healthURL *url.URL
	tickerURL *url.URL

	httpClient *http.Client
}

/*
Get a json http client with default setting

As the the builder pattern is followed, you can chain methods to customize the client.
Use BaseURL, HealthPath, TickerPath etc to customize the client.
*/
func New(baseURL ...string) (*Client, error) {
	var baseurl *url.URL
	var err error

	switch len(baseURL) {
	case 0:
		baseurl, err = url.Parse(types.JSON_BASE_URL)
	
	case 1:
		baseurl, err = url.Parse(baseURL[0])

	default:
		return nil, types.ErrTooManyArguments
	}

	if err != nil {
		return nil, err
	}

	healthurl := baseurl.JoinPath(types.HEALTH_PATH)
	tickerurl := baseurl.JoinPath(types.TICKER_PATH)

	return &Client{
		baseURL:    baseurl,
		healthURL:  healthurl,
		tickerURL:  tickerurl,
		httpClient: http.DefaultClient,
	}, nil
}

// set custom health path
func (c *Client) HealthPath(healthPath string) *Client {
	c.healthURL = c.baseURL.JoinPath(healthPath)
	return c
}

// set custom ticker path
func (c *Client) TickerPath(tickerPath string) *Client {
	c.tickerURL = c.baseURL.JoinPath(tickerPath)
	return c
}

/*
set max timeout for the http client in seconds

Note: this max timeout is applied to all requests
*/
func (c *Client) Timeout(sec int) *Client {
	c.httpClient.Timeout = time.Duration(sec) * time.Second
	return c
}
