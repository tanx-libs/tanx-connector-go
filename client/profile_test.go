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

func TestProfile(t *testing.T) {
	testCases := []struct {
		name         string
		roundTripper http.RoundTripper
		timeout      time.Duration
		want         ProfileResponse
		wantErr      bool
	}{
		{
			name: "Successful roundtrip",
			roundTripper: &MockRoundTripper{
				Response: &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success",
						"message": "Successful",
						"payload": {
						  "name": "USER NAME",
						  "customer_id": "27",
						  "img": null,
						  "username": "guthal",
						  "stark_key": "0x70f41ce6797eb444c9dc95a907....8aa592adf8f1fe3ab75317f7096d38"
						}
					  }`)),
				},
			},
			want: ProfileResponse{
				Status:  "success",
				Message: "Successful",
				Payload: ProfilePayload{
					Name:       "USER NAME",
					CustomerID: "27",
					Img:        nil,
					Username:   "guthal",
					StarkKey:   "0x70f41ce6797eb444c9dc95a907....8aa592adf8f1fe3ab75317f7096d38",
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
			client.refreshToken = "refrsh"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.Profile(ctx)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}
