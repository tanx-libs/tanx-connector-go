package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tanx-libs/tanx-connector-go/crypto_cpp"
)

type StartFastWithdrawalResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		FastWithdrawalID int    `json:"fastwithdrawal_withdrawal_id"`
		MsgHash          string `json:"msg_hash"`
	} `json:"payload"`
}

type StartFastWithdrawalRequest struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"token_id"`
	Network  Network  `json:"network"`
}

func (c *Client) StartFastWithdrawal(ctx context.Context, opt StartFastWithdrawalRequest) (StartFastWithdrawalResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return StartFastWithdrawalResponse{}, err
	}

	requestBody, err := json.Marshal(opt)
	if err != nil {
		return StartFastWithdrawalResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.startFastWithdrawalURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return StartFastWithdrawalResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return StartFastWithdrawalResponse{}, err
	}
	defer resp.Body.Close()

	var startFastWithdrawalResponse StartFastWithdrawalResponse
	err = json.NewDecoder(resp.Body).Decode(&startFastWithdrawalResponse)

	if startFastWithdrawalResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return StartFastWithdrawalResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(startFastWithdrawalResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return StartFastWithdrawalResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(startFastWithdrawalResponse.Message),
			}
		}

		return StartFastWithdrawalResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, startFastWithdrawalResponse.Message)

	} else if err != nil {
		return StartFastWithdrawalResponse{}, &ErrJSONDecoding{Err: err}
	}

	return startFastWithdrawalResponse, nil
}

type ProcessFastWithdrawalRequest struct {
	MsgHash          string     `json:"msg_hash"`
	Signature        WSignature `json:"signature"`
	FastWithdrawalID int        `json:"fastwithdrawal_withdrawal_id"`
}

func (c *Client) ProcessFastWithdrawal(ctx context.Context, opt ProcessFastWithdrawalRequest) (ValidateWithdrawalResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	requestBody, err := json.Marshal(opt)
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.processFastWithdrawalURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}
	defer resp.Body.Close()

	var processFastWithdrawalResponse ValidateWithdrawalResponse
	err = json.NewDecoder(resp.Body).Decode(&processFastWithdrawalResponse)

	if processFastWithdrawalResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return ValidateWithdrawalResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(processFastWithdrawalResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return ValidateWithdrawalResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(processFastWithdrawalResponse.Message),
			}
		}

		return ValidateWithdrawalResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, processFastWithdrawalResponse.Message)

	} else if err != nil {
		return ValidateWithdrawalResponse{}, &ErrJSONDecoding{Err: err}
	}

	return processFastWithdrawalResponse, nil
}

func (c *Client) FastWithdrawal(ctx context.Context, starkPrivateKey string, amount float64, currency Currency, network Network) (ValidateWithdrawalResponse, error) {
	if amount <= 0 {
		return ValidateWithdrawalResponse{}, ErrInvalidAmount
	}

	err := c.CheckAuth()
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	if network == POLYGON {
		networkConfig, err := c.GetNetworkConfig(ctx)
		if err != nil {
			return ValidateWithdrawalResponse{}, err
		}

		polygonConfig := networkConfig.Payload.NetworkConfig[network]
		allowedTokens := polygonConfig.AllowedTokensForDeposit

		if !isPresent(allowedTokens, currency) {
			return ValidateWithdrawalResponse{}, ErrCoinNotFound
		}

	} else {
		coinStatus, err := c.getCoinStatus(ctx)
		if err != nil {
			return ValidateWithdrawalResponse{}, err
		}

		_, err = getCoinStatusPayload(currency, coinStatus.Payload)
		if err != nil {
			return ValidateWithdrawalResponse{}, err
		}

		log.Println("coinstatus", coinStatus)
	}

	// todo
	respStart, err := c.StartFastWithdrawal(ctx, StartFastWithdrawalRequest{
		Amount:   amount,
		Currency: currency,
		Network:  network,
	})
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}
	log.Printf("respStart %+v", respStart)

	if starkPrivateKey[:2] != "0x" {
		starkPrivateKey = "0x" + starkPrivateKey
	}

	log.Printf("msg_hash: %+v", respStart.Payload.MsgHash)

	r, s := crypto_cpp.Sign(starkPrivateKey, respStart.Payload.MsgHash, "0x1")
	log.Println("r", r)
	log.Println("s", s)

	respValidate, err := c.ProcessFastWithdrawal(ctx, ProcessFastWithdrawalRequest{
		MsgHash: respStart.Payload.MsgHash,
		Signature: WSignature{
			R: r,
			S: s,
		},
		FastWithdrawalID: respStart.Payload.FastWithdrawalID,
	})
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	log.Println("respValidate", respValidate)
	return respValidate, nil
}

func (c *Client) ListFastWithdrawal(ctx context.Context, opt ListWithdrawalRequest) (ListWithdrawalResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return ListWithdrawalResponse{}, err
	}

	url := *c.listFastWithdrawalURL
	params := url.Query()

	if opt.Page != 0 {
		params.Set("page", fmt.Sprintf("%d", opt.Page))
	}

	if len(opt.Network) != 0 {
		params.Set("network", string(opt.Network))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return ListWithdrawalResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ListWithdrawalResponse{}, err
	}
	defer resp.Body.Close()

	var listWithdrawalResponse ListWithdrawalResponse
	err = json.NewDecoder(resp.Body).Decode(&listWithdrawalResponse)

	if listWithdrawalResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return ListWithdrawalResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(listWithdrawalResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return ListWithdrawalResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(listWithdrawalResponse.Message),
			}
		}

		return ListWithdrawalResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, listWithdrawalResponse.Message)

	} else if err != nil {
		return ListWithdrawalResponse{}, &ErrJSONDecoding{Err: err}
	}

	return listWithdrawalResponse, nil
}
