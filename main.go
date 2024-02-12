package main

import (
	"log"

	"github.com/tanx-libs/tanx-connector-go/wsclient"
)

func main() {
	var stopCh chan struct{}
	var subUnsubCh chan wsclient.SubUnsubRequest

	ws := wsclient.New()

	wsSubUnsubEventHandler := func(event *wsclient.SubUnsubEvent) {
		log.Printf("SubUnsub event: %+v\n", event)
	}

	wsTradeEventHandler := func(event *wsclient.TradeEvent) {
		log.Printf("Trade event: %+v\n", event)

		// example
		subUnsubCh <- wsclient.SubUnsubTradeTopics(wsclient.SUBSCRIBE,[]string{"btcusdc", "ethusdc"})
	}

	wsErrHandler := func(err error) {
		switch err {
		case err.(*wsclient.ErrSubUnsub):
			log.Println("SubUnsub error: ", err)
			stopCh <- struct{}{}

		default:
			log.Println("Default: ", err)
		}
	}

	doneCh, stopCh, subUnsubCh, err := ws.Trade([]string{"btcusdc", "ethusdc"}, wsTradeEventHandler, wsSubUnsubEventHandler, wsErrHandler)
	if err != nil {
		log.Fatal(err)
	}

	select {
	case <-stopCh:
		log.Println("stopped was called")

	case <-doneCh:
		log.Println("done was called")
	}
}

