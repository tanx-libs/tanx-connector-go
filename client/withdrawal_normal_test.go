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

// todo GetPendingNormalWithdrawalAmountByCoin

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

