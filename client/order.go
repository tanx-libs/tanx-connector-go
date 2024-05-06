package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tanx-libs/tanx-connector-go/crypto_cpp"
)

type OrderNonceRequest struct {
	Market  string  `json:"market"`
	OrdType string  `json:"ord_type"`
	Price   float64 `json:"price"`
	Side    string  `json:"side"`
	Volume  float64 `json:"volume"`
}

type OrderNoncePayload struct {
	Nonce   int    `json:"nonce"`
	MsgHash string `json:"msg_hash"`
}

type OrderNonceResponse struct {
	Status  Status            `json:"status"`
	Message string            `json:"message"`
	Payload OrderNoncePayload `json:"payload"`
}

func (c *Client) orderNonce(ctx context.Context, opt OrderNonceRequest) (OrderNonceResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return OrderNonceResponse{}, err
	}

	requestBody, err := json.Marshal(opt)
	if err != nil {
		return OrderNonceResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.orderNonceURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return OrderNonceResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return OrderNonceResponse{}, err
	}
	defer resp.Body.Close()

	var orderNonceResponse OrderNonceResponse
	err = json.NewDecoder(resp.Body).Decode(&orderNonceResponse)

	if orderNonceResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return OrderNonceResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderNonceResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return OrderNonceResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderNonceResponse.Message),
			}
		}

		return OrderNonceResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, orderNonceResponse.Message)

	} else if err != nil {
		return OrderNonceResponse{}, &ErrJSONDecoding{Err: err}
	}

	return orderNonceResponse, nil
}

type OrderCreateSignature struct {
	R string `json:"r"`
	S string `json:"s"`
}

type OrderCreateRequest struct {
	MsgHash   string               `json:"msg_hash"`
	Signature OrderCreateSignature `json:"signature"`
	Nonce     int                  `json:"nonce"`
}

type OrderPayload struct {
	ID              int    `json:"id"`
	OrderID         int    `json:"order_id"`
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

/*
	{
	  "status": "success",
	  "message": "Created Order Successfully",
	  "payload": {
	    "id": 904,
	    "uuid": "6c79cde4-1e8f-4446-9b4a-ba176d50f08b",
	    "side": "sell",
	    "ord_type": "limit",
	    "price": "29580.51",
	    "avg_price": "0.0",
	    "state": "pending",
	    "market": "btcusdt",
	    "created_at": "2022-05-27T12:36:57+02:00",
	    "updated_at": "2022-05-27T12:36:57+02:00",
	    "origin_volume": "0.015",
	    "remaining_volume": "0.015",
	    "executed_volume": "0.0",
	    "maker_fee": "0.0",
	    "taker_fee": "0.0",
	    "trades_count": 0
	  }
	}
*/
type OrderCreateResponse struct {
	Status  Status       `json:"status"`
	Message string       `json:"message"`
	Payload OrderPayload `json:"payload"`
}

// Create a new order
func (c *Client) OrderCreate(ctx context.Context, starkPrivateKey string, opt OrderNonceRequest) (OrderCreateResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return OrderCreateResponse{}, err
	}

	orderNonceResponse, err := c.orderNonce(ctx, opt)
	if err != nil {
		return OrderCreateResponse{}, err
	}

	if orderNonceResponse.Payload.MsgHash == "" {
		return OrderCreateResponse{}, fmt.Errorf("msg_hash is empty")
	}

	if starkPrivateKey[:2] != "0x" {
		starkPrivateKey = "0x" + starkPrivateKey
	}
	r, s := crypto_cpp.Sign(starkPrivateKey, orderNonceResponse.Payload.MsgHash, "0x1")

	orderCreateRequest := OrderCreateRequest{
		MsgHash: orderNonceResponse.Payload.MsgHash,
		Signature: OrderCreateSignature{
			R: r,
			S: s,
		},
		Nonce: orderNonceResponse.Payload.Nonce,
	}

	requestBody, err := json.Marshal(orderCreateRequest)
	if err != nil {
		return OrderCreateResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.orderCreateURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return OrderCreateResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return OrderCreateResponse{}, err
	}
	defer resp.Body.Close()

	var orderCreateResponse OrderCreateResponse
	err = json.NewDecoder(resp.Body).Decode(&orderCreateResponse)

	if orderCreateResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return OrderCreateResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderCreateResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return OrderCreateResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderCreateResponse.Message),
			}
		}

		return OrderCreateResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, orderCreateResponse.Message)

	} else if err != nil {
		return OrderCreateResponse{}, &ErrJSONDecoding{Err: err}
	}

	return orderCreateResponse, nil
}

