package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}

func TestReverseString2(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString2(s)
	t.Log(string(s))
}
