package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Helper function for signing the payload using your ethereum private key
func SignPayload(payload string, privateKey string) (string, error) {
	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(payload), payload)
	hash := crypto.Keccak256Hash([]byte(message))

	signature, err := crypto.Sign(hash.Bytes(), ecdsaPrivateKey)
	if err != nil {
		return "", err
	}

	signature[64] += 27
	return hexutil.Encode(signature), nil
}

type NonceRequest struct {
	EthAddress string `json:"eth_address"`
}

type NonceResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload string `json:"payload"`
}

func (c *Client) Nonce(ctx context.Context, ethAddress string) (NonceResponse, error) {
	nonceRequest := NonceRequest{EthAddress: ethAddress}

	requestBody, err := json.Marshal(nonceRequest)
	if err != nil {
		return NonceResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.nonceURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return NonceResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return NonceResponse{}, err
	}
	defer resp.Body.Close()

	var nonceResponse NonceResponse
	err = json.NewDecoder(resp.Body).Decode(&nonceResponse)
	if err != nil {
		return NonceResponse{}, err
	}

	return nonceResponse, nil
}

type JWTRequest struct {
	EthAddress   string `json:"eth_address"`
	UserSignaure string `json:"user_signature"`
}

type Token struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}

type JWTPayload struct {
	UID       string `json:"uid"`
	Signature string `json:"signature"`
}

type JWTResponse struct {
	Status  string     `json:"status"`
	Message string     `json:"message"`
	Payload JWTPayload `json:"payload"`
	Token   Token      `json:"token"`
}

/*
Steps followed for login process:-

Step 1: A Nonce is a variable that is generated just once and can be used only one time.
Generation of a Nonce is the first step of the login process.
The payload received in this step will be required in the next one.

Step 2: Payload is signed using the user's private key.

Step 3: The signed payload is sent to the server along with the user's ethereum address. 
This basically logs the user in and returns a JWT token.

Note: jwt token along with the refresh token are already stored in the client object for future use.
So you don't need to store or pass it to any other function separately.
*/
func (c *Client) Login(ctx context.Context, ethAddress string, privateKey string) (NonceResponse, JWTResponse, error) {
	nonceResponse, err := c.Nonce(ctx, ethAddress)
	if err != nil {
		return NonceResponse{}, JWTResponse{}, err
	}

	signedPayload, err := SignPayload(nonceResponse.Payload, privateKey)
	if err != nil {
		return NonceResponse{}, JWTResponse{}, err
	}

	jwtRequest := JWTRequest{
		EthAddress:   ethAddress,
		UserSignaure: signedPayload,
	}

	requestBody, err := json.Marshal(jwtRequest)
	if err != nil {
		return NonceResponse{}, JWTResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.loginURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return NonceResponse{}, JWTResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return NonceResponse{}, JWTResponse{}, err
	}
	defer resp.Body.Close()

	var jwtResponse JWTResponse
	err = json.NewDecoder(resp.Body).Decode(&jwtResponse)
	if err != nil {
		return NonceResponse{}, JWTResponse{}, err
	}

	c.jwtToken = jwtResponse.Token.Access
	c.refreshToken = jwtResponse.Token.Refresh
	return nonceResponse, jwtResponse, nil
}

type RefreshTokenRequest struct {
	Refresh string `json:"refresh"`
}


func (c *Client) Relogin(ctx context.Context) (JWTResponse, error) {
	refreshRequest := RefreshTokenRequest{
		Refresh: c.refreshToken,
	}

	requestBody, err := json.Marshal(refreshRequest)
	if err != nil {
		return JWTResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.refreshTokenURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return JWTResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return JWTResponse{}, err
	}
	defer resp.Body.Close()

	var refreshResponse JWTResponse
	err = json.NewDecoder(resp.Body).Decode(&refreshResponse)
	if err != nil {
		return JWTResponse{}, err
	}

	c.jwtToken = refreshResponse.Token.Access
	c.refreshToken = refreshResponse.Token.Refresh

	return refreshResponse, nil
}
