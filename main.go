package main

import (
	"context"
	"fmt"

	"github.com/tanx-libs/tanx-connector-go/client"
)

func main() {
	c, err := client.New(client.TESTNET)
	if err != nil {
		panic(err)
	}


	e := "0xF58001619C165cDd20B5F7A0EDa072Fd13943002"
	p := ""
	_, _, err = c.Login(context.TODO(), e, p)
	if err != nil {
		panic(err)
	}

	pnl, err := c.PNL(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println(pnl)
}

// func main() {
// 	var stopCh chan struct{}
// 	var subUnsubCh chan wsclient.SubUnsubRequest

// 	ws := wsclient.New()

// 	wsSubUnsubEventHandler := func(event *wsclient.SubUnsubEvent) {
// 		log.Printf("SubUnsub event: %+v\n", event)
// 	}

// 	wsTradeEventHandler := func(event *wsclient.TradeEvent) {
// 		log.Printf("Trade event: %+v\n", event)

// 		// example
// 		subUnsubCh <- wsclient.SubUnsubTradeTopics(wsclient.SUBSCRIBE,[]string{"btcusdc", "ethusdc"})
// 	}

// 	wsErrHandler := func(err error) {
// 		switch err {
// 		case err.(*wsclient.ErrSubUnsub):
// 			log.Println("SubUnsub error: ", err)
// 			stopCh <- struct{}{}

// 		default:
// 			log.Println("Default: ", err)
// 		}
// 	}

// 	doneCh, stopCh, subUnsubCh, err := ws.Trade([]string{"btcusdc", "ethusdc"}, wsTradeEventHandler, wsSubUnsubEventHandler, wsErrHandler)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	select {
// 	case <-stopCh:
// 		log.Println("stopped was called")

// 	case <-doneCh:
// 		log.Println("done was called")
// 	}
// }
