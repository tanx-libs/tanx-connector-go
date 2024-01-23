package types

const (
	BASE_URL    = "https://api.tanx.fi/"
	HEALTH_PATH = "sapi/v1/health/"
	TICKER_PATH = "sapi/v1/market/tickers/"
)

type HealthResponse struct {
	Status string 
	Message string 
	Payload string
}

type Ticker struct {
	At string
	AvgPrice string `json:"avg_price"`
	High string
	Last string
	Low string
	Open string
	PriceChangePercent string `json:"price_change_percent"`
	Volume string
	Amount string
}

type Data struct {
	At string
	Ticker Ticker
}

type Payload struct {
	Symbol Data
}

type TickerResponse struct {
	Status string
	Message string
	Payload Payload
}

