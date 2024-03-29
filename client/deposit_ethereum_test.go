package client

import (
	"context"
	"encoding/hex"
	"io"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/assert"
)

type MyStarkContract struct{}

// DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error)
func (m *MyStarkContract) DepositEth(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int) (*types.Transaction, error) {
	toAddress := common.BigToAddress(starkKey)
	tx := types.NewTransaction(opts.Nonce.Uint64(), toAddress, opts.Value, opts.GasLimit, opts.GasPrice, []byte{})
	return tx, nil
}

func (m *MyStarkContract) DepositERC20(opts *bind.TransactOpts, starkKey *big.Int, assetType *big.Int, vaultId *big.Int, quantizedAmount *big.Int) (*types.Transaction, error) {
	tx := &types.Transaction{}
	tx.SetTime(time.Now())
	return tx, nil
}

func (m *MyStarkContract) GetWithdrawalBalance(opts *bind.CallOpts, ownerKey *big.Int, assetId *big.Int) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (m *MyStarkContract) Withdraw(opts *bind.TransactOpts, ownerKey *big.Int, assetType *big.Int) (*types.Transaction, error) {
	tx := &types.Transaction{}
	tx.SetTime(time.Now())
	return tx, nil
}

func TestDepositFromEthereumNetwork(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	chainID := big.NewInt(1337)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	assert.NoError(t, err)

	balance := "10"

	// setting balance of 10
	balanceInWei := ToWei(balance, 18)

	// simulated backend
	simBack := simulated.NewBackend(types.GenesisAlloc{
		auth.From: {Balance: balanceInWei},
	}, simulated.WithBlockGasLimit(1000000))

	// ethclient
	ethClient := simBack.Client()

	client, err := New(TESTNET)
	assert.NoError(t, err)

	client.jwtToken = "token"
	client.refreshToken = "refresh"
	client.ethClient = ethClient

	t.Run("testing ethereum balance", func(t *testing.T) {
		ethBalance, err := getTokenBalance(context.TODO(), client.ethClient, auth.From.String(), ETH, "18", "")
		assert.NoError(t, err)
		assert.Equal(t, balance, ethBalance.String())
	})

	testCases := []struct {
		name         string
		roundTripper roundTripFunc
		coinStatus   CoinStatusResponse
		amount       float64
		want         CryptoDepositResponse
		timeout      time.Duration
		wantErr      bool
	}{
		{
			name: "Successful deposit",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == VAULTID_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "",
							"payload": {
							  "id": 317,
							  "coin": "eth"
							}
						  }`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": ""
					}`)),
				}, nil
			}),
			coinStatus: CoinStatusResponse{
				Status:  "success",
				Message: "retrieved successfully",
				Payload: map[string]CoinStatusPayload{
					"ethereum": {
						Symbol:            ETH,
						Quanitization:     "10",
						Decimal:           "8",
						BlockchainDecimal: "18",
						TokenContract:     "0x1",
						StarkAssetID:      "0x2",
					},
				},
			},
			amount: 0.001,
			want: CryptoDepositResponse{
				Status:  "success",
				Message: "",
				Payload: "0x4eb4ea28fbe61dc59abd6dfaee7e70201c8200bfcae8c9fcc7e190d72bf9e7f1",
			},
			wantErr: false,
		},

		{
			name: "Insufficient balance",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == VAULTID_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "",
							"payload": {
							  "id": 317,
							  "coin": "eth"
							}
						  }`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": ""
					}`)),
				}, nil
			}),
			coinStatus: CoinStatusResponse{
				Status:  "success",
				Message: "retrieved successfully",
				Payload: map[string]CoinStatusPayload{
					"ethereum": {
						Symbol:            ETH,
						Quanitization:     "10",
						Decimal:           "8",
						BlockchainDecimal: "18",
						TokenContract:     "0x1",
						StarkAssetID:      "0x2",
					},
				},
			},
			amount:  200,
			wantErr: true,
		},
	}

	privateKeyString := hex.EncodeToString(crypto.FromECDSA(privateKey))
	client.starkexContract = &MyStarkContract{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client.coinStatus = tc.coinStatus
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.DepositFromEthereumNetwork(ctx, "", auth.From.String(), privateKeyString, "0x2", tc.amount, ETH)
			if tc.wantErr {
				assert.Error(t, err)
				t.Logf("Error: %v", err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

type MyPolygonContract struct{}

func (m *MyPolygonContract) DepositNative(opts *bind.TransactOpts) (*types.Transaction, error) {
	toAddress := common.HexToAddress("0x1")
	tx := types.NewTransaction(opts.Nonce.Uint64(), toAddress, opts.Value, opts.GasLimit, opts.GasPrice, []byte{})
	return tx, nil
}

func (m *MyPolygonContract) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	toAddress := common.HexToAddress("0x1")
	tx := types.NewTransaction(opts.Nonce.Uint64(), toAddress, opts.Value, opts.GasLimit, opts.GasPrice, []byte{})
	return tx, nil
}
