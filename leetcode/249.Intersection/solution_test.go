package leetcode

import "testing"


func TestIntersection(t *testing.T) {
	r := intersection([]int{1,2,2,1}, []int{2,2})
	r1 := intersection([]int{4,9,5}, []int{9,4,9,8,4})
	r2 := intersection([]int{4,9,5}, []int{})
	t.Log(r, r1, r2)
}
