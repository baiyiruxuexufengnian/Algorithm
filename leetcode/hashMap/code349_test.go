package hashMap

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	nums1 := []int {1,2,3,4,2,3}
	nums2 := []int {2, 3, 3}
	ret := intersection(nums1, nums2)
	t.Log(ret)
}
