package leetcode

import "testing"

func TestStrStr2(t *testing.T) {
	t.Log(strStr("hello", "ll"))
	t.Log(strStr("aaaaa", "bba"))
	t.Log(strStr("", ""))
	t.Log(strStr("bcabcabc", "abc"))
	t.Log(strStr("a", "a"))
}
