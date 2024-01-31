package wsclient

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

// EventHandler handle raw websocket message
type EventHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

type request struct {
	Event   string   `json:"event"`
	Streams []string `json:"streams"`
}

// Config webservice configuration
type Config struct {
	endpoint string
	request  request
	// can add more configurations here in the future
}

var serve = func(cfg *Config, eventHandler EventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, err error) {
	dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := dialer.Dial(cfg.endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	err = c.WriteJSON(cfg.request)
	if err != nil {
		return nil, nil, err
	}

	c.SetReadLimit(655350)
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopCh channel is
		// closed by the client.
		defer close(doneCh)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		// Wait for the stopCh channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopCh:
				silent = true
			case <-doneCh:
			}
			c.Close()
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			eventHandler(message)
		}
	}()
	return
}

/*
- This basically registers a handler for pong messages.
- Then after every "timeout" seconds, it will check if ping pong happend or not.
- If yes then it continues the behaviour thus keeping the connection alive.
- If no then the connection breaks.
*/
func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				c.Close()
				return
			}
		}
	}()
}
