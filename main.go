package main

import (
	"context"
	"fmt"

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
	// starkPrivateKey := "0x7302fa58776da9f8fcf3631f4cb495a4dd0cdfab785e8b72a8a637d4bb14784"

	// rpcURL := "https://sepolia.infura.io/v3/bc9fafffa1f447bab403ee4a8b5090f4"
	_, _, err = c.Login(context.TODO(), ethAddr, ethPrivateKey)
	if err != nil {
		panic(err)
	}

	// resp, err := c.FastWithdrawal(context.TODO(), starkPrivateKey, 10, client.USDC, client.ETHEREUM)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(resp)

	// resp, err := c.ListFastWithdrawal(context.TODO(), client.ListWithdrawalRequest{
	// 	Page: 1,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// log.Printf("%+v", resp)

	// err = c.InitiateNormalWithdrawal(context.TODO(), starkPrivateKey, 1, "usdc")
	// if err != nil {
	// 	panic(err)
	// }

	// ethClient, err := ethclient.Dial(rpcURL)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := c.GetPendingNormalWithdrawalAmountByCoin(context.TODO(), "usdc", ethAddr, ethClient, client.TESTNET)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res)

	// res2, err := c.CompleteNormalWithdrawal(context.TODO(), "usdc", ethPrivateKey, ethAddr, ethClient, client.TESTNET)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(res2)

	// hash, err := c.DepositFromEthereumNetwork(context.TODO(), rpcURL, ethAddr, ethPrivateKey, starkPublicKey, client.TESTNET, client.ETH, 0.001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", hash)

	rpcURL2 := "https://polygon-mumbai-bor-rpc.publicnode.com"
	hash, err := c.DepositFromPolygonNetwork(context.TODO(), rpcURL2, ethAddr, ethPrivateKey, starkPublicKey, client.BTC, 0.001)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", hash)

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
