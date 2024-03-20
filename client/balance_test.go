package client

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBalance(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		currency     Currency
		want         BalanceResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "Working fine!",
						"payload": {
							"currency": "btc",
							"balance": "1.0",
							"locked": "0.0",
							"deposit_address": "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2"
						}
					}`)),
				},
			},
			currency: BTC,
			want: BalanceResponse{
				Status:  "success",
				Message: "Working fine!",
				Payload: BalancePayload{
					Currency:       "btc",
					Balance:        "1.0",
					Locked:         "0.0",
					DepositAddress: "1BvBMSEYstWetqTFn5Au4m4GFg7xJaNVN2",
				},
			},
			wantErr: false,
		},
		{
			name: "HTTP protocol error",
			roundTripper: &MockRoundTripper{
				Err: errors.New("mock error"),
			},
			wantErr: true,
		},
		{
			name: "JSON decoding error",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"status": 1}`)),
				},
			},
			wantErr: true,
		},
		{
			name: "Context timeout error",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"status": "success"}`)),
				},
				Delay: time.Second * 2,
			},
			timeout: time.Nanosecond,
			wantErr: true,
		},

		{
			name: "Invalid currency",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error"}`)),
				},
			},
			currency: "invalid",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client, err := New(TESTNET)
			assert.NoError(t, err)

			client.jwtToken = "token"
			client.refreshToken = "refresh"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.Balance(ctx, tc.currency)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
