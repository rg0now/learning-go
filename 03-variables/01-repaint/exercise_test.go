package repaint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepaintColor(t *testing.T) {
	c, err := repaintColor("vermilion")
	assert.Equal(t, c, "teal")
	assert.NoError(t, err)

	c, err = repaintColor("teal")
	assert.Equal(t, c, "vermilion")
	assert.NoError(t, err)

	c, err = repaintColor("xxx")
	assert.Equal(t, c, "")
	assert.Error(t, err)
}
