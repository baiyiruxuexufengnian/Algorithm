package arrayList

import "testing"

func TestMoveZeroes(t *testing.T) {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	t.Logf("nums: %v", nums)
}
