package leetcode

import "testing"

func TestCanConstruct(t *testing.T) {
	t.Log(canConstruct("a", "b"))
	t.Log(canConstruct("aa", "ab"))
	t.Log(canConstruct("aa", "aab"))
}

func TestCanConstruct2(t *testing.T) {
	t.Log(canConstruct2("a", "b"))
	t.Log(canConstruct2("aa", "ab"))
	t.Log(canConstruct2("aa", "aab"))
}
