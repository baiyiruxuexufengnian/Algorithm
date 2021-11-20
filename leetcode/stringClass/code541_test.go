package stringClass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseStr(t *testing.T)  {
	assert.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	assert.Equal(t, "bacd", reverseStr("abcd", 2))

	assert.Equal(t, "cbad", reverseStr("abcd", 3))
}
