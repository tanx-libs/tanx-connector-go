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

func TestDepositFromPolygonNetwork(t *testing.T) {
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

	// polygonClient
	polygonClient := simBack.Client()

	client, err := New(TESTNET)
	assert.NoError(t, err)

	client.jwtToken = "token"
	client.refreshToken = "refresh"
	client.polygonClient = polygonClient

	t.Run("testing matic balance", func(t *testing.T) {
		maticBalance, err := getTokenBalance(context.TODO(), client.polygonClient, auth.From.String(), MATIC, "18", "")
		assert.NoError(t, err)
		assert.Equal(t, balance, maticBalance.String())
	})

	testCases := []struct {
		name          string
		roundTripper  roundTripFunc
		polygonConfig NetworkConfigData
		amount        float64
		want          CryptoDepositResponse
		timeout       time.Duration
		wantErr       bool
	}{
		{
			name: "Successful deposit",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": ""
					}`)),
				}, nil
			}),
			polygonConfig: NetworkConfigData{
				DepositContract: "0x1",
				Tokens: map[Currency]CoinToken{
					MATIC: {
						BlockchainDecimal:                  "18",
						TokenContract:                      "0x1",
						MaxFastWithdrawalForPlatformPerDay: "100",
						MaxFastWithdrawalForUserPerDay:     "100",
					},
				},
				AllowedTokensForDeposit:         []Currency{MATIC},
				AllowedTokensForDepositFrontend: []Currency{MATIC},
				AllowedTokensForFastWd:          []Currency{MATIC},
				AllowedTokensForFastWdFrontend:  []Currency{MATIC},
			},
			amount: 0.01,
			want: CryptoDepositResponse{
				Status:  "success",
				Message: "",
				Payload: "0x148766d43785848046189d134b65e6c2b7c5f5d2d32171856b0c317374af8d0b",
			},
			wantErr: false,
		},
		{
			name: "Currency not found error",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": ""
					}`)),
				}, nil
			}),
			polygonConfig: NetworkConfigData{
				DepositContract:                 "0x1",
				Tokens:                          map[Currency]CoinToken{},
				AllowedTokensForDeposit:         []Currency{MATIC},
				AllowedTokensForDepositFrontend: []Currency{MATIC},
				AllowedTokensForFastWd:          []Currency{MATIC},
				AllowedTokensForFastWdFrontend:  []Currency{MATIC},
			},
			amount: 0.01,
			want: CryptoDepositResponse{
				Status:  "success",
				Message: "",
				Payload: "0x148766d43785848046189d134b65e6c2b7c5f5d2d32171856b0c317374af8d0b",
			},
			wantErr: true,
		},
	}

	privateKeyString := hex.EncodeToString(crypto.FromECDSA(privateKey))
	client.polygonContract = &MyPolygonContract{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client.polygonConfig = tc.polygonConfig
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, _, err := client.DepositFromPolygonNetwork(ctx, "", auth.From.String(), privateKeyString, "0x2", MATIC, tc.amount)
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
