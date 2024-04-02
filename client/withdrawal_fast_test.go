package client

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartFastWithdrawal(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          StartFastWithdrawalRequest
		want         StartFastWithdrawalResponse
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
			opt: StartFastWithdrawalRequest{
				Currency: ETH,
				Amount:   100,
			},
			want: StartFastWithdrawalResponse{
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
			opt:     StartFastWithdrawalRequest{},
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
			opt:     StartFastWithdrawalRequest{},
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

			got, err := client.startFastWithdrawal(ctx, tc.opt)
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

func TestProcessFastWithdrawal(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          ProcessFastWithdrawalRequest
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
			opt: ProcessFastWithdrawalRequest{
				MsgHash: "0x2",
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
			opt:     ProcessFastWithdrawalRequest{},
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
			opt:     ProcessFastWithdrawalRequest{},
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

			got, err := client.processFastWithdrawal(ctx, tc.opt)
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

// func TestFastWithdrawal(t *testing.T) {
// 	// network config endpoint
// 	// coin endpoint
// 	// start fast withdrawal endpoint
// 	// process fast withdrawal endpoint

// 	testCases := []struct {
// 		name         string
// 		roundTripper roundTripFunc
// 		amount       float64
// 		currency     Currency
// 		network      Network
// 		want         ValidateWithdrawalResponse
// 		timeout      time.Duration
// 		wantErr      bool
// 	}{
// 		{
// 			name: "Successful deposit",
// 			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
// 				if r.URL.Path == NETWORK_CONFIG_ENDPOINT {
// 					return &http.Response{
// 						StatusCode: http.StatusOK,
// 						Body: io.NopCloser(strings.NewReader(`{
// 							"status": "success",
// 							"message": "",
// 							"payload": {
// 							  "id": 317,
// 							  "coin": "eth"
// 							}
// 						  }`)),
// 					}, nil
// 				}

// 				return &http.Response{
// 					StatusCode: http.StatusOK,
// 					Body: io.NopCloser(strings.NewReader(`{
// 						"status": "success",
// 						"message": ""
// 					}`)),
// 				}, nil
// 			}),
// 			coinStatus: CoinStatusResponse{
// 				Status:  "success",
// 				Message: "retrieved successfully",
// 				Payload: map[string]CoinStatusPayload{
// 					"ethereum": {
// 						Symbol:            ETH,
// 						Quanitization:     "10",
// 						Decimal:           "8",
// 						BlockchainDecimal: "18",
// 						TokenContract:     "0x1",
// 						StarkAssetID:      "0x2",
// 					},
// 				},
// 			},
// 			amount: 0.001,
// 			want: CryptoDepositResponse{
// 				Status:  "success",
// 				Message: "",
// 				Payload: "0x4eb4ea28fbe61dc59abd6dfaee7e70201c8200bfcae8c9fcc7e190d72bf9e7f1",
// 			},
// 			wantErr: false,
// 		},
// 	}
// }

func TestListFastWithdrawal(t *testing.T) {
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

			got, err := client.ListFastWithdrawal(ctx, tc.opt)
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
