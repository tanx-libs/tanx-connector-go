package types

const (
	// json api
	JSON_BASE_URL = "https://api.tanx.fi/"
	HEALTH_PATH   = "sapi/v1/health/"
	TICKER_PATH   = "sapi/v1/market/tickers/"

	// websocket
	WS_BASE_URL = "wss://api.tanx.fi/public"
)

// ------------------------- JSON HTTP -------------------------------------
// health
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Payload string `json:"payload"`
}

// 24 hour ticker
type Ticker struct {
	At                 string `json:"at"`
	AvgPrice           string `json:"avg_price"`
	High               string `json:"high"`
	Last               string `json:"last"`
	Low                string `json:"low"`
	Open               string `json:"open"`
	PriceChangePercent string `json:"price_change_percent"`
	Volume             string `json:"volume"`
	Amount             string `json:"amount"`
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
type CandleStickResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Payload [][]int `json:"payload"`
}

type CandleStickOptions struct {
	Limit     string
	Period    int
	StartTime int64
	EndTime   int64
}

// orderBook
type Ask struct {
	ID              int     `json:"id"`
	UUID            string  `json:"uuid"`
	Side            string  `json:"side"`
	OrdType         string  `json:"ord_type"`
	Price           float64 `json:"price"`
	AvgPrice        string  `json:"avg_price"`
	State           string  `json:"state"`
	Market          string  `json:"market"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	OriginVolume    string  `json:"origin_volume"`
	RemainingVolume string  `json:"remaining_volume"`
	ExecutedVolume  string  `json:"executed_volume"`
	MakerFee        string  `json:"maker_fee"`
	TakerFee        string  `json:"taker_fee"`
	TradesCount     int     `json:"trades_count"`
}

type Bid struct {
	ID              int     `json:"id"`
	UUID            string  `json:"uuid"`
	Side            string  `json:"side"`
	OrdType         string  `json:"ord_type"`
	Price           float64 `json:"price"`
	AvgPrice        string  `json:"avg_price"`
	State           string  `json:"state"`
	Market          string  `json:"market"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	OriginVolume    string  `json:"origin_volume"`
	RemainingVolume string  `json:"remaining_volume"`
	ExecutedVolume  string  `json:"executed_volume"`
	MakerFee        string  `json:"maker_fee"`
	TakerFee        string  `json:"taker_fee"`
	TradesCount     int     `json:"trades_count"`
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

// recent trades
type RecentTradePayload struct {
	ID        int     `json:"id"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Total     float64 `json:"total"`
	Market    string  `json:"market"`
	CreatedAt int64   `json:"created_at"`
	TakerType string  `json:"taker_type"`
}

type RecentTradeResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Payload []RecentTradePayload `json:"payload"`
}

// ---------------------- websocket ------------------------------------

/*
{
  "btcusdc.trades": {
    "trades": [
      {
        "amount": "0.001",
        "date": 1659023778,
        "price": "23935.01",
        "taker_type": "buy",
        "tid": 448639
      }
    ]
  }
}
*/

type SubUnsubRequest struct {
	Event   string   `json:"event"`
	Streams []string `json:"streams"`
}

type Trade struct {
	Amount    string `json:"amount"`
	Date      int64  `json:"date"`
	Price     string `json:"price"`
	TakerType string `json:"taker_type"`
	Tid       int64  `json:"tid"`
}

type TradeStreamResponse struct {
	Trades []Trade `json:"trades"`
}
