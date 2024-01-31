package main

import (
	"log"

	"github.com/tanx-libs/tanx-connector-go/wsclient"
)

func main() {
	ws := wsclient.New()

	wsSubUnsubEventHandler := func(event *wsclient.SubUnsubEvent) {
		log.Printf("SubUnsub event: %+v\n", event)
	}

	wsTradeEventHandler := func(event *wsclient.TradeEvent) {
		log.Printf("Trade event: %+v\n", event)
	}

	wsErrHandler := func(err error) {
		log.Printf("Error: %+v\n", err)
	}

	doneCh, stopCh, err := ws.Trade([]string{"btcusdc", "ethusdc"}, wsTradeEventHandler, wsSubUnsubEventHandler, wsErrHandler)
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
