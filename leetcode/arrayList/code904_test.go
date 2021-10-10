package arrayList

import "testing"

func TestTotalFruit(t *testing.T) {
	//in := []int{6,2,1,1,3,6,6}
	//in := []int{3,3,3,1,2,1,1,2,3,3,4}
	//in := []int{0,1,0,2}
	//in := []int{1,0,2,3,4}
	in := []int{1,2,3,2,2}
	r := totalFruit(in)
	t.Logf("result: %v", r)
}
