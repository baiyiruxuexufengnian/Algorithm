package arrayList

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{0,1,2,2,3,0,4,2}
	r := removeElement2(nums, 2)
	t.Logf("r : %v", r)
	t.Logf("nums: %v", nums)
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0}
	r := removeDuplicates2(nums)
	t.Logf("r: %v", r)
	t.Logf("nums: %v", nums)
}
