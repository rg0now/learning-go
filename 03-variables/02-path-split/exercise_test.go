package pathsplit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPath(t *testing.T) {
	assert.Equal(t, splitPath("static/js/jquery.js"), "jquery.js")
	assert.Equal(t, splitPath("multi_langs.js"), "multi_langs.js")
	assert.Equal(t, splitPath("css/style.css"), "style.css")
	assert.Equal(t, splitPath("css/flags.css"), "flags.css")
	assert.Equal(t, splitPath("image/favicon.ico"), "favicon.ico")
}