/*
	{
	  "status": "success",
	  "message": "Order Retrieved Successfully",
	  "payload": {
	    "id": 23157052,
	    "uuid": "a0b17d1f-0f2c-4c93-ad40-2e71077de499",
	    "side": "buy",
	    "ord_type": "market",
	    "price": null,
	    "avg_price": "0.0",
	    "state": "cancel",
	    "market": "btcusdc",
	    "created_at": "2022-11-08T11:53:09+01:00",
	    "updated_at": "2022-11-08T11:53:14+01:00",
	    "origin_volume": "0.0051",
	    "remaining_volume": "0.0051",
	    "executed_volume": "0.0",
	    "maker_fee": "0.001",
	    "taker_fee": "0.001",
	    "trades_count": 0,
	    "trades": []
	  }
	}
*/
type OrderGetResponse struct {
	Status  Status       `json:"status"`
	Message string       `json:"message"`
	Payload OrderPayload `json:"payload"`
}

/*
Retrieve a single order’s details using this endpoint.
Please note that this is a Private 🔒 route which means it needs to be authorised by the account initiating this request.
*/
func (c *Client) OrderGet(ctx context.Context, id int) (OrderGetResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return OrderGetResponse{}, err
	}

	url := c.orderURL.JoinPath(fmt.Sprintf("/%d/", id))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return OrderGetResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return OrderGetResponse{}, err
	}
	defer resp.Body.Close()

	var ordersGetResponse OrderGetResponse
	err = json.NewDecoder(resp.Body).Decode(&ordersGetResponse)
	if ordersGetResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return OrderGetResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(ordersGetResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return OrderGetResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(ordersGetResponse.Message),
			}
		}

		return OrderGetResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, ordersGetResponse.Message)

	} else if err != nil {
		return OrderGetResponse{}, &ErrJSONDecoding{Err: err}
	}

	return ordersGetResponse, nil
}

type OrderListOptions struct {
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	Market    string `json:"market"`
	OrdType   string `json:"ord_type"`
	State     string `json:"state"`
	BaseUnit  string `json:"base_unit"`
	QuoteUnit string `json:"quote_unit"`
	StartTime int    `json:"start_time"`
	EndTime   int    `json:"end_time"`
	Side      string `json:"side"`
}

/*
	{
	  "status": "success",
	  "message": "Orders Retrieved Successfully",
	  "payload": [
	    {
	      "id": 1739939,
	      "uuid": "b658b41f-fe68-4561-801d-dfa94c3ca1b5",
	      "side": "buy",
	      "ord_type": "limit",
	      "price": "450.0",
	      "avg_price": "0.0",
	      "state": "cancel",
	      "market": "ethusdc",
	      "created_at": "2022-06-16T22:41:14+02:00",
	      "updated_at": "2022-06-16T22:43:30+02:00",
	      "origin_volume": "0.001",
	      "remaining_volume": "0.001",
	      "executed_volume": "0.0",
	      "maker_fee": "0.001",
	      "taker_fee": "0.001",
	      "trades_count": 0
	    },
	    {
	      "id": 1739940,
	      "uuid": "19bce015-c143-449f-af2b-fdcfa6384015",
	      "side": "buy",
	      "ord_type": "limit",
	      "price": "1134.0",
	      "avg_price": "1134.0",
	      "state": "done",
	      "market": "ethusdc",
	      "created_at": "2022-06-16T22:41:24+02:00",
	      "updated_at": "2022-06-16T22:41:49+02:00",
	      "origin_volume": "0.001",
	      "remaining_volume": "0.0",
	      "executed_volume": "0.001",
	      "maker_fee": "0.001",
	      "taker_fee": "0.001",
	      "trades_count": 1
	    }
	  ]
	}
*/
type OrdersListResponse struct {
	Status  Status         `json:"status"`
	Message string         `json:"message"`
	Payload []OrderPayload `json:"payload"`
}

