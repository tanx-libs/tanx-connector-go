package wsclient

import (
	"errors"
	"fmt"
)

// Not logged in error
var ErrNotLoggedIn = errors.New("not logged in")

// json unmarshaling error
type ErrJSONUnmarshal struct {
	Msg string
	Err error
}

func (e *ErrJSONUnmarshal) Error() string {
	return fmt.Sprintf("Debug: %s\nError: %s\n", e.Msg, e.Err)
}

// websocket read message error
type ErrWsReadMessage struct {
	Msg string
	Err error
}

func (e *ErrWsReadMessage) Error() string {
	return fmt.Sprintf("Debug: %s\nError: %s\n", e.Msg, e.Err)
}

// Subscribe Unsubscribe erros
type ErrSubUnsub struct {
	Msg string
	Err error
}

func (e *ErrSubUnsub) Error() string {
	return fmt.Sprintf("Debug: %s\nError: %s\n", e.Msg, e.Err)
}