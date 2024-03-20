package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tanx-libs/tanx-connector-go/crypto_lib"
)

func TestSign(t *testing.T) {
    r, s := crypto_lib.Sign("0x1", "0x2", "0x3")
	public_key := crypto_lib.GetPublicKey("0x1")
	ok := crypto_lib.Verify(public_key, "0x2", r, s)
	assert.True(t, ok)
}
