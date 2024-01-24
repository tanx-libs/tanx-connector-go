package client

import (
	"fmt"
	"log"
	"testing"

	"github.com/tanx-libs/tanx-connector-go/types"

	"github.com/stretchr/testify/assert"
)

func TestClient_Health(t *testing.T) {
	client := New(types.BASE_URL).Timeout(10)
	assert.NotNil(t, client)

	got, err := client.Health()
	assert.NoError(t, err)

	want := types.HealthResponse{
		Status:  "success",
		Message: "Working fine!",
		Payload: "",
	}

	assert.Equal(t, want, got)
	t.Logf("got = %+v", got)
}

func BenchmarkClient_Health(b *testing.B) {
	client := New(types.BASE_URL)

	for i := 0; i < b.N; i++ {
		client.Health()
	}
}

func ExampleClient_Health() {
	client := New(types.BASE_URL)
	resp, err := client.Health()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", resp.Status)
	// Output: success
}
