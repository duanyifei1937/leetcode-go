package leetcode

import "testing"

func TestRemoveElement027(t *testing.T) {
	// 	输入：nums = [3,2,2,3], val = 3
	// 输出：2, nums = [2,2]
	t.Log(removeElement027([]int{3, 2, 2, 3}, 3))


	// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
	// 输出：5, nums = [0,1,4,0,3]
	t.Log(removeElement027([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}


func TestRemoveElement027_2(t *testing.T) {
	// 	输入：nums = [3,2,2,3], val = 3
	// 输出：2, nums = [2,2]
	t.Log(removeElement027_2([]int{3, 2, 2, 3}, 3))


	// 输入：nums = [0,1,2,2,3,0,4,2], val = 2
	// 输出：5, nums = [0,1,4,0,3]
	t.Log(removeElement027_2([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}	