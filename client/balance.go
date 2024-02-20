package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type BalancePayload struct {
	Currency       string `json:"currency"`
	Balance        string `json:"balance"`
	Locked         string `json:"locked"`
	DepositAddress string `json:"deposit_address"`
}

type BalanceResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Payload BalancePayload `json:"payload"`
}

// Retrieve details of a specific user’s balance
func (c *Client) Balance(ctx context.Context, currency string) (BalanceResponse, error) {
	if currency == "" {
		return BalanceResponse{}, ErrMarketNotProvided
	}

	url := *c.balanceURL
	params := url.Query()
	params.Set("currency", currency)
	url.RawQuery = params.Encode()

	log.Println(url.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return BalanceResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return BalanceResponse{}, err
	}
	defer resp.Body.Close()

	// Caution: Server is responsing with status code 500
	// ----debug-----
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return BalanceResponse{}, err
	}
	fmt.Println(string(bodyBytes))
	fmt.Println(resp.StatusCode)
	// ---------------

	var balanceResponse BalanceResponse
	err = json.NewDecoder(resp.Body).Decode(&balanceResponse)
	if err != nil {
		return BalanceResponse{}, err
	}

	return balanceResponse, nil
}
