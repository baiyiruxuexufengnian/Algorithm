package stringClass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	assert.Equal(t, "We%20are%20happy.", replaceSpace("We are happy."))
	assert.Equal(t, "%20", replaceSpace(" "))
}
