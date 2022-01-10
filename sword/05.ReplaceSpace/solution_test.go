package sword

import "testing"

func TestReplaceSpace(t *testing.T) {
	t.Log(replaceSpace("a b c"))
}

func TestReplaceSpace2(t *testing.T) {
	t.Log(replaceSpace2("a b c"))
}

func TestReplaceSpace3(t *testing.T) {
	t.Log(replaceSpace3("a b c"))
	t.Log(replaceSpace3("We are happy."))
}
