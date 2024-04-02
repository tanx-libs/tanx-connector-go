package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderBookOptions struct {
	AskLimit int
	BidLimit int
}

type Ask struct {
	ID              int    `json:"id"`
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	OriginVolume    string `json:"origin_volume"`
	RemainingVolume string `json:"remaining_volume"`
	ExecutedVolume  string `json:"executed_volume"`
	MakerFee        string `json:"maker_fee"`
	TakerFee        string `json:"taker_fee"`
	TradesCount     int    `json:"trades_count"`
}

type Bid struct {
	ID              int    `json:"id"`
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	OriginVolume    string `json:"origin_volume"`
	RemainingVolume string `json:"remaining_volume"`
	ExecutedVolume  string `json:"executed_volume"`
	MakerFee        string `json:"maker_fee"`
	TakerFee        string `json:"taker_fee"`
	TradesCount     int    `json:"trades_count"`
}

type OrderBookPayload struct {
	Asks []Ask `json:"asks"`
	Bids []Bid `json:"bids"`
}

/*
{
  "status": "success",
  "message": "Retrieval Successful",
  "payload": {
    "asks": [
      {
        "id": 7495410,
        "uuid": "5c30a97a-39b1-4289-b119-5d0ffd96aaa5",
        "side": "sell",
        "ord_type": "market",
        "price": null,
        "avg_price": "0.0",
        "state": "wait",
        "market": "btcusdc",
        "created_at": "2022-07-20T04:14:08+02:00",
        "updated_at": "2022-07-20T04:14:08+02:00",
        "origin_volume": "0.0006",
        "remaining_volume": "0.0006",
        "executed_volume": "0.0",
        "maker_fee": "0.001",
        "taker_fee": "0.001",
        "trades_count": 0
      }
    ],
    "bids": [
      {
        "id": 7470160,
        "uuid": "795cfe5d-06ab-4150-8f44-06dc2b0e88f8",
        "side": "buy",
        "ord_type": "limit",
        "price": "23490.0",
        "avg_price": "23489.7",
        "state": "wait",
        "market": "btcusdc",
        "created_at": "2022-07-20T00:59:57+02:00",
        "updated_at": "2022-07-20T01:47:44+02:00",
        "origin_volume": "2.1772",
        "remaining_volume": "0.0083",
        "executed_volume": "2.1689",
        "maker_fee": "0.001",
        "taker_fee": "0.001",
        "trades_count": 22
      }
    ]
  }
}
*/
type OrderBookResponse struct {
	Status  Status           `json:"status"`
	Message string           `json:"message"`
	Payload OrderBookPayload `json:"payload"`
}

/*
Retrieve the latest asks and bids created within the tanX orderbook in a specific market using this endpoint
*/
func (c *Client) OrderBook(ctx context.Context, market string, options OrderBookOptions) (OrderBookResponse, error) {
	if len(market) == 0 {
		return OrderBookResponse{}, ErrMarketNotProvided
	}

	orderBookURL := *c.orderBookURL
	params := orderBookURL.Query()

	params.Set("market", market)

	if options.AskLimit != 0 {
		params.Set("ask_limit", fmt.Sprintf("%d", options.AskLimit))
	}

	if options.BidLimit != 0 {
		params.Set("bid_limit", fmt.Sprintf("%d", options.BidLimit))
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
	if orderBookResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return OrderBookResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderBookResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return OrderBookResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderBookResponse.Message),
			}
		}

		return OrderBookResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, orderBookResponse.Message)

	} else if err != nil {
		return OrderBookResponse{}, &ErrJSONDecoding{Err: err}
	}

	return orderBookResponse, nil
}
