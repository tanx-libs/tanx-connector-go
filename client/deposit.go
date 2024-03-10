package client

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
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
	StarkAssetID           string  `json:"stark_asset_id"`
	StarkPublicKey         string  `json:"stark_public_key"`
	DepositBlockchainHash  string  `json:"deposit_blockchain_hash"`
	DepositBlockchainNonce string  `json:"deposit_blockchain_nonce"`
	VaultID                int     `json:"vault_id"`
}

// CryptoDepositStart
func (c *Client) CryptoDepositStart(ctx context.Context, depositReq CryptoDepositRequest) (string, error) {
	err := c.CheckAuth()
	if err != nil {
		return "", err
	}

	requestBody, err := json.Marshal(depositReq)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.cryptoDepositStartURL.String(), bytes.NewReader(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// debug
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("response body", string(body))

	// var response map[string]string
	// err = json.NewDecoder(resp.Body).Decode(&response)
	// if err != nil {
	// 	return "", err
	// }

	return "", nil
}


func getTokenBalance (ctx context.Context, ethClient *ethclient.Client, ethAddress string, currency string, coinStatus CoinStatusPayload) (*big.Float, error) {
	var balance *big.Int
	ethAdr := common.HexToAddress(ethAddress)
	decimal, err := strconv.Atoi(coinStatus.Decimal)
	if err != nil {
		return nil, err
	}
	switch currency {
	case "ethereum":
		balance, err = ethClient.BalanceAt(ctx, ethAdr, nil)
		if err != nil {
			return nil, err
		}

		return ToDecimal(balance, decimal), nil

	default:
		
		ctr, err := contract.NewErc20(common.HexToAddress(coinStatus.TokenContract), ethClient)
		if err != nil {
			return nil, err
		}

		balance, err := ctr.BalanceOf(nil, ethAdr)
		if err != nil {
			return nil, err
		}

		
		return ToDecimal(balance, decimal), nil
	}
}

type DepositContract interface {
	DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error)
	DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error)
}

