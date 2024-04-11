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

func TestFastWithdrawal(t *testing.T) {
	// network config endpoint
	// coin endpoint
	// start fast withdrawal endpoint
	// process fast withdrawal endpoint

	testCases := []struct {
		name         string
		roundTripper roundTripFunc
		amount       float64
		currency     Currency
		network      Network
		want         ValidateWithdrawalResponse
		timeout      time.Duration
		wantErr      bool
	}{
		{
			name: "token not found",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == NETWORK_CONFIG_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "Retrieval Successful",
							"payload": {
							  "network_config": {
								"ETHEREUM": {
								  "allowed_tokens_for_deposit": [],
								  "allowed_tokens_for_deposit_frontend": [],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								},
								"POLYGON": {
								  "deposit_contract": "0x09056dC8E09205eb04C78B9C33df4767B2325cF4",
								  "tokens": {
									"btc": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x8DB9D35eDFdd2fcEe07A0fa60E864dDCBC4eF68e",
									  "max_fast_withdrawal_for_platform_per_day": "10000",
									  "max_fast_withdrawal_for_user_per_day": "4000"
									},
									"matic": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010",
									  "max_fast_withdrawal_for_platform_per_day": "10000",
									  "max_fast_withdrawal_for_user_per_day": "4000"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x4d2548DAbF3d662110d70239Bc3531043984644D",
									  "max_fast_withdrawal_for_platform_per_day": "10000",
									  "max_fast_withdrawal_for_user_per_day": "4000"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"btc",
									"matic",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"btc",
									"matic",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"btc",
									"matic",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"btc",
									"matic",
									"usdt"
								  ]
								},
								"STARKNET": {
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x005a643907b9a4bc6a55e9069c4fd5fd1f5c79a22470690f75556c4736e34426"
									}
								  },
								  "allowed_tokens_for_deposit": [],
								  "allowed_tokens_for_deposit_frontend": [],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								},
								"SCROLL": {
								  "deposit_contract": "0x1C810252902F8e34c9aC724E558Ea99baa2a9a9B",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0xFFFF3e201311e450a0Bf8091124B4B22083514e3"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x69D0377CD7FbA97a662bEA6a7Ae194CDa8C5063A"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc",
									"usdt"
								  ]
								},
								"OPTIMISM": {
								  "deposit_contract": "0x4257742647821d8BbC086caC9c32f8F3328C709A",
								  "tokens": {
									"btc": {
									  "blockchain_decimal": "18",
									  "token_contract": "0xd9C7d45444Fcb9CE39f63478275eF2449380B7F1",
									  "max_fast_withdrawal_for_platform_per_day": "0.03",
									  "max_fast_withdrawal_for_user_per_day": "0.02"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x0b9C59563E7a3727eC363236D349aA526c882a3A"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"btc",
									"usdt"
								  ]
								},
								"ARBITRUM": {
								  "deposit_contract": "0x987405D0e208e801859eb4e9C9923f67F3C487f1",
								  "tokens": {
									"btc": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x3F180fCFd904D2d5047037e33B3f65F5eD312f33"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x7eE00F9D72422e7939f33537aB01e630A3dd95D3"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"btc",
									"usdt"
								  ]
								},
								"LINEA": {
								  "deposit_contract": "0x8Ec96d9bC4fEE06Fb668a0FBca66510a57C42b05",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x2318ED5b67781fa56EC3A677624deb3deE2CB504"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x0f40466610F20ed3C7cba31CBc5D7358C8824FCa"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010",
									  "max_fast_withdrawal_for_platform_per_day": "3",
									  "max_fast_withdrawal_for_user_per_day": "1.5"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ]
								},
								"MODE": {
								  "deposit_contract": "0x886dF268A7AD3dC9F7Ff1Bf197A5dF840DE397F1",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x79378E3E1BE2062861d23635A78A7aA8BE19ccc4"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010",
									  "max_fast_withdrawal_for_platform_per_day": "3",
									  "max_fast_withdrawal_for_user_per_day": "1.5"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x04655c0a6023889B3d8Da5950e6f41AeEc4772bA"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ]
								},
								"ZKPOLY": {
								  "deposit_contract": "",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x28bA3579945A67aD54D83Be0B0F9481C464d4Dc7"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								},
								"BERA": {
								  "deposit_contract": "0xa8a86215102c362ad1E871959522D3830BEd3dAd",
								  "tokens": {
									"bera": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x3736BF60A387585F7349A50BEb58aC0f7142b9c3"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								}
							  }
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
			amount:   0.001,
			currency: "somwthing random",
			network:  POLYGON,
			wantErr:  true,
		},

		{
			name: "Successful roundtrip",
			roundTripper: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if r.URL.Path == NETWORK_CONFIG_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "Retrieval Successful",
							"payload": {
							  "network_config": {
								"ETHEREUM": {
								  "allowed_tokens_for_deposit": [],
								  "allowed_tokens_for_deposit_frontend": [],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								},
								"POLYGON": {
								  "deposit_contract": "0x09056dC8E09205eb04C78B9C33df4767B2325cF4",
								  "tokens": {
									"btc": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x8DB9D35eDFdd2fcEe07A0fa60E864dDCBC4eF68e",
									  "max_fast_withdrawal_for_platform_per_day": "10000",
									  "max_fast_withdrawal_for_user_per_day": "4000"
									},
									"matic": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010",
									  "max_fast_withdrawal_for_platform_per_day": "10000",
									  "max_fast_withdrawal_for_user_per_day": "4000"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x4d2548DAbF3d662110d70239Bc3531043984644D",
									  "max_fast_withdrawal_for_platform_per_day": "10000",
									  "max_fast_withdrawal_for_user_per_day": "4000"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"btc",
									"matic",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"btc",
									"matic",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"btc",
									"matic",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"btc",
									"matic",
									"usdt"
								  ]
								},
								"STARKNET": {
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x005a643907b9a4bc6a55e9069c4fd5fd1f5c79a22470690f75556c4736e34426"
									}
								  },
								  "allowed_tokens_for_deposit": [],
								  "allowed_tokens_for_deposit_frontend": [],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								},
								"SCROLL": {
								  "deposit_contract": "0x1C810252902F8e34c9aC724E558Ea99baa2a9a9B",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0xFFFF3e201311e450a0Bf8091124B4B22083514e3"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x69D0377CD7FbA97a662bEA6a7Ae194CDa8C5063A"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc",
									"usdt"
								  ]
								},
								"OPTIMISM": {
								  "deposit_contract": "0x4257742647821d8BbC086caC9c32f8F3328C709A",
								  "tokens": {
									"btc": {
									  "blockchain_decimal": "18",
									  "token_contract": "0xd9C7d45444Fcb9CE39f63478275eF2449380B7F1",
									  "max_fast_withdrawal_for_platform_per_day": "0.03",
									  "max_fast_withdrawal_for_user_per_day": "0.02"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x0b9C59563E7a3727eC363236D349aA526c882a3A"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"btc",
									"usdt"
								  ]
								},
								"ARBITRUM": {
								  "deposit_contract": "0x987405D0e208e801859eb4e9C9923f67F3C487f1",
								  "tokens": {
									"btc": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x3F180fCFd904D2d5047037e33B3f65F5eD312f33"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x7eE00F9D72422e7939f33537aB01e630A3dd95D3"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"btc",
									"usdt"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"btc",
									"usdt"
								  ]
								},
								"LINEA": {
								  "deposit_contract": "0x8Ec96d9bC4fEE06Fb668a0FBca66510a57C42b05",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x2318ED5b67781fa56EC3A677624deb3deE2CB504"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x0f40466610F20ed3C7cba31CBc5D7358C8824FCa"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010",
									  "max_fast_withdrawal_for_platform_per_day": "3",
									  "max_fast_withdrawal_for_user_per_day": "1.5"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ]
								},
								"MODE": {
								  "deposit_contract": "0x886dF268A7AD3dC9F7Ff1Bf197A5dF840DE397F1",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x79378E3E1BE2062861d23635A78A7aA8BE19ccc4"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010",
									  "max_fast_withdrawal_for_platform_per_day": "3",
									  "max_fast_withdrawal_for_user_per_day": "1.5"
									},
									"usdt": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x04655c0a6023889B3d8Da5950e6f41AeEc4772bA"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc",
									"usdt",
									"eth"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc",
									"usdt",
									"eth"
								  ]
								},
								"ZKPOLY": {
								  "deposit_contract": "",
								  "tokens": {
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x28bA3579945A67aD54D83Be0B0F9481C464d4Dc7"
									},
									"eth": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								},
								"BERA": {
								  "deposit_contract": "0xa8a86215102c362ad1E871959522D3830BEd3dAd",
								  "tokens": {
									"bera": {
									  "blockchain_decimal": "18",
									  "token_contract": "0x0000000000000000000000000000000000001010"
									},
									"usdc": {
									  "blockchain_decimal": "6",
									  "token_contract": "0x3736BF60A387585F7349A50BEb58aC0f7142b9c3"
									}
								  },
								  "allowed_tokens_for_deposit": [
									"usdc"
								  ],
								  "allowed_tokens_for_deposit_frontend": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd": [
									"usdc"
								  ],
								  "allowed_tokens_for_fast_wd_frontend": [
									"usdc"
								  ]
								}
							  }
							}
						  }`)),
					}, nil
				}

				if r.URL.Path == START_FAST_WITHDRAWAL_ENDPOINT {
					return &http.Response{
						StatusCode: http.StatusOK,
						Body: io.NopCloser(strings.NewReader(`{
							"status": "success",
							"message": "",
							"payload": {
							  "msg_hash": "0x2",
							  "fastwithdrawal_withdrawal_id": 1
							}
						}`)),
					}, nil
				}

				// if r.URL.Path == PROCESS_FAST_WITHDRAWAL_ENDPOINT {
				// 	return &http.Response{
				// 		StatusCode: http.StatusOK,
				// 		Body: io.NopCloser(strings.NewReader(`{
				// 			"status": "success",
				// 			"message": ""
				// 		}`)),
				// 	}, nil
				// }

				return &http.Response{
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`{
						"status": "success"
					}`)),
				}, nil
			}),
			amount:   0.001,
			currency: USDT,
			network:  POLYGON,
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

			client.jwtToken = "token"
			client.refreshToken = "refresh"
			client.httpClient.Transport = tc.roundTripper

			ctx := context.Background()
			if tc.timeout != 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			got, err := client.FastWithdrawal(ctx, "0x1", tc.amount, tc.currency, tc.network)
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
