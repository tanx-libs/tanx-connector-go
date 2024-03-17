package client

import "fmt"

var (
	ErrMarketNotProvided = fmt.Errorf("provide a market")
	ErrInvalidAmount     = fmt.Errorf("amount should be greater than 0")
	ErrNotLoggedIn       = fmt.Errorf("not logged in")
	ErrCoinNotFound      = fmt.Errorf("coin not found")
	ErrInvalidNetwork	= fmt.Errorf("invalid network")
	ErrInsufficientBalance = fmt.Errorf("insufficient balance")
	ErrInsufficientAllowance = fmt.Errorf("insufficient allowance")
	ErrAllowanceTooLow = fmt.Errorf("allowance too low")
)


/*
{"status":"error","message":"Buying ETH is disabled in Testnet temporarily","payload":""}
*/
