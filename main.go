package main

import (
	"context"
	"log"

	"github.com/tanx-libs/tanx-connector-go/client"
)

func main() {
	c, err := client.New(client.TESTNET)
	if err != nil {
		panic(err)
	}

	e := "0xf318C11ff6E60115FB3e107bEa2637c060BEbc8C"
	p := "ba169c79340371a9aa4fd516462f939242f92b522081d945c001b0fb3dc3a66f"

	_, err = c.Login(context.TODO(), e, p)
	if err != nil {
		panic(err)
	}

	bal, err := c.Balance(context.TODO(), client.ETH)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v",bal)

	// _, _, err = c.Login(context.TODO(), e, p)
	// if err != nil {
	// 	panic(err)
	// }

	// pnl, err := c.PNL(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(pnl)
}
