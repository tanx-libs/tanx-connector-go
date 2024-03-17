// todo
/*
- test erc20 wala
*/

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
	Network                Network   `json:"network"`
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

func (c *Client) DepositFromPolygonNetworkWithSigner(
	ctx context.Context,
	ethClient *ethclient.Client,
	signerFn func(addr common.Address, tx *types.Transaction) (*types.Transaction, error),
	ethAddress string,
	ethPrivateKey string,
	starkPublicKey string,
	amount float64,
	currency Currency,
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

	networkConfigResp, err := c.GetNetworkConfig(ctx)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	polygonConfig := networkConfigResp.Payload.NetworkConfig[POLYGON]
	allowedTokens := polygonConfig.AllowedTokensForDeposit
	contractAddress := polygonConfig.DepositContract

	if !isPresent(allowedTokens, currency) {
		return CryptoDepositResponse{}, ErrCoinNotFound
	}

	currentCoin := polygonConfig.Tokens[currency]
	decimal, err := strconv.Atoi(currentCoin.BlockchainDecimal)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	polygonAddr := common.HexToAddress(contractAddress)
	ctr, err := contract.NewDepositPolygon(polygonAddr, ethClient)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	balance, err := getTokenBalance(ctx, ethClient, ethAddress, currency, currentCoin.BlockchainDecimal, currentCoin.TokenContract)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Println("balance: ", balance)

	if balance.Cmp(big.NewFloat(amount)) == -1 {
		return CryptoDepositResponse{}, ErrInsufficientBalance
	}

	var transaction *types.Transaction
	var opt *bind.TransactOpts

	if currency == MATIC {
		opt, err = getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, amount, decimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		transaction, err = ctr.DepositNative(opt)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

	} else {
		allowance, err := getAllowance(ethAddress, contractAddress, currentCoin.TokenContract, decimal, ethClient)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		if allowance < amount {
			return CryptoDepositResponse{}, ErrInsufficientAllowance
		}

		opt, err = getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, amount, decimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		quantizedAmount := ToWei(amount, decimal)
		address := common.HexToAddress(currentCoin.TokenContract)
		transaction, err = ctr.Deposit(opt, address, quantizedAmount)
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

func (c *Client) DepositFromPolygonNetwork(ctx context.Context, rpcURL string, ethAddress string, ethPrivateKey string, starkPublicKey string, currency Currency, amount float64) (CryptoDepositResponse, error) {
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return CryptoDepositResponse{}, err
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

	return c.DepositFromPolygonNetworkWithSigner(ctx, ethClient, signerFn, ethAddress, ethPrivateKey, starkPublicKey, amount, currency)
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