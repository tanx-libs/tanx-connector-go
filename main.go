package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tanx-libs/tanx-connector-go/client"
	"github.com/tanx-libs/tanx-connector-go/wsclient"
)

type Entry struct {
	MsgHash string `json:"msg_hash"`
	R       string `json:"r"`
	S       string `json:"s"`
}

func main() {
	jc, err := client.New(client.TESTNET)
	if err != nil {
		panic(err)
	}
	ethAddr := "0xF58001619C165cDd20B5F7A0EDa072Fd13943002"
	ethPrivateKey := "e65c38b42af2e20540fde19d10bec7fb752ab58852e466151747abc08ae2494a"
	_, jwt, err := jc.Login(context.TODO(), ethAddr, ethPrivateKey)
	if err != nil {
		panic(err)
	}

	c := wsclient.New(wsclient.TESTNET)
	c.SetJwt(jwt.Token.Access, jwt.Token.Refresh)

	var ch chan wsclient.SubUnsubRequest

	tradeEventHandler := func(event *wsclient.PrivateEvent) {
		fmt.Printf("%+v\n", event)
	}

	subUnsubEventHandler := func(subUnsub *wsclient.SubUnsubEvent) {
		fmt.Printf("%+v\n", subUnsub)
	}

	errHandler := func(err error) {
		switch e := err.(type) {
		case *wsclient.ErrJSONUnmarshal:
			log.Printf("%+v\n", e)
		case *wsclient.ErrWsReadMessage:
			log.Printf("%+v\n", e)
		case *wsclient.ErrSubUnsub:
			ch <- wsclient.SubUnsubPrivateTopics(wsclient.SUBSCRIBE, []string{"order"})
			log.Printf("%+v\n", e)
		default:
			log.Printf("nahi chala %+v\n", e)
		}
	}

	d, _, ch, err := c.PrivateTrade([]string{"order", "trade"}, tradeEventHandler, subUnsubEventHandler, errHandler)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	// yeh orders wala bakchodi hain. will see it afterwards
	// c, err := client.New(client.TESTNET)
	// if err != nil {
	// 	panic(err)
	// }

	// // first login
	// ethAddr := "0xF58001619C165cDd20B5F7A0EDa072Fd13943002"
	// ethPrivateKey := "e65c38b42af2e20540fde19d10bec7fb752ab58852e466151747abc08ae2494a"
	// _, _, err = c.Login(context.TODO(), ethAddr, ethPrivateKey)
	// if err != nil {
	// 	panic(err)
	// }

	// add correct valueshere to test the order
	starkPrivateKey := "588ad6325783739b3806e27feebeb1c270d4c6875c29517ff1d689e777b13a6"
	order, err := jc.OrderCreate(context.TODO(), starkPrivateKey, client.OrderOptions{
		Market:  "btcusdc",
		OrdType: "market",
		Side:    "buy",
		Volume:  0.0001,
	})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(order)
	<-d
}
