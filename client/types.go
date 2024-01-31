package client

import "fmt"

const (
	// json api
	JSON_BASE_URL    = "https://api.tanx.fi/"
	HEALTH_PATH      = "sapi/v1/health/"
	TICKER_PATH      = "sapi/v1/market/tickers/"
	CANDLESTICK_PATH = "sapi/v1/market/kline/"
	ORDERBOOK_PATH   = "sapi/v1/market/depth/"
	TRADES_PATH      = "sapi/v1/market/trades/"
)

var ErrMarketNotProvided = fmt.Errorf("provide a market")

// health
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload string `json:"payload"`
}

// 24 hour ticker
type Ticker struct {
	Low                string `json:"low"`
	High               string `json:"high"`
	Open               string `json:"open"`
	Last               string `json:"last"`
	Volume             string `json:"volume"`
	Amount             string `json:"amount"`
	Vol                string `json:"vol"`
	AvgPrice           string `json:"avg_price"`
	PriceChangePercent string `json:"price_change_percent"`
	At                 string `json:"at"`
}

type TickerPayload struct {
	At     string `json:"at"`
	Ticker Ticker `json:"ticker"`
}

type TickerResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Payload TickerPayload `json:"payload"`
}

type AllTickerResponse struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Payload map[string]TickerPayload `json:"payload"`
}

// candleStick
const (
	CANDLESTICK_DEFAULT_LIMIT  = 30
	CANDLESTICK_DEFAULT_PERIOD = PERIOD_1MIN
)

type Period int

const (
	PERIOD_1MIN  Period = 1
	PERIOD_5MIN  Period = 5
	PERIOD_15MIN Period = 15
	PERIOD_30MIN Period = 30
	PERIOD_1H    Period = 60
	PERIOD_2H    Period = 120
	PERIOD_4H    Period = 240
	PERIOD_6H    Period = 360
	PERIOD_12H   Period = 720
	PERIOD_1D    Period = 1440
	PERIOD_3D    Period = 4320
	PERIOD_1W    Period = 10080
)

type CandleStickOptions struct {
	Limit     int
	Period    Period
	StartTime int64
	EndTime   int64
}

type CandleStickResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Payload [][]float64 `json:"payload"`
}

// orderBook
const (
	ORDERBOOK_DEFAULT_ASK_LIMIT = 20
	ORDERBOOK_DEFAULT_BID_LIMIT = 20
)

type OrderBookOptions struct {
	AskLimit int
	BidLimit int
}

type Ask struct {
	ID              int    `json:"id"`
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	OriginVolume    string `json:"origin_volume"`
	RemainingVolume string `json:"remaining_volume"`
	ExecutedVolume  string `json:"executed_volume"`
	MakerFee        string `json:"maker_fee"`
	TakerFee        string `json:"taker_fee"`
	TradesCount     int    `json:"trades_count"`
}

type Bid struct {
	ID              int    `json:"id"`
	UUID            string `json:"uuid"`
	Side            string `json:"side"`
	OrdType         string `json:"ord_type"`
	Price           string `json:"price"`
	AvgPrice        string `json:"avg_price"`
	State           string `json:"state"`
	Market          string `json:"market"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	OriginVolume    string `json:"origin_volume"`
	RemainingVolume string `json:"remaining_volume"`
	ExecutedVolume  string `json:"executed_volume"`
	MakerFee        string `json:"maker_fee"`
	TakerFee        string `json:"taker_fee"`
	TradesCount     int    `json:"trades_count"`
}

type OrderBookPayload struct {
	Asks []Ask `json:"asks"`
	Bids []Bid `json:"bids"`
}

type OrderBookResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Payload OrderBookPayload `json:"payload"`
}

// trades
const (
	TRADES_DEFAULT_LIMIT    = 100
	TRADES_DEFAULT_ORDER_BY = DESC
)

type OrderBy string

const (
	ASC  OrderBy = "asc"
	DESC OrderBy = "desc"
)

type TradesOptions struct {
	Limit     int
	Timestamp int64
	OrderBy   OrderBy
}

type TradesPayload struct {
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Total     float64 `json:"total"`
	Market    string  `json:"market"`
	CreatedAt int64   `json:"created_at"`
	TakerType string  `json:"taker_type"`
}

type TradesResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Payload []TradesPayload `json:"payload"`
}
