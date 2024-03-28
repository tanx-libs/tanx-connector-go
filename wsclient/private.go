package wsclient

import (
	"encoding/json"
	"fmt"
	"log"
)

func SubUnsubPrivateTopics(subUnsub SubUnsub, streams []string) SubUnsubRequest {
	return SubUnsubRequest{
		Event:   subUnsub,
		Streams: streams,
	}
}

type PrivateTrade struct {
	Trade
	Market  string `json:"market"`
	OrderID int64  `json:"order_id"`
	Total   string `json:"total"`
}

type PrivateOrder struct {
	At              int64  `json:"at"`
	AvgPrice        string `json:"avg_price"`
	CreatedAt       int64  `json:"created_at"`
	ExecutedVolume  string `json:"executed_volume"`
	ID              int64  `json:"id"`
	Kind            string `json:"kind"`
	Market          string `json:"market"`
	OrdType         string `json:"ord_type"`
	OriginVolume    string `json:"origin_volume"`
	Price           string `json:"price"`
	RemainingVolume string `json:"remaining_volume"`
	Side            string `json:"side"`
	State           string `json:"state"`
	TradesCount     int64  `json:"trades_count"`
	UpdatedAt       int64  `json:"updated_at"`
}

type PrivateEvent struct {
	Trade PrivateTrade `json:"trade"`
	Order PrivateOrder `json:"order"`
}

type PrivateEventHandler func(event *PrivateEvent)

// TradePrivate subscribes to Trade data for authenticated user
/*
1. if jwt not present return ErrWsNotLoggedIn
2. Subscribe payload is different
3. Handle case where jwt expires
*/
func (c *Wsclient) PrivateTrade(symbol []string, privateEventHandler PrivateEventHandler, subUnsubEventHandler SubUnsubEventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, subUnsubCh chan SubUnsubRequest, err error) {
	if c.accessToken == "" {
		return nil, nil, nil, ErrNotLoggedIn
	}

	url := fmt.Sprintf("%s?auth_header=%s", c.privateURL, c.accessToken)
	log.Println(url)
	config := &Config{
		endpoint:        url,
		subUnsubRequest: SubUnsubPrivateTopics(SUBSCRIBE, symbol),
	}

	eventHandler := func(message []byte) {
		var msg map[string]interface{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			errHandler(&ErrJSONUnmarshal{
				Err: err,
			})
			return
		}

		switch {
		case msg[SUB_UNSUB_SUCCESS] != nil:
			var subUnsubEvent *SubUnsubEvent
			err = json.Unmarshal(message, &subUnsubEvent)
			if err != nil {
				errHandler(&ErrJSONUnmarshal{
					Err: err,
				})
				return
			}
			subUnsubEventHandler(subUnsubEvent)

		case msg[SUB_UNSUB_ERROR] != nil:
			errHandler(&ErrSubUnsub{
				Err: fmt.Errorf(msg[SUB_UNSUB_ERROR].(string)),
			})

		default:
			var privateEvent *PrivateEvent
			err = json.Unmarshal(message, &privateEvent)
			if err != nil {
				errHandler(&ErrJSONUnmarshal{
					Err: err,
				})
				return
			}
			privateEventHandler(privateEvent)
		}
	}

	return serve(config, eventHandler, errHandler)
}
