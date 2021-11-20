package stringClass

import "testing"

func TestReverseString(t *testing.T) {
	tt := []byte("hello")
	reverseString(tt)
	t.Log(string(tt))
}
