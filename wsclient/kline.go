package wsclient

import (
	"encoding/json"
	"fmt"
)

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

// helper function to form subscribe or unsubscribe request for kline topics
/*
{
  "event": "subscribe",
  "streams": ["btcusdc.kline-5m"]
}
*/
func SubUnsubKlineTopics(subUnsub SubUnsub, symbol map[string]KlinePeriod) SubUnsubRequest {
	topics := make([]string, len(symbol))
	i := 0
	for s, p := range symbol {
		topics[i] = fmt.Sprintf("%s.kline-%s", s, p)
		i++
	}

	return SubUnsubRequest{
		Event:   subUnsub,
		Streams: topics,
	}
}

/*
{
  "btcusdc.kline-5m": [1659024300, 23935.01, 23935.01, 23935.01, 23935.01, 0]
}
*/
type KlineEvent map[string][]float64

type KlineEventHandler func(event *KlineEvent)

// connects you to kline stream
func (c *Wsclient) Kline(symbol map[string]KlinePeriod, klineEventHandler KlineEventHandler, subUnsubEventHandler SubUnsubEventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, subUnsubCh chan SubUnsubRequest, err error) {
	config := &Config{
		endpoint:        c.publicURL,
		subUnsubRequest: SubUnsubKlineTopics(SUBSCRIBE, symbol),
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
			var klineEvent *KlineEvent
			err = json.Unmarshal(message, &klineEvent)
			if err != nil {
				errHandler(&ErrJSONUnmarshal{
					Err: err,
				})
				return
			}
			klineEventHandler(klineEvent)
		}
	}

	return serve(config, eventHandler, errHandler)
}
