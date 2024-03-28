package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tanx-libs/tanx-connector-go/contract"
)

type CoinToken struct {
	BlockchainDecimal                  string `json:"blockchain_decimal"`
	TokenContract                      string `json:"token_contract"`
	MaxFastWithdrawalForPlatformPerDay string `json:"max_fast_withdrawal_for_platform_per_day"`
	MaxFastWithdrawalForUserPerDay     string `json:"max_fast_withdrawal_for_user_per_day"`
}

type NetworkConfigData struct {
	DepositContract                 string                 `json:"deposit_contract"`
	Tokens                          map[Currency]CoinToken `json:"tokens"`
	AllowedTokensForDeposit         []Currency             `json:"allowed_tokens_for_deposit"`
	AllowedTokensForDepositFrontend []Currency             `json:"allowed_tokens_for_deposit_frontend"`
	AllowedTokensForFastWd          []Currency             `json:"allowed_tokens_for_fast_wd"`
	AllowedTokensForFastWdFrontend  []Currency             `json:"allowed_tokens_for_fast_wd_frontend"`
}

type NetworkConfigResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
		NetworkConfig map[Network]NetworkConfigData `json:"network_config"`
	} `json:"payload"`
}

func (c *Client) GetNetworkConfig(ctx context.Context) (NetworkConfigResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return NetworkConfigResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.networkConfigURL.String(), nil)
	if err != nil {
		return NetworkConfigResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return NetworkConfigResponse{}, err
	}
	defer resp.Body.Close()

	var networkConfigResponse NetworkConfigResponse
	err = json.NewDecoder(resp.Body).Decode(&networkConfigResponse)
	if err != nil {
		return NetworkConfigResponse{}, err
	}

	return networkConfigResponse, nil
}

func isPresent(allowedTokens []Currency, currency Currency) bool {
	for _, v := range allowedTokens {
		if v == currency {
			return true
		}
	}
	return false
}

type CrossChainDepositRequest struct {
	Amount                 string   `json:"amount"`
	Currency               Currency `json:"currency"`
	Network                Network  `json:"network"`
	DepositBlockchainHash  string   `json:"deposit_blockchain_hash"`
	DepositBlockchainNonce string   `json:"deposit_blockchain_nonce"`
}

func (c *Client) CrossChainDepositStart(ctx context.Context, crossChainDepositRequest CrossChainDepositRequest) (CryptoDepositResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	requestBody, err := json.Marshal(crossChainDepositRequest)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, c.crossChainDepositStartURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	defer resp.Body.Close()

	var cryptoDepositResponse CryptoDepositResponse
	err = json.NewDecoder(resp.Body).Decode(&cryptoDepositResponse)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	return cryptoDepositResponse, nil
}

func (c *Client) polygonInit(ctx context.Context, rpcURL string) error {
	if c.polygonClient == nil {
		polygonClient, err := ethclient.Dial(rpcURL)
		if err != nil {
			return err
		}

		// polygon network config
		networkConfigResp, err := c.GetNetworkConfig(ctx)
		if err != nil {
			return err
		}
		polygonConfig := networkConfigResp.Payload.NetworkConfig[POLYGON]
		c.polygonConfig = polygonConfig

		// contract setup
		polygonAddr := common.HexToAddress(c.polygonConfig.DepositContract)
		ctr, err := contract.NewDepositPolygon(polygonAddr, polygonClient)
		if err != nil {
			return err
		}

		c.polygonClient = polygonClient

		c.polygonContract = ctr
	}

	return nil
}

func (c *Client) SetPolygonAllowance(ctx context.Context, rpcURL string, ethPrivateKey string, currency Currency, amount float64) error {
	// one time setup
	err := c.polygonInit(ctx, rpcURL)
	if err != nil {
		return err
	}

	currentCoin := c.polygonConfig.Tokens[currency]
	blockchainDecimal, err := strconv.Atoi(currentCoin.BlockchainDecimal)
	if err != nil {
		return err
	}

	// signer function
	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return err
	}

	signerFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		chainID, err := c.polygonClient.ChainID(context.Background())
		if err != nil {
			return nil, err
		}

		signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}

	opt, err := getTransactionOpt(ctx, c.polygonClient, ethPrivateKey, signerFn, 0, blockchainDecimal)
	if err != nil {
		return err
	}

	contractAddress := c.polygonConfig.DepositContract
	polygonContractAddress := common.HexToAddress(contractAddress)

	err = setAllowance(ctx, c.polygonClient, currentCoin.TokenContract, opt, polygonContractAddress, blockchainDecimal, amount)
	if err != nil {
		return err
	}

	return nil
}

