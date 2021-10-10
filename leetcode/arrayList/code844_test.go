package arrayList

import "testing"

//"bxj##tw"
//"bxo#j##tw"

func TestBackspace(t *testing.T) {
	t1 := "bxj##tw"
	t2 := "bxo#j##tw"
	r := backspaceCompare(t1, t2)
	t.Logf("res: %v", r)
}
