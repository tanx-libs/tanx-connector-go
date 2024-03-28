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

type InternalTransferInitiateRequest struct {
	OrganizationKey    string   `json:"organization_key"`    // required
	ApiKey             string   `json:"api_key"`             // required
	ClientReferenceId  string   `json:"client_reference_id"` // optional
	Currency           Currency `json:"currency"`            // required
	Amount             float64  `json:"amount"`              // required
	DestinationAddress string   `json:"destination_address"` // required
}

type InternalTransferInitiatePayload struct {
	MsgHash string `json:"msg_hash"`
	Nonce   int    `json:"nonce"`
}

type InternalTransferInitiateResponse struct {
	Status  Status                          `json:"status"`
	Message string                          `json:"message"`
	Payload InternalTransferInitiatePayload `json:"payload"`
}

func (c *Client) InternalTransferInitiate(ctx context.Context, opt InternalTransferInitiateRequest) (InternalTransferInitiateResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return InternalTransferInitiateResponse{}, ErrNotLoggedIn
	}

	reqBody, err := json.Marshal(opt)
	if err != nil {
		return InternalTransferInitiateResponse{}, err
	}

	log.Println(c.internalTransferInitiateURL.String())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.internalTransferInitiateURL.String(), bytes.NewReader(reqBody))
	if err != nil {
		return InternalTransferInitiateResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return InternalTransferInitiateResponse{}, err
	}
	defer resp.Body.Close()

	var internalTransferInitiateResponse InternalTransferInitiateResponse
	err = json.NewDecoder(resp.Body).Decode(&internalTransferInitiateResponse)
	log.Println(1)

	if internalTransferInitiateResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return InternalTransferInitiateResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferInitiateResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return InternalTransferInitiateResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferInitiateResponse.Message),
			}
		}

		return InternalTransferInitiateResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, internalTransferInitiateResponse.Message)

	} else if err != nil {
		return InternalTransferInitiateResponse{}, &ErrJSONDecoding{Err: err}
	}

	return internalTransferInitiateResponse, nil
}

// internal transfer process
type Signature struct {
	R string `json:"r"`
	S string `json:"s"`
}

type InternalTransferProcessRequest struct {
	OrganizationKey string    `json:"organization_key"`
	ApiKey          string    `json:"api_key"`
	Signature       Signature `json:"signature"`
	Nonce           int       `json:"nonce"`
	MsgHash         string    `json:"msg_hash"`
}

type InternalTransferProcessResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		ClientReferenceId  string `json:"client_reference_id"`
		Amount             string `json:"amount"`
		Currency           string `json:"currency"`
		FromAddress        string `json:"from_address"`
		DestinationAddress string `json:"destination_address"`
		Status             string `json:"status"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
	} `json:"payload"`
}

func (c *Client) InternalTransferProcess(ctx context.Context, opt InternalTransferProcessRequest) (InternalTransferProcessResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return InternalTransferProcessResponse{}, ErrNotLoggedIn
	}

	reqBody, err := json.Marshal(opt)
	if err != nil {
		return InternalTransferProcessResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.internalTransferProcessURL.String(), bytes.NewReader(reqBody))
	if err != nil {
		return InternalTransferProcessResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return InternalTransferProcessResponse{}, err
	}
	defer resp.Body.Close()

	var internalTransferProcessResponse InternalTransferProcessResponse
	err = json.NewDecoder(resp.Body).Decode(&internalTransferProcessResponse)
	if internalTransferProcessResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return InternalTransferProcessResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferProcessResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return InternalTransferProcessResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferProcessResponse.Message),
			}
		}

		return InternalTransferProcessResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, internalTransferProcessResponse.Message)

	} else if err != nil {
		return InternalTransferProcessResponse{}, &ErrJSONDecoding{Err: err}
	}

	return internalTransferProcessResponse, nil
}

// internal transfer create
func (c *Client) InternalTransferCreate(ctx context.Context, starkPrivateKey string, opt InternalTransferInitiateRequest) (InternalTransferProcessResponse, error) {
	internalTransferInitiateResponse, err := c.InternalTransferInitiate(ctx, opt)
	if err != nil {
		return InternalTransferProcessResponse{}, err
	}

	if starkPrivateKey[:2] != "0x" {
		starkPrivateKey = "0x" + starkPrivateKey
	}

	msgHash := internalTransferInitiateResponse.Payload.MsgHash

	r, s := crypto_cpp.Sign(starkPrivateKey, msgHash, "0x1")

	internalTransferProcessRequest := InternalTransferProcessRequest{
		OrganizationKey: opt.OrganizationKey,
		ApiKey:          opt.ApiKey,
		Signature: Signature{
			R: r,
			S: s,
		},
		Nonce:   internalTransferInitiateResponse.Payload.Nonce,
		MsgHash: internalTransferInitiateResponse.Payload.MsgHash,
	}

	internalTransferProcessResponse, err := c.InternalTransferProcess(ctx, internalTransferProcessRequest)
	if err != nil {
		return InternalTransferProcessResponse{}, err
	}

	return internalTransferProcessResponse, nil
}

// internal transfer get
type InternalTransferGetResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		ClientReferenceId  string `json:"client_reference_id"`
		Amount             string `json:"amount"`
		Currency           string `json:"currency"`
		FromAddress        string `json:"from_address"`
		DestinationAddress string `json:"destination_address"`
		Status             string `json:"status"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
	} `json:"payload"`
}

