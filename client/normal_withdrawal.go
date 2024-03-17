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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tanx-libs/tanx-connector-go/contract"
	"github.com/tanx-libs/tanx-connector-go/crypto-cpp/build/Release/src/starkware/crypto/ffi/crypto_lib"
)

/*
	{
	  "status": "success",
	  "message": "successfully initiated withdrawal",
	  "payload": {
	    "nonce": 74,
	    "msg_hash": "3444275198052187053566742646566392658438268899685086984873448032869787456539"
	  }
	}
*/
type StartNormalWithdrawalResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
		Nonce   int    `json:"nonce"`
		MsgHash string `json:"msg_hash"`
	} `json:"payload"`
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.startFastWithdrawalURL.String(), bytes.NewBuffer(requestBody))
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
	if err != nil {
		return StartNormalWithdrawalResponse{}, fmt.Errorf("error: %s", startNormalWithdrawalResponse.Message)
	}

	return startNormalWithdrawalResponse, nil
}

/*
{"status":"success","message":"successfully initiated withdrawal","payload":{"id":94,"amount":"1.0000000000000000","token_id":"usdc","created_at":"2024-03-17T13:31:12.324080Z","transaction_status":"INITIATED","extras":{"errors":[],"exp_timestamp":3997985,"quantised_amount":1000000}}}
*/
type ValidateWithdrawalResponse struct {
	Status  string `json:"status"`
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
	if err != nil {
		return ValidateWithdrawalResponse{}, fmt.Errorf("error: %s", validateWithdrawalResponse.Message)
	}

	return validateWithdrawalResponse, nil
}

// initiateNormalWithdrawal initiates a normal withdrawal
func (c *Client) InitiateNormalWithdrawal(ctx context.Context, starkPrivateKey string, amount float64, currency Currency) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	err := c.CheckAuth()
	if err != nil {
		return err
	}

	coinStatus, err := c.GetCoinStatus(ctx, currency)
	if err != nil {
		return err
	}
	log.Println("coinstatus", coinStatus)

	respStart, err := c.StartNormalWithdrawal(ctx, StartNormalWithdrawalRequest{
		Amount:   amount,
		Currency: currency,
	})
	if err != nil {
		return err
	}
	log.Println("respStart", respStart)

	if starkPrivateKey[:2] != "0x" {
		starkPrivateKey = "0x" + starkPrivateKey
	}

	numBig := new(big.Int)
	numBig.SetString(respStart.Payload.MsgHash, 10)
	msgHash := hexutil.EncodeBig(numBig)
	log.Println("msgHash", msgHash)

	r, s := crypto_lib.Sign(starkPrivateKey, msgHash, "0x1")
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
		return err
	}

	log.Println("respValidate", respValidate)
	return nil
}

type WithdrawalContract interface {
	GetWithdrawalBalance(opts *bind.CallOpts, ownerKey *big.Int, assetId *big.Int) (*big.Int, error)
	Withdraw(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error)
}

func (c *Client) GetPendingNormalWithdrawalAmountByCoin(ctx context.Context, currency Currency, userPublicEthAddress string, ethClient *ethclient.Client, network string) (string, error) {
	err := c.CheckAuth()
	if err != nil {
		return "", err
	}

	coinStatus, err := c.GetCoinStatus(ctx, currency)
	if err != nil {
		return "", err
	}
	log.Println("coinstatus", coinStatus)

	starkAssetID := coinStatus.StarkAssetID
	blockchainDecimal := coinStatus.BlockchainDecimal

	var starkaddr common.Address
	var ctr WithdrawalContract

	switch network {
	case TESTNET:
		starkaddr = common.HexToAddress(TESTNET_STARK_CONTRACT)
		ctr, err = contract.NewDepositTestnet(starkaddr, ethClient)

	case MAINNET:
		starkaddr = common.HexToAddress(MAINET_STARK_CONTRACT)
		ctr, err = contract.NewDepositMainnet(starkaddr, ethClient)

	default:
		return "", ErrInvalidNetwork
	}
	if err != nil {
		return "", err
	}

	ethAddressBigInt, ok := new(big.Int).SetString(userPublicEthAddress[2:], 16)
	if !ok {
		return "", err
	}

	starkAssetIDBigInt, ok := new(big.Int).SetString(starkAssetID[2:], 16)
	if !ok {
		return "", err
	}
	balance, err := ctr.GetWithdrawalBalance(&bind.CallOpts{
		From:    common.HexToAddress(userPublicEthAddress),
		Context: ctx,
	}, ethAddressBigInt, starkAssetIDBigInt)
	if err != nil {
		return "", err
	}
	log.Printf("balance: %v", balance)

	bd, err := strconv.Atoi(blockchainDecimal)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", ToDecimal(balance, bd)), nil
}

