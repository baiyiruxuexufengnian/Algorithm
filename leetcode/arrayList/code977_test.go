package arrayList

import "testing"

func TestSortedSquares(t *testing.T) {
	nums := []int{-4,-3,-2, -1, 0}
	ret := sortedSquares2(nums)
	t.Logf("ret: %v", ret)
}
