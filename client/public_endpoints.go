package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ---------------------------- Unprotected Endpoints ----------------------------

func (c *Client) Health(ctx context.Context) (HealthResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.healthURL.String(), nil)
	if err != nil {
		return HealthResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return HealthResponse{}, err
	}
	defer resp.Body.Close()

	var healthResponse HealthResponse
	err = json.NewDecoder(resp.Body).Decode(&healthResponse)
	if err != nil {
		return HealthResponse{}, err
	}

	return healthResponse, nil
}

/*
Track 24 hour rolling window price changes to receive stats about a certain market.
For example 24 hour low, high, average price, price change and more.
*/
func (c *Client) Ticker(ctx context.Context, market string) (TickerResponse, error) {
	if len(market) == 0 {
		return TickerResponse{}, ErrMarketNotProvided
	}

	tickerURL := *c.tickerURL
	params := tickerURL.Query()
	params.Set("market", market)
	tickerURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tickerURL.String(), nil)
	if err != nil {
		return TickerResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return TickerResponse{}, err
	}
	defer resp.Body.Close()

	var tickerResponse TickerResponse
	err = json.NewDecoder(resp.Body).Decode(&tickerResponse)
	if err != nil {
		return TickerResponse{}, err
	}

	return tickerResponse, nil
}

// Track 24 hour rolling window price changes to receive stats about all markets.
func (c *Client) AllTickers(ctx context.Context) (AllTickerResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.tickerURL.String(), nil)
	if err != nil {
		return AllTickerResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AllTickerResponse{}, err
	}
	defer resp.Body.Close()

	var allTickerResponse AllTickerResponse
	err = json.NewDecoder(resp.Body).Decode(&allTickerResponse)
	if err != nil {
		return AllTickerResponse{}, err
	}

	return allTickerResponse, nil
}

// TODO make it simpler
// candle stick
func (c *Client) CandleStick(ctx context.Context, market string, options CandleStickOptions) (CandleStickResponse, error) {
	if len(market) == 0 {
		return CandleStickResponse{}, ErrMarketNotProvided
	}

	candleStickURL := *c.candleStickURL
	params := candleStickURL.Query()

	params.Set("market", market)

	if options.Limit != 0 {
		params.Set("limit", fmt.Sprintf("%d", options.Limit))
	} else {
		params.Set("limit", fmt.Sprintf("%d", CANDLESTICK_DEFAULT_LIMIT))
	}

	if options.Period != 0 {
		params.Set("period", fmt.Sprintf("%d", options.Period))
	} else {
		params.Set("period", fmt.Sprintf("%d", CANDLESTICK_DEFAULT_PERIOD))
	}

	if options.StartTime != 0 {
		params.Set("start_time", fmt.Sprintf("%d", options.StartTime))
	}

	if options.EndTime != 0 {
		params.Set("end_time", fmt.Sprintf("%d", options.EndTime))
	}

	candleStickURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, candleStickURL.String(), nil)
	if err != nil {
		return CandleStickResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CandleStickResponse{}, err
	}
	defer resp.Body.Close()

	var candleStickResponse CandleStickResponse
	err = json.NewDecoder(resp.Body).Decode(&candleStickResponse)
	if err != nil {
		return CandleStickResponse{}, err
	}

	return candleStickResponse, nil
}

// orderBook
func (c *Client) OrderBook(ctx context.Context, market string, options OrderBookOptions) (OrderBookResponse, error) {
	if len(market) == 0 {
		return OrderBookResponse{}, ErrMarketNotProvided
	}

	orderBookURL := *c.orderBookURL
	params := orderBookURL.Query()

	params.Set("market", market)

	if options.AskLimit != 0 {
		params.Set("ask_limit", fmt.Sprintf("%d", options.AskLimit))
	} else {
		params.Set("ask_limit", fmt.Sprintf("%d", ORDERBOOK_DEFAULT_ASK_LIMIT))
	}

	if options.BidLimit != 0 {
		params.Set("bid_limit", fmt.Sprintf("%d", options.BidLimit))
	} else {
		params.Set("bid_limit", fmt.Sprintf("%d", ORDERBOOK_DEFAULT_BID_LIMIT))
	}

	orderBookURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, orderBookURL.String(), nil)
	if err != nil {
		return OrderBookResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return OrderBookResponse{}, err
	}
	defer resp.Body.Close()

	var orderBookResponse OrderBookResponse
	err = json.NewDecoder(resp.Body).Decode(&orderBookResponse)
	if err != nil {
		return OrderBookResponse{}, err
	}

	return orderBookResponse, nil
}

// trades
func (c *Client) Trades(ctx context.Context, market string, options TradesOptions) (TradesResponse, error) {
	if len(market) == 0 {
		return TradesResponse{}, ErrMarketNotProvided
	}

	tradesURL := *c.tradesURL
	params := tradesURL.Query()

	params.Set("market", market)

	if options.Limit != 0 {
		params.Set("limit", fmt.Sprintf("%d", options.Limit))
	} else {
		params.Set("limit", fmt.Sprintf("%d", TRADES_DEFAULT_LIMIT))
	}

	if options.Timestamp != 0 {
		params.Set("timestamp", fmt.Sprintf("%d", options.Timestamp))
	}

	if options.OrderBy != "" {
		params.Set("from", string(options.OrderBy))
	} else {
		params.Set("from", string(TRADES_DEFAULT_ORDER_BY))
	}

	tradesURL.RawQuery = params.Encode()
	log.Println(tradesURL.String())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, tradesURL.String(), nil)
	if err != nil {
		return TradesResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return TradesResponse{}, err
	}
	defer resp.Body.Close()

	var tradesResponse TradesResponse
	err = json.NewDecoder(resp.Body).Decode(&tradesResponse)
	if err != nil {
		return TradesResponse{}, err
	}

	return tradesResponse, nil
}
