package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// helper struct for mocking http.RoundTripper
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

func TestClient(t *testing.T) {
	t.Run("basic client initialization", func(t *testing.T) {
		_, err := New(MAINET)
		assert.NoError(t, err)

		_, err = New(TESTNET)
		assert.NoError(t, err)
	})
}
