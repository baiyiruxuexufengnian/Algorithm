package hashMap

import "testing"

func TestFourSumCount(t *testing.T) {
	t.Log(fourSumCount([]int{1,2}, []int{-1,-2}, []int{-1,2}, []int{0,2}))
	t.Log(fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
	t.Log(fourSumCount([]int{-1,-1}, []int{-1,1}, []int{-1,1}, []int{1,-1}))
}
