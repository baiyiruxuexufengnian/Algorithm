package arrayList

import "testing"

func TestSpiralOrder(t *testing.T) {
	// arr := [][]int{{1,2,3,4}, {5,6,7,8}, {9,10,11,12}}
	// arr := [][]int{{1,2,3,4}, {5,6,7,8}}
	arr := [][]int{{1,2,3}, {4,5,6}, {7,8,9}}
	// arr := [][]int{{1,2}}
	// arr := [][]int{{3}, {2}}
	ret := spiralOrder(arr)
	t.Logf("result: %v", ret)
}
