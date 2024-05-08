# Tickers

```go
res, err :=  c.Ticker(context.TODO(), "ethusdc")

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# CandleStick

```go
res, err :=  c.CandleStick(context.TODO(), "ethusdc", client.CandleStickOptions{})

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Trades

```go
res, err :=  c.Trades(context.TODO(), "ethusdc", client.TradesOptions{})

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Login

```go
res, err := c.Login(context.TODO(), <ETH_ADDRESS>, <ETH_PRIVATE_KEY>)

if err != nil {
	panic(err)
}

log.Printf("%+v", res)

```

# PNL

```go
res, err :=  c.PNL(context.TODO())

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)

```

# Balance

```go
res, err :=  c.Balance(context.TODO(), "eth")

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Profile

```go
res, err :=  c.Profile(context.TODO())

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Market Order

```go
order, err := c.OrderCreate(context.TODO(), starkPrivateKey, client.OrderNonceRequest{
		Market: "rplusdt",
		OrdType: "market",
		Side:    "sell",
		Volume:  0.01,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("%v", order)
	
```

# Limit Order

```go
order, err := c.OrderCreate(context.TODO(), starkPrivateKey, client.OrderNonceRequest{
		Market: "rplusdt",
		OrdType: "limit",
		Side:    "buy",
		Volume:  0.01,
		Price: 21.15,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("%v", order)
```

# List Order

```go
orders, err := c.OrdersList(context.TODO(), client.OrderListOptions{})

log.Printf("%v", orders)
```

# Cancel Order

```go
res, err := c.OrderCancel(context.TODO(), 261558257)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Retrieve Order

```go
res, err := c.OrderGet(context.TODO() , 261558257)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Set Allowance

```go
err = c.SetEthereumAllowance(context.TODO(), rpcURL, ethPrivateKey, "usdc", 0)
	if err != nil {
		panic(err)
	}
```

# Deposit ETH

```go
res, err := c.DepositFromEthereumNetwork(context.TODO(), rpcURL, ethAddr, ethPrivateKey, starkPublicKey, 5, client.USDT)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Deposit Polygon

```go
res, err := c.DepositFromPolygonNetwork(context.TODO(),rpcURL,ethAddr,ethPrivateKey,starkPublicKey, client.MATIC, 1)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Deposit Withdrawals

```go
res, err :=  c.ListAllDeposits(context.TODO())

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Normal Withdrawals

```go
res, err := c.InitiateNormalWithdrawal(context.TODO(), starkPrivateKey, 5, client.USDC)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Pending Withdrawals

```go
res, err :=  c.GetPendingNormalWithdrawalAmountByCoin(context.TODO(), rpcURL, ethAddr, client.BTC)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)

```

# Complete Withdrawal

```go
res, err :=  c.CompleteNormalWithdrawal(context.TODO(), rpcURL, ethPrivateKey, ethAddr, client.BTC)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)

```

# Fast Withdrawal

```go
res, err :=  c.FastWithdrawal(context.TODO(), starkPrivateKey, 10, client.USDT, client.POLYGON)

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# List Fast Withdrawal

```go
res, err :=  c.ListFastWithdrawal(context.TODO(), client.ListWithdrawalRequest{})

	if err != nil {
		panic(err)
	}

	log.Printf("%v", res)
```

# Trades WS

```go
package main

import (
	"log"

	"github.com/tanx-libs/tanx-connector-go/wsclient"
)

func main() {

	var stopCh chan struct{}
	var subUnsubCh chan wsclient.SubUnsubRequest

	ws := wsclient.New(wsclient.MAINET)

	wsSubUnsubEventHandler := func(event *wsclient.SubUnsubEvent) {
		log.Printf("SubUnsub event: %+v\n", event)
	}

	wsTradeEventHandler := func(event *wsclient.TradeEvent) {
		log.Printf("Trade event: %+v\n", event)

		// example
		subUnsubCh <- wsclient.SubUnsubTradeTopics(wsclient.SUBSCRIBE, []string{"btcusdc", "ethusdc"})
	}

	wsErrHandler := func(err error) {
		switch err {
		case err.(*wsclient.ErrSubUnsub):
			log.Println("SubUnsub error: ", err)
			stopCh <- struct{}{}

		default:
			log.Println("Default: ", err)
		}
	}

	doneCh, stopCh, subUnsubCh, err := ws.Trade([]string{"btcusdc", "ethusdc"}, wsTradeEventHandler, wsSubUnsubEventHandler, wsErrHandler)
	if err != nil {
		log.Fatal(err)
	}

	select {
	case <-stopCh:
		log.Println("stopped was called")

	case <-doneCh:
		log.Println("done was called")
	}
}

```