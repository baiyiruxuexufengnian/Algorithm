package longestSubstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		testStr string
		expect int
	}{
		{
			testStr: "pwwkew",
			expect: 3,
		},
		{
			testStr: "",
			expect: 0,
		},
		{
			testStr: "abcabcbb",
			expect: 3,
		},
		{
			testStr: "bbbbb",
			expect: 1,
		},
		{
			testStr: "dvdf",
			expect: 3,
		},
		{
			testStr: "tmmzuxt",
			expect: 5,
		},
	}

	for _, tt := range tests {
		actual := Solution(tt.testStr)
		assert.Equal(t, tt.expect, actual)
	}
}