func (c *Client) CompleteNormalWithdrawal(ctx context.Context, currency Currency, ethPrivateKey string, userPublicEthAddress string, ethClient *ethclient.Client, network string) (string, error) {
	err := c.CheckAuth()
	if err != nil {
		return "", err
	}

	coinStatus, err := c.GetCoinStatus(ctx, currency)
	if err != nil {
		return "", err
	}
	log.Println("coinstatus", coinStatus)

	starkAssetID := coinStatus.StarkAssetID

	var starkaddr common.Address
	var ctr WithdrawalContract

	switch network {
	case TESTNET:
		starkaddr = common.HexToAddress(TESTNET_STARK_CONTRACT)
		ctr, err = contract.NewDepositTestnet(starkaddr, ethClient)

	case MAINNET:
		starkaddr = common.HexToAddress(MAINET_STARK_CONTRACT)
		ctr, err = contract.NewDepositMainnet(starkaddr, ethClient)

	default:
		return "", ErrInvalidNetwork
	}
	if err != nil {
		return "", err
	}

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
		chainID, err := ethClient.NetworkID(context.Background())
		if err != nil {
			return nil, err
		}

		signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}
	opt, err := getTransactionOpt(context.Background(), ethClient, ethPrivateKey, signerFn, 0, 0)
	if err != nil {
		return "", err
	}
	res, err := ctr.Withdraw(opt, ethAddressBigInt, starkAssetIDBigInt)
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
	Status  string                      `json:"status"`
	Message string                      `json:"message"`
	Payload ListNormalWithdrawalPayload `json:"payload"`
}

type ListWithdrawalRequest struct {
	Page    int    `json:"page"`
	Network string `json:"network"`
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
		params.Set("network", opt.Network)
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
	if err != nil {
		return ListWithdrawalResponse{}, fmt.Errorf("error: %s", listWithdrawalResponse.Message)
	}

	return listWithdrawalResponse, nil
}

/*
	{
	  "status": "success",
	  "message": "successfully initiated withdrawal",
	  "payload": {
	    "fastwithdrawal_withdrawal_id": 139,
	    "msg_hash": "0x7fed74633a4b8566926b5e3ba5359a07fd82a4f54289e7e9483aa4b5ca9f9e9"
	  }
	}
*/
type StartFastWithdrawalResponse struct {
	Status  string `json:"status"`
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
	if err != nil {
		return StartFastWithdrawalResponse{}, fmt.Errorf("error: %s", startFastWithdrawalResponse.Message)
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
	if err != nil {
		return ValidateWithdrawalResponse{}, fmt.Errorf("error: %s", processFastWithdrawalResponse.Message)
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
		coinStatus, err := c.GetCoinStatus(ctx, currency)
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

	r, s := crypto_lib.Sign(starkPrivateKey, respStart.Payload.MsgHash, "0x1")
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
		params.Set("network", opt.Network)
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
	if err != nil {
		return ListWithdrawalResponse{}, fmt.Errorf("error: %s", listWithdrawalResponse.Message)
	}

	return listWithdrawalResponse, nil
}
