package leetcode

import "testing"

func TestMyStack(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	t.Log(obj.Top())
	t.Log(obj.Pop())
	t.Log(obj.Empty())
}
