package stringClass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	assert.Equal(t, "cdefgab", reverseLeftWords("abcdefg", 2))
	assert.Equal(t, "umghlrlose", reverseLeftWords("lrloseumgh", 6))
}
