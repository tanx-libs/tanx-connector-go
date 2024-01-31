package client

import (
	"net/http"
	"net/url"
)

// Cient stores all the urls and the go standard http client
type Client struct {
	baseURL        *url.URL
	healthURL      *url.URL
	tickerURL      *url.URL
	candleStickURL *url.URL
	orderBookURL   *url.URL
	tradesURL      *url.URL
	httpClient     *http.Client
}

/*
Get a json http client with default setting

As the the builder pattern is followed, you can chain methods to customize the client.
Use BaseURL, HealthPath, TickerPath etc to customize the client.
*/
func New() (*Client, error) {
	baseurl, err := url.Parse(JSON_BASE_URL)
	if err != nil {
		return nil, err
	}

	healthurl := baseurl.JoinPath(HEALTH_PATH)
	tickerurl := baseurl.JoinPath(TICKER_PATH)
	candleStickurl := baseurl.JoinPath(CANDLESTICK_PATH)
	orderBookurl := baseurl.JoinPath(ORDERBOOK_PATH)
	tradesurl := baseurl.JoinPath(TRADES_PATH)

	return &Client{
		baseURL:        baseurl,
		healthURL:      healthurl,
		tickerURL:      tickerurl,
		candleStickURL: candleStickurl,
		orderBookURL:   orderBookurl,
		tradesURL:      tradesurl,
		httpClient:     http.DefaultClient,
	}, nil
}

// TODO make the client configurable