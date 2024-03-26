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

type CoinStatusPayload struct {
	Symbol            Currency `json:"symbol"`
	Quanitization     string   `json:"quanitization"`
	Decimal           string   `json:"decimal"`
	BlockchainDecimal string   `json:"blockchain_decimal"`
	TokenContract     string   `json:"token_contract"`
	StarkAssetID      string   `json:"stark_asset_id"`
}

type CoinStatusResponse struct {
	Status  Status                       `json:"status"`
	Message string                       `json:"message"`
	Payload map[string]CoinStatusPayload `json:"payload"`
}

func (c *Client) getCoinStatus(ctx context.Context) (CoinStatusResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return CoinStatusResponse{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.coinURL.String(), nil)
	if err != nil {
		return CoinStatusResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return CoinStatusResponse{}, err
	}
	defer resp.Body.Close()

	var coinStatusResponse CoinStatusResponse
	err = json.NewDecoder(resp.Body).Decode(&coinStatusResponse)

	if coinStatusResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return CoinStatusResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(coinStatusResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return CoinStatusResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(coinStatusResponse.Message),
			}
		}

		return CoinStatusResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, coinStatusResponse.Message)

	} else if err != nil {
		return CoinStatusResponse{}, &ErrJSONDecoding{Err: err}
	}

	return coinStatusResponse, nil
}

func (c *Client) getCoinStatusPayload(currency Currency) (CoinStatusPayload, error) {
	for _, v := range c.coinStatus.Payload {
		if v.Symbol == currency {
			return v, nil
		}
	}

	return CoinStatusPayload{}, ErrCoinNotFound
}

type VaultResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload struct {
		ID   int    `json:"id"`
		Coin string `json:"coin"`
	} `json:"payload"`
}

type VaultRequest struct {
	Coin Currency `json:"coin"`
}

func (c *Client) getVaultID(ctx context.Context, currency Currency) (int, error) {
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

	if vaultResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return 0, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(vaultResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return 0, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(vaultResponse.Message),
			}
		}

		return 0, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, vaultResponse.Message)

	} else if err != nil {
		return 0, &ErrJSONDecoding{Err: err}
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
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload interface{} `json:"payload"`
}

// CryptoDepositStart
func (c *Client) cryptoDepositStart(ctx context.Context, depositReq CryptoDepositRequest) (CryptoDepositResponse, error) {
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

	if cryptoDepositResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return CryptoDepositResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(cryptoDepositResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return CryptoDepositResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(cryptoDepositResponse.Message),
			}
		}

		return CryptoDepositResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, cryptoDepositResponse.Message)

	} else if err != nil {
		return CryptoDepositResponse{}, &ErrJSONDecoding{Err: err}
	}

	return cryptoDepositResponse, nil
}

func (c *Client) SetAllowance(amount int) {
	c.ethereumNetworkAllowanceSet = false
	c.ethereumNetworkAllowance = amount
}

func (c *Client) DepositFromEthereumNetwork(
	ctx context.Context,
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

	// coin status here
	coinStatus, err := c.getCoinStatusPayload(currency)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// getting vault id here
	vaultID, err := c.getVaultID(ctx, currency)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// contract already build in client

	// Getting balance information here
	balance, err := getTokenBalance(ctx, c.ethClient, ethAddress, currency, coinStatus.BlockchainDecimal, coinStatus.TokenContract)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// a more verbose error here could be possible
	if balance.Cmp(big.NewFloat(amount)) == -1 {
		return CryptoDepositResponse{}, fmt.Errorf("current balance: %v", balance)
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

	// decimal
	decimal, err := strconv.Atoi(coinStatus.Decimal)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// blockchain decimal
	blockchainDecimal, err := strconv.Atoi(coinStatus.BlockchainDecimal)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	// quantization
	quantization, err := strconv.Atoi(coinStatus.Quanitization)
	if err != nil {
		return CryptoDepositResponse{}, nil
	}
	quantizedAmount := ToWei(amount, quantization)

	// signer function
	privateKey, err := crypto.HexToECDSA(ethPrivateKey)
	if err != nil {
		return CryptoDepositResponse{}, err
	}

	signerFn := func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
		// mayhul
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

	if currency == ETH {
		opt, err := getTransactionOpt(ctx, c.ethClient, ethPrivateKey, signerFn, amount, blockchainDecimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		transaction, err = c.starkexContract.DepositEth(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

	} else {
		opt, err := getTransactionOpt(ctx, c.ethClient, ethPrivateKey, signerFn, 0, decimal)
		if err != nil {
			return CryptoDepositResponse{}, err
		}
		opt.GasLimit = 100000

		if !c.ethereumNetworkAllowanceSet {
			err = setAllowance(ctx, c.ethClient, coinStatus.TokenContract, opt, c.starkexContractAddress, blockchainDecimal, c.ethereumNetworkAllowance)
			if err != nil {
				return CryptoDepositResponse{}, err
			}

			c.ethereumNetworkAllowanceSet = true
		}

		allowance, err := getAllowance(c.ethClient, ethAddress, coinStatus.TokenContract, blockchainDecimal, c.starkexContractAddress)
		if err != nil {
			return CryptoDepositResponse{}, err
		}

		if allowance < amount {
			return CryptoDepositResponse{}, fmt.Errorf("current allowance is %v. call SetAllowance function to change the allowance amount", allowance)
		}

		transaction, err = c.starkexContract.DepositERC20(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt, quantizedAmount)
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

	resp, err := c.cryptoDepositStart(ctx, CryptoDepositRequest{
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

type StarkexContract interface {
	DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error)
	DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error)
}

// Note: This function has to be called before using any other deposit function
func (c *Client) DepositFromEthereumNetworkInit(ctx context.Context, rpcURL string) error {
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return err
	}
	c.ethClient = ethClient

	// getting coin information here
	coinStatus, err := c.getCoinStatus(ctx)
	if err != nil {
		return err
	}

	c.coinStatus = coinStatus

	// setting up starkex contract here
	var starkaddr common.Address
	var starkexContract StarkexContract
	switch c.network {
	case TESTNET:
		starkaddr = common.HexToAddress(TESTNET_STARK_CONTRACT)
		starkexContract, err = contract.NewDepositTestnet(starkaddr, c.ethClient)

	case MAINNET:
		starkaddr = common.HexToAddress(MAINET_STARK_CONTRACT)
		starkexContract, err = contract.NewDepositMainnet(starkaddr, c.ethClient)

	default:
		return ErrInvalidNetwork
	}
	if err != nil {
		return err
	}
	c.starkexContract = starkexContract
	c.starkexContractAddress = starkaddr

	return nil
}

// // old allowance code
// setAllowance := func(amount int) error {
// 	tokenContractAddr := common.HexToAddress(coinStatus.TokenContract)

// 	erc20Contract, err := contract.NewErc20(tokenContractAddr, c.ethClient)
// 	if err != nil {
// 		return err
// 	}

// 	opt, err := getTransactionOpt(ctx, c.ethClient, ethPrivateKey, signerFn, 0, decimal)
// 	if err != nil {
// 		return err
// 	}

// 	opt.GasLimit, err = c.ethClient.EstimateGas(ctx, ethereum.CallMsg{
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
