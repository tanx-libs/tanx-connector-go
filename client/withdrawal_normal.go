package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tanx-libs/tanx-connector-go/crypto_cpp"
)

type StartNormalWithdrawalPaylaod struct {
	Nonce   int    `json:"nonce"`
	MsgHash string `json:"msg_hash"`
}

type StartNormalWithdrawalResponse struct {
	Status  Status                       `json:"status"`
	Message string                       `json:"message"`
	Payload StartNormalWithdrawalPaylaod `json:"payload"`
}

type StartNormalWithdrawalRequest struct {
	Amount   float64  `json:"amount"`
	Currency Currency `json:"token_id"`
}

func (c *Client) StartNormalWithdrawal(ctx context.Context, opt StartNormalWithdrawalRequest) (StartNormalWithdrawalResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return StartNormalWithdrawalResponse{}, err
	}

	requestBody, err := json.Marshal(opt)
	if err != nil {
		return StartNormalWithdrawalResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.startNormalWithdrawalURL.String(), bytes.NewBuffer(requestBody))
	if err != nil {
		return StartNormalWithdrawalResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return StartNormalWithdrawalResponse{}, err
	}
	defer resp.Body.Close()

	var startNormalWithdrawalResponse StartNormalWithdrawalResponse
	err = json.NewDecoder(resp.Body).Decode(&startNormalWithdrawalResponse)
	if startNormalWithdrawalResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return StartNormalWithdrawalResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(startNormalWithdrawalResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return StartNormalWithdrawalResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(startNormalWithdrawalResponse.Message),
			}
		}

		return StartNormalWithdrawalResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, startNormalWithdrawalResponse.Message)

	} else if err != nil {
		return StartNormalWithdrawalResponse{}, &ErrJSONDecoding{Err: err}
	}

	return startNormalWithdrawalResponse, nil
}

type ValidateWithdrawalResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		ID                int    `json:"id"`
		Amount            string `json:"amount"`
		TokenID           string `json:"token_id"`
		CreatedAt         string `json:"created_at"`
		TransactionStatus string `json:"transaction_status"`
		Extras            struct {
			Errors          []interface{} `json:"errors"`
			ExpTimestamp    int           `json:"exp_timestamp"`
			QuantisedAmount int           `json:"quantised_amount"`
		} `json:"extras"`
	} `json:"payload"`
}

type WSignature struct {
	R             string `json:"r"`
	S             string `json:"s"`
	RecoveryParam int    `json:"recoveryParam"`
}

type ValidateNormalWithdrawalRequest struct {
	MessageHash string     `json:"msg_hash"`
	Signature   WSignature `json:"signature"`
	Nonce       int        `json:"nonce"`
}

func (c *Client) ValidateNormalWithdrawal(ctx context.Context, opt ValidateNormalWithdrawalRequest) (ValidateWithdrawalResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	requestBody, err := json.Marshal(opt)
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.validateNormalWithdrawalURL.String(), bytes.NewBuffer(requestBody))
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

	var validateWithdrawalResponse ValidateWithdrawalResponse
	err = json.NewDecoder(resp.Body).Decode(&validateWithdrawalResponse)

	if validateWithdrawalResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return ValidateWithdrawalResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(validateWithdrawalResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return ValidateWithdrawalResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(validateWithdrawalResponse.Message),
			}
		}

		return ValidateWithdrawalResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, validateWithdrawalResponse.Message)

	} else if err != nil {
		return ValidateWithdrawalResponse{}, &ErrJSONDecoding{Err: err}
	}

	return validateWithdrawalResponse, nil
}

// initiateNormalWithdrawal initiates a normal withdrawal
func (c *Client) InitiateNormalWithdrawal(ctx context.Context, starkPrivateKey string, amount float64, currency Currency) (ValidateWithdrawalResponse, error) {
	if amount <= 0 {
		return ValidateWithdrawalResponse{}, ErrInvalidAmount
	}

	err := c.CheckAuth()
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	// todo check if currency there or not
	coinStatus, err := c.getCoinStatus(ctx)
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}
	log.Println("coinstatus", coinStatus)

	_, err = getCoinStatusPayload(currency, coinStatus.Payload)
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	respStart, err := c.StartNormalWithdrawal(ctx, StartNormalWithdrawalRequest{
		Amount:   amount,
		Currency: currency,
	})
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}
	log.Println("respStart", respStart)

	if starkPrivateKey[:2] != "0x" {
		starkPrivateKey = "0x" + starkPrivateKey
	}

	log.Println("respStart.Payload.MsgHash", respStart.Payload.MsgHash)

	numBig := new(big.Int)
	numBig.SetString(respStart.Payload.MsgHash, 10)
	msgHash := hexutil.EncodeBig(numBig)
	log.Println("msgHash", msgHash)

	r, s := crypto_cpp.Sign(starkPrivateKey, msgHash, "0x1")
	log.Println("r", r)
	log.Println("s", s)

	respValidate, err := c.ValidateNormalWithdrawal(ctx, ValidateNormalWithdrawalRequest{
		MessageHash: msgHash[2:],
		Signature: WSignature{
			R: r,
			S: s,
		},
		Nonce: respStart.Payload.Nonce,
	})
	if err != nil {
		return ValidateWithdrawalResponse{}, err
	}

	return respValidate, nil
}

