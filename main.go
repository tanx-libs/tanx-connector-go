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
	// // starkPrivateKey := "0x7302fa58776da9f8fcf3631f4cb495a4dd0cdfab785e8b72a8a637d4bb14784"

	rpcURL := "https://sepolia.infura.io/v3/bc9fafffa1f447bab403ee4a8b5090f4"
	_, err = c.Login(context.TODO(), ethAddr, ethPrivateKey)
	if err != nil {
		panic(err)
	}

	// err = c.SetEthereumAllowance(context.TODO(), rpcURL, ethPrivateKey, client.USDC, 1)
	// if err != nil {
	// 	panic(err)
	// }

	// // ethereum desposits
	hash, err := c.DepositFromEthereumNetwork(context.TODO(), rpcURL, ethAddr, ethPrivateKey, starkPublicKey, 0.1, client.USDC)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", hash)

	// // polygon deposits
	// rpcURL2 := "https://polygon-mumbai-bor-rpc.publicnode.com"


	// hash, err := c.DepositFromPolygonNetwork(context.TODO(), rpcURL2, ethAddr, ethPrivateKey, starkPublicKey, client.MATIC, 0.01)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", hash)

	// rpcURL2 := "https://polygon-mumbai-bor-rpc.publicnode.com"
	// hash, err := c.DepositFromPolygonNetwork(context.TODO(), rpcURL2, ethAddr, ethPrivateKey, starkPublicKey, client.BTC, 0.001)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", hash)

	// depos, err := c.ListAllDeposits(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", depos)
}
