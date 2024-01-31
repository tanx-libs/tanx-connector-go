package wsclient

const (
	BASE_URL          = "wss://api.tanx.fi"
	PUBLIC_ENDPOINT   = "/public"
	PRIVATE_ENDPOINT  = "/private"
	SUB_UNSUB_SUCCESS = "success"
	SUB_UNSUB_ERROR   = "error"
	SUBSCRIBE         = "subscribe"
	UNSUBSCRIBE       = "unsubscribe"
)

type Success struct {
	Message string   `json:"message"`
	Streams []string `json:"streams"`
}

type SubUnsubEvent struct {
	Success Success `json:"success"`
	Error   string  `json:"error"`
}

type SubUnsubEventHandler func(event *SubUnsubEvent)

type Trade struct {
	Amount    string `json:"amount"`
	Date      int64  `json:"date"`
	Price     string `json:"price"`
	TakerType string `json:"taker_type"`
	Tid       int64  `json:"tid"`
}

type Trades struct {
	Trades []Trade `json:"trades"`
}

type TradeEvent map[string]Trades

type TradeEventHandler func(event *TradeEvent)

type KlineEvent map[string][]float64

type KlineEventHandler func(event *KlineEvent)

type KlinePeriod string

const (
	PERIOD_1M  KlinePeriod = "1m"
	PERIOD_5M  KlinePeriod = "5m"
	PERIOD_15M KlinePeriod = "15m"
	PERIOD_30M KlinePeriod = "30m"
	PERIOD_1H  KlinePeriod = "1h"
	PERIOD_2H  KlinePeriod = "2h"
	PERIOD_4H  KlinePeriod = "4h"
	PERIOD_6H  KlinePeriod = "6h"
	PERIOD_12H KlinePeriod = "12h"
	PERIOD_1D  KlinePeriod = "1d"
	PERIOD_3D  KlinePeriod = "3d"
	PERIOD_1W  KlinePeriod = "1w"
)
