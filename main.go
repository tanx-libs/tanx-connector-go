package main

import (
	"context"

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
	// starkPrivateKey := "0x7302fa58776da9f8fcf3631f4cb495a4dd0cdfab785e8b72a8a637d4bb14784"
	// starkPublicKey := "0x64211ed550cb37140ef2268cf7b2625aef725d33618c9651765e16318101c17"
	// rpcURL := "https://sepolia.infura.io/v3/bc9fafffa1f447bab403ee4a8b5090f4"

	_, err = c.Login(context.TODO(), ethAddr, ethPrivateKey)
	if err != nil {
		panic(err)
	}

	// resp, err := c.InitiateNormalWithdrawal(context.TODO(), starkPrivateKey, 0.0001, client.ETH)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", resp)

	// resp, err := c.GetPendingNormalWithdrawalAmountByCoin(context.TODO(), rpcURL, ethAddr, client.USDC)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", resp)

	// resp, err := c.FastWithdrawal(context.TODO(), starkPrivateKey, 11, client.USDC, client.ETHEREUM)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v", resp)
}