func (c *Client) DepositFromEthereumNetworkWithStarkKey(
	ctx context.Context,
	rpcURL string,
	ethAddress string,
	network string,
	amount float64,
	currency string,
	starkPublicKey string,
) (string, error) {
	// check amount
	if amount < 0 {
		return "", ErrInvalidAmount
	}

	// check if logged in or not
	err := c.CheckAuth()
	if err != nil {
		return "", err
	}

	// getting coin information here
	currency = strings.ToLower(currency)
	coinStatus, err := c.GetCoinStatus(ctx, currency)
	if err != nil {
		return "", err
	}

	log.Printf("coin %+v\n", coinStatus)

	// quantization
	q, err := strconv.Atoi(coinStatus.Quanitization)
	if err != nil {
		return "", err
	}
	quantizedAmount := int64(amount * math.Pow10(q))
	log.Println("quantized amount", quantizedAmount)

	// getting vault id here
	vaultID, err := c.GetVaultID(ctx, currency)
	if err != nil {
		return "", err
	}
	log.Println("vault id", vaultID)

	// now I have to communicate with the contract which requires rpcURL, contract_address, abi
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", err
	}

	var addr common.Address
	var ctr DepositContract

	switch network {
	case TESTNET:
		addr = common.HexToAddress(TESTNET_STARK_CONTRACT)
		ctr, err = contract.NewDepositTestnet(addr, ethClient)

	case MAINNET:
		addr = common.HexToAddress(MAINET_STARK_CONTRACT)
		ctr, err = contract.NewDepositMainnet(addr, ethClient)

	default:
		return "", ErrInvalidNetwork
	}
	if err != nil {
		return "", err
	}
	log.Println("contract ", ctr)

	amountInWei := ToWei(amount, 18)
	gwei := ToWei(amount, 9)
	log.Println("amountInWei", amountInWei)
	log.Println("gwei", gwei)

	// todo check if this is covered or not
	/*
		const overrides = {
			value: parsedAmount,
			nonce: await getNonce(signer, provider),
		}
	*/


	// getting balance information here
	balance, err := getTokenBalance(ctx, ethClient, ethAddress, currency, coinStatus)
	if err != nil {
		return "", err
	}
	log.Println("balance", balance)

	// a more verbose error here could be possible
	if balance.Cmp(big.NewFloat(amount)) == -1 {
		return "", ErrInsufficientBalance
	}

	var transaction *types.Transaction
	starkPublicKeyBigInt, ok := new(big.Int).SetString(starkPublicKey[2:], 16)
	if !ok {
		return "", fmt.Errorf("failed to convert starkPublicKey to big.Int")
	}

	starkAssetIDBigInt, ok := new(big.Int).SetString(coinStatus.StarkAssetID[2:], 16)
	if !ok {
		return "", fmt.Errorf("failed to convert StarkAssetID to big.Int")
	}

	vaultIDBigInt := big.NewInt(int64(vaultID))
	if currency == "ethereum" {

		ethPrivateKey := "ba169c79340371a9aa4fd516462f939242f92b522081d945c001b0fb3dc3a66f"
		privateKey, err := crypto.HexToECDSA(ethPrivateKey)
		if err != nil {
			return "", err
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return "", fmt.Errorf("error casting public key to ECDSA")
		}
		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := ethClient.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return "", err
		}

		gasPrice, err := ethClient.SuggestGasPrice(context.Background())
		if err != nil {
			return "", err
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

		opt := &bind.TransactOpts{
			From:     fromAddress,
			Nonce:    big.NewInt(int64(nonce)),
			Signer:   signerFn,
			GasPrice: gasPrice,
			Value:    amountInWei,
			Context:  ctx,
		}

		transaction, err = ctr.DepositEth(opt, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt)
		if err != nil {
			return "", err
		}

	} else {
		decimal, err := strconv.Atoi(coinStatus.Decimal)
		if err != nil {
			return "", err
		}

		allowance, err := getAllowance(ethAddress, coinStatus.TokenContract, coinStatus.StarkAssetID, decimal, ethClient)
		if err != nil {
			return "", err
		}

		if allowance < amount {
			return "", ErrInsufficientAllowance
		}

		// todo yaha peh opts shayad reuse hojayega
		_, err = ctr.DepositERC20(&bind.TransactOpts{}, starkPublicKeyBigInt, starkAssetIDBigInt, vaultIDBigInt, big.NewInt(quantizedAmount))
	}

	// todo
	// crpto deposit start wala bh kuch karna hain
	_, err = c.CryptoDepositStart(ctx, CryptoDepositRequest{})
	if err != nil {
		return "", err
	}

	return transaction.Hash().Hex(), nil
}

/*
Can you make this javascript code like in golang

export const getAllowance = async (

	userAddress: string,
	starkContract: string,
	tokenContract: string,
	decimal: number,
	provider: ethers.providers.Provider,

	) => {
	  const contract = new ethers.Contract(
	    tokenContract,
	    CONFIG.ERC20_ABI,
	    provider,
	  )
	  const allowance = await contract.allowance(userAddress, starkContract)
	  return dequantize(Number(allowance), decimal)
	}
*/
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

// todo
func DepositFromEthereumNetwork(ctx context.Context, rpcURL string, ethPrivateKey string, network string, currency string, amount float64) {
}

// todo docs
/*
Steps

1. Check for amount
2. Check if logged in or not
3. Now get quanitization, decimal, tokenContract, starkAssetId from coin status endpoint
4. Get vault id from vault id endpoint
5. Now connect to the smart contract using rpcURL, contract_address, abi
6. now make a getBalanceFunction
7.
*/

// ethAddr := "0xf318C11ff6E60115FB3e107bEa2637c060BEbc8C"
// ethPrivateKey := "ba169c79340371a9aa4fd516462f939242f92b522081d945c001b0fb3dc3a66f"

// starkPublicKey := "0x64211ed550cb37140ef2268cf7b2625aef725d33618c9651765e16318101c17"
// starkPrivateKey := "0x7302fa58776da9f8fcf3631f4cb495a4dd0cdfab785e8b72a8a637d4bb14784"
