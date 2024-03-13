package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tanx-libs/tanx-connector-go/contract"
)

type CoinStatusPayload struct {
	Quanitization string `json:"quanitization"`
	Decimal       string `json:"decimal"`
	TokenContract string `json:"token_contract"`
	StarkAssetID  string `json:"stark_asset_id"`
}

type CoinStatusResponse struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Payload map[string]CoinStatusPayload `json:"payload"`
}

func (c *Client) GetCoinStatus(ctx context.Context, currency string) (CoinStatusPayload, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.coinURL.String(), nil)
	if err != nil {
		return CoinStatusPayload{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CoinStatusPayload{}, err
	}
	defer resp.Body.Close()

	var coinStatusResponse CoinStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&coinStatusResponse)
	if err != nil {
		return CoinStatusPayload{}, err
	}

	val, ok := coinStatusResponse.Payload[currency]
	if !ok {
		return CoinStatusPayload{}, ErrCoinNotFound
	}

	return val, nil
}

type VaultResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
		ID   int    `json:"id"`
		Coin string `json:"coin"`
	} `json:"payload"`
}

type VaultRequest struct {
	Coin string `json:"coin"`
}

func (c *Client) GetVaultID(ctx context.Context, currency string) (int, error) {
	err := c.CheckAuth()
	if err != nil {
		return 0, nil
	}

	r := VaultRequest{
		Coin: currency,
	}
	requestBody, err := json.Marshal(r)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.vaultIDURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return 0, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var vaultResponse VaultResponse
	err = json.NewDecoder(resp.Body).Decode(&vaultResponse)
	if err != nil {
		return 0, err
	}

	return vaultResponse.Payload.ID, nil
}

type CryptoDepositRequest struct {
	Amount                 float64 `json:"amount"`
	StarkAssetID           string  `json:"token_id"`
	StarkPublicKey         string  `json:"stark_key"`
	DepositBlockchainHash  string  `json:"deposit_blockchain_hash"`
	DepositBlockchainNonce string  `json:"deposit_blockchain_nonce"`
	VaultID                int     `json:"vault_id"`
}

type CryptoDepositResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// CryptoDepositStart
func (c *Client) CryptoDepositStart(ctx context.Context, depositReq CryptoDepositRequest) (CryptoDepositResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	requestBody, err := json.Marshal(depositReq)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.cryptoDepositStartURL.String(), bytes.NewReader(requestBody))
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

func getTokenBalance(ctx context.Context, ethClient *ethclient.Client, ethAddress string, currency string, decimal string, contractAddress string) (*big.Float, error) {
	var balance *big.Int

	ethAdr := common.HexToAddress(ethAddress)
	d, err := strconv.Atoi(decimal)
	if err != nil {
		return nil, err
	}

	switch currency {
	case "ethereum", "matic":
		balance, err = ethClient.BalanceAt(ctx, ethAdr, nil)
		if err != nil {
			return nil, err
		}

	default:
		tokenAddr := common.HexToAddress(contractAddress)
		ctr, err := contract.NewErc20(tokenAddr, ethClient)
		if err != nil {
			return nil, err
		}

		balance, err = ctr.BalanceOf(nil, ethAdr)
		if err != nil {
			return nil, err
		}
	}

	return ToDecimal(balance, d), nil
}

type DepositContract interface {
	DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error)
	DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error)
}

func getAllowance(
	userAddress string,
	starkContract string,
	tokenContract string,
	decimal int,
	provider *ethclient.Client,
) (float64, error) {
	contractAddress := common.HexToAddress(tokenContract)
	contract, err := contract.NewErc20(contractAddress, provider)
	if err != nil {
		return 0, err
	}

	allowance, err := contract.Allowance(nil, common.HexToAddress(userAddress), common.HexToAddress(starkContract))
	if err != nil {
		return 0, err
	}

	return dequantize(allowance, decimal), nil
}

func dequantize(number *big.Int, decimals int) float64 {
	factor := math.Pow10(decimals)
	return float64(number.Int64()) / factor
}

