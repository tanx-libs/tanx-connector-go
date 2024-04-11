package wsclient

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type SubUnsub string

const (
	SUBSCRIBE   SubUnsub = "subscribe"
	UNSUBSCRIBE SubUnsub = "unsubscribe"
)

type Base string

const (
	MAINET  Base = "wss://api.tanx.fi"
	TESTNET Base = "wss://api-testnet.tanx.fi"
)

const (
	PUBLIC_PATH  = "/public"
	PRIVATE_PATH = "/private"

	SUB_UNSUB_SUCCESS = "success"
	SUB_UNSUB_ERROR   = "error"
)

type Wsclient struct {
	baseURL    string
	publicURL  string
	privateURL string

	accessToken  string
	refreshToken string
}


func New(baseURL Base) *Wsclient {
	return &Wsclient{
		baseURL:    string(baseURL),
		publicURL:  string(baseURL) + PUBLIC_PATH,
		privateURL: string(baseURL) + PRIVATE_PATH,

		accessToken:  "",
		refreshToken: "",
	}
}

func (c *Wsclient) SetJwt(jwtToken string, refreshToken string) {
	c.accessToken = jwtToken
	c.refreshToken = refreshToken
}

var (
	WebsocketTimeout = time.Second * 60

	WebsocketKeepalive = false
)

type SubUnsubRequest struct {
	Event   SubUnsub `json:"event"`
	Streams []string `json:"streams"`
}

type Config struct {
	endpoint        string
	subUnsubRequest SubUnsubRequest
}

type EventHandler func(message []byte)

type ErrHandler func(err error)

type Success struct {
	Message string   `json:"message"`
	Streams []string `json:"streams"`
}

type SubUnsubEvent struct {
	Success Success `json:"success"`
}

type SubUnsubEventHandler func(event *SubUnsubEvent)

func serve(cfg *Config, eventHandler EventHandler, errHandler ErrHandler) (doneCh, stopCh chan struct{}, subUnsubCh chan SubUnsubRequest, err error) {
	dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := dialer.Dial(cfg.endpoint, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	err = c.WriteJSON(cfg.subUnsubRequest)
	if err != nil {
		return nil, nil, nil, &ErrSubUnsub{
			Msg: "Error when sending sub/unsub request to the server",
			Err: err,
		}
	}

	c.SetReadLimit(655350)
	doneCh = make(chan struct{})
	stopCh = make(chan struct{})
	subUnsubCh = make(chan SubUnsubRequest, 1)
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
			select {
			case subUnsubRequest := <-subUnsubCh:
				err := c.WriteJSON(subUnsubRequest)
				if err != nil && !silent {
					errHandler(&ErrSubUnsub{
						Msg: "Error when sending sub/unsub request to the server",
						Err: err,
					})
				}
			default:
				_, message, err := c.ReadMessage()
				if err != nil {
					if !silent {
						errHandler(&ErrWsReadMessage{
							Msg: "Error when reading message from the server",
							Err: err,
						})
					}
					return
				}
				eventHandler(message)
			}
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
