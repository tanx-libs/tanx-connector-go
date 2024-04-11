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
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/assert"
)

func TestStartNormalWithdrawal(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          StartNormalWithdrawalRequest
		want         StartNormalWithdrawalResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "",
						"payload": {
						  "msg_hash": "0x2",
						  "nonce": 1
						}
					  }`)),
				},
			},
			opt: StartNormalWithdrawalRequest{
				Currency: ETH,
				Amount:   100,
			},
			want: StartNormalWithdrawalResponse{
				Status:  SUCCESS,
				Message: "",
				Payload: StartNormalWithdrawalPaylaod{
					MsgHash: "0x2",
					Nonce:   1,
				},
			},
			wantErr: false,
		},

		// 404 status code
		{
			name: "404 status code",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error", "message": "Not found"}`)),
				},
			},
			opt:     StartNormalWithdrawalRequest{},
			wantErr: true,
		},

		// 500 status code
		{
			name: "500 status code",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error", "message": "Not found"}`)),
				},
			},
			opt:     StartNormalWithdrawalRequest{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := New(TESTNET)
			assert.NoError(t, err)

			client.jwtToken = "jwt_token"
			client.refreshToken = "refresh_token"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.startNormalWithdrawal(ctx, tc.opt)
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

func TestValidateNormalWithdrawal(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          ValidateNormalWithdrawalRequest
		want         ValidateWithdrawalResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": ""
					  }`)),
				},
			},
			opt: ValidateNormalWithdrawalRequest{
				MessageHash: "0x2",
				Nonce:       1,
			},
			want: ValidateWithdrawalResponse{
				Status:  SUCCESS,
				Message: "",
			},
			wantErr: false,
		},

		// 404 status code
		{
			name: "404 status code",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error", "message": "Not found"}`)),
				},
			},
			opt:     ValidateNormalWithdrawalRequest{},
			wantErr: true,
		},

		// 500 status code
		{
			name: "500 status code",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error", "message": "Not found"}`)),
				},
			},
			opt:     ValidateNormalWithdrawalRequest{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := New(TESTNET)
			assert.NoError(t, err)

			client.jwtToken = "jwt_token"
			client.refreshToken = "refresh_token"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.validateNormalWithdrawal(ctx, tc.opt)
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

// todo InitiateNormalWithdrawal
func TestInitiateNormalWithdrawal(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper roundTripFunc
		timeout      time.Duration
		opt          ValidateNormalWithdrawalRequest
		want         ValidateWithdrawalResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: roundTripFunc(func(req *http.Request) (*http.Response, error) {
				if req.URL.Path == COIN_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "Retrieval Successful",
							"payload": {
							  "ethereum": {
								"name": "Ethereum",
								"symbol": "eth",
								"type": "crypto",
								"fee": "0.0001",
								"quanitization": "10",
								"stark_asset_id": "0x2705737cd248ac819034b5de474c8f0368224f72a0fda9e031499d519992d9e",
								"token_contract": "0xcfc50728f5b5ee794ff6c19c61e235f1a8de2fa4",
								"news_name": "ETH",
								"trade_fee": "0.00",
								"chart_name": "ethereum",
								"minimum_order": "0.0002",
								"minimum_withdraw": "0.04",
								"maximum_order": "2",
								"minimum_deposit": "0.01",
								"frontend_visibility": true,
								"starkex_deposits_enabled": true,
								"starkex_deposits_enabled_frontend": true,
								"fast_withdrawals_enabled": true,
								"fast_withdrawal_fee": "0.0001",
								"max_fast_withdrawal_for_platform_per_day": "5",
								"max_fast_withdrawal_for_user_per_day": "2",
								"min_fast_withdrawal": "0.001",
								"decimal": "8",
								"color": "#617EEA",
								"trade_decimal": "4",
								"blockchain_decimal": "18",
								"chart": "BITFINEX:ETHUSD|12M",
								"logo": "https://brine-mainnet-fe-assets.s3.ap-southeast-1.amazonaws.com/ETHEREUM.png",
								"description": "Ethereum is an open platform that enables developers to build and deploy decentralized applications such as smart contracts and other complex legal and financial applications. "
							  },
							  "tether": {
								"name": "Tether",
								"symbol": "usdt",
								"decimal": "6",
								"fee": "10",
								"quanitization": "6",
								"type": "crypto",
								"stark_asset_id": "0x1b6cf04337900caeab4f74bdfe9d5faa83a4a02a0494f8899f1d5a3715df8a9",
								"token_contract": "0x817fF250EB13e85d3B6Dd7Da3B901573Cf1D28fb",
								"trade_fee": "0.00",
								"color": "#26A17B",
								"minimum_withdraw": "0",
								"news_name": "USDT",
								"chart_name": "tether",
								"chart": "KRAKEN:USDTUSD|12M",
								"minimum_order": "1",
								"maximum_order": "2500",
								"trade_decimal": "2",
								"blockchain_decimal": "18",
								"minimum_deposit": "1",
								"frontend_visibility": true,
								"starkex_deposits_enabled": true,
								"starkex_deposits_enabled_frontend": true,
								"fast_withdrawals_enabled": true,
								"fast_withdrawal_fee": "1",
								"max_fast_withdrawal_for_platform_per_day": "5000",
								"max_fast_withdrawal_for_user_per_day": "1000",
								"min_fast_withdrawal": "10",
								"logo": "https://brine-mainnet-fe-assets.s3.ap-southeast-1.amazonaws.com/TETHER.png",
								"description": "Tether is a controversial cryptocurrency with tokens issued by Tether Limited. It formerly claimed that each token was backed by one United States dollar, but on 14 March 2019 changed the backing to include loans to affiliate companies."
							  },
							  "bitcoin": {
								"name": "Bitcoin",
								"symbol": "btc",
								"decimal": "8",
								"fee": "0.0006",
								"quanitization": "8",
								"type": "crypto",
								"stark_asset_id": "0x16dd3888b66b40e6da2c971208cbd4688e8aa8b96e13c1b6f7e17048f0d8c2c",
								"token_contract": "0x29F53FFC222CD75805712c2dF76fA683c4b1e967",
								"trade_fee": "0.00",
								"color": "#F7931A",
								"minimum_withdraw": "0.001",
								"news_name": "BTC",
								"chart_name": "bitcoin",
								"chart": "BITFINEX:BTCUSD|12M",
								"minimum_order": "0.00001",
								"maximum_order": "10",
								"trade_decimal": "4",
								"blockchain_decimal": "18",
								"minimum_deposit": "1",
								"frontend_visibility": true,
								"starkex_deposits_enabled": true,
								"starkex_deposits_enabled_frontend": true,
								"fast_withdrawals_enabled": true,
								"fast_withdrawal_fee": "0.0006",
								"max_fast_withdrawal_for_platform_per_day": "0.19",
								"max_fast_withdrawal_for_user_per_day": "0.038",
								"min_fast_withdrawal": "0.00076",
								"logo": "https://brine-mainnet-fe-assets.s3.ap-southeast-1.amazonaws.com/bitcoin.png",
								"description": "The world’s first cryptocurrency, Bitcoin is stored and exchanged securely on the internet through a digital ledger known as a blockchain. Bitcoins are divisible into smaller units known as satoshis — each satoshi is worth 0.00000001 bitcoin."
							  },
							  "usdc": {
								"name": "USDC",
								"symbol": "usdc",
								"decimal": "6",
								"fee": "10",
								"quanitization": "6",
								"type": "crypto",
								"stark_asset_id": "0x38962d21e77a22fae39b01463d8df1f5b22b91bd6f97ec5574d47c6f69328cf",
								"token_contract": "0x6858aF063582c7d355bFbc711251ADF4Efbc4fd0",
								"trade_fee": "0.00",
								"color": "#26A17B",
								"minimum_withdraw": "0",
								"news_name": "USDC",
								"chart_name": "usdc",
								"chart": "KRAKEN:USDCUSD|12M",
								"minimum_order": "1",
								"maximum_order": "250000",
								"trade_decimal": "2",
								"blockchain_decimal": "18",
								"minimum_deposit": "1",
								"frontend_visibility": true,
								"starkex_deposits_enabled": true,
								"starkex_deposits_enabled_frontend": true,
								"fast_withdrawals_enabled": true,
								"fast_withdrawal_fee": "1",
								"max_fast_withdrawal_for_platform_per_day": "5000",
								"max_fast_withdrawal_for_user_per_day": "1000",
								"min_fast_withdrawal": "10",
								"logo": "https://brine-mainnet-fe-assets.s3.ap-southeast-1.amazonaws.com/USDC.png",
								"description": "USD Coin (known by its ticker USDC) is a stablecoin that is pegged to the U.S. dollar on a 1:1 basis. Every unit of this cryptocurrency in circulation is backed up by $1 that is held in reserve, in a mix of cash and short-term U.S. Treasury bonds. The Centre consortium, which is behind this asset, says USDC is issued by regulated financial institutions."
							  },
							  "matic": {
								"name": "Matic",
								"symbol": "matic",
								"decimal": "8",
								"fee": "10",
								"quanitization": "8",
								"type": "crypto",
								"stark_asset_id": "0x2adb57bf5c9f3613b96fd1d7517cac044d30a1f8b543ec74e2683290981c9d2",
								"token_contract": "0x68e80f254761010DaF2B45E698E059073F2dab82",
								"trade_fee": "0.00",
								"color": "#26A17B",
								"minimum_withdraw": "0",
								"news_name": "MATIC",
								"chart_name": "matic",
								"chart": "KRAKEN:MATICUSD|12M",
								"minimum_order": "1",
								"maximum_order": "250000",
								"trade_decimal": "2",
								"blockchain_decimal": "18",
								"minimum_deposit": "1",
								"frontend_visibility": true,
								"starkex_deposits_enabled": true,
								"starkex_deposits_enabled_frontend": true,
								"fast_withdrawals_enabled": true,
								"fast_withdrawal_fee": "28.6",
								"max_fast_withdrawal_for_platform_per_day": "20000",
								"max_fast_withdrawal_for_user_per_day": "10000",
								"min_fast_withdrawal": "0.01",
								"logo": "https://brine-mainnet-fe-assets.s3.ap-southeast-1.amazonaws.com/MATIC.png",
								"description": "MATIC is the native cryptocurrency of Polygon, a layer 2 scaling solution for Ethereum. It facilitates fast and low-cost transactions within the Polygon ecosystem, providing an efficient alternative to Ethereum's mainnet. As an ERC-20 token, MATIC is used for transaction fees, governance participation, and interaction with decentralized applications (dApps) on the Polygon network. Its purpose is to enhance scalability and usability, making it attractive for developers and users seeking a more seamless blockchain experience."
							  }
							}
						  }`)),
					}, nil
				}

				if req.URL.Path == START_NORMAL_WITHDRAWAL_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "",
						"payload": {
						  "msg_hash": "0x2",
						  "nonce": 1
						}
					  }`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success"
					  }`)),
				}, nil
			}),

			opt: ValidateNormalWithdrawalRequest{
				MessageHash: "0x2",
				Nonce:       1,
			},
			want: ValidateWithdrawalResponse{
				Status: SUCCESS,
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := New(TESTNET)
			assert.NoError(t, err)

			client.jwtToken = "jwt_token"
			client.refreshToken = "refresh_token"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.validateNormalWithdrawal(ctx, tc.opt)
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

// todo GetPendingNormalWithdrawalAmountByCoin
func TestGetPendingNormalWithdrawalAmountByCoin(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	chainID := big.NewInt(1337)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	assert.NoError(t, err)

	balance := "10"

	// setting balance of 10
	balanceInWei := toWei(balance, 18)

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
		name       string
		currency   Currency
		coinStatus CoinStatusResponse
		want       string
		wantErr    bool
	}{
		{
			name:     "success",
			currency: ETH,
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
			want:    "0.000000",
			wantErr: false,
		},
	}

	// privateKeyString := hex.EncodeToString(crypto.FromECDSA(privateKey))
	client.starkexContract = &MyStarkContract{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client.coinStatus = tc.coinStatus

			got, err := client.GetPendingNormalWithdrawalAmountByCoin(context.Background(), "url", auth.From.String(), tc.currency)
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

func TestCompleteNormalWithdrawal(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)

	chainID := big.NewInt(1337)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	assert.NoError(t, err)

	balance := "10"

	// setting balance of 10
	balanceInWei := toWei(balance, 18)

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

func TestListNormalWithdrawal(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          ListWithdrawalRequest
		want         ListWithdrawalResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": ""
					  }`)),
				},
			},
			opt: ListWithdrawalRequest{
				Page:    1,
				Network: ETHEREUM,
			},
			want: ListWithdrawalResponse{
				Status:  SUCCESS,
				Message: "",
			},
			wantErr: false,
		},

		// 404 status code
		{
			name: "404 status code",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error", "message": "Not found"}`)),
				},
			},
			opt:     ListWithdrawalRequest{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := New(TESTNET)
			assert.NoError(t, err)

			client.jwtToken = "token"
			client.refreshToken = "token"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.ListNormalWithdrawal(ctx, tc.opt)
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
