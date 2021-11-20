package stringClass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"))
	assert.Equal(t, "world hello", reverseWords("  hello world  "))
}
