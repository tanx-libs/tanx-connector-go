package client

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tanx-libs/tanx-connector-go/types"
)

// ---------------------------- Unprotected Endpoints ----------------------------

/*
Test your connectivity to our REST APIs using a health check endpoint

TODO: add context support
*/
func (c *Client) Health(ctx context.Context) (types.HealthResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.healthURL.String(), nil)
	if err != nil {
		return types.HealthResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
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


// TODO yeh sala change hona padega. Bohot chudap sa hain
/*
Track 24 hour rolling window price changes to receive stats about a certain market.
For example 24 hour low, high, average price, price change and more.
*/
func (c *Client) Ticker(ctx context.Context, market string) (types.TickerResponse, error) {
	tickerURL := *c.tickerURL

	if len(market) == 1 {
		params := tickerURL.Query()
		params.Set("market", market)
		tickerURL.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tickerURL.String(), nil)
	if err != nil {
		return types.TickerResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.TickerResponse{}, err
	}
	defer resp.Body.Close()

	var tickerResponse types.TickerResponse
	err = json.NewDecoder(resp.Body).Decode(&tickerResponse)
	if err != nil {
		return types.TickerResponse{}, err
	}

	return tickerResponse, nil
}

// Track 24 hour rolling window price changes to receive stats about all markets.
func (c *Client) AllTickers(ctx context.Context) (types.AllTickerResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.tickerURL.String(), nil)
	if err != nil {
		return types.AllTickerResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return types.AllTickerResponse{}, err
	}
	defer resp.Body.Close()

	var allTickerResponse types.AllTickerResponse
	err = json.NewDecoder(resp.Body).Decode(&allTickerResponse)
	if err != nil {
		return types.AllTickerResponse{}, err
	}

	return allTickerResponse, nil
}

