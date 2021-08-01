package GreedyAlgorithm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindContentChildren(t *testing.T)  {
	tests := []struct{
		children []int
		cookies  []int
		expect 	 int
	}{
		{
			children: []int{2, 1},
			cookies:  []int{3, 2, 1},
			expect: 2,
		},
		{
			children: []int{2, 1, 3},
			cookies:  []int{3, 2, 1, 4, 5},
			expect: 3,
		},
		{
			children: []int{2, 1, 3},
			cookies:  []int{3},
			expect: 1,
		},
	}

	for _, tt := range tests {
		actual := findContentChildren(tt.children, tt.cookies)
		assert.Equal(t, tt.expect, actual)
	}
}
