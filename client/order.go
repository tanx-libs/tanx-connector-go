package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tanx-libs/tanx-connector-go/crypto-cpp/build/Release/src/starkware/crypto/ffi/crypto_lib"
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
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Payload OrderNoncePayload `json:"payload"`
}

type OrderOptions struct {
	Market  string  `json:"market"`
	OrdType string  `json:"ord_type"`
	Price   float64 `json:"price"`
	Side    string  `json:"side"`
	Volume  float64 `json:"volume"`
}

func (c *Client) OrderNonce(ctx context.Context, opt OrderOptions) (OrderNonceResponse, error) {
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
	if err != nil {
		return OrderNonceResponse{}, fmt.Errorf("error: %s", orderNonceResponse.Message)
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

type OrderCreateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
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
	} `json:"payload"`
}

func (c *Client) OrderCreate(ctx context.Context, starkPrivateKey string, opt OrderOptions) (OrderCreateResponse, error) {
	orderNonceResponse, err := c.OrderNonce(ctx, opt)
	if err != nil {
		return OrderCreateResponse{}, err
	}

	if starkPrivateKey[:1] != "0x" {
		starkPrivateKey = "0x" + starkPrivateKey
	}

	r, s := crypto_lib.Sign(starkPrivateKey, orderNonceResponse.Payload.MsgHash, "0x1")

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
	if err != nil {
		return OrderCreateResponse{}, fmt.Errorf("error: %s", orderCreateResponse.Message)
	}

	return orderCreateResponse, nil
}

// type OrdersGetResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Payload struct {
// 		ID              int    `json:"id"`
// 		UUID            string `json:"uuid"`
// 		Side            string `json:"side"`
// 		OrdType         string `json:"ord_type"`
// 		Price           string `json:"price"`
// 		AvgPrice        string `json:"avg_price"`
// 		State           string `json:"state"`
// 		Market          string `json:"market"`
// 		CreatedAt       string `json:"created_at"`
// 		UpdatedAt       string `json:"updated_at"`
// 		OriginVolume    string `json:"origin_volume"`
// 		RemainingVolume string `json:"remaining_volume"`
// 		ExecutedVolume  string `json:"executed_volume"`
// 		MakerFee        string `json:"maker_fee"`
// 		TakerFee        string `json:"taker_fee"`
// 		TradesCount     int    `json:"trades_count"`

// 		// todo
// 		Trades []interface{} `json:"trades"`
// 	} `json:"payload"`
// }

// func (c *Client) OrdersGet(ctx context.Context, id int) (OrdersGetResponse, error) {
// 	path := fmt.Sprintf("%s/%d/", c.orderURL.String(), id)

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
// 	if err != nil {
// 		return OrdersGetResponse{}, err
// 	}

// 	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return OrdersGetResponse{}, err
// 	}
// 	defer resp.Body.Close()

// 	// debug
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return OrdersGetResponse{}, err
// 	}

// 	log.Println("OrdersGet: ", string(body))

// 	var ordersGetResponse OrdersGetResponse
// 	err = json.NewDecoder(resp.Body).Decode(&ordersGetResponse)
// 	if err != nil {
// 		return OrdersGetResponse{}, fmt.Errorf("error: no orders found")
// 	}

// 	return ordersGetResponse, nil
// }

// type OrdersListResponse struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// 	Payload []struct {
// 		ID              int    `json:"id"`
// 		UUID            string `json:"uuid"`
// 		Side            string `json:"side"`
// 		OrdType         string `json:"ord_type"`
// 		Price           string `json:"price"`
// 		AvgPrice        string `json:"avg_price"`
// 		State           string `json:"state"`
// 		Market          string `json:"market"`
// 		CreatedAt       string `json:"created_at"`
// 		UpdatedAt       string `json:"updated_at"`
// 		OriginVolume    string `json:"origin_volume"`
// 		RemainingVolume string `json:"remaining_volume"`
// 		ExecutedVolume  string `json:"executed_volume"`
// 		MakerFee        string `json:"maker_fee"`
// 		TakerFee        string `json:"taker_fee"`
// 		TradesCount     int    `json:"trades_count"`
// 	} `json:"payload"`
// }

// type OrderListOptions struct {
// 	Limit     int    `json:"limit"`
// 	Page      int    `json:"page"`
// 	Market    string `json:"market"`
// 	OrdType   string `json:"ord_type"`
// 	State     string `json:"state"`
// 	BaseUnit  string `json:"base_unit"`
// 	QuoteUnit string `json:"quote_unit"`
// 	StartTime int    `json:"start_time"`
// 	EndTime   int    `json:"end_time"`
// 	Side      string `json:"side"`
// }

// func (c *Client) OrdersList(ctx context.Context, opt OrderListOptions) (OrdersListResponse, error) {

// 	url := *c.orderURL
// 	params := url.Query()

// 	if opt.Limit != 0 {
// 		params.Add("limit", fmt.Sprintf("%d", opt.Limit))
// 	}

// 	if opt.Page != 0 {
// 		params.Add("page", fmt.Sprintf("%d", opt.Page))
// 	}

// 	if len(opt.Market) != 0 {
// 		params.Add("market", opt.Market)
// 	}

// 	if len(opt.OrdType) != 0 {
// 		params.Add("ord_type", opt.OrdType)
// 	}

// 	if len(opt.State) != 0 {
// 		params.Add("state", opt.State)
// 	}

// 	if len(opt.BaseUnit) != 0 {
// 		params.Add("base_unit", opt.BaseUnit)
// 	}

// 	if len(opt.QuoteUnit) != 0 {
// 		params.Add("quote_unit", opt.QuoteUnit)
// 	}

// 	if opt.StartTime != 0 {
// 		params.Add("start_time", fmt.Sprintf("%d", opt.StartTime))
// 	}

// 	if opt.EndTime != 0 {
// 		params.Add("end_time", fmt.Sprintf("%d", opt.EndTime))
// 	}

// 	if len(opt.Side) != 0 {
// 		params.Add("side", opt.Side)
// 	}

// 	url.RawQuery = params.Encode()

// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
// 	if err != nil {
// 		return OrdersListResponse{}, err
// 	}

// 	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return OrdersListResponse{}, err
// 	}
// 	defer resp.Body.Close()

// 	var ordersListResponse OrdersListResponse
// 	err = json.NewDecoder(resp.Body).Decode(&ordersListResponse)
// 	if err != nil {
// 		return OrdersListResponse{}, fmt.Errorf("error: no orders found")
// 	}

// 	return ordersListResponse, nil
// }