/*
Retrieve details of all orders for a user using this endpoint.
Please note that this is a Private 🔒 route which means it needs to be authorized by the account initiating this request.
*/
func (c *Client) OrdersList(ctx context.Context, opt OrderListOptions) (OrdersListResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return OrdersListResponse{}, err
	}

	url := *c.orderURL
	params := url.Query()

	if opt.Limit != 0 {
		params.Add("limit", fmt.Sprintf("%d", opt.Limit))
	}

	if opt.Page != 0 {
		params.Add("page", fmt.Sprintf("%d", opt.Page))
	}

	if len(opt.Market) != 0 {
		params.Add("market", opt.Market)
	}

	if len(opt.OrdType) != 0 {
		params.Add("ord_type", opt.OrdType)
	}

	if len(opt.State) != 0 {
		params.Add("state", opt.State)
	}

	if len(opt.BaseUnit) != 0 {
		params.Add("base_unit", opt.BaseUnit)
	}

	if len(opt.QuoteUnit) != 0 {
		params.Add("quote_unit", opt.QuoteUnit)
	}

	if opt.StartTime != 0 {
		params.Add("start_time", fmt.Sprintf("%d", opt.StartTime))
	}

	if opt.EndTime != 0 {
		params.Add("end_time", fmt.Sprintf("%d", opt.EndTime))
	}

	if len(opt.Side) != 0 {
		params.Add("side", opt.Side)
	}

	url.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return OrdersListResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return OrdersListResponse{}, err
	}
	defer resp.Body.Close()

	var ordersListResponse OrdersListResponse
	err = json.NewDecoder(resp.Body).Decode(&ordersListResponse)

	if ordersListResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return OrdersListResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(ordersListResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return OrdersListResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(ordersListResponse.Message),
			}
		}

		return OrdersListResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, ordersListResponse.Message)

	} else if err != nil {
		return OrdersListResponse{}, &ErrJSONDecoding{Err: err}
	}

	return ordersListResponse, nil
}

type OrderCancelRequest struct {
	OrderID int `json:"order_id"`
}

/*
	{
	  "status": "success",
	  "message": "Cancelled Order Successfully",
	  "payload": {
	    "order_id": 3,
	    "uuid": "3626e3a1-c15f-4979-97f2-30ce8845578d",
	    "side": "buy",
	    "ord_type": "limit",
	    "price": "1663.0",
	    "avg_price": "1663.0",
	    "state": "wait",
	    "market": "ethusdc",
	    "created_at": "2022-08-05T12:20:22+02:00",
	    "updated_at": "2022-08-05T12:21:28+02:00",
	    "origin_volume": "0.003",
	    "remaining_volume": "0.0021",
	    "executed_volume": "0.0009",
	    "maker_fee": "0.001",
	    "taker_fee": "0.001",
	    "trades_count": 3
	  }
	}
*/
type OrderCancelResponse struct {
	Status  Status       `json:"status"`
	Message string       `json:"message"`
	Payload OrderPayload `json:"payload"`
}