func getTransactionOpt(ctx context.Context, ethClient *ethclient.Client, ethPrivateKey string, signerFn bind.SignerFn, amount float64, decimal int) (*bind.TransactOpts, error) {
	if ethPrivateKey[:2] == "0x" {
		ethPrivateKey = ethPrivateKey[2:]
	}

	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	amountInWei := ToWei(amount, decimal)
	return &bind.TransactOpts{
		From:     fromAddress,
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   signerFn,
		GasPrice: gasPrice,
		Value:    amountInWei,
		Context:  ctx,
	}, nil
}

func (c *Client) DepositFromEthereumNetworkWithStarkKey(
	ctx context.Context,
	ethClient *ethclient.Client,
	signerFn bind.SignerFn,
	network string,
	ethAddress string,
	ethPrivateKey string,
	starkPublicKey string,
	amount float64,
	currency string,
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

	// getting coin information here
	currency = strings.ToLower(currency)
	coinStatus, err := c.GetCoinStatus(ctx, currency)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	log.Printf("coin %+v\n", coinStatus)

	// quantization
	q, err := strconv.Atoi(coinStatus.Quanitization)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	quantizedAmount := int64(amount * math.Pow10(q))
	log.Println("quantized amount", quantizedAmount)

	// getting vault id here
	vaultID, err := c.GetVaultID(ctx, currency)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Println("vault id", vaultID)

	// building contract here
	var starkaddr common.Address
	var ctr DepositContract

	switch network {
	case TESTNET:
		starkaddr = common.HexToAddress(TESTNET_STARK_CONTRACT)
		ctr, err = contract.NewDepositTestnet(starkaddr, ethClient)

	case MAINNET:
		starkaddr = common.HexToAddress(MAINET_STARK_CONTRACT)
		ctr, err = contract.NewDepositMainnet(starkaddr, ethClient)

	default:
		return CryptoDepositResponse{}, ErrInvalidNetwork
	}
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Println("contract ", ctr)

	amountInWei := ToWei(amount, 18)
	gwei := ToWei(amount, 9)
	log.Println("amountInWei", amountInWei)
	log.Println("gwei", gwei)

	// getting balance information here
	balance, err := getTokenBalance(ctx, ethClient, ethAddress, currency, coinStatus.Decimal, coinStatus.TokenContract)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Println("balance", balance)

	// a more verbose error here could be possible
	if balance.Cmp(big.NewFloat(amount)) == -1 {
		return CryptoDepositResponse{}, ErrInsufficientBalance
	}

	var transaction *types.Transaction
	starkPublicKeyBigInt, ok := new(big.Int).SetString(starkPublicKey[2:], 16)
	if !ok {
		return CryptoDepositResponse{}, fmt.Errorf("failed to convert starkPublicKey to big.Int")
	}

	starkAssetIDBigInt, ok := new(big.Int).SetString(coinStatus.StarkAssetID[2:], 16)
	if !ok {
		return CryptoDepositResponse{}, fmt.Errorf("failed to convert StarkAssetID to big.Int")
	}

	vaultIDBigInt := big.NewInt(int64(vaultID))

	var opt *bind.TransactOpts
	if currency == "eth" {
		opt, err = getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, amount, 18)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		transaction, err = ctr.DepositEth(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		log.Println("transaction", transaction)
	} else {
		decimal, err := strconv.Atoi(coinStatus.Decimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		allowance, err := getAllowance(ethAddress, starkaddr.String(), coinStatus.TokenContract, decimal, ethClient)
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

		transaction, err = ctr.DepositERC20(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt, big.NewInt(quantizedAmount))
		if err != nil {
			return CryptoDepositResponse{}, err
		}
	}

	var starkAssetId string
	if coinStatus.StarkAssetID[0:3] == "0x0" {
		starkAssetId = "0x" + coinStatus.StarkAssetID[3:]
	} else {
		starkAssetId = coinStatus.StarkAssetID
	}

	// todo what amount goes into this
	resp, err := c.CryptoDepositStart(ctx, CryptoDepositRequest{
		Amount:                 amount,
		StarkAssetID:           starkAssetId,
		StarkPublicKey:         starkPublicKey,
		DepositBlockchainHash:  transaction.Hash().Hex(),
		DepositBlockchainNonce: fmt.Sprintf("%v", transaction.Nonce()),
		VaultID:                vaultID,
	})
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	resp.Payload = transaction.Hash().Hex()
	return resp, nil
}

// todo
func (c *Client) DepositFromEthereumNetwork(ctx context.Context, rpcURL string, ethAddress string, ethPrivateKey string, starkPublicKey string, network string, currency string, amount float64) (CryptoDepositResponse, error) {
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

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}

	return c.DepositFromEthereumNetworkWithStarkKey(ctx, ethClient, signerFn, network, ethAddress, ethPrivateKey, starkPublicKey, amount, currency)
}

