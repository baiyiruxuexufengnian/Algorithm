package arrayList

import "testing"

func TestSolution(t *testing.T) {
	test := []int {1,0,2,3}
	res := Solution(test, 3)
	t.Log(res)
}
