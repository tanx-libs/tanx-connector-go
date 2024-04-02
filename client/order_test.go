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
	"github.com/tanx-libs/tanx-connector-go/crypto_cpp"
)

func TestSign(t *testing.T) {
	r, s := crypto_cpp.Sign("0x1", "0x2", "0x3")
	public_key := crypto_cpp.GetPublicKey("0x1")
	ok := crypto_cpp.Verify(public_key, "0x2", r, s)
	assert.True(t, ok)
}

func TestOrderNonce(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         OrderNonceResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "order nonce created successfully, please sign the msg_hash to place the order",
						"payload": {
						  "nonce": 931,
						  "msg_hash": "0x5f128faa6e96360629b1fcb4f24fe9bc547f0d79c7a3aed056d824baa786755"
						}
					  }`)),
				},
			},
			want: OrderNonceResponse{
				Status:  "success",
				Message: "order nonce created successfully, please sign the msg_hash to place the order",
				Payload: OrderNoncePayload{
					Nonce:   931,
					MsgHash: "0x5f128faa6e96360629b1fcb4f24fe9bc547f0d79c7a3aed056d824baa786755",
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

			got, err := client.orderNonce(ctx, OrderNonceRequest{})
			if tc.wantErr {
				assert.Error(t, err)
				t.Logf("%+v", err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestOrderCreate(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper roundTripFunc
		params       string
		timeout      time.Duration
		want         OrderCreateResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == ORDER_NONCE_ENDPOINT {
					t.Logf("URL: %v", r.URL.Path)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "order nonce created successfully, please sign the msg_hash to place the order",
							"payload": {
							  "nonce": 931,
							  "msg_hash": "0x5f128faa6e96360629b1fcb4f24fe9bc547f0d79c7a3aed056d824baa786755"
							}
						  }`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "Created Order Successfully",
						"payload": {
						  "id": 904,
						  "uuid": "6c79cde4-1e8f-4446-9b4a-ba176d50f08b",
						  "side": "sell",
						  "ord_type": "limit",
						  "price": "29580.51",
						  "avg_price": "0.0",
						  "state": "pending",
						  "market": "btcusdt",
						  "created_at": "2022-05-27T12:36:57+02:00",
						  "updated_at": "2022-05-27T12:36:57+02:00",
						  "origin_volume": "0.015",
						  "remaining_volume": "0.015",
						  "executed_volume": "0.0",
						  "maker_fee": "0.0",
						  "taker_fee": "0.0",
						  "trades_count": 0
						}
					  }`)),
				}, nil
			}),
			params: "0x1",

			want: OrderCreateResponse{
				Status:  "success",
				Message: "Created Order Successfully",
				Payload: OrderPayload{
					ID:              904,
					UUID:            "6c79cde4-1e8f-4446-9b4a-ba176d50f08b",
					Side:            "sell",
					OrdType:         "limit",
					Price:           "29580.51",
					AvgPrice:        "0.0",
					State:           "pending",
					Market:          "btcusdt",
					CreatedAt:       "2022-05-27T12:36:57+02:00",
					UpdatedAt:       "2022-05-27T12:36:57+02:00",
					OriginVolume:    "0.015",
					RemainingVolume: "0.015",
					ExecutedVolume:  "0.0",
					MakerFee:        "0.0",
					TakerFee:        "0.0",
					TradesCount:     0,
				},
			},
			wantErr: false,
		},

		{
			name: "Successful roundtrip",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == ORDER_NONCE_ENDPOINT {
					t.Logf("URL: %v", r.URL.Path)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "order nonce created successfully, please sign the msg_hash to place the order",
							"payload": {
							  "nonce": 931,
							  "msg_hash": "0x5f128faa6e96360629b1fcb4f24fe9bc547f0d79c7a3aed056d824baa786755"
							}
						  }`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "Created Order Successfully",
						"payload": {
						  "id": 904,
						  "uuid": "6c79cde4-1e8f-4446-9b4a-ba176d50f08b",
						  "side": "sell",
						  "ord_type": "limit",
						  "price": "29580.51",
						  "avg_price": "0.0",
						  "state": "pending",
						  "market": "btcusdt",
						  "created_at": "2022-05-27T12:36:57+02:00",
						  "updated_at": "2022-05-27T12:36:57+02:00",
						  "origin_volume": "0.015",
						  "remaining_volume": "0.015",
						  "executed_volume": "0.0",
						  "maker_fee": "0.0",
						  "taker_fee": "0.0",
						  "trades_count": 0
						}
					  }`)),
				}, nil
			}),
			params: "01",

			want: OrderCreateResponse{
				Status:  "success",
				Message: "Created Order Successfully",
				Payload: OrderPayload{
					ID:              904,
					UUID:            "6c79cde4-1e8f-4446-9b4a-ba176d50f08b",
					Side:            "sell",
					OrdType:         "limit",
					Price:           "29580.51",
					AvgPrice:        "0.0",
					State:           "pending",
					Market:          "btcusdt",
					CreatedAt:       "2022-05-27T12:36:57+02:00",
					UpdatedAt:       "2022-05-27T12:36:57+02:00",
					OriginVolume:    "0.015",
					RemainingVolume: "0.015",
					ExecutedVolume:  "0.0",
					MakerFee:        "0.0",
					TakerFee:        "0.0",
					TradesCount:     0,
				},
			},
			wantErr: false,
		},

		{
			name: "Invalid msg hash",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == ORDER_NONCE_ENDPOINT {
					t.Logf("URL: %v", r.URL.Path)
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "order nonce created successfully, please sign the msg_hash to place the order",
							"payload": {
							  "nonce": 931,
							  "msg_hash": ""
							}
						  }`)),
					}, nil
				}

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "Created Order Successfully",
						"payload": {
						  "id": 904,
						  "uuid": "6c79cde4-1e8f-4446-9b4a-ba176d50f08b",
						  "side": "sell",
						  "ord_type": "limit",
						  "price": "29580.51",
						  "avg_price": "0.0",
						  "state": "pending",
						  "market": "btcusdt",
						  "created_at": "2022-05-27T12:36:57+02:00",
						  "updated_at": "2022-05-27T12:36:57+02:00",
						  "origin_volume": "0.015",
						  "remaining_volume": "0.015",
						  "executed_volume": "0.0",
						  "maker_fee": "0.0",
						  "taker_fee": "0.0",
						  "trades_count": 0
						}
					  }`)),
				}, nil
			}),
			params:  "0x1",
			wantErr: true,
		},



		// status code 404
		{
			name: "404 status code",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       io.NopCloser(strings.NewReader(`{"status": "error", "message": "Not found"}`)),
				}, nil
			}),
			params:  "0x1",
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

			got, err := client.OrderCreate(ctx, tc.params, OrderNonceRequest{})
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

// get order
func TestOrderGet(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         OrderGetResponse
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
							"id": 1
						}
					}`)),
				},
			},
			want: OrderGetResponse{
				Status:  "success",
				Message: "Working fine!",
				Payload: OrderPayload{
					ID: 1,
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

			got, err := client.OrderGet(ctx, 1)
			if tc.wantErr {
				assert.Error(t, err)
				t.Logf("%+v", err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

// list orders
func TestOrdersList(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         OrdersListResponse
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
								"id": 1
							}
						]
					}`)),
				},
			},
			want: OrdersListResponse{
				Status:  "success",
				Message: "Working fine!",
				Payload: []OrderPayload{
					{
						ID: 1,
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

			got, err := client.OrdersList(ctx, OrderListOptions{})
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

// cancel orders
func TestOrderCancel(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         OrderCancelResponse
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
							"id": 1
						}
					}`)),
				},
			},
			want: OrderCancelResponse{
				Status:  "success",
				Message: "Working fine!",
				Payload: OrderPayload{
					ID: 1,
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

			got, err := client.OrderCancel(ctx, 1)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

// list trades
func TestTradesList(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         TradesListResponse
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
								"id": 1
							},
							{
								"id": 2
							}
						]
					}`)),
				},
			},
			want: TradesListResponse{
				Status:  "success",
				Message: "Working fine!",
				Payload: []TradePayload{
					{
						ID: 1,
					},
					{
						ID: 2,
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

			got, err := client.TradesList(ctx, TradesListOptions{})
			if tc.wantErr {
				assert.Error(t, err)
				t.Logf("%+v", err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
