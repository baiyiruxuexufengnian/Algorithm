package arrayList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	nums := []int{3, 5, 7, 9, 10}
	target := 8
	r := searchInsert(nums, target)
	assert.Equal(t, 3, r)
}
