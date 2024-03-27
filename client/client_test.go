package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// using to simulate a roundtrip for a http network call
type MockRoundTripper struct {
	Response *http.Response
	Delay    time.Duration
	Err      error
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	select {
	case <-req.Context().Done():
		return m.Response, req.Context().Err()

	case <-time.After(m.Delay):
		return m.Response, m.Err
	}
}

// using when a funtion is dependent on multiple API calls underneath and we want to test the function without actually making the API calls
type roundTripFunc func (r *http.Request) (*http.Response, error)

func (s roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return s(r)
}


func TestClient(t *testing.T) {
	_, err := New(MAINNET)
	assert.NoError(t, err)

	c, err := New(TESTNET)
	assert.NoError(t, err)

	err = c.CheckAuth()
	assert.Error(t, err)
}
