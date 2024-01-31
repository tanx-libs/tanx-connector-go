package wsclient

import (
	"encoding/json"
	"fmt"
)

// filled with default values by default
type Wsclient struct {
	baseURL    string
	publicURL  string
	privateURL string
}

func New() *Wsclient {
	return &Wsclient{
		baseURL:    BASE_URL,
		publicURL:  BASE_URL + PUBLIC_ENDPOINT,
		privateURL: BASE_URL + PRIVATE_ENDPOINT,
	}
}

func (c *Wsclient) BaseURL(baseURL string) *Wsclient {
	c.baseURL = baseURL
	return c
}

func (c *Wsclient) PublicPath(publicPath string) *Wsclient {
	c.publicURL = c.baseURL + publicPath
	return c
}

func (c *Wsclient) PrivatePath(privatePath string) *Wsclient {
	c.privateURL = c.baseURL + privatePath
	return c
}

// TODO: handle unsubscribe (look for different ways)
func (c *Wsclient) Trade(symbol []string, tradeEventHandler TradeEventHandler, subUnsubEventHandler SubUnsubEventHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	topics := make([]string, len(symbol))
	for i, s := range symbol {
		topics[i] = fmt.Sprintf("%s.trades", s)
	}
	config := &Config{
		endpoint: c.publicURL,
		request: request{
			Event:   SUBSCRIBE,
			Streams: topics,
		},
	}

	eventHandler := func(message []byte) {
		var msg map[string]interface{}
		err := json.Unmarshal(message, &msg)
		if err != nil {
			errHandler(err)
			return
		}

		switch {
		case msg[SUB_UNSUB_SUCCESS] != nil || msg[SUB_UNSUB_ERROR] != nil:
			var subUnsubEvent *SubUnsubEvent
			err = json.Unmarshal(message, &subUnsubEvent)
			if err != nil {
				errHandler(err)
				return
			}
			subUnsubEventHandler(subUnsubEvent)

		default:
			var tradeEvent *TradeEvent
			err = json.Unmarshal(message, &tradeEvent)
			if err != nil {
				errHandler(err)
				return
			}
			tradeEventHandler(tradeEvent)
		}
	}

	return serve(config, eventHandler, errHandler)
}


// TODO: Rest of them will follow the same mechanic 

// func (c *Wsclient) Kline(symbol string, interval KlinePeriod, klineEventHandler KlineEventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
// 	topic := fmt.Sprintf("%s.kline-%s", strings.ToLower(symbol), interval)
// 	config := &Config{
// 		endpoint: c.publicURL,
// 		request: request{
// 			Event:   "subscribe",
// 			Streams: []string{topic},
// 		},
// 	}
// 	log.Printf("%+v", config)

// 	// here data parsing is done
// 	eventHandler := func(message []byte) {
// 		var event *KlineEvent
// 		var success *successResponse

// 		err := json.Unmarshal(message, &event)
// 		if err == nil {
// 			klineEventHandler(event)
// 		} else {
// 			err = json.Unmarshal(message, &success)
// 			if err != nil {
// 				errHandler(err)
// 			} else {
// 				log.Printf("%+v", success)
// 			}
// 		}
// 	}
// 	return serve(config, eventHandler, errHandler)
// }
