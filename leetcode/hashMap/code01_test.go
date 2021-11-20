package hashMap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTowSum(t *testing.T) {
	in := []int {2,7,11,15}
	target := 9
	assert.Equal(t, []int{0,1}, twoSum(in, target))

	in = []int {3,2,4}
	target = 6
	assert.Equal(t, []int{1,2}, twoSum(in, target))

	in = []int {3,3}
	target = 6
	assert.Equal(t, []int{0,1}, twoSum(in, target))

	in = []int {3,3,3}
	target = 6
	assert.Equal(t, []int{0,2}, twoSum(in, target))
}
