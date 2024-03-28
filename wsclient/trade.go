// TODO add msg for extra debugging

package wsclient

import (
	"encoding/json"
	"fmt"
)

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

// SubUnsubTradeTopics returns a SubUnsubRequest to subscribe or unsubscribe to Trade data
func SubUnsubTradeTopics(subUnsub SubUnsub, symbol []string) SubUnsubRequest {
	streams := make([]string, len(symbol))
	for i, s := range symbol {
		streams[i] = fmt.Sprintf("%s.trades", s)
	}

	return SubUnsubRequest{
		Event:   subUnsub,
		Streams: streams,
	}
}

// Trade subscribes to Trade data
func (c *Wsclient) Trade(symbol []string, tradeEventHandler TradeEventHandler, subUnsubEventHandler SubUnsubEventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, subUnsubCh chan SubUnsubRequest, err error) {
	config := &Config{
		endpoint:        c.publicURL,
		subUnsubRequest: SubUnsubTradeTopics(SUBSCRIBE, symbol),
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
			var tradeEvent *TradeEvent
			err = json.Unmarshal(message, &tradeEvent)
			if err != nil {
				errHandler(&ErrJSONUnmarshal{
					Err: err,
				})
				return
			}
			tradeEventHandler(tradeEvent)
		}
	}

	return serve(config, eventHandler, errHandler)
}
