package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Period int

const (
	PERIOD_1MIN  Period = 1
	PERIOD_5MIN  Period = 5
	PERIOD_15MIN Period = 15
	PERIOD_30MIN Period = 30
	PERIOD_1H    Period = 60
	PERIOD_2H    Period = 120
	PERIOD_4H    Period = 240
	PERIOD_6H    Period = 360
	PERIOD_12H   Period = 720
	PERIOD_1D    Period = 1440
	PERIOD_3D    Period = 4320
	PERIOD_1W    Period = 10080
)

type CandleStickOptions struct {
	Limit     int
	Period    Period
	StartTime int64
	EndTime   int64
}

/*
{
  "status": "success",
  "message": "Retrieval Successful",
  "payload": [
    [1660634520, 1697, 1697, 1697, 1697, 0],
    [1660634520, 1697, 1697, 1697, 1697, 0],
    [1660634520, 1697, 1697, 1697, 1697, 0]
  ]
}

Note: Kline payload item as array of numbers

Example: [1660634520, 0.0839, 0.0921, 0.0781, 0.0845, 0.5895]

1. Timestamp.

2. Open price.

3. Max price.

4. Min price.

5. Last price.

6. Period volume
*/
type CandleStickResponse struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Payload [][]float64 `json:"payload"`
}

/*
Retrieve K-line/Candlestick data for the specified by making use of this endpoint.
You can define the limit, period, start time and end time for the same.
*/
func (c *Client) CandleStick(ctx context.Context, market string, options CandleStickOptions) (CandleStickResponse, error) {
	if len(market) == 0 {
		return CandleStickResponse{}, ErrMarketNotProvided
	}

	candleStickURL := *c.candleStickURL
	params := candleStickURL.Query()

	params.Set("market", market)
	if options.Limit != 0 {
		params.Set("limit", fmt.Sprintf("%d", options.Limit))
	}

	if options.Period != 0 {
		params.Set("period", fmt.Sprintf("%d", options.Period))
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
	
	if candleStickResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return CandleStickResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(candleStickResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return CandleStickResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(candleStickResponse.Message),
			}
		}

		return CandleStickResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, candleStickResponse.Message)

	} else if err != nil {
		return CandleStickResponse{}, &ErrJSONDecoding{Err: err}
	}

	return candleStickResponse, nil
}