func (c *Client) InternalTransferGet(ctx context.Context, clientReferenceId string) (InternalTransferGetResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return InternalTransferGetResponse{}, ErrNotLoggedIn
	}

	url := c.internalTransferGetURL.JoinPath(clientReferenceId)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return InternalTransferGetResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return InternalTransferGetResponse{}, err
	}
	defer resp.Body.Close()

	var internalTransferGetResponse InternalTransferGetResponse
	err = json.NewDecoder(resp.Body).Decode(&internalTransferGetResponse)

	if internalTransferGetResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return InternalTransferGetResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferGetResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return InternalTransferGetResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferGetResponse.Message),
			}
		}

		return InternalTransferGetResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, internalTransferGetResponse.Message)

	} else if err != nil {
		return InternalTransferGetResponse{}, &ErrJSONDecoding{Err: err}
	}

	return internalTransferGetResponse, nil
}

// internal transfer user
type InternalTransferUserResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		DestinationAddress string `json:"destination_address"`
		Exists             bool   `json:"exists"`
	} `json:"payload"`
}

type InternalTransferUserRequest struct {
	OrganizationKey    string `json:"organization_key"`
	ApiKey             string `json:"api_key"`
	DestinationAddress string `json:"destination_address"`
}

func (c *Client) InternalTransferUser(ctx context.Context, opt InternalTransferUserRequest) (InternalTransferUserResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return InternalTransferUserResponse{}, ErrNotLoggedIn
	}

	reqBody, err := json.Marshal(opt)
	if err != nil {
		return InternalTransferUserResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.internalTransferUserURL.String(), bytes.NewReader(reqBody))
	if err != nil {
		return InternalTransferUserResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return InternalTransferUserResponse{}, err
	}
	defer resp.Body.Close()

	var internalTransferUserResponse InternalTransferUserResponse
	err = json.NewDecoder(resp.Body).Decode(&internalTransferUserResponse)
	if internalTransferUserResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return InternalTransferUserResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferUserResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return InternalTransferUserResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferUserResponse.Message),
			}
		}

		return InternalTransferUserResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, internalTransferUserResponse.Message)

	} else if err != nil {
		return InternalTransferUserResponse{}, &ErrJSONDecoding{Err: err}
	}

	return internalTransferUserResponse, nil
}

// internal transfer list
type InternalTransferListResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		InternalTransfers []struct {
			ClientReferenceId  string `json:"client_reference_id"`
			Amount             string `json:"amount"`
			Currency           string `json:"currency"`
			FromAddress        string `json:"from_address"`
			DestinationAddress string `json:"destination_address"`
			Status             string `json:"status"`
			CreatedAt          string `json:"created_at"`
			UpdatedAt          string `json:"updated_at"`
		} `json:"internal_transfers"`
	} `json:"payload"`
}

type InternalTransferListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (c *Client) InternalTransferList(ctx context.Context, opt InternalTransferListRequest) (InternalTransferListResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return InternalTransferListResponse{}, ErrNotLoggedIn
	}

	reqBody, err := json.Marshal(opt)
	if err != nil {
		return InternalTransferListResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.internalTransferListURL.String(), bytes.NewReader(reqBody))
	if err != nil {
		return InternalTransferListResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return InternalTransferListResponse{}, err
	}
	defer resp.Body.Close()

	var internalTransferListResponse InternalTransferListResponse
	err = json.NewDecoder(resp.Body).Decode(&internalTransferListResponse)

	if internalTransferListResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return InternalTransferListResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferListResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return InternalTransferListResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(internalTransferListResponse.Message),
			}
		}

		return InternalTransferListResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, internalTransferListResponse.Message)

	} else if err != nil {
		return InternalTransferListResponse{}, &ErrJSONDecoding{Err: err}
	}

	return internalTransferListResponse, nil
}
