package wsclient

import (
	"context"
	"encoding/json"
	"log"
	"net/url"

	"github.com/tanx-libs/tanx-connector-go/types"
	"nhooyr.io/websocket"
)

type WsClient struct {
	baseURL     *url.URL
	closeCh     chan struct{}
	TradeStream chan types.TradeStreamResponse
	// OrderBookStream   chan types.OrderBookStreamResponse
	// CandlestickStream chan types.CandlestickStreamResponse
	ws          *websocket.Conn
}

/*
This checks connection to the websocket server and return error if it fails
*/
func New(ctx context.Context) (*WsClient, error) {
	url, err := url.Parse(types.WS_BASE_URL)
	if err != nil {
		return nil, err
	}

	c, _, err := websocket.Dial(ctx, url.String(), nil)
	if err != nil {
		return nil, err
	}

	// TODO
	wsclient := &WsClient{
		baseURL:     url,
		closeCh:     make(chan struct{}),
		TradeStream: make(chan types.TradeStreamResponse),
		ws:          c,
	}

	go wsclient.listen(ctx)

	return wsclient, nil
}

func (c *WsClient) BaseURL(baseURL string) (*WsClient, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c.baseURL = url
	return c, nil
}

func (c *WsClient) Subscribe(ctx context.Context, streams []string) error {
	req := types.SubUnsubRequest{
		Event:   "subscribe",
		Streams: streams,
	}

	err := c.checkSubUnsubStatus(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *WsClient) UnSubscribe(ctx context.Context, streams []string) error {
	req := types.SubUnsubRequest{
		Event:   "unsubscribe",
		Streams: streams,
	}

	err := c.checkSubUnsubStatus(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *WsClient) checkSubUnsubStatus(ctx context.Context, req types.SubUnsubRequest) error {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = c.ws.Write(ctx, websocket.MessageText, reqBytes)
	if err != nil {
		return err
	}

	return nil
}

func (c *WsClient) listen(ctx context.Context) {
	go func() {
		for {
			// logging success message
			_, data, err := c.ws.Read(ctx)
			if err != nil {
				log.Println(err)
				c.closeCh <- struct{}{}
			}

			log.Println(string(data))
		}
	}()

	for {
		select {
		// TODO
		case <-c.closeCh:
			log.Println("closing websocket connection")
			c.ws.Close(websocket.StatusNormalClosure, "closing websocket connection")
			return

		// TODO
		case <-ctx.Done():
			log.Println("context done triggered")
			c.ws.Close(websocket.StatusNormalClosure, "context done")
			return
		}
	}
}

func (c *WsClient) Close() {
	c.closeCh <- struct{}{}
}
