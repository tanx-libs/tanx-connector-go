package main

import (
	"context"
	"log"

	"github.com/tanx-libs/tanx-connector-go/wsclient"
)

func main() {
	wsclient, err := wsclient.New(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	wsclient.Subscribe(context.TODO(), []string{"btcusdc.trades", "btcusdc.ob-inc", "btcusdc.kline-5m"})

	select{}
}
