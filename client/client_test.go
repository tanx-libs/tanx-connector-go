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
		_, err := New()
		assert.NoError(t, err)

		// assert.Equal(t, JSON_BASE_URL, client.baseURL.String())
		// assert.Equal(t, JSON_BASE_URL+HEALTH_PATH, client.healthURL.String())
		// assert.Equal(t, JSON_BASE_URL+TICKER_PATH, client.tickerURL.String())

		// client = client.HealthPath("/health").TickerPath("/ticker")
		// assert.Equal(t, JSON_BASE_URL+"health", client.healthURL.String())
		// assert.Equal(t, JSON_BASE_URL+"ticker", client.tickerURL.String())
	})
}
