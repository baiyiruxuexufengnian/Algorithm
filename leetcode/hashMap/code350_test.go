package hashMap

import "testing"

func TestIntersect(t *testing.T) {
	nums1 := []int {4,9,5}
	nums2 := []int {9,4,9,8,4}
	ret := intersect(nums1, nums2)
	t.Log(ret)
}
