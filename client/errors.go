package client

import "fmt"

var ErrMarketNotProvided = fmt.Errorf("provide a market")

/*
{"status":"error","message":"Buying ETH is disabled in Testnet temporarily","payload":""}
*/