func (c *Client) GetPendingNormalWithdrawalAmountByCoin(ctx context.Context, rpcURL string, userPublicEthAddress string, currency Currency) (string, error) {
	err := c.CheckAuth()
	if err != nil {
		return "", err
	}

	err = c.ethereumInit(ctx, rpcURL)
	if err != nil {
		return "", err
	}

	coinStatusPayload, err := getCoinStatusPayload(currency, c.coinStatus.Payload)
	if err != nil {
		return "", err
	}
	starkAssetID := coinStatusPayload.StarkAssetID
	blockchainDecimal := coinStatusPayload.BlockchainDecimal

	ethAddressBigInt, ok := new(big.Int).SetString(userPublicEthAddress[2:], 16)
	if !ok {
		return "", err
	}

	starkAssetIDBigInt, ok := new(big.Int).SetString(starkAssetID[2:], 16)
	if !ok {
		return "", err
	}

	balance, err := c.starkexContract.GetWithdrawalBalance(
		&bind.CallOpts{
			From: common.HexToAddress(userPublicEthAddress),
		},
		ethAddressBigInt,
		starkAssetIDBigInt,
	)
	if err != nil {
		return "", err
	}
	log.Printf("balance: %v", balance)

	bd, err := strconv.Atoi(blockchainDecimal)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", ToDecimal(balance, bd)), nil
}

func (c *Client) CompleteNormalWithdrawal(ctx context.Context, rpcURL string, ethPrivateKey string, userPublicEthAddress string, currency Currency) (string, error) {
	err := c.CheckAuth()
	if err != nil {
		return "", err
	}

	err = c.ethereumInit(ctx, rpcURL)
	if err != nil {
		return "", err
	}

	coinStatusPayload, err := getCoinStatusPayload(currency, c.coinStatus.Payload)
	if err != nil {
		return "", err
	}

	starkAssetID := coinStatusPayload.StarkAssetID

	ethAddressBigInt, ok := new(big.Int).SetString(userPublicEthAddress[2:], 16)
	if !ok {
		return "", err
	}

	starkAssetIDBigInt, ok := new(big.Int).SetString(starkAssetID[2:], 16)
	if !ok {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return "", err
	}

	signerFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		chainID, err := c.ethClient.ChainID(context.Background())
		if err != nil {
			return nil, err
		}

		signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}

	opt, err := getTransactionOpt(context.Background(), c.ethClient, ethPrivateKey, signerFn, 0, 0)
	if err != nil {
		return "", err
	}

	res, err := c.starkexContract.Withdraw(opt, ethAddressBigInt, starkAssetIDBigInt)
	if err != nil {
		return "", err
	}

	log.Printf("res: %v", res)

	return res.Hash().Hex(), nil
}

type ListNormalWithdrawalPayload struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		ID                int    `json:"id"`
		Amount            string `json:"amount"`
		TokenID           string `json:"token_id"`
		CreatedAt         string `json:"created_at"`
		TransactionStatus string `json:"transaction_status"`
		Extras            struct {
			Errors          []interface{} `json:"errors"`
			ExpTimestamp    int           `json:"exp_timestamp"`
			QuantisedAmount int           `json:"quantised_amount"`
		} `json:"extras"`
	} `json:"results"`
}

type ListWithdrawalResponse struct {
	Status  Status                      `json:"status"`
	Message string                      `json:"message"`
	Payload ListNormalWithdrawalPayload `json:"payload"`
}

type ListWithdrawalRequest struct {
	Page    int    `json:"page"`
	Network Network `json:"network"`
}

func (c *Client) ListNormalWithdrawal(ctx context.Context, opt ListWithdrawalRequest) (ListWithdrawalResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return ListWithdrawalResponse{}, err
	}

	url := *c.listNormalWithdrawalURL
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
