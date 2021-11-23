package leetcode

import "testing"

func TestSearchInsert035(t *testing.T) {
	t.Log(searchInsert035([]int{1, 3, 5, 6}, 5))
	t.Log(searchInsert035([]int{1,3,5,6}, 2))
	t.Log(searchInsert035([]int{1,3,5,6}, 7))
	t.Log(searchInsert035([]int{1,3,5,6}, 0))
	t.Log(searchInsert035([]int{1}, 0))
}
