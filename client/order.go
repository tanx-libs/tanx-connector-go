package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
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

func (c *Client) OrderNonce(ctx context.Context, market string, ordType string, price float64, side string, volume float64) (OrderNonceResponse, error) {
	orderNonceRequest := OrderNonceRequest{
		Market:  market,
		OrdType: ordType,
		Price:   price,
		Side:    side,
		Volume:  volume,
	}

	requestBody, err := json.Marshal(orderNonceRequest)
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

func (c *Client) OrderCreate(ctx context.Context, starkPrivateKey string, market string, ordType string, price float64, side string, volume float64) (OrderCreateResponse, error) {
	orderNonceResponse, err := c.OrderNonce(ctx, market, ordType, price, side, volume)
	if err != nil {
		return OrderCreateResponse{}, err
	}

	nonce := fmt.Sprintf("0x%x", orderNonceResponse.Payload.Nonce)
	log.Println(nonce)
	
	r, s := crypto_lib.Sign(starkPrivateKey, orderNonceResponse.Payload.MsgHash, nonce)
	// log.Printf("r: %s, s: %s", r, s)

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

	// debug
	var orderCreateResponse OrderCreateResponse
	err = json.NewDecoder(resp.Body).Decode(&orderCreateResponse)
	if err != nil {
		return OrderCreateResponse{}, fmt.Errorf("error: %s", orderCreateResponse.Message)
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
type OrdersGetResponse struct {
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

		// todo
		Trades []interface{} `json:"trades"`
	} `json:"payload"`
}

/*
{baseurl}/sapi/v1/orders/23157052/
*/
// func (c *Client) OrdersGet(ctx context.Context, id int) (OrdersGetResponse, error) {

// 	req, err := http.NewRequestWithContext(ctx, , nil)
// 	if err != nil {
// 		return OrdersGetResponse{}, err
// 	}

// 	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return OrdersGetResponse{}, err
// 	}
// 	defer resp.Body.Close()

// 	var ordersGetResponse OrdersGetResponse
// 	err = json.NewDecoder(resp.Body).Decode(&ordersGetResponse)
// 	if err != nil {
// 		return OrdersGetResponse{}, fmt.Errorf("error: %s", ordersGetResponse.Message)
// 	}

// 	return ordersGetResponse, nil
// }
