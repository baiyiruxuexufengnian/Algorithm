package hashMap

import "testing"

func TestIsAnagram(t *testing.T) {
	s := "rat"
	tt := "car"
	t.Logf("result: %v", isAnagram(s, tt))
}
