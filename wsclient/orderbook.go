package wsclient

import (
	"encoding/json"
	"fmt"
)

type OrderbookData struct {
	Asks     []string `json:"asks"`
	Bids     []string `json:"bids"`
	Sequence int64    `json:"sequence"`
}

type OrderbookEvent map[string]OrderbookData

type OrderbookEventHandler func(event *OrderbookEvent)

// SubUnsubOrderbookTopics returns a SubUnsubRequest to subscribe or unsubscribe to Orderbook data
func SubUnsubOrderbookTopics(subUnsub SubUnsub, symbol []string) SubUnsubRequest {
	topics := make([]string, len(symbol))
	for i, s := range symbol {
		topics[i] = fmt.Sprintf("%s.ob-inc", s)
	}

	return SubUnsubRequest{
		Event:   subUnsub,
		Streams: topics,
	}
}

// Orderbook subscribes to Orderbook data
func (c *Wsclient) Orderbook(symbol []string, orderbookEventHandler OrderbookEventHandler, subUnsubEventHandler SubUnsubEventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, subUnsubCh chan SubUnsubRequest, err error) {
	config := &Config{
		endpoint:        c.publicURL,
		subUnsubRequest: SubUnsubOrderbookTopics(SUBSCRIBE, symbol),
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
				Msg: msg[SUB_UNSUB_ERROR].(string),
				Err: fmt.Errorf(""),
			})

		default:
			var orderbookEvent *OrderbookEvent
			err = json.Unmarshal(message, &orderbookEvent)
			if err != nil {
				errHandler(&ErrJSONUnmarshal{
					Err: err,
				})
				return
			}
			orderbookEventHandler(orderbookEvent)
		}
	}

	return serve(config, eventHandler, errHandler)
}
