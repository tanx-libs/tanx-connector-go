package wsclient

import (
	"encoding/json"
	"fmt"
)

// helper function to form subscribe or unsubscribe request for private topics
/*
{ "event": "subscribe", "streams": ["trade", "order"] }
*/
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

/*
responses for trade

	{
	  "trade": {
	    "amount": "0.001",
	    "created_at": 1659024869,
	    "id": 448640,
	    "market": "btcusdc",
	    "order_id": 8840859,
	    "price": "23829.22",
	    "side": "buy",
	    "taker_type": "buy",
	    "total": "23.82922"
	  }
	}

responses for order
{
  "order": {
    "at": 1659024868,
    "avg_price": "23829.22",
    "created_at": 1659024868,
    "executed_volume": "0.001",
    "id": 8840859,
    "kind": "bid",
    "market": "btcusdc",
    "ord_type": "limit",
    "origin_volume": "0.001",
    "price": "24000.0",
    "remaining_volume": "0.0",
    "side": "buy",
    "state": "done",
    "trades_count": 1,
    "updated_at": 1659024869
  }
}
*/
type PrivateEvent struct {
	Trade PrivateTrade `json:"trade"`
	Order PrivateOrder `json:"order"`
}

type PrivateEventHandler func(event *PrivateEvent)

// subscribes to Trade data for authenticated user
func (c *Wsclient) PrivateTrade(symbol []string, privateEventHandler PrivateEventHandler, subUnsubEventHandler SubUnsubEventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, subUnsubCh chan SubUnsubRequest, err error) {
	if c.accessToken == "" {
		return nil, nil, nil, ErrNotLoggedIn
	}

	url := fmt.Sprintf("%s?auth_header=%s", c.privateURL, c.accessToken)
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
