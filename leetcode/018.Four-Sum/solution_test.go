package leetcode

import "testing"

func TestFourSum(t *testing.T) {
	t.Log(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	t.Log(fourSum([]int{2, 2, 2, 2, 2}, 8))
	t.Log(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
	t.Log(fourSum2([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
