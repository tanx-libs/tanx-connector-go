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

func TestPNL(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         PNLResponse
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
						"payload": [
							{
								"currency": "btc",
								"pnl_currency": "usdc",
								"total_credit": "30.2311686",
								"total_debit": "8.6779552",
								"total_credit_value": "951376.399690608",
								"total_debit_value": "263030.247889309",
								"average_buy_price": "31470.05040653995757213302",
								"average_sell_price": "30310.16429876349211851197",
								"average_balance_price": "31446.4028484392969497",
								"total_balance_value": "677771.031254780024102"
							}
						]
					}`)),
				},
			},
			want: PNLResponse{
				Status:  "success",
				Message: "Working fine!",
				Payload: []PNLPayload{
					{
						Currency:            "btc",
						PnlCurrency:         "usdc",
						TotalCredit:         "30.2311686",
						TotalDebit:          "8.6779552",
						TotalCreditValue:    "951376.399690608",
						TotalDebitValue:     "263030.247889309",
						AverageBuyPrice:     "31470.05040653995757213302",
						AverageSellPrice:    "30310.16429876349211851197",
						AverageBalancePrice: "31446.4028484392969497",
						TotalBalanceValue:   "677771.031254780024102",
					},
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

		// status 400
		{
			name: "Status 400",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "error",
						"message": "Bad Request",
						"payload": ""
					}`)),
				},
			},
			wantErr: true,
		},

		// status 500
		{
			name: "Status 500",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "error",
						"message": "Internal Server Error",
						"payload": ""
					}`)),
				},
			},
			wantErr: true,
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

			got, err := client.PNL(ctx)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

