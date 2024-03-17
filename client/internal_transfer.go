package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tanx-libs/tanx-connector-go/crypto-cpp/build/Release/src/starkware/crypto/ffi/crypto_lib"
)

// create internal transfer
// todo currency ka issue hain
type InternalTransferInitiateRequest struct {
	OrganizationKey    string  `json:"organization_key"`    // required
	ApiKey             string  `json:"api_key"`             // required
	ClientReferenceId  string  `json:"client_reference_id"` // optional
	Currency           string  `json:"currency"`            // required
	Amount             float64 `json:"amount"`              // required
	DestinationAddress string  `json:"destination_address"` // required
}

type InternalTransferInitiateResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
		MsgHash string `json:"msg_hash"`
		Nonce   int    `json:"nonce"`
	} `json:"payload"`
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
	if err != nil {
		return InternalTransferInitiateResponse{}, fmt.Errorf("error: %s", internalTransferInitiateResponse.Message)
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
	Status  string `json:"status"`
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
	if err != nil {
		return InternalTransferProcessResponse{}, fmt.Errorf("error: %s", internalTransferProcessResponse.Message)
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

	r, s := crypto_lib.Sign(starkPrivateKey, msgHash, "0x1")

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
	Status  string `json:"status"`
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
	if err != nil {
		return InternalTransferGetResponse{}, fmt.Errorf("error: %s", internalTransferGetResponse.Message)
	}

	return InternalTransferGetResponse{}, nil
}

// internal transfer user
type InternalTransferUserResponse struct {
	Status  string `json:"status"`
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
	if err != nil {
		return InternalTransferUserResponse{}, fmt.Errorf("error: %s", internalTransferUserResponse.Message)
	}

	return internalTransferUserResponse, nil
}

// internal transfer list
type InternalTransferListResponse struct {
	Status  string `json:"status"`
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
	if err != nil {
		return InternalTransferListResponse{}, fmt.Errorf("error: %s", internalTransferListResponse.Message)
	}

	return internalTransferListResponse, nil
}
