package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Ticker struct {
	Low                string `json:"low"`
	High               string `json:"high"`
	Open               string `json:"open"`
	Last               string `json:"last"`
	Volume             string `json:"volume"`
	Amount             string `json:"amount"`
	Vol                string `json:"vol"`
	AvgPrice           string `json:"avg_price"`
	PriceChangePercent string `json:"price_change_percent"`
	At                 string `json:"at"`
}

type TickerPayload struct {
	At     string `json:"at"`
	Ticker Ticker `json:"ticker"`
}

/*
{
  "status": "success",
  "message": "Retrieval Successful",
  "payload": {
    "ethusdc": {
      "at": "1660636224",
      "ticker": {
        "at": "1660636224",
        "avg_price": "0.0",
        "high": "0.0",
        "last": "1697.0",
        "low": "0.0",
        "open": "0.0",
        "price_change_percent": "+0.00%",
        "volume": "0.0",
        "amount": "0.0"
      }
    }
  }
}
*/
type TickerResponse struct {
	Status  Status        `json:"status"`
	Message string        `json:"message"`
	Payload TickerPayload `json:"payload"`
}


/*
{
  "status": "success",
  "message": "Retrieval Successful",
  "payload": {
    "ethusdc": {
      "at": "1660636224",
      "ticker": {
        "at": "1660636224",
        "avg_price": "0.0",
        "high": "0.0",
        "last": "1697.0",
        "low": "0.0",
        "open": "0.0",
        "price_change_percent": "+0.00%",
        "volume": "0.0",
        "amount": "0.0"
      }
    }
  }
}
*/
type AllTickerResponse struct {
	Status  Status                   `json:"status"`
	Message string                   `json:"message"`
	Payload map[string]TickerPayload `json:"payload"`
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
	if tickerResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return TickerResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(tickerResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return TickerResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(tickerResponse.Message),
			}
		}

		return TickerResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, tickerResponse.Message)

	} else if err != nil {
		return TickerResponse{}, &ErrJSONDecoding{Err: err}
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
	if allTickerResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return AllTickerResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(allTickerResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return AllTickerResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(allTickerResponse.Message),
			}
		}

		return AllTickerResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, allTickerResponse.Message)

	} else if err != nil {
		return AllTickerResponse{}, &ErrJSONDecoding{Err: err}
	}

	return allTickerResponse, nil
}
