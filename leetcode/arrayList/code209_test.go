package arrayList

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	test := []int{2,3,1,2,4,3}
	//test := []int{1,2,3,4,5}
	res := minSubArrayLen(7, test)
	t.Logf("res: %v", res)
}