type PolygonContract interface {
	DepositNative(opts *bind.TransactOpts) (*types.Transaction, error)
	Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error)
}

func (c *Client) DepositFromPolygonNetwork(
	ctx context.Context,
	rpcURL string,
	ethAddress string,
	ethPrivateKey string,
	starkPublicKey string,
	currency Currency,
	amount float64,
) (CryptoDepositResponse, error) {
	// check amount
	if amount < 0 {
		return CryptoDepositResponse{}, ErrInvalidAmount
	}

	// check if logged in or not
	err := c.CheckAuth()
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	err = c.polygonInit(ctx, rpcURL)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// network config already setup

	allowedTokens := c.polygonConfig.AllowedTokensForDeposit
	contractAddress := c.polygonConfig.DepositContract

	if !isPresent(allowedTokens, currency) {
		return CryptoDepositResponse{}, ErrCoinNotFound
	}

	currentCoin := c.polygonConfig.Tokens[currency]
	blockchainDecimal, err := strconv.Atoi(currentCoin.BlockchainDecimal)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// contract already setup

	balance, err := getTokenBalance(ctx, c.polygonClient, ethAddress, currency, currentCoin.BlockchainDecimal, currentCoin.TokenContract)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	if balance.Cmp(big.NewFloat(amount)) == -1 {
		return CryptoDepositResponse{}, fmt.Errorf("current balance of %v is less than deposit amount %v", balance, amount)
	}

	var transaction *types.Transaction

	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	signerFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		chainID, err := c.polygonClient.ChainID(context.Background())
		if err != nil {
			return nil, err
		}

		signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}

	if currency == MATIC {
		opt, err := getTransactionOpt(ctx, c.polygonClient, ethPrivateKey, signerFn, amount, blockchainDecimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		transaction, err = c.polygonContract.DepositNative(opt)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

	} else {
		opt, err := getTransactionOpt(ctx, c.polygonClient, ethPrivateKey, signerFn, 0, blockchainDecimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		polygonContractAddress := common.HexToAddress(contractAddress)
		allowance, err := getAllowance(c.polygonClient, ethAddress, currentCoin.TokenContract, blockchainDecimal, polygonContractAddress)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		if allowance < amount {
			return CryptoDepositResponse{},
				&ErrInsufficientAllowance{
					Allowance: allowance,
					Amount:    amount,
				}
		}

		quantizedAmount := ToWei(amount, blockchainDecimal)
		address := common.HexToAddress(currentCoin.TokenContract)
		transaction, err = c.polygonContract.Deposit(opt, address, quantizedAmount)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
	}

	resp, err := c.CrossChainDepositStart(ctx, CrossChainDepositRequest{
		Amount:                 fmt.Sprintf("%f", amount),
		Currency:               currency,
		Network:                POLYGON,
		DepositBlockchainHash:  transaction.Hash().Hex(),
		DepositBlockchainNonce: fmt.Sprintf("%v", transaction.Nonce()),
	})
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	resp.Payload = transaction.Hash().Hex()

	return resp, nil
}

type ListDepositsResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
		Count    int         `json:"count"`
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
		Results  []struct {
			TokenID               string `json:"token_id"`
			Network               string `json:"network"`
			Amount                string `json:"amount"`
			CreatedAt             string `json:"created_at"`
			BlockchainDepositHash string `json:"deposit_blockchain_hash"`
		} `json:"results"`
	} `json:"payload"`
}

func (c *Client) ListAllDeposits(ctx context.Context) (ListDepositsResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return ListDepositsResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.listDepositsURL.String(), nil)
	if err != nil {
		return ListDepositsResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ListDepositsResponse{}, err
	}
	defer resp.Body.Close()

	var listDepositsResponse ListDepositsResponse
	err = json.NewDecoder(resp.Body).Decode(&listDepositsResponse)
	if err != nil {
		return ListDepositsResponse{}, err
	}

	return listDepositsResponse, nil
}
