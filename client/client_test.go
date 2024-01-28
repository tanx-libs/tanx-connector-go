package client

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tanx-libs/tanx-connector-go/types"
)

// helper struct for mocking http.RoundTripper
type MockRoundTripper struct {
	Response *http.Response
	Delay   time.Duration
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
		client, err := New()
		assert.NoError(t, err)

		assert.Equal(t, types.JSON_BASE_URL, client.baseURL.String())
		assert.Equal(t, types.JSON_BASE_URL + types.HEALTH_PATH, client.healthURL.String())
		assert.Equal(t, types.JSON_BASE_URL + types.TICKER_PATH, client.tickerURL.String())

		client = client.HealthPath("/health").TickerPath("/ticker")
		assert.Equal(t, types.JSON_BASE_URL + "health", client.healthURL.String())
		assert.Equal(t, types.JSON_BASE_URL + "ticker", client.tickerURL.String())
	})
}
