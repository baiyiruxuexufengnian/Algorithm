package hashMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonChars(t *testing.T) {
	words := []string{"bella","label","roller"}
	assert.Equal(t, []string{"e", "l", "l"}, commonChars(words))

	words = []string{"cool","lock","cook"}
	assert.Equal(t, []string{"c", "o"}, commonChars(words))

	words = []string{"abcabc","ab","a"}
	assert.Equal(t, []string{"a"}, commonChars(words))

	words = []string{"ab","a", "abcabc"}
	assert.Equal(t, []string{"a"}, commonChars(words))
	var cnt1 [26]int = [26]int{1,2}
	var cnt2 [26]int = [26]int{1,2}
	t.Log(cnt1 == cnt2)
}

