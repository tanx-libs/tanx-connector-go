package client

import (
	"net/http"
	"net/url"
)

// Specify mainet or testnet
const (
	MAINNET  = "https://api.tanx.fi"
	TESTNET = "https://api-testnet.tanx.fi"

	HEALTH_ENDPOINT      = "sapi/v1/health/"
	TICKER_ENDPOINT      = "sapi/v1/market/tickers/"
	CANDLESTICK_ENDPOINT = "sapi/v1/market/kline/"
	ORDERBOOK_ENDPOINT   = "sapi/v1/market/orderbook/"
	TRADES_ENDPOINT      = "sapi/v1/market/trades/"

	NONCE_ENDPOINT         = "sapi/v2/auth/nonce/"
	LOGIN_ENDPOINT         = "sapi/v2/auth/login/"
	REFRESH_TOKEN_ENDPOINT = "sapi/v2/auth/token/refresh/"
	PROFILE_ENDPOINT       = "sapi/v1/user/profile/"
	BALANCE_ENDPOINT       = "sapi/v1/user/balance/"
	PNL_ENDPOINT           = "sapi/v1/user/pnl/"

	ORDER_BASE_ENDPOINT   = "sapi/v1/orders/"
	ORDER_NONCE_ENDPOINT  = "sapi/v1/orders/nonce/"
	ORDER_CREATE_ENDPOINT = "sapi/v1/orders/create/"
	ORDER_CANCEL_ENDPOINT = "sapi/v1/orders/cancel/"

	COIN_ENDPOINT           = "main/stat/v2/coins/"
	VAULTID_ENDPOINT        = "main/user/create_vault/"
	NETWORK_CONFIG_ENDPOINT = "main/stat/v2/app-and-markets/"
	CRYPTO_DEPOSIT_START_ENDPOINT = "sapi/v1/payment/stark/start/"


	MAINET_STARK_CONTRACT   = "0x1390f521A79BaBE99b69B37154D63D431da27A07"
	TESTNET_STARK_CONTRACT  = "0xA2eC709125Ea693f5522aEfBBC3cb22fb9146B52"
)

/*
Maybe we can think of configuring client using a builder pattern. But that can be done later
*/
type Client struct {
	httpClient   *http.Client
	jwtToken     string
	refreshToken string

	// public
	baseURL        *url.URL
	healthURL      *url.URL
	tickerURL      *url.URL
	candleStickURL *url.URL
	orderBookURL   *url.URL
	tradesURL      *url.URL

	// login
	nonceURL        *url.URL
	loginURL        *url.URL
	refreshTokenURL *url.URL

	// protected
	profileURL *url.URL
	balanceURL *url.URL
	pnlURL     *url.URL

	orderURL       *url.URL
	orderNonceURL  *url.URL
	orderCreateURL *url.URL
	orderCancelURL *url.URL

	coinURL          *url.URL
	vaultIDURL       *url.URL
	networkConfigURL *url.URL
	cryptoDepositStartURL *url.URL
}

func New(base string) (*Client, error) {
	baseurl, err := url.Parse(string(base))
	if err != nil {
		return nil, err
	}

	healthurl := baseurl.JoinPath(HEALTH_ENDPOINT)
	tickerurl := baseurl.JoinPath(TICKER_ENDPOINT)
	candleStickurl := baseurl.JoinPath(CANDLESTICK_ENDPOINT)
	orderBookurl := baseurl.JoinPath(ORDERBOOK_ENDPOINT)
	tradesurl := baseurl.JoinPath(TRADES_ENDPOINT)

	nonceurl := baseurl.JoinPath(NONCE_ENDPOINT)
	loginurl := baseurl.JoinPath(LOGIN_ENDPOINT)
	refreshTokenurl := baseurl.JoinPath(REFRESH_TOKEN_ENDPOINT)

	profileurl := baseurl.JoinPath(PROFILE_ENDPOINT)
	balanceurl := baseurl.JoinPath(BALANCE_ENDPOINT)
	pnlurl := baseurl.JoinPath(PNL_ENDPOINT)

	orderURL := baseurl.JoinPath(ORDER_BASE_ENDPOINT)
	ordernonceurl := baseurl.JoinPath(ORDER_NONCE_ENDPOINT)
	ordercreateurl := baseurl.JoinPath(ORDER_CREATE_ENDPOINT)
	ordercancelurl := baseurl.JoinPath(ORDER_CANCEL_ENDPOINT)

	coinurl := baseurl.JoinPath(COIN_ENDPOINT)
	vaultidurl := baseurl.JoinPath(VAULTID_ENDPOINT)
	networkConfigurl := baseurl.JoinPath(NETWORK_CONFIG_ENDPOINT)
	cryptoDepositStarturl := baseurl.JoinPath(CRYPTO_DEPOSIT_START_ENDPOINT)

	return &Client{
		httpClient:   http.DefaultClient,
		jwtToken:     "",
		refreshToken: "",

		baseURL:        baseurl,
		healthURL:      healthurl,
		tickerURL:      tickerurl,
		candleStickURL: candleStickurl,
		orderBookURL:   orderBookurl,
		tradesURL:      tradesurl,

		nonceURL:        nonceurl,
		loginURL:        loginurl,
		refreshTokenURL: refreshTokenurl,

		profileURL: profileurl,
		balanceURL: balanceurl,
		pnlURL:     pnlurl,

		orderURL:       orderURL,
		orderNonceURL:  ordernonceurl,
		orderCreateURL: ordercreateurl,
		orderCancelURL: ordercancelurl,

		coinURL:          coinurl,
		vaultIDURL:       vaultidurl,
		networkConfigURL: networkConfigurl,
		cryptoDepositStartURL: cryptoDepositStarturl,
	}, nil
}

// check auth
func (c *Client) CheckAuth() error {
	if c.jwtToken == "" || c.refreshToken == "" {
		// throw error
		return ErrNotLoggedIn
	}

	return nil
}