type CoinToken struct {
	BlockchainDecimal                  string `json:"blockchain_decimal"`
	TokenContract                      string `json:"token_contract"`
	MaxFastWithdrawalForPlatformPerDay string `json:"max_fast_withdrawal_for_platform_per_day"`
	MaxFastWithdrawalForUserPerDay     string `json:"max_fast_withdrawal_for_user_per_day"`
}

type NetworkConfigData struct {
	DepositContract                 string               `json:"deposit_contract"`
	Tokens                          map[string]CoinToken `json:"tokens"`
	AllowedTokensForDeposit         []string             `json:"allowed_tokens_for_deposit"`
	AllowedTokensForDepositFrontend []string             `json:"allowed_tokens_for_deposit_frontend"`
	AllowedTokensForFastWd          []string             `json:"allowed_tokens_for_fast_wd"`
	AllowedTokensForFastWdFrontend  []string             `json:"allowed_tokens_for_fast_wd_frontend"`
}

type NetworkConfigResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload struct {
		NetworkConfig map[string]NetworkConfigData `json:"network_config"`
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

func isPresent(allowedTokens []string, currency string) bool {
	for _, v := range allowedTokens {
		if v == currency {
			return true
		}
	}
	return false
}

type CrossChainDepositRequest struct {
	Amount                 string `json:"amount"`
	Currency               string `json:"currency"`
	Network                string `json:"network"`
	DepositBlockchainHash  string `json:"deposit_blockchain_hash"`
	DepositBlockchainNonce string `json:"deposit_blockchain_nonce"`
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
	currency string,
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

	polygonConfig := networkConfigResp.Payload.NetworkConfig["POLYGON"]
	allowedTokens := polygonConfig.AllowedTokensForDeposit
	contractAddress := polygonConfig.DepositContract
	log.Printf("polygonConfig: %+v", polygonConfig)
	log.Printf("allowedTokens: %+v", allowedTokens)
	log.Printf("constractAddress: %s", contractAddress)

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
	log.Printf("contract: %+v", ctr)

	balance, err := getTokenBalance(ctx, ethClient, ethAddress, currency, currentCoin.BlockchainDecimal, currentCoin.TokenContract)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Println("balance", balance)

	if balance.Cmp(big.NewFloat(amount)) == -1 {
		return CryptoDepositResponse{}, ErrInsufficientBalance
	}

	var transaction *types.Transaction
	var opt *bind.TransactOpts

	if currency == "matic" {
		opt, err = getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, amount, decimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		log.Printf("opt: %+v", opt)

		transaction, err = ctr.DepositNative(opt)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		log.Printf("transaction: %+v", transaction)

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
		Network:                "POLYGON",
		DepositBlockchainHash:  transaction.Hash().Hex(),
		DepositBlockchainNonce: fmt.Sprintf("%v", transaction.Nonce()),
	})
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	resp.Payload = transaction.Hash().Hex()

	return resp, nil
}

func (c *Client) DepositFromPolygonNetwork(ctx context.Context, rpcURL string, ethAddress string, ethPrivateKey string, starkPublicKey string, currency string, amount float64) (CryptoDepositResponse, error) {
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

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return nil, err
		}

		return signedTx, nil
	}

	return c.DepositFromPolygonNetworkWithSigner(ctx, ethClient, signerFn, ethAddress, ethPrivateKey, starkPublicKey, amount, currency)
}

