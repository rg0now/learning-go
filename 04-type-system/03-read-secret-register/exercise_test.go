package readsecretregister

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseChannelControlRegister(t *testing.T) {
	a,b,c,d := parseChannelControlRegister(0x82abba19)
	assert.Equal(t, []byte{0x82, 0xab, 0xba, 0x19}, []byte{a,b,c,d})

	a,b,c,d = parseChannelControlRegister(0xdeadbeef)
	assert.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, []byte{a,b,c,d})

	a,b,c,d = parseChannelControlRegister(0x01234567)
	assert.Equal(t, []byte{0x01, 0x23, 0x45, 0x67}, []byte{a,b,c,d})
}
