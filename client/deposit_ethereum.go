// todo
/*
- test if normal wala chalra hain ya nahi
- make it clean
*/

package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
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

type CoinStatusPayload struct {
	Symbol            Currency `json:"symbol"`
	Quanitization     string   `json:"quanitization"`
	Decimal           string   `json:"decimal"`
	BlockchainDecimal string   `json:"blockchain_decimal"`
	TokenContract     string   `json:"token_contract"`
	StarkAssetID      string   `json:"stark_asset_id"`
}

type CoinStatusResponse struct {
	Status  string                       `json:"status"`
	Message string                       `json:"message"`
	Payload map[string]CoinStatusPayload `json:"payload"`
}

func (c *Client) GetCoinStatus(ctx context.Context, currency Currency) (CoinStatusPayload, error) {
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

	for _, v := range coinStatusResponse.Payload {
		if v.Symbol == currency {
			return v, nil
		}
	}

	return CoinStatusPayload{}, ErrCoinNotFound
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
	Coin Currency `json:"coin"`
}

func (c *Client) GetVaultID(ctx context.Context, currency Currency) (int, error) {
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
	Amount                 string `json:"amount"`
	StarkAssetID           string `json:"token_id"`
	StarkPublicKey         string `json:"stark_key"`
	DepositBlockchainHash  string `json:"deposit_blockchain_hash"`
	DepositBlockchainNonce string `json:"deposit_blockchain_nonce"`
	VaultID                int    `json:"vault_id"`
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

func getTokenBalance(ctx context.Context, ethClient *ethclient.Client, ethAddress string, currency Currency, decimal string, contractAddress string) (*big.Float, error) {
	var balance *big.Int

	ethAdr := common.HexToAddress(ethAddress)

	d, err := strconv.Atoi(decimal)
	if err != nil {
		return nil, err
	}

	switch currency {
	case ETH, MATIC:
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

	log.Println("balance before decimal call", balance)
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
	contract, err := contract.NewErc20(common.HexToAddress(tokenContract), provider)
	if err != nil {
		return 0, err
	}

	allowance, err := contract.Allowance(nil, common.HexToAddress(userAddress), common.HexToAddress(starkContract))
	if err != nil {
		return 0, err
	}

	res, _ := ToDecimal(allowance, decimal).Float64()

	return res, nil
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

	amountInWei := ToWei(amount, decimal)

	tipCap, _ := ethClient.SuggestGasTipCap(context.Background())
	feeCap, _ := ethClient.SuggestGasPrice(context.Background())

	return &bind.TransactOpts{
		From:      fromAddress,
		Nonce:     big.NewInt(int64(nonce)),
		Signer:    signerFn,
		Value:     amountInWei,
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
		Context:   ctx,
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

	// getting coin information here
	coinStatus, err := c.GetCoinStatus(ctx, currency)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Printf("coinstatus %+v\n", coinStatus)

	// getting vault id here
	vaultID, err := c.GetVaultID(ctx, currency)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Printf("vault id: %+v\n", vaultID)

	// building contract here
	var starkaddr common.Address
	var ctr DepositContract

	switch network {
	case string(TESTNET):
		starkaddr = common.HexToAddress(TESTNET_STARK_CONTRACT)
		ctr, err = contract.NewDepositTestnet(starkaddr, ethClient)

	case string(MAINNET):
		starkaddr = common.HexToAddress(MAINET_STARK_CONTRACT)
		ctr, err = contract.NewDepositMainnet(starkaddr, ethClient)

	default:
		return CryptoDepositResponse{}, ErrInvalidNetwork
	}
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// Getting balance information here
	balance, err := getTokenBalance(ctx, ethClient, ethAddress, currency, coinStatus.BlockchainDecimal, coinStatus.TokenContract)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Printf("balance: %+v\n", balance)

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

	// decimal
	decimal, err := strconv.Atoi(coinStatus.Decimal)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Printf("decimal: %d\n", decimal)

	// blockchain decimal
	blockchainDecimal, err := strconv.Atoi(coinStatus.BlockchainDecimal)
	if err != nil {
		return CryptoDepositResponse{}, err
	}
	log.Printf("blockchain decimal: %+v\n", blockchainDecimal)

	// quantization
	quantization, err := strconv.Atoi(coinStatus.Quanitization)
	if err != nil {
		return CryptoDepositResponse{}, nil
	}
	quantizedAmount := ToWei(amount, quantization)
	log.Printf("quantized amount: %+v\n", quantizedAmount)

	if currency == ETH {
		opt, err = getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, amount, blockchainDecimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		log.Printf("opt %+v", opt)

		transaction, err = ctr.DepositEth(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		log.Println("transaction", transaction)

	} else {
		allowance, err := getAllowance(ethAddress, starkaddr.String(), coinStatus.TokenContract, blockchainDecimal, ethClient)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		log.Printf("allowance: %f\n", allowance)

		// // allowance code from here
		// setAllowance := func(amount int) error {
		// 	tokenContractAddr := common.HexToAddress(coinStatus.TokenContract)

		// 	erc20Contract, err := contract.NewErc20(tokenContractAddr, ethClient)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	opt, err := getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, 0, decimal)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	opt.GasLimit, err = ethClient.EstimateGas(ctx, ethereum.CallMsg{
		// 		From: tokenContractAddr,
		// 	})
		// 	if err != nil {
		// 		return err
		// 	}

		// 	amountInWei := ToWei(float64(amount), blockchainDecimal)

		// 	_, err = erc20Contract.Approve(opt, starkaddr, amountInWei)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	log.Println("gas price: ", opt.GasPrice)
		// 	log.Println("gas limit: ", opt.GasLimit)

		// 	return nil
		// }

		// err = setAllowance(100)
		// if err != nil {
		// 	return CryptoDepositResponse{}, err
		// }

		if allowance < amount {
			return CryptoDepositResponse{}, fmt.Errorf("current allowance is %v", allowance)
		}

		opt, err = getTransactionOpt(ctx, ethClient, ethPrivateKey, signerFn, 0, decimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		opt.GasLimit = 100000
		log.Printf("opt: %+v\n", opt)

		transaction, err = ctr.DepositERC20(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt, quantizedAmount)
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

	resp, err := c.CryptoDepositStart(ctx, CryptoDepositRequest{
		Amount:                 quantizedAmount.String(),
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
func (c *Client) DepositFromEthereumNetwork(ctx context.Context, rpcURL string, ethAddress string, ethPrivateKey string, starkPublicKey string, network string, currency Currency, amount float64) (CryptoDepositResponse, error) {
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

	return c.DepositFromEthereumNetworkWithStarkKey(ctx, ethClient, signerFn, network, ethAddress, ethPrivateKey, starkPublicKey, amount, currency)
}

