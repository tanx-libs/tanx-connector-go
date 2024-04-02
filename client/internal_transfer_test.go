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

func TestInternalTransferInitiate(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          InternalTransferInitiateRequest
		want         InternalTransferInitiateResponse
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
			opt: InternalTransferInitiateRequest{
				OrganizationKey:    "org_key",
				ApiKey:             "api_key",
				ClientReferenceId:  "123456",
				Currency:           ETH,
				Amount:             100,
				DestinationAddress: "0x3",
			},
			want: InternalTransferInitiateResponse{
				Status:  SUCCESS,
				Message: "",
				Payload: InternalTransferInitiatePayload{
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
			opt:     InternalTransferInitiateRequest{},
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
			opt:     InternalTransferInitiateRequest{},
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

			got, err := client.internalTransferInitiate(ctx, tc.opt)
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

func TestInternalTransferProcess(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          InternalTransferProcessRequest
		want         InternalTransferProcessResponse
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
			opt: InternalTransferProcessRequest{
				OrganizationKey: "org_key",
				ApiKey:          "api_key",
				Signature: Signature{
					R: "0x1",
					S: "0x2",
				},
				Nonce:   1,
				MsgHash: "0x2",
			},
			want: InternalTransferProcessResponse{
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
			opt:     InternalTransferProcessRequest{},
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

			got, err := client.internalTransferProcess(ctx, tc.opt)
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

func TestInternalTransferCreate(t *testing.T) {
	testCases := []struct {
		name            string
		roundTripper    roundTripFunc
		timeout         time.Duration
		starkPrivateKey string
		req             InternalTransferInitiateRequest
		want            InternalTransferProcessResponse
		wantErr         bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == INTERNAL_TRANSFER_INITIATE_ENDPOINT {
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
					Body:       io.NopCloser(strings.NewReader(`{"status": "success"}`)),
				}, nil
			}),
			starkPrivateKey: "0x1",
			req: InternalTransferInitiateRequest{
				OrganizationKey:    "org_key",
				ApiKey:             "api_key",
				ClientReferenceId:  "1",
				Currency:           ETH,
				Amount:             0.01,
				DestinationAddress: "0x2",
			},
			want: InternalTransferProcessResponse{
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
			client.refreshToken = "refresh"

			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.InternalTransferCreate(ctx, tc.starkPrivateKey, tc.req)
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

func TestInternalTransferGet(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		clientID     string
		want         InternalTransferGetResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"status": "success"}`)),
				},
			},
			clientID: "1",
			want: InternalTransferGetResponse{
				Status: SUCCESS,
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

			got, err := client.InternalTransferGet(ctx, tc.clientID)
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

func TestInternalTransferUser(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          InternalTransferUserRequest
		want         InternalTransferUserResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"status": "success"}`)),
				},
			},
			opt: InternalTransferUserRequest{
				OrganizationKey:    "prg_key",
				ApiKey:             "api_key",
				DestinationAddress: "0x1",
			},
			want: InternalTransferUserResponse{
				Status: SUCCESS,
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

			got, err := client.InternalTransferUser(ctx, tc.opt)
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

func TestInternalTransferList(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		opt          InternalTransferListRequest
		want         InternalTransferListResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"status": "success"}`)),
				},
			},
			opt: InternalTransferListRequest{
				Limit: 1,
				Offset: 2,
			},
			want: InternalTransferListResponse{
				Status: SUCCESS,
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

			got, err := client.InternalTransferList(ctx, tc.opt)
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