/*
This endpoint is used to cancel a limit order which hasn’t already been executed at a given time.
Please note that this is a Private 🔒 route which means it needs to be authorised by the account initiating this request.
*/
func (c *Client) OrderCancel(ctx context.Context, orderID int) (OrderCancelResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return OrderCancelResponse{}, err
	}

	orderCancelRequest := OrderCancelRequest{
		OrderID: orderID,
	}

	requestBody, err := json.Marshal(orderCancelRequest)
	if err != nil {
		return OrderCancelResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.orderCancelURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return OrderCancelResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return OrderCancelResponse{}, err
	}
	defer resp.Body.Close()

	var orderCancelResponse OrderCancelResponse
	err = json.NewDecoder(resp.Body).Decode(&orderCancelResponse)

	if orderCancelResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return OrderCancelResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderCancelResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return OrderCancelResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(orderCancelResponse.Message),
			}
		}

		return OrderCancelResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, orderCancelResponse.Message)

	} else if err != nil {
		return OrderCancelResponse{}, &ErrJSONDecoding{Err: err}
	}

	return orderCancelResponse, nil
}


type TradesListOptions struct {
	Limit     int
	Page      int
	Market    string
	StartTime int
	EndTime   int
	OrderBy   string
}

type TradePayload struct {
	ID          int    `json:"id"`
	Price       string `json:"price"`
	Amount      string `json:"amount"`
	Total       string `json:"total"`
	FeeCurrency string `json:"fee_currency"`
	Fee         string `json:"fee"`
	FeeAmount   string `json:"fee_amount"`
	Market      string `json:"market"`
	CreatedAt   string `json:"created_at"`
	TakerType   string `json:"taker_type"`
	Side        string `json:"side"`
	OrderID     int    `json:"order_id"`
}

/*
{
  "status": "success",
  "message": "Trades Retrieved Successfully",
  "payload": [
    {
      "id": 7281,
      "price": "1134.0",
      "amount": "0.001",
      "total": "1.134",
      "fee_currency": "eth",
      "fee": "0.001",
      "fee_amount": "0.000001",
      "market": "ethusdc",
      "created_at": "2022-06-16T22:41:49+02:00",
      "taker_type": "sell",
      "side": "buy",
      "order_id": 1739940
    },
    {
      "id": 7280,
      "price": "1134.0",
      "amount": "0.001",
      "total": "1.134",
      "fee_currency": "eth",
      "fee": "0.001",
      "fee_amount": "0.000001",
      "market": "ethusdc",
      "created_at": "2022-06-16T22:37:56+02:00",
      "taker_type": "sell",
      "side": "buy",
      "order_id": 1739937
    }
  ]
}
*/
type TradesListResponse struct {
	Status  Status         `json:"status"`
	Message string         `json:"message"`
	Payload []TradePayload `json:"payload"`
}

/*
Retrieve details of all trades executed for a user by making use of certain filters.
Please note that this is a Private 🔒 route which means it needs to be authorized by the account initiating this request.
*/
func (c *Client) TradesList(ctx context.Context, opt TradesListOptions) (TradesListResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return TradesListResponse{}, err
	}

	url := *c.tradesListURL
	params := url.Query()

	if opt.Limit != 0 {
		params.Add("limit", fmt.Sprintf("%d", opt.Limit))
	}

	if opt.Page != 0 {
		params.Add("page", fmt.Sprintf("%d", opt.Page))
	}

	if len(opt.Market) != 0 {
		params.Add("market", opt.Market)
	}

	if opt.StartTime != 0 {
		params.Add("start_time", fmt.Sprintf("%d", opt.StartTime))
	}

	if opt.EndTime != 0 {
		params.Add("end_time", fmt.Sprintf("%d", opt.EndTime))
	}

	if len(opt.OrderBy) != 0 {
		params.Add("order_by", opt.OrderBy)
	}

	url.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return TradesListResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return TradesListResponse{}, err
	}
	defer resp.Body.Close()

	var tradesListResponse TradesListResponse
	err = json.NewDecoder(resp.Body).Decode(&tradesListResponse)
	if tradesListResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return TradesListResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(tradesListResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return TradesListResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(tradesListResponse.Message),
			}
		}

		return TradesListResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, tradesListResponse.Message)

	} else if err != nil {
		return TradesListResponse{}, &ErrJSONDecoding{Err: err}
	}

	return tradesListResponse, nil
}
