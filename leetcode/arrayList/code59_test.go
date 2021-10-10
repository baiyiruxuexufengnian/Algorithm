package arrayList

import "testing"

func TestGenerateMatrix(t *testing.T) {
	ret := generateMatrix(4)
	t.Logf("ret: %v", ret)
}
