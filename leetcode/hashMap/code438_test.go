package hashMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	//s := "cbaebabacd"
	//p := "abc"
	//assert.Equal(t, []int{0,6}, findAnagrams(s, p))
	//
	//s = "abab"
	//p = "ab"
	//assert.Equal(t, []int{0,1,2}, findAnagrams(s, p))
	//
	//s = "cbaacb"
	//p = "abc"
	//assert.Equal(t, []int{0,3}, findAnagrams(s, p))

	s := "aa"
	p := "bb"
	assert.Equal(t, []int{}, findAnagrams(s, p))
}
