package client

import (
	"net/http"
	"net/url"
)

// Specify mainet or testnet
type Base string

const (
	MAINET  Base = "https://api.tanx.fi"
	TESTNET Base = "https://api-testnet.tanx.fi"

	HEALTH_PATH      = "sapi/v1/health/"
	TICKER_PATH      = "sapi/v1/market/tickers/"
	CANDLESTICK_PATH = "sapi/v1/market/kline/"
	ORDERBOOK_PATH   = "sapi/v1/market/orderbook/"
	TRADES_PATH      = "sapi/v1/market/trades/"

	NONCE_PATH         = "sapi/v2/auth/nonce/"
	LOGIN_PATH         = "sapi/v2/auth/login/"
	REFRESH_TOKEN_PATH = "sapi/v2/auth/token/refresh/"
	PROFILE_PATH       = "sapi/v1/user/profile/"
	BALANCE_PATH       = "sapi/v1/user/balance/"
	PNL_PATH           = "sapi/v1/user/pnl/"

	ORDER_NONCE_PATH  = "sapi/v1/orders/nonce/"
	ORDER_CREATE_PATH = "sapi/v1/orders/create/"
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

	orderNonceURL  *url.URL
	orderCreateURL *url.URL
}

func New(base Base) (*Client, error) {
	baseurl, err := url.Parse(string(base))
	if err != nil {
		return nil, err
	}

	healthurl := baseurl.JoinPath(HEALTH_PATH)
	tickerurl := baseurl.JoinPath(TICKER_PATH)
	candleStickurl := baseurl.JoinPath(CANDLESTICK_PATH)
	orderBookurl := baseurl.JoinPath(ORDERBOOK_PATH)
	tradesurl := baseurl.JoinPath(TRADES_PATH)

	nonceurl := baseurl.JoinPath(NONCE_PATH)
	loginurl := baseurl.JoinPath(LOGIN_PATH)
	refreshTokenurl := baseurl.JoinPath(REFRESH_TOKEN_PATH)

	profileurl := baseurl.JoinPath(PROFILE_PATH)
	balanceurl := baseurl.JoinPath(BALANCE_PATH)
	pnlurl := baseurl.JoinPath(PNL_PATH)

	ordernonceurl := baseurl.JoinPath(ORDER_NONCE_PATH)
	ordercreateurl := baseurl.JoinPath(ORDER_CREATE_PATH)

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

		orderNonceURL:  ordernonceurl,
		orderCreateURL: ordercreateurl,
	}, nil
}


func (c *Client) SetJWTToken(jwtToken string) {
	c.jwtToken = jwtToken
}

func (c *Client) SetRefreshToken(refreshToken string) {
	c.refreshToken = refreshToken
}
