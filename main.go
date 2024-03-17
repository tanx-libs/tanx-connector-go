package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tanx-libs/tanx-connector-go/client"
)

type Entry struct {
	MsgHash string `json:"msg_hash"`
	R       string `json:"r"`
	S       string `json:"s"`
}

func main() {
	c, err := client.New(client.TESTNET)
	if err != nil {
		panic(err)
	}

	ethAddr := "0xf318C11ff6E60115FB3e107bEa2637c060BEbc8C"
	ethPrivateKey := "ba169c79340371a9aa4fd516462f939242f92b522081d945c001b0fb3dc3a66f"
	starkPublicKey := "0x64211ed550cb37140ef2268cf7b2625aef725d33618c9651765e16318101c17"

	rpcURL := "https://sepolia.infura.io/v3/bc9fafffa1f447bab403ee4a8b5090f4"
	// rpcURLPolygon := "https://polygon-testnet.public.blastapi.io"
	_, jwt, err := c.Login(context.TODO(), ethAddr, ethPrivateKey)
	if err != nil {
		panic(err)
	}
	log.Printf("jwt %+v", jwt)

	hash, err := c.DepositFromEthereumNetwork(context.TODO(), rpcURL, ethAddr, ethPrivateKey, starkPublicKey, client.TESTNET, client.USDC, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", hash)

	// opt := client.InternalTransferInitiateRequest{
	// 	OrganizationKey:    "+b^GSDqS#X79rE@_cNG!gk43HHz3V!-hhu+KFu*%b-ym6X#V^L?_jSCVtAy-yCec",
	// 	ApiKey:             "BRINE_SALT",
	// 	ClientReferenceId:  "123",
	// 	Currency:           "usdc",
	// 	Amount:             2,
	// 	DestinationAddress: ethAddr2,
	// }
	// r, err := c.InternalTransferCreate(context.TODO(), "0x2df627ebbf926c7c7f30dc69d201e8865e6cc01e2dc136cd7e26904a47e4abe", opt)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", hash)

	// opt := client.InternalTransferListRequest{
	// 	Limit:  2,
	// 	Offset: 0,
	// }
	// r1, err := c.InternalTransferList(context.TODO(), opt)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("%+v", r1)

	// depos, err := c.ListAllDeposits(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", depos)

	// // TODO chal nahi raha yeh patani kyu ?
	// jc, err := client.New(client.TESTNET)
	// if err != nil {
	// 	panic(err)
	// }
	// ethAddr := "0xF58001619C165cDd20B5F7A0EDa072Fd13943002"
	// ethPrivateKey := "e65c38b42af2e20540fde19d10bec7fb752ab58852e466151747abc08ae2494a"
	// _, jwt, err := jc.Login(context.TODO(), ethAddr, ethPrivateKey)
	// if err != nil {
	// 	panic(err)
	// }

	// c := wsclient.New(wsclient.TESTNET)
	// c.SetJwt(jwt.Token.Access, jwt.Token.Refresh)

	// var ch chan wsclient.SubUnsubRequest

	// tradeEventHandler := func(event *wsclient.PrivateEvent) {
	// 	fmt.Printf("%+v\n", event)
	// }

	// subUnsubEventHandler := func(subUnsub *wsclient.SubUnsubEvent) {
	// 	fmt.Printf("%+v\n", subUnsub)
	// }

	// errHandler := func(err error) {
	// 	switch e := err.(type) {
	// 	case *wsclient.ErrJSONUnmarshal:
	// 		log.Printf("%+v\n", e)
	// 	case *wsclient.ErrWsReadMessage:
	// 		log.Printf("%+v\n", e)
	// 	case *wsclient.ErrSubUnsub:
	// 		ch <- wsclient.SubUnsubPrivateTopics(wsclient.SUBSCRIBE, []string{"order"})
	// 		log.Printf("%+v\n", e)
	// 	default:
	// 		log.Printf("nahi chala %+v\n", e)
	// 	}
	// }

	// d, _, ch, err := c.PrivateTrade([]string{"order", "trade"}, tradeEventHandler, subUnsubEventHandler, errHandler)
	// if err != nil {
	// 	log.Fatalf("%+v\n", err)
	// }

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
	// starkPrivateKey := "588ad6325783739b3806e27feebeb1c270d4c6875c29517ff1d689e777b13a6"
	// order, err := jc.OrderCreate(context.TODO(), starkPrivateKey, client.OrderOptions{
	// 	Market:  "btcusdc",
	// 	OrdType: "market",
	// 	Side:    "buy",
	// 	Volume:  0.0001,
	// })
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(order)
	// <-d
}
