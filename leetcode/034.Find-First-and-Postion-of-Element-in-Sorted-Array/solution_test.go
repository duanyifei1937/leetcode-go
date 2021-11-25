package leetcode

import "testing"

func TestSearchRange034(t *testing.T) {
	t.Log(searchRange034([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(searchRange034([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(searchRange034([]int{}, 0))
}
