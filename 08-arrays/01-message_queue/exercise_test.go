package messagequeue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageQueue(t *testing.T) {
	x := messageQueue("Hello", "World", "Again")
	assert.Equal(t, x[0], "Hello")
	assert.Equal(t, x[1], "Again")
	assert.Equal(t, x[2], "World")
}
