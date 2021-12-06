package leetcode

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	t.Log(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	t.Log(minSubArrayLen(4, []int{1, 4, 4}))
}
