package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanx-libs/tanx-connector-go/crypto_cpp"
)

func TestSign(t *testing.T) {
    r, s := crypto_cpp.Sign("0x1", "0x2", "0x3")
	public_key := crypto_cpp.GetPublicKey("0x1")
	ok := crypto_cpp.Verify(public_key, "0x2", r, s)
	assert.True(t, ok)
}
